package quicksort_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/quicksort"
)

func TestQuicksort(t *testing.T) {

	t.Run("Quicksort using Lomuto Parition Scheme", func(t *testing.T) {
		s := []int{2, 1, 5, 7, 3, 8, 4, 9, 0, 6}
		fmt.Printf("Input : %v\n", s)

		quicksort.LomutoQuicksort(s, 0, len(s)-1)
		fmt.Printf("Output : %v\n", s)

		if sort.IntsAreSorted(s) {
			t.Log("Slice sorted w/ Quicksort successfully!")
		}
	})

	t.Run("Quicksort using Hoare Partition Scheme", func(t *testing.T) {
		s := []int{2, 1, 5, 7, 3, 8, 4, 9, 0, 6}
		fmt.Printf("Input : %v\n", s)

		quicksort.HoareQuicksort(s, 0, len(s)-1)
		fmt.Printf("Output : %v\n", s)

		if sort.IntsAreSorted(s) {
			t.Log("Slice sorted w/ Hoare Partition Scheme !")
		}
	})

}
