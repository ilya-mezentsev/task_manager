FROM golang:1.13.5-alpine3.10 AS build
# Support CGO and SSL
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR ${GOPATH}
RUN go get github.com/gorilla/mux
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/dgrijalva/jwt-go

COPY . .
EXPOSE 8080
