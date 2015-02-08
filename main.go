package helloGoLangWeb

import (
	"net/http"
	"goatstone/form/admin"
)

func init() {
	http.HandleFunc("/admin", admin.HandleTemplate)
}


