package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate := template.Must(template.ParseFiles("src/templates/home.html"))
	homeTemplate.Execute(w, nil)

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is contact page")
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This are the projects i have done and working on!")
}
