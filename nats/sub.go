package main

import (
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

const (
	natsServer   = "localhost:4222"
	natsSubTopic = "json-test"
)

func main() {
	// NATS Init
	log.Printf("Connect to NATS Server: %s", natsServer)
	nc, err := nats.Connect(natsServer)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("NATS Subscribe Topic: %s", natsSubTopic)
	nc.Subscribe(natsSubTopic, func(m *nats.Msg) {
		log.Printf("Received on [%s]: '%s'", m.Subject, string(m.Data))
	})
	nc.Flush()

	runtime.Goexit()
}
