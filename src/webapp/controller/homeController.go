package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"webapp/model"
)

type home struct{}

func (h *home) registerRoutes() {
	// Handle home route
	http.HandleFunc("/home/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tasks := model.GetTasks()
	reqPath := r.URL.Path[len("/home/"):]
	if strings.EqualFold(reqPath, "") {
		reqPath = "index.html"
	}
	templatePath := strings.Join([]string{"src", "webapp", "templates", reqPath}, string(os.PathSeparator))
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}

	err = t.Execute(w, tasks)
	if err != nil {
		fmt.Println(err)
	}
}
