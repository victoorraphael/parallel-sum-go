// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := make([]M, 500_000)
	for i := range data {
		data[i] = M{
			Tipo:  rand.Intn(5),
			Venda: rand.ExpFloat64(),
		}
	}
	//cria arquivo de pprof
	//f, err := os.Create("parallel-sum.prof")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	t := time.Now()
	total := ParallelSumGeneric(MVenda{ListaM: data})
	fmt.Println("TOTAL =", total, time.Since(t))

	t = time.Now()
	total2 := Sum(data)
	fmt.Println("TOTAL =", total2, time.Since(t))

	t = time.Now()
	total3 := MVenda{ListaM: data}.Sum()
	fmt.Println("TOTAL =", total3, time.Since(t))
}
