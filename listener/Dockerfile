FROM golang:alpine
WORKDIR /code
COPY . .
RUN export GOPATH=`pwd` \
  && go get github.com/nats-io/go-nats \
  && go run main.go
