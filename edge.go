package graph

import "fmt"

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

func (e *Edge) String() string {
	return fmt.Sprintf("%V:%V::%d", e.From(), e.To(), e.Cost())
}
