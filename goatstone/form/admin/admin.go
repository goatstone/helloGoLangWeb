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
)

type templateData struct {
	Title      string
	Legend     string
	Inputs     []input
	Message    string
	AHref      string
}
type input struct {
	Label     string
	Name      string
	Value     string
	inputType string
	Disabled  string
}

var inputs = []input{
	{"Label 0", "name0", "val0", "text", ""},
	{"Label 1", "name1", "val1", "text", ""},
	{"Label 2", "name2", "val2", "text", ""},
	{"Label 3", "name3", "val3", "text", ""},
	{"Label 4", "name4", "val4", "text", ""},
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {

	cwd, _ := os.Getwd()
	var (
		templates = template.Must(
	template.ParseFiles(
	filepath.Join(cwd, templatePath)))
	)
	td := templateData{Title:title }
	if r.Method == "POST" {
		log.Print("admin::: ", r.FormValue("name0"))
		td.Legend = legend
		// set inputs to disabled
		inputs[0].Disabled = "disabled"
		inputs[0].Value = r.FormValue("name0")
		td.Inputs = inputs
		td.Message = " Return to edit form"
		td.AHref = "/admin"
	}
	if r.Method != "POST" {
		inputs[0].Disabled = ""
		td.Inputs = inputs
		td.Legend = legend
	}
	out := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(out, templateName, td); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out.WriteTo(w)
	return
}
