package graph

func CycleDFS(g *Graph) (traces [][]interface{}) {
	V := g.V()

	for _, u := range V {
		u.predecessor = nil
		u.color = White
	}

	cyclicEdges := []*Edge{}

	for _, u := range V {
		if u.color == White { // Note yet discovered
			cyclicEdges = append(cyclicEdges, cycleDFSVisit(g, u)...)
		}
	}

	return unravelCycle(cyclicEdges)
}

func unravelCycle(el []*Edge) (trace [][]interface{}) {
	for _, e := range el {
		vx := e.from
		t1 := []interface{}{}

		for vx != nil {
			t1 = append(t1, vx.Data())
			vx = vx.predecessor
		}

		trace = append(trace, t1)
	}

	return
}

func cycleDFSVisit(g *Graph, src *Vertex) (cyclicEdges []*Edge) {

	src.color = Gray

	adjV := g.Adj(src.Data())

	for _, v := range adjV {
		if v.color == White {
			v.predecessor = src
			cyclicEdges = append(cyclicEdges, cycleDFSVisit(g, v)...)
		} else if v.color == Gray {
			cyclicEdges = append(cyclicEdges, &Edge{from: src, to: v})
		}
	}

	src.color = Black

	return
}
