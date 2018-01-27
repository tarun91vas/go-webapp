package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"webapp/model"
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
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		email := r.Form.Get("username")
		password := r.Form.Get("password")
		if model.IsValidUser(email, password) {
			http.Redirect(w, r, "/static", http.StatusTemporaryRedirect)
			return
		}
	}

	//on GET serve login page
	t := l.templates.Lookup("login.html")

	err := t.Execute(w, nil)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
