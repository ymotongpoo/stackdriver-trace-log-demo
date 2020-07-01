FROM golang:1.14 as builder
RUN apt-get -y install --no-install-recommends \
  ca-certificates git
ENV PROJECT github.com/ymotongpoo/stackdriver-trace-log-demo/src/arrayparse
WORKDIR /go/src/$PROJECT
COPY . .
# restore dependencies
ENV GO111MODULE=on
ENV GOPROXY="https://proxy.golang.org,direct"
RUN go build -mod=vendor -o /go/bin/arrayparse
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
        chmod +x /bin/grpc_health_probe

FROM gcr.io/distroless/base-debian10 as release
WORKDIR /arrayparse
COPY --from=builder /go/bin/arrayparse /arrayparse/server
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
EXPOSE 5050
ENTRYPOINT ["/arrayparse/server"]
