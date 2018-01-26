package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"webapp/model"
)

type home struct {
	templates *template.Template
}

func (h *home) registerRoutes() {
	// Handle home route
	http.HandleFunc("/home/", h.homeHandler)
}

func (h *home) homeHandler(w http.ResponseWriter, r *http.Request) {
	tasks := model.GetTasks()
	reqPath := r.URL.Path[len("/home/"):]
	if strings.EqualFold(reqPath, "") {
		reqPath = "index.html"
	}

	t := h.templates.Lookup(reqPath)
	if t == nil {
		http.NotFound(w, r)
		return
	}

	err := t.Execute(w, tasks)
	if err != nil {
		fmt.Println(err)
	}
}
