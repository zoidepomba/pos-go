package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHorario int
}

func main() {

	curso := Curso{"Go", 40} //inicio com os valores
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - CargaHorario: {{.CargaHorario}}")) //fazendo a alteração com o must
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
