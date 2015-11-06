package graph

import (
	"sync"
)

type Color uint
type Identity uint

const (
	Black = Color(1) << iota
	White
	Gray
)

const (
	Empty = Identity(1) << iota
	Dense
	Sparse
)

func (i *Identity) String() string {
	switch *i {
	case Empty:
		return "empty"
	case Dense:
		return "dense"
	case Sparse:
		return "sparse"
	}
	return "unknown-identity"
}

type Vertex struct {
	data        interface{}
	color       Color
	predecessor *Vertex
	meta        interface{}
}

type VisitInfo struct {
	discoveryTime int64
	finishTime    int64
}

type Edge struct {
	from, to *Vertex
	cost     float64
}

type Graph struct {
	mu       sync.Mutex
	etag     int64
	vertices map[interface{}]*Vertex
	edges    AdjMap
}

type AdjMap map[interface{}]map[interface{}]*Edge
