module github.com/ymotongpoo/stackdriver-trace-log-demo/src/addnumber

go 1.14

replace cloud.google.com/go => ../../../google-cloud-go

require (
	cloud.google.com/go v0.57.0
	contrib.go.opencensus.io/exporter/stackdriver v0.8.0
	github.com/golang/protobuf v1.4.2
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.2.0
	github.com/ymotongpoo/stackdriver-trace-log-demo/src/arrayparse v0.0.0-20190121061800-f1fd8a58ac42
	go.opencensus.io v0.22.3
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/grpc v1.29.1
)
