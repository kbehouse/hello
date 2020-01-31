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

	type Command struct {
		Cmd     string `json:"cmd,omitempty"`
		ImgPath string `json:"imgPath,omitempty"`
	}

	response := ""
	me := &Command{Cmd: "readImage", ImgPath: "./example.jpg"}
	c.Request(natsPubTopic, me, &response, time.Second)

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%v', response: '%s'\n", natsPubTopic, me, response)
		decode2File(response, "example-response.jpg")
	}
}
