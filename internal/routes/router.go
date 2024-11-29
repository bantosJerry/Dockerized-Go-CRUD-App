package routes

import (
	"go-crud-app/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	authHandler := handlers.AuthHandler{DB: db}

	// Public routes
	router.HandleFunc("/api/register", authHandler.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/login", authHandler.LoginHandler).Methods(http.MethodPost)

	// Protected routes (CRUD operations)
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(authHandler.Authenticate)
	protected.HandleFunc("/users", authHandler.GetAllUsersHandler).Methods(http.MethodGet)
	protected.HandleFunc("/user/{id}", authHandler.UpdateUserHandler).Methods(http.MethodPut)
	protected.HandleFunc("/user/{id}", authHandler.DeleteUserHandler).Methods(http.MethodDelete)

	return router
}
