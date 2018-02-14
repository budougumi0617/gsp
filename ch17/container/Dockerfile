FROM golang:1.9.4-alpine3.7

RUN apk add --update openssl-dev pcre-dev git gcc musl-dev linux-headers sudo

WORKDIR /go/src/github.com/budougumi0617/gsp/ch17/container/
ADD _main.go ./main.go
ADD rootfs ./rootfs/
RUN go get -v .
RUN go build main.go
CMD echo "test"
