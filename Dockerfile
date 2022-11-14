FROM golang:1.19.2 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://goproxy.cn
WORKDIR /app
ADD . /app
RUN cd /app/src \
    && buildflags="-X 'main.BuildTime=`date`' -X 'main.GitHead=`git rev-parse --short HEAD`' -X 'main.GoVersion=$(go version)'" \
    && go build -ldflags "$buildflags" -tags netgo -a -v -o /app/kubesphere-webhook-proxy-go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/kubesphere-webhook-proxy-go /app
CMD ["./kubesphere-webhook-proxy-go"]