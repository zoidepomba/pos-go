package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	/*mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello world!"))
	})*/ //posso passar igual a seguinte forma porem essa e uma função anonima

	mux.HandleFunc("/", HomeHandler) //mesma coisa que a de cima
	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello guilherme"))
	})
	http.ListenAndServe(":8081", mux2)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

type blog struct {
	title string
}
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
