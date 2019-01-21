// Copyright 2019 Yoshi Yamaguchi
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
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	pb "github.com/ymotongpoo/stackdriver-trace-log-demo/src/arrayparse/genproto"
)

const (
	listenPort       = "4040"
	initMaxRetry     = 3
	traceLogFieldKey = "logging.googleapis.com/trace"
)

var (
	logger *logrus.Logger
)

func main() {
	initLogger()
	go initTracing()
	go initProfiling("addnumber", "1.0.0")

	port := listenPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := new(addNumberServiceServer)
	pb.RegisterAddNumberServiceServer(srv, svc)
	healthpb.RegisterHealthServer(srv, svc)
	logger.Infof("starting to listen on tcp: %q", lis.Addr().String())
	err = srv.Serve(lis)
	logger.Fatal(err)
}

// ---- gRPC ----

type addNumberServiceServer struct{}

func (an *addNumberServiceServer) Add(ctx context.Context, ar *pb.AddRequest) (*pb.AddResult, error) {
	// Extract TraceID from parent context
	// https://cloud.google.com/logging/docs/agent/configuration#special-fields
	span := trace.FromContext(ctx)
	sc := span.SpanContext()
	l := logger.WithField(traceLogFieldKey, sc.TraceID)

	l.Info("Start Add")
	nums := ar.GetNumbers()
	total := int64(0)
	for _, n := range nums {
		total += n
	}
	l.Infof("End Add: %v", nums)
	result := pb.AddResult{
		Number: total,
	}
	return &result, nil
}

func (an *addNumberServiceServer) Check(ctx context.Context, r *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{
		Status: healthpb.HealthCheckResponse_SERVING,
	}, nil
}

func (an *addNumberServiceServer) Watch(r *healthpb.HealthCheckRequest, _ healthpb.Health_WatchServer) error {
	return nil
}

// ---- init funcitons ----

func initLogger() {
	logger = logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	}
	logger.Out = os.Stdout
}

func initStats(log logrus.FieldLogger, exporter *stackdriver.Exporter) {
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
		log := logger.WithField("retry", i)
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			logger.Fatalf("failed to initialize stackdriver exporter: %+v", err)
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
		log := logger.WithField("retry", i)
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
