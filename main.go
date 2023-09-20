package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	getEnv := godotenv.Load()

	if getEnv != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Running development server on port " + portString + "...")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/test", handlerReadiness)
	v1Router.Get("/error", handlerErr)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
} 