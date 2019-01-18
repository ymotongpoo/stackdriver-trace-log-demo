FROM golang:1.11.4-alpine as builder
RUN apk add --no-cache ca-certificates git build-base

ENV PROJECT github.com/ymotongpoo/stackdriver-trace-log-demo/src/frontend
WORKDIR /go/src/$PROJECT
COPY . .

# restore dependencies
ENV GO111MODULE=on
RUN go get
RUN go build

FROM alpine as release
RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools
WORKDIR /frontend
COPY --from=builder /go/bin/frontend /frontend/server
EXPOSE 8080
ENTRYPOINT ["/frontend/server"]
