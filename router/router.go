package router

import (
	"net/http"

	"tix-proxy/handler"

	"github.com/gorilla/mux"
)

func InternalRouter() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/internal/v1").Subrouter()
	api.HandleFunc("/aviator/reviews", handler.AviatorReviewHandler).Methods(http.MethodGet)
	api.HandleFunc("/aviator/photos", handler.AviatorPhotoHandler).Methods(http.MethodGet)
	api.HandleFunc("/aviator/availability/date", handler.AviatorAvailableDateHandler).Methods(http.MethodGet)
	api.HandleFunc("/aviator/event/detail", handler.AviatorEventDetailHandler).Methods(http.MethodGet)
	return router
}
