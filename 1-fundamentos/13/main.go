package main

import "fmt"

// struct
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// metodo
func (e Endereco) Mudanca() {
	fmt.Println(e.Logradouro)
	e.Logradouro = "avenida jose fonseca e silva"
	fmt.Println(e.Logradouro)
}

// interface que chama meu metodo mudança
type Mudando interface {
	Mudanca()
}

// funçao que chama minha interface para alteração
func novoEndereco(end Mudando) {
	end.Mudanca()
}


// função main
func main() {
	neEndereco := Endereco{}

	fmt.Printf(neEndereco.Logradouro)

	novoEndereco(neEndereco)
	
}
