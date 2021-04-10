package myNats

import (
	"github.com/nats-io/nats.go"
)

func Publish(topic string, data []byte){
	myNatsObj.Publish(topic, data)
}

func Subscribe(topic string, callback nats.MsgHandler){
	myNatsObj.Subscribe(topic, callback)
}

func Request(topic string, interf interface{}, callback nats.MsgHandler){
	myNatsObj.Request(topic, interf, callback)
}

func Close(){
	myNatsObj.Close()
}