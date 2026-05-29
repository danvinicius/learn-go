package main

import (
	"html/template"
	"log"
	"net/http"
)

type Loja struct {
	Nome            string
	TituloHeader    string
	DescricaoHeader string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Olá mundo"))

		loja := Loja{"Mono", "Minimalismo que veste atitude.", "Peças premium com design limpo, caimento moderno e estética sofisticada para o cotidiano."}
		templates.ExecuteTemplate(w, "index.html", loja)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
