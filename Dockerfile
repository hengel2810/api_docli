FROM golang:latest
RUN mkdir shared
WORKDIR /go/src
COPY ./src .
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/docker/docker/client"
RUN go get "golang.org/x/net/context"
RUN go build main.go
ENV GOPATH=/go/src
CMD ["/go/src/main"]