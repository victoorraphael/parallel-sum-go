package main

type M struct {
	Tipo  int
	Venda float64
	Lucro float64
}

func (m *M) V() float64 {

	switch m.Tipo {
	case 1:
		return -m.Venda
	default:
		return m.Venda
	}
}

type ListaM []M

type MVenda struct {
	ListaM
}

func (m MVenda) Sum() float64 {

	n := len(m.ListaM)
	var total float64
	switch n {
	case 0:
		return total
	case 1:
		return m.ListaM[0].Venda
	default:
		for i := 0; i < n; i++ {
			total += m.ListaM[i].V()
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
