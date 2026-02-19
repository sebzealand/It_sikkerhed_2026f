package main

import (
	"fmt"
	"net/http"
	"go-rest-api/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.UserHandler)

	http.HandleFunc("/docs", handlers.DocsHandler)

	fmt.Println("Serveren starter på http://localhost:8000")
	fmt.Println("Se dokumentation på http://localhost:8000/docs")

	http.ListenAndServe(":8000",  nil)
}
