package benchmarks

import (
	"testing"

	"github.com/lcslima45/matrix/calc"
)

func BenchmarkDetLaplace(b *testing.B) {
	matriz2 := [][]float64{
		{4, 3, 6},
		{7, 8, 9},
		{1, 2, 3},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = calc.DetLaplace(matriz2)
	}
}

func BenchmarkDetGauss(b *testing.B) {
	matriz2 := [][]float64{
		{4, 3, 6},
		{7, 8, 9},
		{1, 2, 3},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = calc.DetGauss(matriz2)
	}
}
