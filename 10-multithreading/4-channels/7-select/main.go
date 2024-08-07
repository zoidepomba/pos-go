package main

import (
	"fmt"
	"time"
)

type Message struct{
	id int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	//rabbit mq processa
	go func() {
		time.Sleep(time.Second)
		msg := Message{1, "Hello from RabbitMQ"}
		c1 <- msg
	}()

	//kafka
	go func() {
		time.Sleep(time.Second * 4)
		msg := Message{2, "Hello from kafka"}
		c2 <- msg
	}()
	for{
		select {
		case msg := <-c1: //rabitmq
			fmt.Printf("received frrom Rabbit: %s\n", msg.Msg)
		
		case msg := <-c2: //kafka
			fmt.Printf("received from kafka: %s\n", msg.Msg)

		case <-time.After(time.Second *3):
			println("timeout")
		}
	}
}