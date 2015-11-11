package graph

import (
	"fmt"
	"testing"
)

func TestTopSort(t *testing.T) {
	connections := map[string][]string{
		"undershorts": []string{"pants", "shoes"},
		"pants":       []string{"belt", "shoes"},
		"socks":       []string{"shoes"},
		"watch":       []string{},
		"belt":        []string{"jacket"},
		"shirt":       []string{"belt", "tie"},
		"tie":         []string{"jacket"},
	}

	g := NewGraph()

	for key, conns := range connections {
		for _, c := range conns {
			g.UpdateEdge(key, c, 0)
		}

		{
			fmt.Printf("%v => %v\n", key, conns)
		}
	}

	listing := TopologicalSortDFS(g)
	fmt.Println("after top sort", listing)
}
