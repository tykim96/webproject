package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tykim96/webproject/src/handlers"
)

func SetupUserRoute(router *mux.Router, userHandler handlers.UserHandler) {
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
}
