//refer: https://docs.nats.io/developing-with-nats/receiving/structure
package main

import (
	"log"
	"sync"

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
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("NATS Subscribe Topic: %s", natsSubTopic)
	type person struct {
		Name    string
		Address string
		Age     int
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	// c.Subscribe(natsSubTopic, func(p *person) {
	// 	log.Printf("Received a person! %+v\n", p)
	// 	nc.Publish(s.Reply, []byte("there is no s.Reply field available"))
	// })

	c.Subscribe(natsSubTopic, func(subj, reply string, p *person) {
		// log.Printf("Reply-Topic: %v", reply)
		if reply != "" {
			nc.Publish(reply, []byte(encodeFromFile("./example.jpg")))
		}
		log.Printf("Received a person on subject %s! %+v\n", subj, p)
	})

	// Wait for a message to come in
	wg.Wait()
}
