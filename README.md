# example-nats-go-docker
Example docker setup to run a nats server, a nats-go listener, and a nats-go requester

# Usage
1. run nats server and listener: `docker-compose up`
2. Verify subscriptions=1 at: http://localhost:8222/subsz
3. test with sample publisher
```
apt-get install golang # https://wiki.ubuntu.com/Go
export GOPATH=`pwd`
go get github.com/nats-io/go-nats
go run publish.go
```

# Related
[example-nats-cli-docker-compose](https://github.com/shadiakiki1986/example-nats-cli-docker-compose)

# References
Here are links that helped me
* https://jacobmartins.com/2016/06/06/practical-golang-getting-started-with-nats-and-related-patterns/
* https://github.com/Logimethods/docker-nats-connector-spark/blob/d1d83ebc4fc281250487ccb48751a536464021e8/compose/docker-compose.no-cluster.yml
* https://github.com/nats-io/go-nats
* http://nats.io/documentation/clients/nats-clients/
* https://hub.docker.com/_/nats/

