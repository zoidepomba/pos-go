package main // Define o pacote principal do programa

import "fmt" // Importa o pacote fmt, que implementa funções de formatação I/O

// Definição de uma estrutura chamada 'Estudante' com quatro campos:
// Nome, Matricula, Curso e Periodo, todos do tipo string ou int.
type Estudante struct {
	Nome      string
	Matricula int
	Curso     string
	Periodo   string
}

// Método 'alteraCurso' associado à estrutura 'Estudante'.
// Este método imprime o curso atual, altera o curso para "Sistemas De Informação",
// e depois imprime a alteração realizada.
func (e Estudante) alteraCurso() {
	fmt.Printf("O curso atual é %s\n", e.Curso)
	e.Curso = "Sistemas De Informação"
	fmt.Printf("Realizada a alteração do curso para %s\n", e.Curso)
}

// Método 'alteraPeriodo' associado à estrutura 'Estudante'.
// Este método imprime o período atual, altera o período para "Noturno",
// e depois imprime a alteração realizada.
func (e Estudante) alteraPeriodo() {
	fmt.Printf("O período atual é %s\n", e.Periodo)
	e.Periodo = "Noturno"
	fmt.Printf("O período atual é %s\n", e.Periodo)
}

// Definição de uma interface chamada 'chamaAlteraCursos'.
// Esta interface exige que qualquer tipo que a implemente possua
// os métodos 'alteraCurso' e 'alteraPeriodo'.
type chamaAlteraCursos interface {
	alteraCurso()
	alteraPeriodo()
}

// Função 'novoCurso' que aceita um argumento do tipo 'chamaAlteraCursos'.
// Esta função chama os métodos 'alteraCurso' e 'alteraPeriodo' do argumento passado.
func novoCurso(curso chamaAlteraCursos) {
	curso.alteraCurso()
	curso.alteraPeriodo()
}

// Função 'novoPeriodo' que aceita um argumento do tipo 'chamaAlteraCursos'.
// Esta função chama o método 'alteraPeriodo' do argumento passado.
func novoPeriodo(curso chamaAlteraCursos) {
	curso.alteraPeriodo()
}

// Função principal onde a execução do programa começa.
func main() {
	fmt.Println("Iniciando o programa") // Imprime uma mensagem inicial
	
	// Criação de uma nova instância de 'Estudante' com valores padrão (vazios/zero).
	estudante := Estudante{}
	
	// Impressão do curso atual do estudante (valor padrão, que é uma string vazia).
	fmt.Printf("O curso atual é %s\n", estudante.Curso)
	
	// Chamada da função 'novoPeriodo' passando a instância 'estudante' como argumento.
	// Isso chama o método 'alteraPeriodo' na instância 'estudante'.
	novoPeriodo(estudante)
}