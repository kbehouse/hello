package main

import (
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
)

const (
	natsServer   = "localhost:4222"
	natsPubTopic = "json-test"
)

func main() {
	// NATS Init
	log.Printf("Connect to NATS Server: %s", natsServer)
	nc, err := nats.Connect(natsServer)
	if err != nil {
		log.Fatal(err)
	}

	// https://github.com/nats-io/nats.go/blob/6dfc35499d/example_test.go
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	type person struct {
		Name    string
		Address string
		Age     int
	}

	response := ""
	me := &person{Name: "derek", Age: 22, Address: "85 Second St"}
	c.Request(natsPubTopic, me, &response, time.Second)

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%v', response: '%s'\n", natsPubTopic, me, response)

	}
}
