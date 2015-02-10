package admin

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"goatstone/data"
	"appengine"
	"log"
)

var (
	title               = "goatstone : go : admin"
	legend              = "Admin Values!"
	templatePath string = "./template/admin.html"
	templateName string = "admin.html"
)

func populateData(ctx appengine.Context) {

	prop := map[string]string{"Name":"title", "Value":"Goatstone : Go", }
	data.AddSiteProp(ctx, prop)
	prop = map[string]string{"Name":"heading", "Value":"Welcome!", }
	data.AddSiteProp(ctx, prop)

}
func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	data.StoreLog(ctx, "HandleTemplate")
	method := "get"
	if r.Method == "POST" {
		method = "update"
	}
	// RUN ON INITIALIZATION
	//populateData(ctx)
	cwd, _ := os.Getwd()
	var (
		templates = template.Must(template.ParseFiles(filepath.Join(cwd, templatePath)))
	)
	templatedata := data.TemplateData{}
	var siteProps  []data.SiteProp
	siteProps, err := data.GetSiteProps(ctx)
	if err != nil {
		log.Print("ERROR : GetSiteProps :  ", err)
		http.Error(w, "Problem Getting Site Properties.", 500)
		return
	}
	if method == "update" {
		log.Print("update  :  ")
		args := []string{
			r.FormValue("0"), r.FormValue("1"),
			r.FormValue("2"), r.FormValue("3"),
			r.FormValue("4"), }
		_ = args
		templatedata.Legend = "Posted Values"
		// set inputs to disabled
		for k, v := range siteProps {
			siteProps[k].Disabled = "disabled"
			log.Print("k: ", k, " v: ", v.Disabled, v.Value)
		}
		templatedata.Inputs = siteProps
		templatedata.Message = " Return to edit form"
		templatedata.AHref = "/admin"
	} else if method == "get" {
		templatedata.Inputs = siteProps
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
