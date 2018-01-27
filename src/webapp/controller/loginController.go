package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type login struct {
	templates *template.Template
}

func (l *login) registerRoutes() {
	http.HandleFunc("/login", l.loginHandler)
}

func (l *login) loginHandler(w http.ResponseWriter, r *http.Request) {
	//on POST handle login
	if r.Method == http.MethodPost {
		fmt.Println(r)
		return
	}

	//on GET serve login page
	t := l.templates.Lookup("login.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
