package main

import "github.com/zoidepomba/pos-go/tree/main/9/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
