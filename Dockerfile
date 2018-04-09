FROM golang:latest
WORKDIR /go/src
RUN mkdir shared
COPY ./src .
RUN go get "github.com/gorilla/mux"
RUN go get -u -v "github.com/docker/docker/client"
RUN go get "golang.org/x/net/context"
RUN go build main.go
ENV GOPATH=/go/src
CMD ["/go/src/main"]