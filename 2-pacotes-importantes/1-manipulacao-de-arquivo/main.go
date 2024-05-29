package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo no arquivo "))
	if err != nil {
		panic(err)
	}
	fmt.Printf("arquivo criado com sucesso e o seu tamnho e: %d bytes\n", tamanho)
	f.Close()

	///leitura de arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	//leitura de arquivo pouco a pouco

	arquivo2, err := os.Open("arquivo.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 20)

	for {
		n, err :=reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	
	os.Remove("arquivo.txt")
	print("Removendo o arquivo\n")
	if err != nil{
		panic(err)
	}
}
