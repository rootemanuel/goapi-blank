package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"

	. "github.com/rootemanuel/goapi-blank/entity"
)

type RootDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "root"
)

func (m *RootDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *RootDAO) FindAll() ([]RootEntity, error) {
	var roots []RootEntity
	err := db.C(COLLECTION).Find(nil).All(&roots)
	if err != nil {
		log.Println(err)
	}
	return roots, err
}
