package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", Hello).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	handler := router
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}


func Hello(w http.ResponseWriter, r *http.Request) {
	data := "Hello World"
	json.NewEncoder(w).Encode(data)
}