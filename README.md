# parallel-sum-go

O objetivo é conseguir somar N structs e o campo a ser utilizado para a soma pode variar de acordo com a demanda ou a struct.

comando usado como referência para comparação:
```bash
go test -bench=. -count 3 -cpu=3 -benchtime=5s -benchmem -memprofile=mem.out -cpuprofile=cpu.out
```
