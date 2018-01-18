package httpproxy

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

const DocPathPrefix string = "/doc/"

var (
	docDir = flag.String("docDir", "./doc/", "the dir of doc")
)

// Be careful, JsonPath should be the same as SwaggerUIBundle.url defined in swagger-ui-dist/index.html
const JsonPath string = "/swagger.json"

func AddDocHandlers(router *mux.Router) {
	log.Infoln(*docDir)
	fs := http.FileServer(http.Dir(filepath.Join(*docDir, "swagger-ui-dist")))
	router.Methods("GET").PathPrefix(DocPathPrefix).Handler(http.StripPrefix(DocPathPrefix, fs))
	router.Methods("GET").Path(JsonPath).HandlerFunc(ServeSwagger)
}

func ServeSwagger(w http.ResponseWriter, r *http.Request) {
	fn := "doc.yaml"
	log.Infoln(filepath.Join(*docDir, fn))
	file, _ := ioutil.ReadFile(filepath.Join(*docDir, fn))
	byte_data, _ := yaml.YAMLToJSON(file)
	var m map[string]interface{}
	json.Unmarshal(byte_data, &m)
	m["host"] = r.Host

	writeJson(r.Context(), w, http.StatusOK, m)
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
