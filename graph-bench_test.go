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

func dfsRun() {
	g := NewGraph()

	nodes := map[int][]int{
		2:  []int{0, 2, 9, 10, 12},
		8:  []int{9, 3, 1},
		1:  []int{},
		11: []int{12, 19, 8},
		-1: []int{2, 8, 1, 11, -1, 3, 9, 12, -1000},
	}

	for src, nbs := range nodes {
		for _, nb := range nbs {
			g.UpdateEdge(src, nb, -1)
		}
	}

	DFSOnFresh(g)
}

func bfsRun() {
	g := NewGraph()

	nodes := map[int][]int{
		2:  []int{0, 2, 9, 10, 12},
		8:  []int{9, 3, 1},
		1:  []int{},
		11: []int{12, 19, 8},
		-1: []int{2, 8, 1, 11, -1, 3, 9, 12, -1000},
	}

	for src, nbs := range nodes {
		for _, nb := range nbs {
			g.UpdateEdge(src, nb, -1)
		}
	}

	for src, _ := range nodes {
		BFSOnFresh(g, src)
	}
}

func BenchmarkDFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dfsRun()
	}
}

func BenchmarkBFS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bfsRun()
	}
}
