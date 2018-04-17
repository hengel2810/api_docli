FROM golang:latest

RUN mkdir -p /go/src/github.com/hengel2810/api_docli
WORKDIR /go/src/github.com/hengel2810/api_docli
RUN mkdir shared
COPY . .
RUN curl -L -s https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 -o $GOPATH/bin/dep
RUN chmod +x $GOPATH/bin/dep
RUN dep ensure
RUN rm -rf $GOPATH/src/github.com/docker/docker/vendor/github.com/docker/go-connections
RUN go build main.go
ENV GOPATH=/go/src
CMD ["/go/src/github.com/hengel2810/api_docli/main"]