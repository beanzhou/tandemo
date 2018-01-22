package httpproxy

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/beanzhou/tandemo/storage"

	"github.com/gorilla/mux"
)

type RelationShip struct {
	UserId int64
	State  string
	Type   string
}

func GetRelationships(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		writeJson(context.Background(), w, 400, err)
	}

	rsManager := storage.RelationshipManager{}
	list, err := rsManager.GetRelationships(int64(userid))
	if err != nil {
		writeJson(context.Background(), w, 500, err)
		return
	}
	rsList := make([]RelationShip, len(list))
	for index, rs := range list {
		rsList[index].UserId = rs.OtherId
		rsList[index].State = rs.State
		rsList[index].Type = rs.Type
	}
	writeJson(context.Background(), w, 200, rsList)
}

// TODO: validate userid / state type
func UpdateRelationship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		writeJson(context.Background(), w, 400, err)
	}

	otherUserid, err := strconv.Atoi(vars["other_user_id"])
	if err != nil {
		writeJson(context.Background(), w, 400, err)
	}

	decoder := json.NewDecoder(r.Body)
	var t RelationShip
	err = decoder.Decode(&t)
	if err != nil {
		writeJson(context.Background(), w, 400, err)
	}
	defer r.Body.Close()

	rsManager := storage.RelationshipManager{}
	rs, err := rsManager.Insert(storage.Relationship{
		UserId:  int64(userid),
		OtherId: int64(otherUserid),
		Type:    "relationship",
		State:   t.State,
	})
	if err != nil {
		writeJson(context.Background(), w, 500, err)
		return
	}
	writeJson(context.Background(), w, 200, rs)
}
