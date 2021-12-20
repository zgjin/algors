package algors

type DeepFirstSearch struct {
	marked []bool
	count  int
}

func NewDeepFirstSearch(graph *Graph, s int) *DeepFirstSearch {
	marked := make([]bool, graph.V())
	dfs := &DeepFirstSearch{
		marked: marked,
		count:  0,
	}
	dfs.dfs(graph, s)

	return dfs
}

func (dfs *DeepFirstSearch) dfs(graph *Graph, v int) {
	dfs.marked[v] = true
	dfs.count++
	adfs := graph.Adj(v)
	for e := adfs.list.Front(); e != nil; e = e.Next() {
		if !dfs.marked[e.Value.(int)] {
			dfs.dfs(graph, e.Value.(int))
		}
	}
}
