package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nats"
)

func main() {
	c, err := nats.Connect(os.Args[1])
	if err != nil {
		panic(err)
	}
	ec, err := nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	ec.Subscribe("balance.curr", func(sub, reply string, cl *ClientId) {
		fmt.Printf("Req for client %+v\n", cl)
		ec.Publish(reply, &Balance{
			ClientId: cl.Id,
			Balance:  cl.Id * 1e3,
		})
	})
	<-make(chan struct{})
}

type ClientId struct {
	Id int `json:"id"`
}

type Balance struct {
	ClientId int `json:"clientId"`
	Balance  int `json:"balance"`
}
