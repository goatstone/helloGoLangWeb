package goatstone

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"log"
	"path/filepath"
)

var yourId int = 0
var I = 12;

func  retI()( i int){
	i = 9
	return
}
func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	log.Print("inits 1", yourId)
	if r.URL.Path != "/admin" {
		http.NotFound(w, r)
		return
	}
	type AStrct struct {
		a, b, c int
	}
	cwd, _ := os.Getwd()
	adminFilePath := filepath.Join(cwd, "./template/admin.html")
	var (
		templates = template.Must(template.ParseFiles(adminFilePath))
	)
	m := AStrct{a:1}
	m.b = yourId
	if r.Method != "POST" {
		out := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(out, "admin.html", m); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		out.WriteTo(w)
		yourId++
		return
	}
}
