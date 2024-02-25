package main

import (
	"html/template"
	"net/http"

	"github.com/leoisrael/GoLangPokerMonteCarlo/server"
)

type PageVariables struct {
	Title       string
	Description string
}

func main() {
	server.Api("8080")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	pageTitle := "Exemplo de Template com Go"
	pageDescription := "Uma página simples usando html/template em Go"

	p := PageVariables{Title: pageTitle, Description: pageDescription}

	t, _ := template.ParseFiles("template.html")

	t.Execute(w, p)
}
