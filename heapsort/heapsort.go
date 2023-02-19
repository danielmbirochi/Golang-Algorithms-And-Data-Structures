package heap

func Heapsort(A *Heap) {
	BuildMaxHeap(A)
	for i := A.Len() - 1; i > 0; i-- {
		A.Swap(0, i)
		A.size = A.size - 1
		MaxHeapify(A, 0)
	}
}

func BuildMaxHeap(A *Heap) {
	A.size = A.Len()
	for i := A.Len() / 2; i >= 0; i-- {
		MaxHeapify(A, i)
	}
}

func MaxHeapify(A *Heap, i int) {
	l := Left(i)
	r := Right(i)

	var highest int
	if l < A.size && A.values[l] > A.values[i] {
		highest = l
	} else {
		highest = i
	}
	if r < A.size && A.values[r] > A.values[highest] {
		highest = r
	}

	if highest != i {
		A.Swap(highest, i)
		MaxHeapify(A, highest)
	}
}
