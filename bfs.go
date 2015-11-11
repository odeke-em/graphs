package graph

import "math"

func BFS(g *Graph, source interface{}) []interface{} {
	return bfs(g, source)
}

func BFSOnFresh(g *Graph, source interface{}) (*Graph, []interface{}) {
	copy := g.DeepCopy()

	traversed := bfs(copy, source)

	return copy, traversed
}

func bfs(g *Graph, source interface{}) (walkOrder []interface{}) {
	sourceVertex := g.VertexByKey(source)
	if sourceVertex == nil {
		return
	}

	vertices := g.V()

	for _, v := range vertices {
		if v != sourceVertex {
			v.color = White
			v.predecessor = nil
			v.meta = VisitInfo{depth: math.MaxInt64}
		}
	}

	sourceVertex.meta = VisitInfo{depth: 0}
	sourceVertex.predecessor = nil
	sourceVertex.color = Gray

	Q := []*Vertex{sourceVertex}
	walkOrder = append(walkOrder, sourceVertex.Data())
	for {
		l := len(Q)
		if l < 1 {
			break
		}

		// Dequeue Q
		u := Q[l-1]
		Q = Q[0 : l-1]

		adjV := g.Adj(u.Data())

		uvi, ok := u.meta.(VisitInfo)
		if !ok {
			panic("invalid attr saved in meta, not of type 'VisitInfo'")
		}

		for _, v := range adjV {
			if v.color == White {
				v.color = Gray
				v.meta = VisitInfo{depth: uvi.depth + 1}
				v.predecessor = u

				// Enqueue v
				Q = append([]*Vertex{v}, Q...)
				walkOrder = append(walkOrder, v.Data())
			}
		}

		u.color = Black
	}
	return
}
