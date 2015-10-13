package main

import (
	"fmt"

	"github.com/odeke-em/graphs/src"
)

func main() {
	connections := map[int][]int{
		7:  []int{11, 8},
		5:  []int{11},
		3:  []int{8, 10},
		11: []int{2, 9, 10},
		8:  []int{9},
	}

	g := graph.NewGraph()
	for from, toNodes := range connections {
		for _, to := range toNodes {
			g.UpdateEdge(from, to, 0)
		}
	}

	_, visitOrder := graph.DFSOnFresh(g)

	for i, it := range visitOrder {
		fmt.Printf("%d: %v\n", i, it)
	}
}
