FROM golang:latest

RUN mkdir -p /go/src/github.com/hengel2810/api_docli
WORKDIR /go/src/github.com/hengel2810/api_docli
RUN mkdir shared
COPY . .
RUN go get "github.com/gorilla/mux"
RUN go get -u -v "github.com/docker/docker/client"
RUN go get "golang.org/x/net/context"
RUN go build main.go
ENV GOPATH=/go/src
CMD ["/go/src/github.com/hengel2810/api_docli/main"]