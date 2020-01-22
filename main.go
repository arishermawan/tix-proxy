package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"tix-proxy/config"
	"tix-proxy/handler"

	"github.com/gorilla/mux"
)

func main() {
	config.Load()
	router := mux.NewRouter()
	api := router.PathPrefix("/internal/v1").Subrouter()
	api.HandleFunc("/aviator/reviews", handler.AviatorReviewHandler).Methods(http.MethodGet)
	api.HandleFunc("/aviator/photos", handler.AviatorPhotoHandler).Methods(http.MethodGet)
	api.HandleFunc("/aviator/availability/date", handler.AviatorAvailableDateHandler).Methods(http.MethodGet)
	log.Info("GoTix Proxy API Server start...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
