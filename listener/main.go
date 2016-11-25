package main

import (
    "github.com/nats-io/go-nats"
    "fmt"
    "os"
//    "sync"
)

func main() {
  // url := nats.DefaultURL
  // url := "nats://nats-main:4222"
  url := os.Getenv("NATS_URI")
  nc, _ := nats.Connect(url)

  fmt.Printf("Starting listener on: %s\n", url)

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

  fmt.Printf("Listening for messages\n")
  select {} // Block forever
}
