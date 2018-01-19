package storage

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	Id   int64
	Name string
	Type string
}

var db *pg.DB

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

func Insert() {
	user := User{Name: "test", Type: "user"}
	err := db.Insert(&user)
	if err != nil {
		log.Errorln(err)
	}
	log.Info(user)
}

func Update() {
	db.Update(&User{Id: 1, Name: "tt"})
}

func Delete() {
	db.Delete(&User{Id: 1})
}

func GetOne() {
	user := User{Id: 2}
	db.Select(&user)
	log.Printf("%v", user)
}

func GetList() {
	var users []User
	db.Model(&users).Limit(3).Offset(4).Select()
	log.Printf("%v", users)
}
