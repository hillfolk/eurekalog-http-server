FROM golang:1.11 as goimage
ENV SRC=/go/src/
ENV GO111MODULE=on
RUN mkdir -p /go/src/
WORKDIR /go/src/github.com/hillfolk/eurekalog-http-server
RUN git clone -b master --single-branch https://github.com/hillfolk/eurekalog-http-server.git /go/src/github.com/hillfolk/eurekalog-http-server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o bin/eurekalog-http-server


FROM alpine:3.9 as baseimagealp
RUN apk add --no-cache bash
ENV WORK_DIR=/docker/bin
WORKDIR $WORK_DIR
COPY --from=goimage /go/src/github.com/hillfolk/eurekalog-http-server/bin/ ./
ENTRYPOINT /docker/bin/eurekalog-http-server
EXPOSE 8080