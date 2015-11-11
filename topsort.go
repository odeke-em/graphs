package graph

func TopologicalSortDFS(g *Graph) []interface{} {
	// CLR Page 613 Chapter 22.4
	// Call DFS(G) to compute finishing times v.f for each vertex v
	// as each vertex is finished, insert it onto the front of a linked list
	// return the linked list of vertices

	dataListing := []interface{}{}
	onEachVertexFinish := func(d interface{}) {
		dataListing = append([]interface{}{d}, dataListing...)
	}

	DFSOnFreshWithOnFinish(g, onEachVertexFinish)

	return dataListing
}
