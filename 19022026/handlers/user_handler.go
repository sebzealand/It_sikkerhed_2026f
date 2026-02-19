package handlers

import (
	"encoding/json"
	"go-rest-api/models"
	"go-rest-api/services"
	"net/http"
	"strconv"
	"fmt"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	users, err := services.LoadUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Ugyldig JSON", http.StatusBadRequest)
			return
		}

		users.Users = append(users.Users, newUser)
		services.SaveUsers(users)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	
	case http.MethodPut:
		var updatedUser models.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, "Ugyldig JSON", http.StatusBadRequest)
			return
		}

		if services.UpdateUser(&users, updatedUser) {
			services.SaveUsers(users)
			json.NewEncoder(w).Encode(updatedUser)
		} else {
			http.Error(w, "Bruger ikke fundet", http.StatusNotFound)
		}

	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID skal være et tal", http.StatusBadRequest)
			return
		}

		if services.DeleteUser(&users, id) {
			services.SaveUsers(users)
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "Bruger ikke fundet", http.StatusNotFound)
		}

	default:
		http.Error(w, "Metode ikke tilladt", http.StatusMethodNotAllowed)
	}
}

func DocsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, `
        <h1>Go User API - Dokumentation</h1>
        <ul>
            <li><strong>GET /users</strong> - Lister alle brugere</li>
            <li><strong>GET /users?id=X</strong> - Henter en specifik bruger</li>
            <li><strong>POST /users</strong> - Opretter en ny bruger (JSON body påkrævet)</li>
            <li><strong>PUT /users</strong> - Opdaterer en bruger (JSON body påkrævet)</li>
            <li><strong>DELETE /users?id=X</strong> - Sletter en bruger</li>
        </ul>
    `)
}
