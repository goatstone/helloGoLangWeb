package helloGoLangWeb

import (
	"net/http"
	"log"
	 "goatstone/form/admin"
)
func init() {
	http.HandleFunc("/admin", admin.HandleTemplate)
	log.Print("inits :  " , 12)
}
