package main

type Conta struct {
	saldo int
}

func NewConta() *Conta{
	
	return &Conta{saldo: 0}
	
}

type criaConta interface{}

func (c *Conta) simular(valor int) int{
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {

	conta := Conta{saldo: 100}
	conta.simular(200)
	NewConta()
	println(conta.saldo)
	
}
