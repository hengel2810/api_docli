FROM golang:latest

RUN mkdir -p /go/src/github.com/hengel2810/api_docli
WORKDIR /go/src/github.com/hengel2810/api_docli
RUN mkdir shared
COPY . .
RUN go get "github.com/gorilla/mux"
RUN echo "##########################"
RUN echo "##########################"
RUN echo "##########################"
RUN go get "github.com/docker/docker/client"
RUN echo "##########################"
RUN echo "##########################"
RUN rm -rf /go/src/github.com/docker/docker/vendor/github.com/docker/go-connections
RUN echo "##########################"
RUN echo "##########################"
RUN echo "##########################"
RUN go get "golang.org/x/net/context"
RUN go get "gopkg.in/mgo.v2"
RUN go get "github.com/auth0/go-jwt-middleware"
RUN go get "github.com/codegangsta/negroni"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/satori/go.uuid"
RUN go get "github.com/Pallinder/sillyname-go"
RUN go get "github.com/docker/go-connections/nat"
RUN go get "github.com/pkg/errors"
RUN go build main.go
ENV GOPATH=/go/src
CMD ["/go/src/github.com/hengel2810/api_docli/main"]