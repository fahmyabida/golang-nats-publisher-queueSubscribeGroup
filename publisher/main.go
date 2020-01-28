package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

type Message struct {
	Id		string
	Content string
}

func main() {
	var (
		host = "localhost"
		port = "4222"
	)
	// Connect to NATS
	natsConn, errConNats := nats.Connect("nats://" + host + ":" + port)
	if errConNats != nil {
		log.Fatal(errConNats)
	}
	natsJSON, errEncodeConnection := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	if errEncodeConnection != nil {
		log.Fatal("error encode connection JSON nats server : " + errEncodeConnection.Error())
	}

	message := Message{
		Id:      "99",
		Content: "Hello world!",
	}
	var reply []byte
	for z:=0;z<9 ;z++  {
		err := natsJSON.Request("lets_change", message, &reply, 5*time.Second)
		if err !=nil {
			fmt.Println("waah error publish :"+err.Error())
		}
		fmt.Println(string(reply))
	}
	runtime.Goexit()
}
