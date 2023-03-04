package unionfind

type WeightedQuickUnionPathCompression struct {
	parent []int
	size   []int
	count  int
}

func NewWeightedQuickUnionPathCompression(n int) WeightedQuickUnionPathCompression {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = i
	}
	return WeightedQuickUnionPathCompression{
		parent: parent,
		size:   size,
		count:  n,
	}
}

// Count returns the number of sets - connected components.
func (w *WeightedQuickUnionPathCompression) Count() int {
	return w.count
}

func (w *WeightedQuickUnionPathCompression) Validate(p int) {
	if p < 0 || p >= len(w.parent) {
		panic("invalid index")
	}
}

// Root iterate each parent node until it reaches the
// root node - where the node points to itself. After we
// find the root node, we move all the touched nodes in the path
// to point to the root node - that way we compress the path
// of each touched node to the root.
func (w *WeightedQuickUnionPathCompression) Root(p int) int {
	w.Validate(p)
	root := p
	for root != w.parent[root] {
		root = w.parent[root]
	}
	for p != root {
		newp := w.parent[p]
		w.parent[p] = root
		p = newp
	}
	return root
}

func (w *WeightedQuickUnionPathCompression) Connected(p, q int) bool {
	return w.Root(q) == w.Root(p)
}

func (w *WeightedQuickUnionPathCompression) Union(p, q int) {
	i := w.Root(p)
	j := w.Root(q)
	if i == j {
		return
	}
	if w.size[i] < w.size[j] {
		w.parent[i] = j
		w.size[j] += w.size[i]
	} else {
		w.parent[j] = i
		w.size[i] += w.size[j]
	}
	w.count--
}
