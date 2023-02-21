package quicksort

type Vector []int

func NewVector(capacity int) *Vector {
	v := make(Vector, capacity)
	return &v
}

func Quicksort(A *Vector, p, r int) {
	if p < r {
		q := LomutoPartition(A, p, r)
		Quicksort(A, p, q-1)
		Quicksort(A, q+1, r)
	}
}

func LomutoPartition(A *Vector, p, r int) int {
	x := (*A)[r]
	i := p - 1
	for j := p; j <= (r - 1); j++ {
		if (*A)[j] <= x {
			i += 1
			Swap(A, i, j)
		}
	}
	Swap(A, i+1, r)
	return i + 1
}

func Swap(A *Vector, i, j int) {
	(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
}

func (v *Vector) Len() int {
	return len(*v) - 1
}

func (v *Vector) Less(i, j int) bool {
	return (*v)[i] < (*v)[j]
}

func (v *Vector) Swap(i, j int) {
	Swap(v, i, j)
}
