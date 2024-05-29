package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...)) //os tres pontos significa que e uma função variatica, ou seja ela pode ter varias variaveis
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 30},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}

}
