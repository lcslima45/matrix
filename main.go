package main

import (
	"fmt"
	"net/http"

	"github.com/lcslima45/matrix/handlers"
)

func main() {
	http.HandleFunc("/determinant", handlers.HandleMatrix)
	http.HandleFunc("/linearsystem", handlers.HandleLinearSystem)
	http.HandleFunc("/sum", handlers.HandlerSum)
	http.HandleFunc("/product", handlers.HandlerProduct)

	fmt.Println("Servidor rodando na porta :8080")
	http.ListenAndServe(":8080", nil)
}
