package router

import (
	"github.com/Divak-ar/mongo_go_api/controller"
	"github.com/gorilla/mux"
)

// since this router will be imported , it must return something to be used in main.go or main func
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("api/v1/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("api/v1/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("api/v1/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("api/v1/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}