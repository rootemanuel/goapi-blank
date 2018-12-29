package dao

import (
	"log"

	. "github.com/rootemanuel/goapi-blank/entity"
)

type RootDao struct{}

const (
	DB_ROOT   = "root"
	COLL_ROOT = "root"
)

func (m *RootDao) FindAll() ([]RootEntity, error) {
	var roots []RootEntity
	session := GetSession()
	defer session.Close()

	err := session.DB(DB_ROOT).C(COLL_ROOT).Find(nil).All(&roots)
	if err != nil {
		log.Println(err)
	}

	return roots, err
}
