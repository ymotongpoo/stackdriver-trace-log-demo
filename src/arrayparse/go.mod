module github.com/ymotongpoo/stackdriver-trace-log-demo/src/arrayparse

go 1.14

replace cloud.google.com/go => ../../../google-cloud-go

require (
	cloud.google.com/go v0.57.0
	contrib.go.opencensus.io/exporter/stackdriver v0.8.0
	github.com/golang/protobuf v1.4.2
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.2.0
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/smartystreets/goconvey v0.0.0-20181108003508-044398e4856c // indirect
	go.opencensus.io v0.22.3
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/grpc v1.29.1
)
