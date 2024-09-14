package router

import (
	"github.com/backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/createUser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/getAllUsers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/createGame", controller.CreateGame).Methods("POST")
	router.HandleFunc("/getOneGame/{id}", controller.GetOneGame).Methods("GET")
	router.HandleFunc("/getOneUser/{name}", controller.GetOneUser).Methods("GET")
	return router

}