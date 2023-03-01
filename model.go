package main

type M struct {
	Tipo  int
	Venda float64
	Lucro float64
}

type ListaM []M

type MVenda struct {
	ListaM
}

func (m MVenda) Sum() float64 {
	if len(m.ListaM) == 0 {
		return 0
	}

	if len(m.ListaM) == 1 {
		return m.ListaM[0].Venda
	}

	var total float64
	for _, v := range m.ListaM {
		if v.Tipo == 1 {
			total -= v.Venda
		} else {
			total += v.Venda
		}
	}

	return total
}

func (m MVenda) Len() int {
	return len(m.ListaM)
}

func (m MVenda) Partial(start, end int) Summable {
	mv := MVenda{}
	mv.ListaM = m.ListaM[start:end]
	return mv
}

type Summable interface {
	Sum() float64
	Len() int
	Partial(start, end int) Summable
}
