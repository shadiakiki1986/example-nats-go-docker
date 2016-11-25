package main

import (
    "github.com/nats-io/go-nats"
"time"
)

func main() {
  nc, _ := nats.Connect(nats.DefaultURL)

  // Make a request
  nc.Request("foo", []byte("help me"), 10*time.Millisecond)
}
