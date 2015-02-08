package admin

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"log"
	"path/filepath"
)

var (
	title               = "goatstone : go : admin"
	legend              = "Admin Values!"
	templatePath string = "./template/admin.html"
	templateName string = "admin.html"
	count int           = 0
)

type templateData struct {
	Title      string
	Legend     string
	Inputs     []input
}
type input struct {
	Label     string
	name      string
	value     string
	inputType string
}

var inputs = []input{
	{"Label 0", "name0", "val0", "text"},
	{"Label 1", "name1", "val1", "text"},
	{"Label 2", "name2", "val2", "text"},
	{"Label 3", "name3", "val3", "text"},
	{"Label 4", "name4", "val4", "text"},
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	log.Print("admin::: ", count)
	td := templateData{Title:title }
	td.Inputs = inputs
	td.Legend = legend
	cwd, _ := os.Getwd()
	var (
		templates = template.Must(
	template.ParseFiles(
	filepath.Join(cwd, templatePath)))
	)
	if r.Method != "POST" {
		out := &bytes.Buffer{}
		if err := templates.ExecuteTemplate(out, templateName, td); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		out.WriteTo(w)
		count++
		return
	}
}
