package graph

import (
	"fmt"
	"testing"
)

func TestDFSInit(t *testing.T) {
	g := NewGraph()

	nodes := map[int][]int{
		2:  []int{0, 2, 9, 10, 12},
		8:  []int{9, 3, 1},
		1:  []int{},
		11: []int{12, 19, 8},
	}

	for src, nbs := range nodes {
		for _, nb := range nbs {
			g.UpdateEdge(src, nb, -1)
		}
	}

	dfsdG, walkOrder := DFSOnFresh(g)
	if dfsdG == g {
		t.Errorf("a fresh copy should have been made before DFS operation")
	}

	fmt.Println("walkOrder", walkOrder)
}
