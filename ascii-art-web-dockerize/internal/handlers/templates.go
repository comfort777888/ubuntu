package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// templateExecution function execute applies a parsed template to the specified
// data object, writing the output to w.
func templateExecution(w http.ResponseWriter, parseFile string, data interface{}) {
	html, err := template.ParseFiles(parseFile)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500InternalServerError.html", http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "500InternalServerError.html", http.StatusInternalServerError)
		return
	}
}
