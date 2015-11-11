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

func (c *Color) String() string {
	switch *c {
	case Black:
		return "black"
	case White:
		return "white"
	case Gray:
		return "gray"
	}

	return "unknown"
}

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
	depth         int64
	discoveryTime int64
	finishTime    int64
}

func (vi *VisitInfo) DiscoveryTime() int64 {
	return vi.discoveryTime
}

func (vi *VisitInfo) FinishTime() int64 {
	return vi.finishTime
}

func (vi *VisitInfo) Depth() int64 {
	return vi.depth
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
