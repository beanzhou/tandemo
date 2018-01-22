package httpproxy

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type RelationShip struct {
	UserId int64
	State  string
	Type   string
}

func GetRelationships(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, []RelationShip{{UserId: 1212, State: "like", Type: "relationship"}})
	log.Println("GetUsers")
}

func UpdateRelationship(w http.ResponseWriter, r *http.Request) {
	writeJson(context.Background(), w, 200, RelationShip{UserId: 1212, State: "like", Type: "relationship"})
	log.Println("GetUsers")
}
