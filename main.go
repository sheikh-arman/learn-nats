package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {

	wait := make(chan bool)
	//opt := server.Options{
	//	//Host: "nats://127.0.0.1:4222",
	//}
	//
	//s, err := server.NewServer(&opt)
	//
	//if err != nil {
	//	panic(err)
	//}
	//go s.Start()
	//fmt.Println("test")
	//if !s.ReadyForConnections(4 * time.Second) {
	//	panic("not ready for connection")
	//}

	//nc, err := nats.Connect(s.ClientURL())
	nc, err := nats.Connect("nats://0.0.0.0:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	//err = nc.Publish("foo", []byte("Hello World"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("published")
	fmt.Println(nats.DefaultURL)
	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("subscribed")
	<-wait
}
