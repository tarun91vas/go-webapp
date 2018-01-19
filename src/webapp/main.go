package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

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
