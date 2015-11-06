package graph

func (e *Edge) From() interface{} {
	return e.from.Data()
}

func (e *Edge) To() interface{} {
	return e.to.Data()
}

func (e *Edge) Cost() float64 {
	return e.cost
}

func (e *Edge) SetCost(c float64) bool {
	if e == nil {
		return false
	}

	e.cost = c
	return true
}
