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

func Transpose(m [][]float64) [][]float64 {
	var t [][]float64
	for i := 0; i < len(m); i++ {
		t = append(t, Columns(m,i))
	}
	return t
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
	matriz1 := [][]float64{
        {1, 2, 3},
        {4, 8, 6},
        {7, 8, 10},
    }
	matriz2 := [][]float64{
        {2, 3, 4},
        {5, 6, 7},
        {8, 9, 11},
    }

	WriteMatrix(SumMatrix(matriz1, matriz2))
	WriteMatrix(matriz1)
	fmt.Println("Transposta")
	WriteMatrix(Transpose(matriz1))
}