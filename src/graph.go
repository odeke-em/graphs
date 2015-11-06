package graph

import (
	"time"
)

func NewGraph() *Graph {
	g := Graph{
		vertices: make(map[interface{}]*Vertex),
		edges:    make(AdjMap),
		etag:     time.Now().Unix(),
	}

	return &g
}

func (g *Graph) EdgeCount() int {
	eCount := 0
	for _, neighbours := range g.edges {
		eCount += len(neighbours)
	}
	return eCount
}

func (g *Graph) VertexCount() int {
	return len(g.vertices)
}

func (g *Graph) V() []*Vertex {
	v := make([]*Vertex, 0, g.VertexCount())
	for _, it := range g.vertices {
		v = append(v, it)
	}

	return v
}

func (g *Graph) Identity() Identity {
	v := g.VertexCount()
	v2 := v * v
	if v2 == 0 {
		return Empty
	}

	ratio := float32(g.EdgeCount()) / float32(v2)
	if ratio >= 0.75 {
		return Dense
	}

	return Sparse
}

func (g *Graph) AddVertex(data interface{}) (*Vertex, bool) {
	prev, found := g.vertices[data]
	if !found {
		prev = NewVertex(data)
		g.vertices[data] = prev
	}

	return prev, !found
}

func (g *Graph) UpdateEdge(fromData, toData interface{}, cost float64) (*Edge, bool) {
	uVert, _ := g.AddVertex(fromData)
	vVert, _ := g.AddVertex(toData)

	uAdjMap := g.edges[fromData]
	if uAdjMap == nil {
		uAdjMap = make(map[interface{}]*Edge)
	}

	edge, found := uAdjMap[toData]
	if !found {
		edge = &Edge{from: uVert, to: vVert}
		uAdjMap[toData] = edge
		g.edges[fromData] = uAdjMap
	}

	// Update the cost
	edge.cost = cost

	return edge, !found
}

func (g *Graph) Edge(fromData, toData interface{}) *Edge {
	from, fromOk := g.edges[fromData]
	if !fromOk {
		return nil
	}

	return from[toData]
}

func (g *Graph) DeepCopy() *Graph {
	copy := NewGraph()
	copy.etag = g.etag

	for fromData, neighbours := range g.edges {
		for toData, edge := range neighbours {
			copy.UpdateEdge(fromData, toData, edge.cost)
		}
	}

	return copy
}

func (g *Graph) Adj(v interface{}) (adjl []*Vertex) {
	edgeList, ok := g.edges[v]
	if !ok {
		return
	}

	adjl = make([]*Vertex, 0, len(edgeList))
	for _, edge := range edgeList {
		adjl = append(adjl, edge.to)
	}

	return
}

func (g *Graph) VerticesByChan() chan interface{} {
	vChan := make(chan interface{})
	go func() {
		for rawData, _ := range g.vertices {
			vChan <- rawData
		}
	}()

	return vChan
}

func (g *Graph) Vertices() (vertices []interface{}) {
	for rawData, _ := range g.vertices {
		vertices = append(vertices, rawData)
	}

	return
}
