package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"webapp/model"
)

type task struct {
	templates *template.Template
}

func (t *task) registerRoutes() {
	http.HandleFunc("/task/", t.taskHandler)
}

func (t *task) taskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/task/"):]
	viewTask, err := model.GetTaskByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tpl := t.templates.Lookup("task.html")
	if tpl == nil {
		http.NotFound(w, r)
		return
	}

	err = tpl.Execute(w, viewTask)
	if err != nil {
		fmt.Println(err)
	}

}
