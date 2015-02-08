package helloGoLangWeb

import (
	"net/http"
	"log"
	 "goatstone"
)
func init() {
	http.HandleFunc("/admin", goatstone.HandleTemplate)
	log.Print("inits :  " , goatstone.I)
}
