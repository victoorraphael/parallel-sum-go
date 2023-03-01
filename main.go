// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	data := make([]M, 500_000)
	for i := range data {
		data[i] = M{
			Tipo:  rand.Intn(5),
			Venda: rand.ExpFloat64(),
		}
	}

	total := ParallelSumGeneric(MVenda{ListaM: data})
	fmt.Println("TOTAL =", total)
}
