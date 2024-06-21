package main

import (
	"fmt"

	"github.com/zoidepomba/pos-go/tree/main/5-Packaging/1/math"
)

func main() {
	m := math.Match{A: 1, B: 2}
	fmt.Println(m.Add())
}
