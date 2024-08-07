package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data{
		fmt.Printf("worker %d received %d\n",x, data)
		time.Sleep(time.Second)
	}
}

func main() {
	data:= make(chan int)
	QtdWorkers := 10000

	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++{
		data <- i
	}
}