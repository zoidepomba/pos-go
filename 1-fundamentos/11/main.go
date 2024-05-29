package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	guilherme := Cliente{
		Nome:  "gui",
		Idade: 25,
		Ativo: false,
	}

	guilherme.Ativo = true
	fmt.Printf("O nome e %s ele tem a idade de %d e ele esta %t", guilherme.Nome, guilherme.Idade, guilherme.Ativo)
}
