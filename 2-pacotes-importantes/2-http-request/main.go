package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://api-prd.clarobrasil.mobi/address/viability/86200981/260")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	print(string(res))
}
