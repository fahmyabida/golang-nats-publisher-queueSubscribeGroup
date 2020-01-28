package mq

import (
	"errors"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

var natInstance INats

// INats interface is used to create object nats command
type INats interface {
	Publish(topic string, data interface{}) error
	Subscribe(topic string, field interface{}) error
	Request(topic string, field interface{}, reply interface{}, timeout time.Duration) error
	Queue(topic, wait string, reply interface{}) error
	New(host, port string) (INats, error)
}

// Nats is used to create object nats
type nat struct {
	natsEncode *nats.EncodedConn
}

// Publish is used to create command to publish
func (n nat) Publish(topic string, data interface{}) error {
	errorResult := n.natsEncode.Publish(topic, data)
	return errorResult

}

// Subscribe is used to create command to subcribe
func (n nat) Subscribe(topic string, field interface{}) error {
	_, errorResult := n.natsEncode.Subscribe(topic, field)
	return errorResult
}

// Request is used to create command to subcribe
func (n nat) Request(topic string, field interface{}, reply interface{}, timeout time.Duration) error {
	errorResult := n.natsEncode.Request(topic, field, reply, timeout)
	return errorResult
}

// Queue is used to create command to queuesubcribe
func (n nat) Queue(topic, wait string, reply interface{}) error {
	_, errorResult := n.natsEncode.QueueSubscribe(topic, wait, reply)
	return errorResult

}

// New is used to create new connectio Nats
func (n nat) New(host, port string) (INats, error) {
	connectionNATS, errorConnectionNATS := nats.Connect("nats://" + host + ":" + port)
	if errorConnectionNATS != nil {
		return nil, errors.New("error connect nats server")
	}
	natsJSON, errEncodeConnection := nats.NewEncodedConn(connectionNATS, nats.JSON_ENCODER)
	if errEncodeConnection != nil {
		return nil, errors.New("error encode connection JSON nats server")

	}
	iNat := nat{
		natsEncode: natsJSON,
	}
	return iNat, errEncodeConnection
}

// MultiJob is used to excute multi job
func MultiJob(worker ...func(waitGroup *sync.WaitGroup)) {

	var waitGroup sync.WaitGroup
	for _, work := range worker {
		waitGroup.Add(1)
		go work(&waitGroup)
	}
	waitGroup.Wait()
}

// SingleJob is used to excute single job
func SingleJob(worker func()) {

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go worker()
	waitGroup.Wait()
}

//Connect is used to create object nats engine
func Connect(host, port string) (INats, error) {
	tempNats := nat{}
	natInstance = tempNats
	natOnject, errorIntialize := natInstance.New(host, port)
	return natOnject, errorIntialize
}
