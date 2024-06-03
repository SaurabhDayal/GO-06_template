package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	Title   string
	Heading string
	Items   []string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the absolute path of the current directory
		cwd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Define the path to the template file relative to the project root
		tmplPath := filepath.Join(cwd, "templates", "template.html")

		// Parse the template
		tmpl := template.Must(template.ParseFiles(tmplPath))

		// Sample data to pass to the template
		data := PageData{
			Title:   "My Page Title",
			Heading: "Welcome to My Page",
			Items:   []string{"Item 1", "Item 2", "Item 3"},
		}

		// Execute the template with the data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Start the web server
	http.ListenAndServe(":8080", nil)
}
