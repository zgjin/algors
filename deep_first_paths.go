package algors

type DeepFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDeepFirstPaths(g *Graph, s int) *DeepFirstPaths {
	dfp := &DeepFirstPaths{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}

	dfp.dps(g, s)

	return dfp
}

func (dfp *DeepFirstPaths) dps(g *Graph, v int) {
	dfp.marked[v] = true
	adfs := g.Adj(v)
	for e := adfs.list.Front(); e != nil; e = e.Next() {
		if !dfp.marked[e.Value.(int)] {
			dfp.edgeTo[e.Value.(int)] = v
			dfp.dps(g, e.Value.(int))
		}
	}
}

func (dfp *DeepFirstPaths) HasPathTo(v int) bool {
	return dfp.marked[v]
}

func (dfp *DeepFirstPaths) PathTo(v int) *Bag {
	if !dfp.HasPathTo(v) {
		return nil
	}

	b := NewBag()
	for x := v; x != dfp.s; x = dfp.edgeTo[x] {
		b.Add(x)
	}
	b.Add(dfp.s)

	return b
}
