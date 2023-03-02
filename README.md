# parallel-sum-go

O objetivo é conseguir somar N structs e o campo a ser utilizado para a soma pode variar de acordo com a demanda ou a struct.

comando usado como referência para comparação:
```bash
go test -bench=. -count 3 -cpu=3 -benchtime=5s -benchmem -memprofile=mem.out -cpuprofile=cpu.out -race
```

Depois de rodar o teste pela primeira vez da para usar o pprof para otimizar. go > 1.20

```bash
cp cpu.out cpu2.out
go test -bench=. -count 3 -cpu=3 -benchtime=5s -benchmem -memprofile=mem.out -cpuprofile=cpu.out -pgo=cpu2.out
```

E no build

```bash
go build . -pgo=cpu.out
```
