package main

func main() {
	for i := 0; i < 10; i++ {
		print(i)
	}

	numeros := []string{"um", "dois", "Tres"}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 10 {
		print(i)
		i++
	}

	for {
		print("Hello, world")
	}
}
