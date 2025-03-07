package main

import (
	"testing"
)

func BenchmarkGaussMethod(b *testing.B) {
	matriz2:= [][]float64{
		{4,3,6},
		{7,8,9},
		{1,2,3},
	}
	a := []float64{13,14,21}	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GaussMethod(matriz2, a)
	}
}

func BenchmarkGaussJordanMethod(b *testing.B) {
	matriz2:= [][]float64{
		{4,3,6},
		{7,8,9},
		{1,2,3},
	}
	a := []float64{13,14,21}	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GaussJordanMethod(matriz2, a)
	}
}