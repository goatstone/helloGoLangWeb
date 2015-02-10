package data

import (
	"log"
	"appengine"
	"appengine/datastore"
	"math/rand"
	"time"
	"appengine/user"

)

type appLog struct {
	Name      string
	Id        int
	TimeStamp time.Time
	Account   string
}
type SiteInfo struct {
	Title           string
	Header          string
	Message         string
	Color           string
	BackgroundColor string
}
type TemplateData struct {
	Title      string
	Legend     string
	Inputs     []SiteProp
	Message    string
	AHref      string
}
type SiteProp struct {
	Name          string
	Value         string
	FormLabel     string
	InputType     string
	Disabled      string
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

func AddSiteProp(ctx appengine.Context, prop map[string]string) (err error) {

	keyName := prop["Name"]

	sp := &SiteProp{Disabled:""}
	sp.Name = prop[ "Name"]
	sp.Value = prop["Value"]
	sp.FormLabel = prop["Name"]

	dataKey := datastore.NewKey(ctx, "SiteProperties", keyName, 0, nil)
	if _, err := datastore.Put(ctx, dataKey, sp); err != nil {
		log.Print("err::  ", err)
	}
	return
}
func GetSiteProps(ctx appengine.Context) (siteProps []SiteProp , err error) {
	siteProps = []SiteProp{}
	q := datastore.NewQuery("SiteProperties")
	_, err = q.GetAll(ctx, &siteProps)
	if err != nil {
		log.Print("ERROR :  ", err)
		return
	}
	return
}
func StoreSiteInfo(ctx appengine.Context, args []string) (err error) {
	dataKey := datastore.NewKey(ctx, "SiteInfo", "siteinfo", 0, nil)
	si := &SiteInfo{args[0], args[1], args[2], args[3], args[4]}
	if _, err := datastore.Put(ctx, dataKey, si); err != nil {
		log.Print("err:  ", err)
	}
	return
}
func GetSiteInfo(ctx appengine.Context) (si *SiteInfo, err error) {
	key := datastore.NewKey(ctx, "SiteInfo", "siteinfo", 0, nil)
	si = &SiteInfo{}
	err = datastore.Get(ctx, key, si)
	return
}

//func LoadCredentials(ctx appengine.Context) (cred *Credentials, err error) {
//	key := datastore.NewKey(ctx, "Credentials", "main", 0, nil)
//	cred = &Credentials{}
//	err = datastore.Get(ctx, key, cred)
//	return
//}
//func StoreCredentials(cred *Credentials, ctx appengine.Context) (err error) {
//	key := datastore.NewKey(ctx, "Credentials", "main", 0, nil)
//	_, err = datastore.Put(ctx, key, cred)
//	return
//}
func Load() {
	log.Print("Load")
}
