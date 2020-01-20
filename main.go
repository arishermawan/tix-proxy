package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"tix-proxy/config"
	"tix-proxy/handler"

	"github.com/gorilla/mux"
)

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func main() {
	config.Load()

	fmt.Println("------------------->>>", config.GetValue("PORT"))
	fmt.Println("------------------->>>", config.GetValue("AVIATOR_BASE_URL"))

	r := mux.NewRouter()
	api := r.PathPrefix("/internal/v1").Subrouter()
	api.HandleFunc("/aviator/reviews", handler.AviatorReviewHandler).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
