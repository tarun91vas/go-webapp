package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"webapp/model"
)

func main() {

	// Handle view task
	http.HandleFunc("/view/", viewHandler)

	//Serve static folder
	staticDir := strings.Join([]string{"src", "webapp", "static"}, string(os.PathSeparator))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	//Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The server is up and running!!"))
	})

	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	tasks := model.GetTasks()
	reqPath := r.URL.Path[len("/view/"):]
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
