package main

import (
	"testing"
)

func BenchmarkDetLaplace(b *testing.B) {
	matriz2:= [][]float64{
		{4,3,6},
		{7,8,9},
		{1,2,3},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_  = DetLaplace(matriz2)
	}
}

func BenchmarkDetGauss(b *testing.B) {
	matriz2:= [][]float64{
		{4,3,6},
		{7,8,9},
		{1,2,3},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_  = DetGauss(matriz2)
	}
}