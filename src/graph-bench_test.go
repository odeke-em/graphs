package graph

import (
	"testing"
)

func BenchmarkNewGraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGraph()
	}
}

func BenchmarkSingleVertexInsertion(b *testing.B) {
	g := NewGraph()
	for i := 0; i < b.N; i++ {
		g.AddVertex(g)
	}
}

func BenchmarkMultiVertexInsertion(b *testing.B) {
	g := NewGraph()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			g.AddVertex(j)
		}
	}
}
