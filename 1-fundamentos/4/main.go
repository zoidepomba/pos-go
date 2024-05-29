package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 1
	d string  = "ola"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	b = true
	fmt.Printf("o tipo de E e %T %T %T %T %T %T", a, b, c, d, e, f)
}
