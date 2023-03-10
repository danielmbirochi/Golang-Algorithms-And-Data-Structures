package quicksort

func LomutoQuicksort(A []int, p, r int) {
	if p < r {
		q := LomutoPartition(A, p, r)
		LomutoQuicksort(A, p, q-1)
		LomutoQuicksort(A, q+1, r)
	}
}

func LomutoPartition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j <= (r - 1); j++ {
		if A[j] <= x {
			i += 1
			Swap(A, i, j)
		}
	}
	Swap(A, i+1, r)
	return i + 1
}

// This function takes first element as pivot, and places
// all the elements smaller than the pivot on the left side
// and all the elements greater than the pivot on
// the right side. It returns the index of the last element
// on the smaller side
func HoarePartition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low-1, high+1
	for {
		// Find leftmost element greater than
		// or equal to pivot
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		// Find rightmost element smaller than
		// or equal to pivot
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		// If two pointers met.
		if i >= j {
			return j
		}

		Swap(arr, i, j)
	}
}

func HoareQuicksort(A []int, p, r int) {
	if p < r {
		q := HoarePartition(A, p, r)
		HoareQuicksort(A, p, q)
		HoareQuicksort(A, q+1, r)
	}
}

func Swap(A []int, i, j int) {
	A[i], A[j] = A[j], A[i]
}
