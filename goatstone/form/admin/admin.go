package admin

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"goatstone/data"
	"appengine"
	//		"log"
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
	{"Title", "0", "val0", "text", ""},
	{"Heading", "1", "val1", "text", ""},
	{"Message", "2", "val2", "text", ""},
	{"Color", "3", "val3", "text", ""},
	{"Background Color 4", "4", "val4", "text", ""},
}
func populateData(ctx appengine.Context){

	prop := map[string]string{"Name":"title", "Value":"Goatstone : Go", }
	data.AddSiteProp(ctx, prop)
	prop = map[string]string{"Name":"heading", "Value":"Welcome!", }
	data.AddSiteProp(ctx, prop)

}
func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	data.StoreLog(ctx, "HandleTemplate")

	// RUN ON INITIALIZATION
	//populateData(ctx)

	cwd, _ := os.Getwd()
	var (
		templates = template.Must(template.ParseFiles(
	filepath.Join(cwd, templatePath)))
	)
	templatedata := templateData{Title:title }
	if r.Method == "POST" {

		args := []string{
			r.FormValue("0"), r.FormValue("1"),
			r.FormValue("2"), r.FormValue("3"),
			r.FormValue("4"), }
		if err := data.StoreSiteInfo(ctx, args); err != nil {
			http.Error(w, "Problem Storing Site Infromation.", 500)
		}

		templatedata.Legend = "Posted Values"
		// set inputs to disabled
		for ip := range inputs {
			inputs[ip].Disabled = "disabled"
			inputs[ip].Value = r.FormValue(inputs[ip].Name) // TODO Get data from DB
		}
		templatedata.Inputs = inputs
		templatedata.Message = " Return to edit form"
		templatedata.AHref = "/admin"

	}
	if r.Method != "POST" {
		si, err := data.GetSiteInfo(ctx);
		if err != nil {
			http.Error(w, "Problem Getting Site Infromation.", 500)
		}
		_ = si

		//log.Print("hello ", si.Title)
		//		si = data.GetSiteInfo(ctx)
		// set inputs to active
		for ip := range inputs {
			//log.Print("inputs: ", ip, inputs[ip])
			inputs[ip].Disabled = ""
		}
		//inputs[0][1] = si.Title
		templatedata.Inputs = inputs
		templatedata.Legend = "GET!"
	}
	out := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(out, templateName, templatedata); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out.WriteTo(w)
	return
}
