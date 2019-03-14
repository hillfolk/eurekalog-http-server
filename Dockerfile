FROM golang:1.11 as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/eurekalog
RUN git clone -b master --single-branch https://github.com/hillfolk/eurekalog-http-server.git /go/src/eurekalog/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o bin/eurekalog


FROM alpine:3.9 as baseimagealp
RUN apk add --no-cache bash
ENV WORK_DIR=/docker/bin
WORKDIR $WORK_DIR
COPY --from=goimage /go/src/eurekalog/bin/ ./
ENTRYPOINT /docker/bin/eurekalog
EXPOSE 8080