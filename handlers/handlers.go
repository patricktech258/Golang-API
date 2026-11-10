package handlers

import (
	"database/sql"
	"encoding/json"
	model "example/main/models"
	"fmt"
	"net/http"
	"strconv"
)

func AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "JSON inválido ou mal formatado", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT email from users WHERE email = ?", user.Email)

		if err != nil {
			http.Error(w, "Erro no banco", 500)
			return
		}

		defer rows.Close()

		db.Query("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)

		w.WriteHeader(http.StatusCreated)
	}
}

func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT id, name, email FROM users")

		if err != nil {
			http.Error(w, "Erro na Query", 400)
		}

		users := make([]model.User, 0)

		for rows.Next() {
			var user model.User

			err := rows.Scan(&user.Id, &user.Name, &user.Email)

			if err != nil {
				fmt.Println("Erro ao scannear a linha %w", err)
			}

			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Erro ao iterar os usuários", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Created_at)

		if err == sql.ErrNoRows {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "Erro no banco", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "JSON inválido ou mal formatado", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		result, err := db.Exec(
			"UPDATE users SET name = ?, email = ? WHERE id = ? ",
			user.Name, user.Email, id,
		)

		if err != nil {
			http.Error(w, "Erro no banco", http.StatusInternalServerError)
			return
		}

		affected, err := result.RowsAffected()

		if err != nil {
			http.Error(w, "Erro ao tentar alterar este usuário", http.StatusExpectationFailed)
			return
		}

		if affected == 0 {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
		}

		result, err := db.Exec("DELETE FROM users WHERE id = ?", id)

		if err != nil {
			http.Error(w, "Erro na query", http.StatusInternalServerError)
		}

		affected, err := result.RowsAffected()

		if affected == 0 {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "Erro ao tentar deletar este usuário", http.StatusExpectationFailed)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
