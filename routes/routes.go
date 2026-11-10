package routes

import (
	"database/sql"
	handlers "example/main/handlers"
	"net/http"
)

func SetupRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", handlers.GetUsers(db))
	mux.HandleFunc("POST /create", handlers.AddUser(db))
	mux.HandleFunc("GET /users/{id}", handlers.GetUser(db))
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser(db))
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser(db))
	mux.HandleFunc("POST /upload", handlers.UploadHandler(db))

	return mux
}
