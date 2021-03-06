package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading environment variables.")
	}
}

func main() {
	ApiHost := os.Getenv("API_HOST")
	app := mux.NewRouter()

	app.HandleFunc("/posts/{page}/", handlers.PostsHandler).Methods("GET")
	app.HandleFunc("/cat/{catSlug}/{page}/", handlers.PostsByCatHandler).Methods("GET")
	app.HandleFunc("/cats/{page}/", handlers.CategoriesHandler).Methods("GET")

	server := &http.Server{
		Handler:      app,
		Addr:         ApiHost + ":" + os.Getenv("API_PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
