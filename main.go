package main

import (
	"go-crud-app/internal/handlers"
	"go-crud-app/internal/models"
	"go-crud-app/internal/routes"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()

	// Initialize the router
	router := routes.SetupRoutes(models.DB)

	// Serve login.html for the root route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/login.html")
	})

	// Serve other static files from the 'frontend' directory
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	// Auth routes
	authHandler := &handlers.AuthHandler{DB: models.DB}
	router.HandleFunc("/api/register", authHandler.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/api/users", authHandler.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/user/{id}", authHandler.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/user/{id}", authHandler.DeleteUserHandler).Methods("DELETE")

	// Start the server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
