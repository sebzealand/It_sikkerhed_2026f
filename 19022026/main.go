package main

import (
	"fmt"
	"go-rest-api/handlers"
	"net/http"
)

func main() {
	// Beskyttede endpoints
	http.HandleFunc("/users", handlers.ValidateTokenHandler(handlers.UserHandler))

	// Offentlige endpoints
	http.HandleFunc("/docs", handlers.DocsHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	fmt.Println("Serveren starter på http://localhost:8000")
	fmt.Println("Se dokumentation på http://localhost:8000/docs")

	// Åbner port 8000 og lytter og reagere på http requests
	http.ListenAndServe(":8000", nil)
}
