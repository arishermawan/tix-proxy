package client

import(
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LogRequest(source string, req *http.Request) {
	log.Info(source + req.Method + " " + req.URL.Path)
}
