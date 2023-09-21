package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/zulfikar-ditya/go-course/internal/database"
	_"github.com/lib/pq"
)

type APIConfig struct {
	DB *database.Queries
}

func main() {
	getEnv := godotenv.Load()

	if getEnv != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("$DB_URL must be set")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	apiConfig := APIConfig{
		DB: database.New(conn),
	}

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
	v1Router.Post("/users", apiConfig.handleCreateNewUser)
	
	router.Mount("/v1", v1Router)
	
	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	
	fmt.Println("Running development server on port " + portString + "...")
	err = server.ListenAndServe()
	
	if err != nil {
		log.Fatal(err)
	}
} 