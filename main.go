package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
)

func OutputMatrix(m [][]float64) {
	file, _ := os.OpenFile("outputMatrix.txt", os.O_CREATE|os.O_WRONLY, 0644)
	for _, line := range m {
		fmt.Fprintln(file, line)
	}
	file.Close()
}

func WriteMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

func Identity(n int) [][]float64 {
	id := make([][]float64, n)
	for i := 0; i < n; i++ {
		id[i] = make([]float64, n)
		id[i][i] = 1
	}

	return id
}

func Dot(lambda float64, m [][]float64) [][]float64 {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			m[i][j] *= lambda
		}
	}
	return m
}

func Transpose(m [][]float64) [][]float64 {
	var t [][]float64
	for i := 0; i < len(m); i++ {
		t = append(t, Columns(m,i))
	}
	return t
}


func Inverse(m [][]float64) [][]float64 {
	det := DetLaplace(m)
	if det == 0 {
		return nil 
	}
	var result [][]float64 
	result = Dot((1/det),Adjoint(m))
	return result
}


func LowerTriangular(m [][]float64) ([][]float64, float64) {
	if len(m) == 0 {
		return nil, 0
	}
	changes := float64(1)

	for p := len(m) - 1; p >= 0; p-- {
		if p == 0 {
			break
		}

		pivot := m[p][p]
		aux := p

		for pivot == 0 && aux > 0 {
			aux--
			pivot = m[aux][p]
		}

		if pivot == 0 {
			return nil, 0
		}

		if aux != p {
			m[p], m[aux] = m[aux], m[p]
			changes *= -1
		}

		for i := p - 1; i >= 0; i-- {
			lambda := m[i][p] / m[p][p]
			for j := p; j >= 0; j-- { // Corrigido para j=p até 0 (apenas as colunas relevantes)
				m[i][j] -= lambda * m[p][j]
			}
		}
	}
	return m, changes
}


func UpperTriangular(m [][]float64) ([][]float64, float64) {
	if len(m) == 0 {
		return nil, 0
	}
	changes := float64(1)

	for p := 0; p < len(m); p++ {
		if p == len(m)-1 {
			break
		}
		
		pivot := m[p][p]
		aux := p

		for pivot == 0 && aux < len(m)-1 {
			aux++
			pivot = m[aux][p]
		}

		if pivot == 0 {
			return nil, 0
		}

		if aux != p {
			m[p], m[aux] = m[aux], m[p]
			changes *= -1
		}

		for i := p + 1; i < len(m); i++ {
			lambda := m[i][p] / m[p][p]
			for j := p; j < len(m[0]); j++ { // Corrigido para j=p (não precisa alterar colunas já zeradas)
				m[i][j] -= lambda * m[p][j]
			}
		}
	}
	return m, changes
}

func ForwardSubstitution(L [][]float64, b []float64) []float64 {
	n := len(b)
	y := make([]float64, n)

	for i := 0; i < n; i++ {  // Processa cada linha de cima para baixo
		sum := 0.0
		for j := 0; j < i; j++ {  // Considera todos os elementos antes da diagonal
			sum += L[i][j] * y[j]  // Multiplica pelo valor já calculado de y[j]
		}
		y[i] = (b[i] - sum) / L[i][i]  // Resolve o sistema com o termo independente
	}
	return y
}


func BackwardSubstitution(U [][]float64, b []float64) []float64 {
	n := len(b)
	x := make([]float64, n)

	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ { // Corrigido: começa de i+1
			sum += U[i][j] * x[j]  // Corrigido: multiplica por x[j]
		}
		x[i] = (b[i] - sum) / U[i][i] // Divide pelo elemento da diagonal
	}
	return x
}

func CofactorsMatrix(m [][]float64) [][]float64 {
	var result [][]float64 
	for i := 0; i < len(m); i++ {
		var l []float64
		for j := 0; j < len(m); j++ {
			l = append(l, math.Pow(-1, float64((j+1) + (i+1))) * DetLaplace(Cofactors(m, i, j)))
		}
		result = append(result, l)
	}
	return result
}

