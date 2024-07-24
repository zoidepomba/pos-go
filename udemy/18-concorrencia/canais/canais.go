package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	/*for { /// para funcionamento em função que tem for
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/

	mensagem := <- canal //quando eu uso assim e porque eu estou recebendo um valor no canal
	


	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // com a seta apontando para o canal ele esta recebendo
		time.Sleep(time.Second)
	}

	close(canal)
}
