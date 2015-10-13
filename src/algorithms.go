package graph

func DFSOnFresh(g *Graph) (*Graph, []interface{}) {
	copy := g.DeepCopy()

	traversed := dfs(copy)

	return copy, traversed
}

func DFS(g *Graph) []interface{} {
	return dfs(g)
}

func dfs(g *Graph) (visitOrder []interface{}) {
	V := g.V()

	for _, u := range V {
		u.predecessor = nil
		u.color = White
	}

	t := int64(0)

	for _, u := range V {
		if u.color == White {
			dfsVisit(g, u, &t, &visitOrder)
		}
	}

	return
}

func dfsVisit(g *Graph, u *Vertex, t *int64, visitOrder *[]interface{}) {
	*t += 1

	visitInfo := VisitInfo{discoveryTime: *t}
	u.color = Gray

	adjV := g.Adj(u)

	for _, v := range adjV {
		if v.color == White {
			v.predecessor = u
			dfsVisit(g, v, t, visitOrder)
		}
	}

	*t += 1
	visitInfo.finishTime = *t
	u.meta = visitInfo
	*visitOrder = append(*visitOrder, u.Data())
}
