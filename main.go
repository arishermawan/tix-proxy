package main

import (
	"log"
	"net/http"

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
	log.Fatal(http.ListenAndServe(":8080", router))
}
