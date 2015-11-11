package graph

import (
	"fmt"
	"testing"
)

func TestNoCycles(t *testing.T) {
	g := NewGraph()

	linking := map[int][]int{
		1: []int{2},
		2: []int{3},
	}

	for from, values := range linking {
		for _, to := range values {
			g.UpdateEdge(from, to, 0)
		}
	}

	cyclicEdges := CycleDFS(g)
	if len(cyclicEdges) > 0 {
		t.Errorf("g has no cycles, yet got %v", cyclicEdges)
	}
}

func TestWithCycles(t *testing.T) {
	g := NewGraph()

	linking := map[int][]int{
		1: []int{2, 3},
		2: []int{3},
		3: []int{4, 5, 1},
		4: []int{1},
		8: []int{5, 1},
	}

	for from, values := range linking {
		for _, to := range values {
			g.UpdateEdge(from, to, 0)
		}
	}

	cyclicEdges := CycleDFS(g)
	fmt.Println(cyclicEdges)
	if len(cyclicEdges) < 1 {
		t.Errorf("g has cycles, yet got %v of len(%v)", cyclicEdges, len(cyclicEdges))
	}
}
