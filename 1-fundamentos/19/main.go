package main

import "fmt"

func main() {
	var minhaVar interface{} = "Guilherme Felipe"

	print(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res e %v e o resultado de ok e %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("o valor de res2 e %v", res2)
}
