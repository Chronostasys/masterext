#build stage
FROM golang:alpine AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v main.go



FROM alpine:3.6 as alpine
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add -U --no-cache ca-certificates
COPY --from=builder /go/bin/app /bin/app
COPY . /
ENTRYPOINT ["/entry.sh"]