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
	Account  string
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
func store(ctx appengine.Context,) {

	dataKey := datastore.NewIncompleteKey(ctx, "Entity", nil)

	entity := &Entity{}
	entity.Value = "value one" // random value here
	entity.Count = rand.Intn(1000)

	if _, err := datastore.Put(ctx, dataKey, entity); err != nil {
		log.Print("err:  ", err)
		return
	}
	log.Print("store", entity.Count)
}

//func StoreCredentials(cred *Credentials, ctx appengine.Context) (err error) {
//	key := datastore.NewKey(ctx, "Credentials", "main", 0, nil)
//	_, err = datastore.Put(ctx, key, cred)
//	return
//}
func Load() {
	log.Print("Load")
}
