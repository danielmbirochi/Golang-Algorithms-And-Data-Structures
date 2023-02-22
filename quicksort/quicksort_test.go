package quicksort_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/quicksort"
)

func TestQuicksort(t *testing.T) {

	t.Run("Quicksort using Lomuto Parition Scheme", func(t *testing.T) {
		input := []int{2, 1, 5, 7, 3, 8, 4, 9, 0, 6}
		v := quicksort.NewVector(len(input))
		copy(*v, input)
		fmt.Printf("Input : %v\n", *v)

		quicksort.LomutoQuicksort(v, 0, len(*v)-1)
		fmt.Printf("Output : %v\n", *v)

		if sort.IsSorted(v) {
			t.Log("Slice sorted w/ Quicksort successfully!")
		}
	})

	t.Run("Quicksort using Hoare Partition Scheme", func(t *testing.T) {
		input := []int{2, 1, 5, 7, 3, 8, 4, 9, 0, 6}
		v := quicksort.NewVector(len(input))
		copy(*v, input)
		fmt.Printf("Input : %v\n", *v)

		quicksort.HoareQuicksort(v, 0, len(*v)-1)
		fmt.Printf("Output : %v\n", *v)

		if sort.IsSorted(v) {
			t.Log("Slice sorted w/ Hoare Partition Scheme !")
		}
	})

}
