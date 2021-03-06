FROM golang:1.11.4-alpine as builder
RUN apk add --no-cache ca-certificates git build-base

ENV PROJECT github.com/ymotongpoo/stackdriver-trace-log-demo/src/addnumber
WORKDIR /go/src/$PROJECT
COPY . .

# restore dependencies
ENV GO111MODULE=on
RUN go get
RUN go build -mod=vendor

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
        chmod +x /bin/grpc_health_probe
WORKDIR /addnumber
COPY --from=builder /go/bin/addnumber /addnumber/server
EXPOSE 4040
ENTRYPOINT ["/addnumber/server"]
