package main

import "fmt"

func main(){
	canal := make(chan string) //vazio

	//thread 2
	go func(){
		canal <- "ola mundo"//esta cheio
	}()
	//thread 3
	msg := <-canal //canal esvazia
	fmt.Println(msg)
}