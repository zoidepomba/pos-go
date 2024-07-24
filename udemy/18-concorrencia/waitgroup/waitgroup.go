package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//concorrencia != paralelismo
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("ola mundo")
		waitGroup.Done()
	}()
	go func() {
		escrever("mundo ola")
		waitGroup.Done()

	}()
	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 5; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
