package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var todos = []Todo{
	{ID: 1, Title: "Learn Go", Status: "pending"},
	{ID: 2, Title: "Build a REST API", Status: "completed"},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contect-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}

func getTodoByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for _, item := range todos {

		if item.ID == id {
			json.NewEncoder(w).Encode(item)

			return
		}

	}
	http.Error(w, "Todo not found", http.StatusNotFound)

}

func createTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var todo Todo

	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	json.NewEncoder(w).Encode(todo)

}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, item := range todos {

		if item.ID == id {

			todos = append(todos[:i], todos[i+1:]...)

			var todo Todo
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return

		}
	}

	http.Error(w, "Todo ot found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, item := range todos {

		if item.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(todos)

}
