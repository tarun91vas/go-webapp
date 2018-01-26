package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"webapp/controller"
	"webapp/lib"
)

func main() {
	// Compile templates
	templateDir := strings.Join([]string{"src", "webapp", "templates"}, string(os.PathSeparator))
	templates := lib.PopulateTemplates(templateDir)

	// Setup controller
	controller.Setup(templates)

	fmt.Println("Server listening on 8080...")
	http.ListenAndServe(":8080", nil)
}
