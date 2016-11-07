package main

import (
	"fmt"
	"os"
	"time"

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

	id := 0
	for _ = range time.Tick(500 * time.Millisecond) {
		id++
		client := ClientId{id}
		var balance Balance
		if err = ec.Request("balance.curr", client, &balance, 100*time.Millisecond); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Got balance %+v\n", balance)
		}
	}
}

type ClientId struct {
	Id int `json:"id"`
}

type Balance struct {
	ClientId int `json:"clientId"`
	Balance  int `json:"balance"`
}
