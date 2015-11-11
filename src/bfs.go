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
			v.meta = int64(math.MaxInt64)
		}
	}

	sourceVertex.meta = 0
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

		for _, v := range adjV {
			if v.color == White {
				v.color = Gray
				v.meta = plusOneInt64(u.meta)
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

func plusOneInt64(v interface{}) int64 {
	var value int64 = math.MaxInt64
	if cast, ok := v.(int64); ok {
		value = cast + 1
	}
	return value
}
