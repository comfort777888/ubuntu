package handlers

import (
	"fmt"
	"net/http"

	ascii "ascii-art-web/internal/ascii"
)

type Ascii struct {
	AsciiFont string
}

// GETHandler func receives only GET request and displays main page.
func GETHandler(w http.ResponseWriter, r *http.Request) {
	status := checkMethodAndPath(r, "/", http.MethodGet)
	if status != 200 {
		w.WriteHeader(status)
		templateExecution(w, "./ui/html/"+fmt.Sprint(status)+".html", status)
		return
	}
	templateExecution(w, "./ui/html/home.html", nil)
}

// Post handler responces only post request and processes data we receive through FromValue
// checks if text is correct and do not contain cyrilic alphabet, correct new lines,checks if font name is correct
// and if file that contains format font haven't been modified through HashSum & ConverFont func.
func POSTHandler(w http.ResponseWriter, r *http.Request) {
	method := checkMethodAndPath(r, "/ascii-art", http.MethodPost)
	if method != 200 {
		w.WriteHeader(method)
		templateExecution(w, "./ui/html/"+fmt.Sprint(method)+".html", method)
		return
	}
	text := r.FormValue("formtext") // receives text to format
	font := r.FormValue("fonts")    // receives font's name
	status := CheckUserInput(text, font)
	if status != 200 {
		w.WriteHeader(status)
		templateExecution(w, "./ui/html/"+fmt.Sprint(status)+".html", status) //+http.StatusText(status)
		return
	}
	// Converting ascii output result  and saves it in string
	result := ascii.OutputAscii(text, font)
	ascii := Ascii{
		AsciiFont: result,
	}
	templateExecution(w, "./ui/html/home.html", ascii)
}
