package main

import (
	"math/rand"
	"testing"
)

func BenchmarkParallelSum(b *testing.B) {
	data := make([]M, 500_000)
	for i := range data {
		data[i] = M{
			Tipo:  rand.Intn(5),
			Venda: rand.ExpFloat64(),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParallelSum(data)
	}
}

func BenchmarkSum(b *testing.B) {
	data := make([]M, 500_000)
	for i := range data {
		data[i] = M{
			Tipo:  rand.Intn(5),
			Venda: rand.ExpFloat64(),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data)
	}
}

func BenchmarkParallelSumGeneric(b *testing.B) {
	data := make([]M, 500_000)
	for i := range data {
		data[i] = M{
			Tipo:  rand.Intn(5),
			Venda: rand.ExpFloat64(),
		}
	}

	realData := MVenda{ListaM: data}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParallelSumGeneric(realData)
	}
}

func BenchmarkParallelSumGeneric2(b *testing.B) {
	data := make([]M, 500_000)
	for i := range data {
		data[i] = M{
			Tipo:  rand.Intn(5),
			Venda: rand.ExpFloat64(),
		}
	}

	realData := MVenda{ListaM: data}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		realData.Sum()
	}
}
