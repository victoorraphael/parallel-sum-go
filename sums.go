package main

import (
	"math"
	"runtime"
)

func ParallelSumGeneric(movs Summable) float64 {
	numOfCores := runtime.NumCPU()
	result := make(chan float64, numOfCores)

	partialSize := movs.Len() / numOfCores
	for i := 0; i < numOfCores; i++ {
		start := i * partialSize
		end := start + partialSize

		if i == numOfCores-1 {
			end = movs.Len()
		}

		go func(m Summable, results chan<- float64) {
			results <- m.Sum()
		}(movs.Partial(start, end), result)
	}

	var total float64
	for i := 0; i < numOfCores; i++ {
		total += <-result
	}
	return total
}

func Sum(movs []M) float64 {
	n := len(movs)
	var total float64
	switch n {
	case 0:
		return total
	case 1:
		return movs[0].Venda
	default:
		for i := 0; i < n; i++ {
			total += movs[i].V()
		}
	}
	return total
}

func ParallelSum(movs []M) float64 {
	numOfCores := runtime.NumCPU()
	result := make(chan float64, numOfCores)

	partialSize := int(math.Round(float64(len(movs)) / float64(numOfCores)))
	for i := 0; i < numOfCores; i++ {
		start := i * partialSize
		end := start + partialSize

		if i == numOfCores-1 {
			end = len(movs)
		}

		go func(m []M, results chan<- float64) {
			results <- Sum(m)
		}(movs[start:end], result)
	}

	var total float64
	for i := 0; i < numOfCores; i++ {
		total += <-result
	}
	return total
}

func ParallelSumGeneric2(movs Summable) float64 {
	return movs.Sum()
}
