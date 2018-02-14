FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc sudo

RUN wget -P /tmp https://storage.googleapis.com/golang/go1.9.4.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.9.4.linux-amd64.tar.gz
RUN rm /tmp/go1.9.4.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH
WORKDIR /go/src/github.com/budougumi0617/gsp/ch17/container/
ADD main.go .
ADD rootfs ./rootfs/
RUN go get -v .
RUN go build main.go
CMD echo "test"
