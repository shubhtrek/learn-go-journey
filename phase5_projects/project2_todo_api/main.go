package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: "1", Title: "Learn Go variables", Done: true},
	{ID: "2", Title: "Master Go Concurrency", Done: false},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", getTodos)

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
