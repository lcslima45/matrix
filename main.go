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

func CofactorMatrix(m [][]float64, i, j int) [][]float64 {
	var n [][]float64
	for i2 := 0; i2 < len(m); i2++ {
		if i2 == i {
			continue 
		}
		var l []float64
		for j2 := 0; j2 < len(m); j2++ {
			if j2 == j {
				continue
			}
			l = append(l,m[i2][j2])
		}
		n = append(n, l)
	}
	return n
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
		det += math.Pow(-1, float64((j+1) + (i+1))) * m[i][j] * DetLaplace(CofactorMatrix(m, i, j))
	}
	return det  
}

func main() {
	OutputMatrix(Identity(100))
}