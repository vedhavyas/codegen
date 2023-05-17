package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks", CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/tasks/{id}/toggle", ToggleTaskComplete).Methods(http.MethodPatch)

	return router
}