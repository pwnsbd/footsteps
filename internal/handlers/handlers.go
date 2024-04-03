package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my page")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the contact page")
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This are the projects i have done and working on!")
}
