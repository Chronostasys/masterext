#build stage
FROM golang:1.17 AS builder
WORKDIR /go/src/app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v main.go



FROM ubuntu:18.04
RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get clean
RUN apt-get update
RUN apt-get install -y libssl-dev
RUN apt-get install -y libgssapi-krb5-2
RUN apt-get install -y ca-certificates
COPY --from=builder /go/bin/app /bin/app
COPY . /
RUN chmod +x /entry.sh
RUN chmod -R +x /fastgithub_linux-x64
ENTRYPOINT ["/entry.sh"]