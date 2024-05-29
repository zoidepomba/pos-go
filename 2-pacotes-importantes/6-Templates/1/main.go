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
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} - CargaHorario: {{.CargaHorario}}")
	err:= tmp.Execute(os.Stdout, curso)
	if err != nil{
		panic(err)
	}

}
