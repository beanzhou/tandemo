package httpproxy

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Id   int64
	Name string
	Type string
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, []User{{Id: int64(1), Name: "nnnn", Type: "user"}})
	log.Println("GetUsers")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, User{Id: int64(1), Name: "nnnn", Type: "user"})
	log.Println("GetUsers")
}