func Adjoint(m [][]float64) [][]float64 {
	return Transpose(CofactorsMatrix(m))
}

func Lines(m [][]float64, l int) []float64 {
	return m[l]
}

func Columns(m [][]float64, c int) []float64 {
	var column []float64
	for l := 0; l < len(m); l++ {
		column = append(column, m[l][c])
	}
	return column
}

func CrossProduct(a1, a2 []float64) float64 {
	if len(a1) != len(a2) {
		return math.NaN()
	}
	var sum float64
	for i := 0; i < len(a1); i++ {
		sum += a1[i] * a2[i] 
	}
	return sum 
}

func SumNMatrix(m1 ...[][]float64) [][]float64 {
	if len(m1) == 0 {
		return nil 
	}
	result := m1[0]

	for i := 1; i < len(m1); i++ {
		result = SumMatrix(result, m1[i])
	}

	return result
}

func SumMatrix(m1, m2 [][]float64) [][]float64 {
	if len(m1) != len(m2) || len(m1[0]) != len(m2[0]) {
		return nil
	}
	var result [][]float64
	for i := 0; i < len(m1); i++ {
		var l []float64
		for j := 0; j < len(m1[0]); j++ {
			l = append(l, m1[i][j] + m2[i][j]) 
		}
		result = append(result, l)
	}

	return result
}

func MatrixProduct(m1, m2 [][]float64) [][]float64 {
	if len(m1[0]) != len(m2) {
		return nil
	}
	var result [][]float64
	for i := 0; i < len(m1); i++ {
		var l []float64
		for j := 0; j < len(m2[0]); j++ {
			l = append(l, CrossProduct(Lines(m1,i), Columns(m2,j)))
		}
		result = append(result, l)
	}

	return result 
}

func MatrixNProduct(m1 ...[][]float64) [][]float64 {
	if len(m1) == 0 {
		return nil
	}
	result := m1[0]

	for i := 1; i < len(m1); i++ {
		result = MatrixProduct(result, m1[i])
	}

	return result
}

func Cofactors(m [][]float64, i, j int) [][]float64 {
	var n [][]float64
	for i2 := 0; i2 < len(m); i2++ {
		if i2 == i {
			continue // Ignora a linha 'i'
		}
		var l []float64
		for j2 := 0; j2 < len(m); j2++ {
			if j2 == j {
				continue // Ignora a coluna 'j'
			}
			l = append(l, m[i2][j2])
		}
		n = append(n, l)
	}
	return n
}

func DetGauss(m [][]float64) float64{
	det := 1.0 
	m1, changes := UpperTriangular(m)
	for i := 0; i < len(m); i++ {
		det *= m1[i][i]
	}
	return det * changes  
}

func DetLaplace(m [][]float64) float64 {
	if len(m) != len(m[0]) {
		return math.NaN()
	}
	if len(m[0]) == 0 {
		return math.NaN()
	}
	if len(m[0]) == 1 {
		return m[0][0]
	}

	det := 0.0
	n := len(m)
	i := 0
	for j := 0; j < n; j++ {
		det += math.Pow(-1, float64((j+1) + (i+1))) * m[i][j] * DetLaplace(Cofactors(m, i, j))
	}
	return det  
}

func LinearlyDependent(m [][]float64) bool {
	return DetGauss(m) == 0 
}

func LinearlyIndependent(m [][]float64) bool {
	return !LinearlyDependent(m)
}


func LowerTriangularSystem(m [][]float64, b []float64) ([][]float64, []float64) {
	if len(m) == 0 {
		return nil, nil
	}

	changes := 1 // Para controle de trocas de linhas

	for p := len(m) - 1; p >= 0; p-- {
		if p == len(m)-1 {
			break
		}
		pivot := m[p][p]
		aux := p
		for pivot == 0 && aux > 0 {
			aux--
			pivot = m[aux][p]
		}
		if pivot == 0 {
			return nil, nil 
		}
		if aux != p {
			m[p], m[aux] = m[aux], m[p]
			b[p], b[aux] = b[aux], b[p] // Aplicando a troca no vetor b
			changes *= -1
		}
		for i := p - 1; i >= 0; i-- {
			lambda := m[i][p] / m[p][p]
			for j := p; j >= 0; j-- { 
				m[i][j] -= lambda * m[p][j]
			}
			b[i] -= lambda * b[p]
		}
	}
	return m, b
}


