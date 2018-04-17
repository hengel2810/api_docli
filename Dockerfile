FROM golang:latest

RUN mkdir -p /go/src/github.com/hengel2810/api_docli
WORKDIR /go/src/github.com/hengel2810/api_docli
RUN mkdir shared
RUN go get github.com/gorilla/mux
RUN go get github.com/docker/docker/client
RUN rm -rf ../../docker/docker/vendor/github.com/docker/go-connections
RUN ls -la
RUN go get github.com/docker/go-connections/nat
RUN go get github.com/pkg/errors
RUN go get golang.org/x/net/context
RUN go get gopkg.in/mgo.v2
RUN go get github.com/auth0/go-jwt-middleware
RUN go get github.com/codegangsta/negroni
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/satori/go.uuid
RUN go get github.com/Pallinder/sillyname-go
COPY . .
RUN go build main.go
ENV GOPATH=/go/src
CMD ["/go/src/github.com/hengel2810/api_docli/main"]