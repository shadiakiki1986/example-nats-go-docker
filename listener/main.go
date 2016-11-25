package main

import (
    "github.com/nats-io/go-nats"
    "fmt"
//    "sync"
)

func main() {
  nc, _ := nats.Connect(nats.DefaultURL)

  // Simple Publisher
  //nc.Publish("foo", []byte("Hello World"))

  // Simple Async Subscriber
  nc.Subscribe("foo", func(m *nats.Msg) {
      fmt.Printf("Received a message: %s\n", string(m.Data))
  })

  // Close connection
//  nc := nats.Connect("nats://localhost:4222")
//  nc.Close();

//  wg := sync.WaitGroup{}
//  wg.Wait()

  select {} // Block forever
}
