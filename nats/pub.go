package main

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

const (
	natsServer   = "localhost:4222"
	natsPubTopic = "test"
)

func main() {
	// NATS Init
	log.Printf("Connect to NATS Server: %s", natsServer)
	nc, err := nats.Connect(natsServer)
	if err != nil {
		log.Fatal(err)
	}

	msg := "hello, abc"
	nc.Publish(natsPubTopic, []byte(msg))

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", natsPubTopic, msg)
	}
}
