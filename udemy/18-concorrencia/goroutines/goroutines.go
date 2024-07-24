package main

import (
	"fmt"
	"time"
)

func main() {
	//concorrencia != paralelismo
	go escrever("ola mundo")
	escrever("mundo ola")
}


func escrever(texto string){
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}