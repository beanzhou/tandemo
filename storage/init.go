package storage

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	db *pg.DB
)

func init() {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "root",
		Database: "testdb",
		Addr:     "127.0.0.1:5432",
	})

	err := db.CreateTable(&User{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Errorln(err)
	}
}
