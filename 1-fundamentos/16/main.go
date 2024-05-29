package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVariavel1 := 10
	minhaVariavel2 := 20

	soma(&minhaVariavel1, &minhaVariavel2)
	println(minhaVariavel1)
	println(minhaVariavel2)

}
