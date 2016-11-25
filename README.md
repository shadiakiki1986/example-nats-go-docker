# example-nats-go-docker
Example docker setup to run a nats server, a nats-go listener, and a nats-go requester

# Usage
1. run nats server and listener: `docker-composer up`
2. test with sample publisher
```
apt-get install golang # https://wiki.ubuntu.com/Go
export GOPATH=`pwd`
go get github.com/nats-io/go-nats
go run publish.go
```
