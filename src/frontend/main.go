// Copyright 2019 Yoshi Yamaguchi, 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
	pb "github.com/ymotongpoo/stackdriver-trace-log-demo/src/frontend/genproto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	port         = "8080"
	cookieMaxAge = 60 * 60 * 24
	initMaxRetry = 3
)

var (
	logger *zap.SugaredLogger

	arrayParseSvcAddr string
	arrayParseSvcConn *grpc.ClientConn

	addNumberSvcAddr string
	addNumberSvcConn *grpc.ClientConn
)

func main() {
	ctx := context.Background()
	initLogger()
	go initTracing()
	go initProfiling("frontend", "1.0.0")

	srvPort := port
	if p := os.Getenv("PORT"); p != "" {
		srvPort = p
	}
	addr := os.Getenv("LISTEN_ADDR")
	mustMapEnv(&arrayParseSvcAddr, "ARRAY_PARSE_SERVICE_ADDR")
	mustMapEnv(&addNumberSvcAddr, "ADD_NUMBER_SERVICE_ADDR")
	mustConnGRPC(ctx, &arrayParseSvcConn, arrayParseSvcAddr)
	mustConnGRPC(ctx, &addNumberSvcConn, addNumberSvcAddr)

	e := echo.New()
	e.GET("/", homeGetHandler)
	e.HEAD("/", homeHeadHandler)
	e.GET("/_healthz", healthHandler)

	e.Start("0.0.0.0:" + srvPort)
	logger.Infof("starting server on " + addr + ":" + srvPort)
}

// ---- handlers ----

func homeGetHandler(c echo.Context) error {
	numStr := c.QueryParam("num")
	if numStr == "" {
		return c.String(http.StatusBadRequest, "add `num` URL parameter with comma separated int values")
	}

	// Start a trace span.
	ctx, span := trace.StartSpan(c.Request().Context(), "frontend")
	defer span.End()

	parseReq := &pb.ParseRequest{
		TargetStr: numStr,
	}
	logger.Infof("[homeGetHandler] call arrayparse service: %v", parseReq.String())
	apSvc := pb.NewArrayParseServiceClient(arrayParseSvcConn)
	pa, err := apSvc.Parse(ctx, parseReq)
	if err != nil {
		logger.Errorf("[homeGetHandler] %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	addReq := &pb.AddRequest{
		Numbers: pa.GetNumbers(),
	}
	logger.Infof("[homeGetHandler] call addnumber service: %v", addReq.String())
	anSvc := pb.NewAddNumberServiceClient(addNumberSvcConn)
	ar, err := anSvc.Add(ctx, addReq)
	if err != nil {
		logger.Errorf("[homeGetHandler] %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	arStr := strconv.FormatInt(ar.GetNumber(), 10)
	return c.String(http.StatusOK, arStr)
}

func homeHeadHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// ---- init funcitons ----

func initLogger() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "severity",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			TimeKey:        "timestamp",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	var err error
	var l *zap.Logger
	l, err = cfg.Build()
	if err != nil {
		log.Fatal("failed to create logger: %v", err)
	}
	logger = l.Sugar()
}

func initStats(log *zap.SugaredLogger, exporter *stackdriver.Exporter) {
	view.SetReportingPeriod(60 * time.Second)
	view.RegisterExporter(exporter)
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Warn("Error registering http default server views")
	} else {
		log.Info("Registered http default server views")
	}
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Warn("Error registering grpc default client views")
	} else {
		log.Info("Registered grpc default client views")
	}
}

func initTracing() {
	// This is a demo app with low QPS. trace.AlwaysSample() is used here
	// to make sure traces are available for observation and analysis.
	// In a production environment or high QPS setup please use
	// trace.ProbabilitySampler set at the desired probability.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	for i := 1; i <= initMaxRetry; i++ {
		log := logger.With("retry", i)
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			log.Fatalf("failed to initialize stackdriver exporter: %+v", err)
		} else {
			trace.RegisterExporter(exporter)
			log.Info("registered stackdriver tracing")

			// Register the views to collect server stats.
			initStats(logger, exporter)
			return
		}
		d := time.Second * 20 * time.Duration(i)
		log.Debugf("sleeping %v to retry initializing stackdriver exporter", d)
		time.Sleep(d)
	}
	logger.Warn("could not initialize stackdriver exporter after retrying, giving up")
}

func initProfiling(service, version string) {
	for i := 1; i <= initMaxRetry; i++ {
		log := logger.With("retry", i)
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			// ProjectID must be set if not running on GCP.
			// ProjectID: "my-project",
		}); err != nil {
			log.Warnf("warn: failed to start profiler: %+v", err)
		} else {
			log.Info("started stackdriver profiler")
			return
		}
		d := time.Second * 10 * time.Duration(i)
		log.Debugf("sleeping %v to retry initializing stackdriver profiler", d)
		time.Sleep(d)
	}
	logger.Warn("warning: could not initialize stackdriver profiler after retrying, giving up")
}

// ---- helpers ----

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		log.Fatalf("environment variable %q not set", envKey)
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
