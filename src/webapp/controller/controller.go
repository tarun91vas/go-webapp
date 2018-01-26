package controller

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

//Setup register all the routes of the app
func Setup(templates *template.Template) {
	// Register controllers
	homeController := &home{templates: templates}
	homeController.registerRoutes()

	taskController := &task{templates: templates}
	taskController.registerRoutes()

	//Serve static folder
	staticDir := strings.Join([]string{"src", "webapp", "static"}, string(os.PathSeparator))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	//Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The server is up and running!!"))
	})

}
