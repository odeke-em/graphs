package graph

import (
	"testing"
)

func TestInit(t *testing.T) {
	g := NewGraph()

	if g == nil {
		t.Errorf("initialized graph cannot be nil")
	}

	if g.Identity() != Empty {
		t.Errorf("expecting an empty graph")
	}
}

func TestIdentities(t *testing.T) {
	g := NewGraph()

	connections := map[string][]interface{}{
		"github":      []interface{}{},
		"los angeles": []interface{}{"hub city"},
		"facebook":    []interface{}{"suncore", "android", "edge"},
		"main":        []interface{}{"printf", "return", 0, -1, "exit", "cluster"},
	}

	type edgeKV struct {
		key  string
		edge *Edge
	}

	edgeMappings := []edgeKV{}

	originalCost := float64(-1)
	for src, neighbours := range connections {
		for _, nb := range neighbours {
			edge, inserted := g.UpdateEdge(src, nb, originalCost)
			if edge == nil {
				t.Errorf("edge cannot be nil")
			}
			if ec := edge.Cost(); ec != originalCost {
				t.Errorf("expecting edge cost of %v, got %v", originalCost, ec)
			}
			if !inserted {
				t.Errorf("src: %v neighbour: %v expected not previously encountered but in vain", src, nb)
			}

			edgeMappings = append(edgeMappings, edgeKV{key: src, edge: edge})
		}
	}

	updatedCost := float64(10)
	updatedEdgeMappings := []edgeKV{}

	for src, neighbours := range connections {
		for _, nb := range neighbours {
			edge, inserted := g.UpdateEdge(src, nb, updatedCost)
			if inserted {
				t.Errorf("src: %v neighbour: %v expected a previous insertion", src, nb)
			}
			if uec := edge.Cost(); uec != updatedCost {
				t.Errorf("expecting updated edge cost of %v, got %v", updatedCost, uec)
			}
			updatedEdgeMappings = append(updatedEdgeMappings, edgeKV{key: src, edge: edge})
		}
	}

	emLen, uemLen := len(edgeMappings), len(updatedEdgeMappings)
	if emLen != uemLen {
		t.Errorf("edgeMapping lengths do not match: original len: %v, updated len: %v", emLen, uemLen)
	}
}

func TestIdentity(t *testing.T) {
	g := NewGraph()
	if g.Identity() != Empty {
		t.Errorf("No insertion has been made into the graph yet")
	}

	edges := map[int][]int{
		2: []int{9, 10, 12},
		3: []int{2, 5, 6},
		5: []int{1, 3},
		6: []int{2, 3, 5, 9, 10, 12},
	}

	for onlySrc, _ := range edges {
		g.AddVertex(onlySrc)
	}

	if idt := g.Identity(); idt != Sparse {
		t.Errorf("expected sparseness, no edge connections yet, got %q", idt.String())
	}

	pivot := 2
	onlyPivotConnections, _ := edges[pivot]
	for _, dest := range onlyPivotConnections {
		g.UpdateEdge(pivot, dest, 0)
	}

	if idt := g.Identity(); idt != Sparse {
		t.Errorf("expected sparseness, some edge connections but not most, got %q", idt.String())
	}

	// Let's make it dense now
	srcs := []int{}
	for src, nbs := range edges {
		srcs = append(srcs, src)
		for _, nb := range nbs {
			g.UpdateEdge(src, nb, 10)
			g.UpdateEdge(nb, src, 10)
		}
	}

	vertices := g.Vertices()
	for i, n := 0, len(vertices); i < n; i++ {
		src := vertices[0]
		for j := 0; j < n; j++ {
			dest := vertices[j]
			g.UpdateEdge(src, dest, -1)
		}
	}

	for i, n := 0, len(srcs); i < n; i++ {
		for j := 0; j < n; j++ {
			otherNbs, _ := edges[srcs[j]]

			g.UpdateEdge(srcs[i], srcs[j], 20)
			g.UpdateEdge(srcs[j], srcs[i], 20)

			for _, otherNb := range otherNbs {
				g.UpdateEdge(srcs[i], otherNb, 20)
				g.UpdateEdge(otherNb, srcs[i], 20)
			}
		}
	}

	if idt := g.Identity(); idt != Dense {
		t.Errorf("expected a dense graph, got %q", idt.String())
	}
}
