package data

import (
	"log"
	"appengine"
	"appengine/datastore"
	"math/rand"
	"time"
	"appengine/user"
)

type Entity struct {
	Value string
	Count int
}
type appLog struct {
	Name      string
	Id        int
	TimeStamp time.Time
	Account   string
}
type siteInfo struct {
	Title           string
	Header          string
	Message         string
	Color           string
	BackgroundColor string
}

func StoreLog(ctx appengine.Context, name string) {
	var entityKind = "appLog"
	dataKey := datastore.NewIncompleteKey(ctx, entityKind, nil)

	alog := &appLog{}
	alog.Name = name
	alog.Id = rand.Intn(1000)
	alog.TimeStamp = time.Now()
	alog.Account = user.Current(ctx).String()

	if _, err := datastore.Put(ctx, dataKey, alog); err != nil {
		log.Print("err:  ", err)
		return
	}

}
func StoreSiteInfo(ctx appengine.Context, args []string) {
	dataKey := datastore.NewKey(ctx, "SiteInfo", "siteinfo", 0, nil)
	si := &siteInfo{args[0], args[1], args[2], args[3], args[4]}
	if _, err := datastore.Put(ctx, dataKey, si); err != nil {
		log.Print("err:  ", err)
		return
	}
	//log.Print("store", si.Title)
}

//func StoreCredentials(cred *Credentials, ctx appengine.Context) (err error) {
//	key := datastore.NewKey(ctx, "Credentials", "main", 0, nil)
//	_, err = datastore.Put(ctx, key, cred)
//	return
//}
func Load() {
	log.Print("Load")
}
