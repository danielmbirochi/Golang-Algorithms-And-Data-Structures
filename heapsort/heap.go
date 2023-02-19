package heap

type Heap struct {
	values []int
	size   int
}

func NewHeap(capacity int) *Heap {
	h := Heap{
		values: make([]int, 0, capacity),
		size:   0,
	}
	return &h
}

func (h *Heap) Init(values []int) *Heap {
	h.values = values
	h.size = len(values)
	return h
}

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*i + 2
}

func (h *Heap) Values() []int {
	return h.values
}

func (h *Heap) Len() int {
	return len(h.values)
}

func (h *Heap) Swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

// Insert appends a new value into the heap, and call IncreaseKey for
func (h *Heap) Insert(x int) {
	h.values = append(h.values, x)
	h.IncreaseKey(h.size, x)
	h.size += 1
}

func (h *Heap) Maximum() int {
	return h.values[0]
}

// ExtractMax removes the maximum value from the Heap, and perform the maintainance of
// the Max Heap property.
func (h *Heap) ExtractMax() int {
	if h.size < 1 {
		panic("heap underflow")
	}

	max := h.Maximum()

	h.values[0] = h.values[h.size-1]
	h.values = h.values[:h.Len()-1]
	h.size = h.size - 1
	MaxHeapify(h, 0)

	return max
}

// IncreaseKey assign the new value v for the element in the given index,
// and perform the maintainance of the Max Heap property.
func (h *Heap) IncreaseKey(i int, v int) {
	if v < h.values[i] {
		panic("new value is lower than current value")
	}

	h.values[i] = v
	for i > 0 && h.values[Parent(i)] < h.values[i] {
		h.Swap(i, Parent(i))
		i = Parent(i)
	}
}
