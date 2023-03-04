package unionfind

// QuickUnion implements a lazy solution for the Union Find problem
// it performs Union operation just changing one element in the tree array
// but it won't scale for tall trees.
type QuickUnion struct {
	parent []int
	count  int
}

func NewQuickUnion(n int) QuickUnion {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return QuickUnion{
		parent: parent,
		count:  n,
	}
}

// Count returns the number of sets - connected components.
func (qu *QuickUnion) Count() int {
	return qu.count
}

// Root iterate each parent node until it reaches the
// root node - where the node points to itself.
func (qu *QuickUnion) Root(p int) int {
	qu.Validate(p)
	for p != qu.parent[p] {
		p = qu.parent[p]
	}
	return p
}

func (qu *QuickUnion) Validate(p int) {
	if p < 0 || p >= len(qu.parent) {
		panic("invalid index")
	}
}

func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.Root(q) == qu.Root(p)
}

func (qu *QuickUnion) Union(p, q int) {
	i := qu.Root(p)
	j := qu.Root(q)
	if i == j {
		return
	}
	qu.parent[i] = j
	qu.count -= 1
}
