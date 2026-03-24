package handlers

import (
	"encoding/json"
	"go-rest-api/internal/crypto"
	"go-rest-api/models"
	"go-rest-api/services"
	"net/http"
	"strings"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. modtag email og password fra request body
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userList, err := services.GetUsers()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 2. Tjek email og password
	var foundUser *models.User
	for i := range userList.Users {
		if userList.Users[i].Email == req.Email {
			foundUser = &userList.Users[i]
			break
		}
	}

	if foundUser == nil || !crypto.CheckPasswordHash([]byte(req.Password), foundUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 3. Lav JWT token
	tokenString, err := services.GenerateToken(foundUser.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// 4. Returner JWT token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: tokenString})
}

func ValidateTokenHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "No token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if err := services.ValidateToken(tokenString); err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
