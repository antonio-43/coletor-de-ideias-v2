package router

import (
	"cdi/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/add", middleware.CreateIdea)
	router.HandleFunc("/api/see", middleware.ShowData)

	return router
}
