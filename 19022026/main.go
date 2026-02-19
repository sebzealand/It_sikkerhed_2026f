package main

import (
	"fmt"
	"net/http"
	"go-rest-api/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.UserHandler)

	fmt.Println("Serveren starter på http://localhost:8000")
	http.ListenAndServe(":8000",  nil)
}
