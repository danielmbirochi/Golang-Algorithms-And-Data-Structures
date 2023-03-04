package unionfind

type WeightedQuickUnion struct {
	parent []int
	size   []int // keeps the number of objects in the tree rooted at i.
	count  int
}

func NewWeightedQuickUnion(n int) WeightedQuickUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = i
	}
	return WeightedQuickUnion{
		parent: parent,
		size:   size,
		count:  n,
	}
}

// Count returns the number of sets - connected components.
func (w *WeightedQuickUnion) Count() int {
	return w.count
}

func (w *WeightedQuickUnion) Validate(p int) {
	if p < 0 || p >= len(w.parent) {
		panic("invalid index")
	}
}

// Root iterate each parent node until it reaches the
// root node - where the node points to itself.
func (qu *WeightedQuickUnion) Root(p int) int {
	qu.Validate(p)
	for p != qu.parent[p] {
		p = qu.parent[p]
	}
	return p
}

func (w *WeightedQuickUnion) Connected(p, q int) bool {
	return w.Root(q) == w.Root(p)
}

func (w *WeightedQuickUnion) Union(p, q int) {
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
