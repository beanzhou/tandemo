package httpproxy

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUsers")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUsers")
}

func GetRelationships(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUsers")
}

func UpdateRelationship(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUsers")
}
