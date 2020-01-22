package client

import(
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LogRequest(req *http.Request) {
	log.Info("[AVIATOR] " + req.Method + " " + req.URL.Path)
}
