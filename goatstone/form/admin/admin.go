package admin

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"goatstone/data"
	"appengine"
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

	ctx := appengine.NewContext(r)
	data.StoreLog(ctx, "HandleTemplate")

	cwd, _ := os.Getwd()
	var (
		templates = template.Must(template.ParseFiles(
	filepath.Join(cwd, templatePath)))
	)
	td := templateData{Title:title }
	if r.Method == "POST" {
		td.Legend = legend
		// set inputs to disabled
		for ip := range inputs {
			inputs[ip].Disabled = "disabled"
			inputs[ip].Value = r.FormValue(inputs[ip].Name)
		}
		td.Inputs = inputs
		td.Message = " Return to edit form"
		td.AHref = "/admin"
		args := []string{
			r.FormValue("name0"), r.FormValue("name1"),
			r.FormValue("name2"), r.FormValue("name3"),
			r.FormValue("name4"), }
		data.StoreSiteInfo(ctx, args)
	}
	if r.Method != "POST" {
		// set inputs to active
		for ip := range inputs {
			inputs[ip].Disabled = ""
		}
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