func UpperTriangularSystem(m [][]float64, b []float64) ([][]float64, []float64) {
	if len(m) == 0 {
		return nil, nil
	}

	changes := 1 // Para controle de trocas de linhas

	for p := 0; p < len(m); p++ {
		if p == len(m)-1 {
			break
		}
		pivot := m[p][p]
		aux := p
		for pivot == 0 && aux < len(m)-1 {
			aux++
			pivot = m[aux][p]
		}
		if pivot == 0 {
			return nil, nil 
		}
		if aux != p {
			m[p], m[aux] = m[aux], m[p]
			b[p], b[aux] = b[aux], b[p] // Aplicando a troca no vetor b
			changes *= -1
		}
		for i := p + 1; i < len(m); i++ {
			lambda := m[i][p] / m[p][p]
			for j := p; j < len(m[0]); j++ {
				m[i][j] -= lambda * m[p][j]
			}
			b[i] -= lambda * b[p]
		}
	}
	return m, b
}

func GaussMethodForward(m [][]float64, b []float64) []float64 {
	m, b = LowerTriangularSystem(m, b)
	if m == nil || b == nil {
		return nil 
	}
	return ForwardSubstitution(m,b)
}

func GaussMethod(m [][]float64, b []float64) []float64 {
	m, b = UpperTriangularSystem(m, b)
	if m == nil || b == nil {
		return nil 
	}
	return BackwardSubstitution(m,b)
}

func DiagonalizeLinearSystem(m [][]float64, b []float64) ([][]float64, []float64) {
	for i := len(m) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
				lambda := m[j][i] / m[i][i]
				m[j][i] -= lambda * m[i][i]
				b[j] -= lambda * b[i]
		}
	}
	return m, b
}

func GaussJordanMethod(m [][]float64, b []float64) []float64 {
	m, b = UpperTriangularSystem(m, b)
	if m == nil || b == nil {
		return nil
	}
	x := make([]float64, len(b))
	m, b = DiagonalizeLinearSystem(m,b) 
	for i := 0; i < len(m); i++ {
		x[i] = b[i] / m[i][i]
	}
	return x
}

func LU(m [][]float64) ([][]float64, [][]float64) {
	l := Identity(len(m))
	u := m 
	if len(u) == 0 {
		return nil, nil 
	}
	changes := float64(1)

	for p := 0; p < len(u); p++ {
		if p == len(u)-1 {
			break
		}
		
		pivot := u[p][p]
		aux := p

		for pivot == 0 && aux < len(u)-1 {
			aux++
			pivot = u[aux][p]
		}

		if pivot == 0 {
			return nil, nil 
		}

		if aux != p {
			u[p], u[aux] = u[aux], u[p]
			changes *= -1
		}

		for i := p + 1; i < len(u); i++ {
			lambda := u[i][p] / u[p][p]
			l[i][p] = lambda
			for j := p; j < len(m[0]); j++ { // Corrigido para j=p (não precisa alterar colunas já zeradas)
				u[i][j] -= lambda * u[p][j]
			}
		}
	}

	return l, u 
}


type Matrix struct {
	Matrix [][]float64 `json:"matrix"`
}

type LinearSystem struct {
	Matrix [][]float64 `json:"matrix"`
	B []float64 `json:"b"`
}

func handleMatrix(w http.ResponseWriter, r *http.Request) {
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
		det := DetGauss(requestData.Matrix)
		response := map[string]float64{"determinant": det}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func handleLinearSystem(w http.ResponseWriter, r *http.Request) {
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
		result := GaussMethod(requestData.Matrix, requestData.B)
		fmt.Println("result:", result)
		response := map[string][]float64{"result": result}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	http.HandleFunc("/determinant", handleMatrix)
	http.HandleFunc("/linearsystem", handleLinearSystem)

	fmt.Println("Servidor rodando na porta :8080")
	http.ListenAndServe(":8080", nil)
}	