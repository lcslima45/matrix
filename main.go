package main

import (
	"fmt"
	"math"
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


func TriangularMatrix(m [][]float64) ([][]float64, int) {
	if len(m) == 0 {
		return nil, 0
	}
	changes := 1

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

func ReverseSubstitution(U [][]float64, b []float64) []float64 {
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
	m1, _ := TriangularMatrix(m)
	for i := 0; i < len(m); i++ {
		det *= m1[i][i]
	}
	return det 
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
	return DetLaplace(m) == 0 
}

func LinearlyIndependent(m [][]float64) bool {
	return !LinearlyDependent(m)
}

func TriangularizeLinearSystem(m [][]float64, b []float64) ([][]float64, []float64) {
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

func GaussMethod(m [][]float64, b []float64) []float64 {
	m, b = TriangularizeLinearSystem(m, b)
	if m == nil || b == nil {
		return nil 
	}
	return ReverseSubstitution(m,b)
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
	m, b = TriangularizeLinearSystem(m, b)
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


func main() {
	matriz2:= [][]float64{
		{4,3,6},
		{7,8,9},
		{1,2,3},
	}
	b := []float64{13,14,21}
	fmt.Println(GaussMethod(matriz2, b))
}	