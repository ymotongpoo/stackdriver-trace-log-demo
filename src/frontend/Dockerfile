FROM golang:1.14 as builder
RUN apt-get -y install --no-install-recommends \
  ca-certificates git

ENV PROJECT github.com/ymotongpoo/stackdriver-trace-log-demo/src/frontend
WORKDIR /go/src/$PROJECT
COPY . .
# restore dependencies
ENV GO111MODULE=on
ENV GOPROXY="https://proxy.golang.org,direct"
RUN go build -mod=vendor -o /go/bin/frontend

FROM gcr.io/distroless/base-debian10 as release
WORKDIR /frontend
COPY --from=builder /go/bin/frontend /frontend/server
EXPOSE 8080
ENTRYPOINT ["/frontend/server"]
