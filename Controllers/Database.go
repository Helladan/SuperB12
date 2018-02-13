package Controllers

import (
	"github.com/globalsign/mgo"
)

type Database struct {
	Host         string
	DatabaseName string

	Database *mgo.Database
	Session  *mgo.Session

	err error
}

func NewDatabase(host string, databaseName string) Database {
	db := Database{
		Host:         host,
		DatabaseName: databaseName,
	}

	db.GetSession()
	db.GetDatabase()

	return db
}

func (db *Database) GetSession() error {
	db.Session, db.err = mgo.Dial(db.Host)

	return db.err
}

func (db *Database) GetDatabase() {
	db.Database = db.Session.DB(db.DatabaseName)
}
