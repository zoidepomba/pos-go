package main

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
	println(a, b, c, d, e, f)
}
