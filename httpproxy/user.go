package httpproxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/beanzhou/tandemo/storage"
)

type User struct {
	Id   int64
	Name string
	Type string
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userManager := storage.UserManager{}
	users, err := userManager.GetUsers()
	if err != nil {
		writeJson(context.Background(), w, 500, err)
		return
	}
	writeJson(context.Background(), w, 200, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var t User
	err := decoder.Decode(&t)
	if err != nil {
		writeJson(context.Background(), w, 400, err)
	}
	defer r.Body.Close()

	userManager := storage.UserManager{}
	user, err := userManager.Insert(storage.User{Name: t.Name, Type: "user"})
	if err != nil {
		writeJson(context.Background(), w, 500, err)
		return
	}

	writeJson(context.Background(), w, 200, user)
}
