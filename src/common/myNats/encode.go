package myNats

import (
	"github.com/davyxu/ulog"
	
	"time"
)


func (self *MyNats) Subscribe4Encode(topic string, callback interface{}) bool {

	_, err := self.encodeConn.Subscribe(topic, callback)
	if err != nil {
		ulog.Errorf("nats subscribe err: %v.", err)	
	}

	return err == nil

}


func (self *MyNats) Request4Encode(topic string, req interface{}, ack interface{}) bool {

	err := self.encodeConn.Request(topic, req, ack, 3*time.Second)
	if err != nil {
		ulog.Errorf("nats encode request err: %v.", err)
	}

	return err == nil

}

func (self *MyNats) Publish4Encode(topic string, obj interface{})bool{

	err := self.encodeConn.Publish(topic, obj)
	if err != nil {
		ulog.Errorf("nats encode publish err: %v.", err)
	}

	return err == nil 
}