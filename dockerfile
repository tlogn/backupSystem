FROM golang:latest
WORKDIR /go/src/app
ADD . /go/src/app
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy
EXPOSE 8090
EXPOSE 10001
CMD go run main.go