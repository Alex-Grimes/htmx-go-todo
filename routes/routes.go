package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	// sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// sendTodos(w)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	// sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/mark", markTodo).Methods("PUT")
	mux.HandleFunc("/delete", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", mux))
}
