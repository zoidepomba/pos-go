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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
	fmt.Println(c.Ativo)
	fmt.Println("teste")
}

func main() {
	guilherme := Cliente{
		Nome:  "gui",
		Idade: 25,
		Ativo: false,
	}
	fmt.Println(guilherme.Ativo)
	guilherme.Logradouro = "Rua jerusalém"
	guilherme.Ativo = true
	fmt.Println(guilherme.Ativo)
	guilherme.Desativar()
	fmt.Println(guilherme.Ativo)
	fmt.Printf("O nome e %s ele tem a idade de %d e ele esta %t e o endereco e %s", guilherme.Nome, guilherme.Idade, guilherme.Ativo, guilherme.Logradouro)
}
