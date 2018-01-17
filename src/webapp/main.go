package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	// Serve public folder
	publicDir := strings.Join([]string{"src", "webapp", "public"}, string(os.PathSeparator))
	http.Handle("/", http.FileServer(http.Dir(publicDir)))

	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
