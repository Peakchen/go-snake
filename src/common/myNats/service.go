package myNats

import (
	"github.com/nats-io/nats.go"
	"github.com/davyxu/ulog"
	
	"time"
	"errors"
)

type MyNats struct {
	natsConn *nats.Conn
}

var (
	myNatsObj *MyNats
)

func CreateMyNats(addr string){
	
	conn, err := nats.Connect(addr) //nats.DefaultURL
	if err != nil {
	
		ulog.Errorf("nats connect fail, err: %v.", err)
		return
	}

	conn.Flush()

	myNatsObj = &MyNats{
		natsConn: conn,
	}

}


func (self *MyNats) GetConn()*nats.Conn{
	
	return self.natsConn
}

func (self *MyNats) Close(){

	self.natsConn.Close()
}

func (self *MyNats) Publish(topic string, data []byte) (err error) {
	
	if !self.natsConn.IsConnected(){
		
		err = errors.New("nats connect is disconnect")
		ulog.Errorf("nats connect is disconnect, topic: %v, data: %v.", topic, data)
		return
	}

	err = self.natsConn.Publish(topic, data)
	if err != nil {

		ulog.Errorf("nats publish data fail, topic: %v, data: %v, err: %v.", topic, data, err)
		return
	}

	return
}

func (self *MyNats) Subscribe(topic string, callback nats.MsgHandler){

	_, err := self.natsConn.Subscribe(topic, callback) // return ?
	if err != nil {

		ulog.Errorf("nats subscribe err: %v.", err)	
	}

}

func (self *MyNats) Request(topic string, interf interface{}, callback nats.MsgHandler){

	encryptData, err := Codec().Marshal(interf)
	if encryptData == nil || err != nil {

		ulog.Errorf("codec marshal fail, topic: %v, err: %v.", topic, err)
		return
	}

	msg, err := self.natsConn.Request(topic, encryptData, 2*time.Second)
	if err != nil {

		ulog.Errorf("nats conn request fail, topic: %v, err: %v.",topic,err)
		return
	}

	callback(msg)
	
}