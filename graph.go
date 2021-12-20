package algors

type Graph struct {
	v   int    // number of vertices
	e   int    // number of edges
	adj []*Bag // adjacency lists
}

func NewGraph(V int) *Graph {
	g := &Graph{
		v: V,
		e: 0,
	}

	g.adj = make([]*Bag, V)
	for i := 0; i < V; i++ {
		g.adj[i] = NewBag()
	}

	return g
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) AddEdge(v int, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.e++
}

func (g *Graph) Adj(v int) *Bag {
	return g.adj[v]
}
