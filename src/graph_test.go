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
