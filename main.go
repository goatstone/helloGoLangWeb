package helloweb

import (
	"bytes"
	"html/template"
	"net/http"
)

var yourId int  = 0

var (
	templates = template.Must(template.ParseFiles(
		"gotemplate.html",
	))
)

func init() {
	http.HandleFunc("/templates", handleTemplates)
}

 func handleTemplates(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/templates" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "POST" {
		out := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(out, "gotemplate.html", yourId); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		out.WriteTo(w)
		yourId++
		return
	}
}
