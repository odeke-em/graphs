package graph

func noopOnVertex(v interface{}) {
	return
}

func DFSOnFreshWithOnFinish(g *Graph, fn func(interface{})) (*Graph, []interface{}) {
	return dfsOnFresh(g, fn)
}

func DFSOnFresh(g *Graph) (*Graph, []interface{}) {
	return dfsOnFresh(g, noopOnVertex)
}

func dfsOnFresh(g *Graph, fn func(interface{})) (*Graph, []interface{}) {
	copy := g.DeepCopy()

	traversed := dfs(copy, fn)

	return copy, traversed
}

func DFS(g *Graph) []interface{} {
	return dfs(g, noopOnVertex)
}

func DFSWithOnFinish(g *Graph, onEachVertexFinish func(interface{})) []interface{} {
	return dfs(g, onEachVertexFinish)
}

func dfs(g *Graph, onEachVertexFinish func(interface{})) (visitOrder []interface{}) {
	V := g.V()

	for _, u := range V {
		u.predecessor = nil
		u.color = White
	}

	t := int64(0)

	for _, u := range V {
		if u.color == White {
			dfsVisit(g, u, &t, &visitOrder, onEachVertexFinish)
		}
	}

	return
}

func dfsVisit(g *Graph, u *Vertex, t *int64, visitOrder *[]interface{}, onFinish func(interface{})) {
	*t += 1

	u.color = Gray
	visitInfo := VisitInfo{discoveryTime: *t}

	adjV := g.Adj(u.Data())

	for _, v := range adjV {
		if v.color == White {
			v.predecessor = u
			dfsVisit(g, v, t, visitOrder, onFinish)
		}
	}

	*t += 1

	curVertexData := u.Data()
	*visitOrder = append(*visitOrder, curVertexData)
	visitInfo.finishTime = *t

	u.meta = visitInfo
	u.color = Black

	onFinish(curVertexData)
}
