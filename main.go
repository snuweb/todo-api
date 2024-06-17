package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodoByID).Methods("/GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/{id}", deleteTodo).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
