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
	type Command struct {
		Cmd     string `json:"cmd,omitempty"`
		ImgPath string `json:"imgPath,omitempty"`
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	// c.Subscribe(natsSubTopic, func(p *person) {
	// 	log.Printf("Received a person! %+v\n", p)
	// 	nc.Publish(s.Reply, []byte("there is no s.Reply field available"))
	// })

	c.Subscribe(natsSubTopic, func(subj, reply string, cmd *Command) {
		// log.Printf("Reply-Topic: %v", reply)
		if cmd.Cmd == "readImage" && reply != "" {
			nc.Publish(reply, []byte(encodeFromFile(cmd.ImgPath)))
		}
		log.Printf("Received a person on subject %s! %+v\n", subj, cmd)
	})

	// Wait for a message to come in
	wg.Wait()
}
