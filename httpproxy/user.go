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

type RelationShip struct {
	UserId int64
	State  string
	Type   string
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, []User{{Id: int64(1), Name: "nnnn", Type: "user"}})
	log.Println("GetUsers")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, User{Id: int64(1), Name: "nnnn", Type: "user"})
	log.Println("GetUsers")
}

func GetRelationships(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, []RelationShip{{UserId: 1212, State: "like", Type: "relationship"}})
	log.Println("GetUsers")
}

func UpdateRelationship(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, RelationShip{UserId: 1212, State: "like", Type: "relationship"})
	log.Println("GetUsers")
}
