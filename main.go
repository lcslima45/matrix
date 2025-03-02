package main

import (
	"fmt"
	"math"
)
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
	fmt.Println(DetLaplace([][]float64{
		{8, 9, 7, 6, 8},
		{10, 8, 3, 6, 6},
		{6, 1, 7, 2, 3},
		{8, 2, 10, 8, 3},
		{3, 10, 4, 4, 7},
	}))
	fmt.Println(DetLaplace([][]float64{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}))
	fmt.Println(DetLaplace([][]float64{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}))
}