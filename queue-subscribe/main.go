package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"sync"
)

func main(){
	service := flag.String("service", "A","ok")
	flag.Parse()
	var (
		host = "localhost"
		port = "4222"
		wg = &sync.WaitGroup{}
	)
	// Connect to NATS
	natsConn, errConNats := nats.Connect("nats://" + host + ":" + port)
	if errConNats != nil {
		log.Fatal(errConNats)
	}
	natsJSON, errEncodeConnection := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	if errEncodeConnection != nil {
		log.Fatal("error encode connection JSON nats server : "+errEncodeConnection.Error())

	}
	fmt.Println("listening")
	wg.Add(1)
	go queGroup(wg, natsJSON, *service)
	wg.Wait()
}

func queGroup (wg *sync.WaitGroup, natsJSON *nats.EncodedConn, service string){
	i := 0
	sub, err := natsJSON.QueueSubscribe("lets_change", "queueGroup" ,func(msg *nats.Msg) {
		i++
		fmt.Println(i)
		fmt.Println(msg.Reply)
		fmt.Println(string(msg.Data))
		fmt.Println(msg.Subject)
		iString := fmt.Sprintf("%v", i)
		msg.Respond([]byte(service+" aku bales ya "+iString))
		fmt.Println(msg)
		printMsg(msg, i)
	})
	fmt.Println(sub, err)
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s] Queue[%s] Pid[%d]: '%s'", i, m.Subject, m.Sub.Queue, os.Getpid(), string(m.Data))
}