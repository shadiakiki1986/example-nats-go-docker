FROM golang:alpine
RUN apk update \
  && apk add git
WORKDIR /code
COPY . .
ENV GOPATH=/code
RUN go get github.com/nats-io/go-nats
ENTRYPOINT sleep 5 && go run main.go
