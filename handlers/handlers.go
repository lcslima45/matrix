package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lcslima45/matrix/calc"
)

type Matrix struct {
	Matrix [][]float64 `json:"matrix"`
}

type LinearSystem struct {
	Matrix [][]float64 `json:"matrix"`
	B      []float64   `json:"b"`
}

func HandleMatrix(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		var requestData Matrix
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println(requestData)
		det := calc.DetGauss(requestData.Matrix)
		response := map[string]float64{"determinant": det}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func HandleLinearSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		var requestData LinearSystem
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println("matrix:", requestData.Matrix)
		fmt.Println("b:", requestData.B)
		result := calc.GaussMethod(requestData.Matrix, requestData.B)
		fmt.Println("result:", result)
		response := map[string][]float64{"result": result}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

type MatrixSum struct {
	MatrixA [][]float64 `json:"matrixA"`
	MatrixB [][]float64 `json:"matrixB"`
}

func HandlerSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		var requestData MatrixSum
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println("matrix:", requestData.MatrixA)
		fmt.Println("b:", requestData.MatrixB)
		result := calc.SumMatrix(requestData.MatrixA, requestData.MatrixB)
		fmt.Println("result:", result)
		response := map[string][][]float64{"result": result}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func HandlerProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		var requestData MatrixSum
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println("matrix:", requestData.MatrixA)
		fmt.Println("b:", requestData.MatrixB)
		result := calc.MatrixProduct(requestData.MatrixA, requestData.MatrixB)
		fmt.Println("result:", result)
		response := map[string][][]float64{"result": result}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
