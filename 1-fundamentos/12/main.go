package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	guilherme := Cliente{
		Nome:  "gui",
		Idade: 25,
		Ativo: false,
	}
	guilherme.Logradouro = "Rua jerusalém"
	guilherme.Ativo = true
	fmt.Printf("O nome e %s ele tem a idade de %d e ele esta %t e o endereco e %s", guilherme.Nome, guilherme.Idade, guilherme.Ativo, guilherme.Logradouro)
}
