package routes

import (
	"github.com/gorilla/mux"
	"github.com/viveksingh-01/ginger-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/restaurants", handlers.HandleRestaurants)
}
