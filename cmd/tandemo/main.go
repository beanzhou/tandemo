package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/beanzhou/tandemo/httpproxy"
	"github.com/gorilla/handlers"
)

var (
	flagBind = flag.String("bind", ":8061", "http bind port")
)

func main() {
	flag.Parse()

	router := httpproxy.NewRouter()
	handler := handlers.CompressHandler(router)

	log.Info("listen at ", *flagBind)
	panic(http.ListenAndServe(*flagBind, handler))
}
