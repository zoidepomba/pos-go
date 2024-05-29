package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Registro() {
	fmt.Printf("entrou aqui\n")
	p.Nome = "Guilherme Felipe"
	p.Idade = 25
	fmt.Printf(p.Nome)
}

type Registrando interface {
	Registro()
}

func registrandoPessoa(name Registrando){
	name.Registro()
}

func main() {
	newPessoa := Pessoa{}
	fmt.Println(newPessoa.Nome)
	registrandoPessoa(newPessoa)

}
