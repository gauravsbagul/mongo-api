package router

import (
	"github.com/gauravsbagul/mongo-api/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsReadMovies).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovies).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
