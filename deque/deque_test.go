package deque_test

import (
	"testing"

	"github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/deque"
	"github.com/stretchr/testify/require"
)

func TestDeque(t *testing.T) {

	t.Run("Test Dynamic Sized Deque", func(t *testing.T) {
		dsd := deque.DynamicSizedDeque{}

		deque.EnqueueFront(&dsd, 1)
		deque.EnqueueFront(&dsd, 2)
		deque.EnqueueRear(&dsd, 3)
		deque.EnqueueFront(&dsd, 4)
		deque.EnqueueRear(&dsd, 5)

		require.ElementsMatch(t, dsd.Values(), []int{4, 2, 1, 3, 5}, "expected elements [4, 2, 1, 3, 5] got ", dsd.Values())

		f := deque.DequeueFront(&dsd)
		if f != 4 {
			t.Fatalf("Expected 4 got %d", f)
		}

		r := deque.DequeueRear(&dsd)
		if r != 5 {
			t.Fatalf("Expected 5 got %d", r)
		}

		require.ElementsMatch(t, dsd.Values(), []int{2, 1, 3}, "expected elements [2, 1, 3] got ", dsd.Values())

		deque.DequeueRear(&dsd)
		deque.DequeueRear(&dsd)
		deque.DequeueRear(&dsd)
		if !deque.IsEmpty(&dsd) {
			t.Fatalf("Expected empty Deque. Got %+v", dsd.Values())
		}

		underflow := deque.DequeueFront(&dsd)
		if underflow != -1 {
			t.Fatalf("Expected Deque to underflow. Got %+v", underflow)
		}
	})

}
