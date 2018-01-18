package httpproxy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetUsers",
		"GET",
		"/v1/users",
		GetUsers,
	},

	Route{
		"CreateUser",
		"POST",
		"/v1/users",
		CreateUser,
	},

	Route{
		"GetRelationships",
		"GET",
		"/v1/users/{user_id}/relationships",
		GetRelationships,
	},

	Route{
		"UpdateRelationship",
		"PUT",
		"/v1/users/{user_id}/relationships/{other_user_id}",
		UpdateRelationship,
	},
}

func writeJson(ctx context.Context, w http.ResponseWriter, statusCode int, v interface{}) {
	requestId, _ := ctx.Value("requestId").(string)
	header := w.Header()
	header.Set("x-request-id", requestId)
	header.Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Error(err)
	}
}

func createHandler(route Route) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				conn, _, _ := w.(http.Hijacker).Hijack()
				conn.Close()
			}
		}()
		route.HandlerFunc.ServeHTTP(w, r)
		log.Info("Name:", route.Name)
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(createHandler(route))
	}
	AddDocHandlers(router)
	return router
}
