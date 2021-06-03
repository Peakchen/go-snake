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

func Subscribe4Encode(topic string, callback interface{}) bool {
	return myNatsObj.Subscribe4Encode(topic, callback)
}

func Request4Encode(topic string, req interface{}, ack interface{}) bool {
	return myNatsObj.Request4Encode(topic, req, ack)
}

func Publish4Encode(topic string, obj interface{}) bool {
	return myNatsObj.Publish4Encode(topic, obj)
}


func Close(){
	myNatsObj.Close()
}