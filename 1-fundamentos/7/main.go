package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{"Guilherme": 1000, "Maria": 2000, "Jose": 3000}

	fmt.Println(salarios["Guilherme"])
	delete(salarios, "Guilherme")
	salarios["Yes"] = 1500
	fmt.Println(salarios["Yes"])

	//	sal := make(map[string]int)

	for nome, salario := range salarios {
		fmt.Printf("o salario de %s e %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("o salario e %d\n", salario)
	}
}
