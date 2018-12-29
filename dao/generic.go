package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {

	if session == nil {
		sessionA, err := mgo.Dial("127.0.0.1")
		if err != nil {
			log.Fatal(err)
		}

		session = sessionA
		return sessionA.Copy()
	}

	return session.Copy()
}
