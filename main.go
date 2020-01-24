package main

import (
	"net/http"

	"tix-proxy/config"
	"tix-proxy/router"

	log "github.com/sirupsen/logrus"
)

func main() {
	config.Load()
	router := router.InternalRouter()
	log.Info("GoTix Proxy API Server start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
