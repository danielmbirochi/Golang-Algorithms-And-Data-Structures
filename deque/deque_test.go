package deque_test

import (
	"testing"

	"github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/deque"
	"github.com/stretchr/testify/require"
)

func TestDeque(t *testing.T) {

	t.Run("Test Dynamic Sized Deque", func(t *testing.T) {
		dsd := deque.DynamicSizedDeque{}

		dsd.EnqueueFront(1)
		dsd.EnqueueFront(2)
		dsd.EnqueueBack(3)
		dsd.EnqueueFront(4)
		dsd.EnqueueBack(5)

		require.ElementsMatch(t, dsd.Values(), []int{4, 2, 1, 3, 5}, "expected elements [4, 2, 1, 3, 5] got ", dsd.Values())

		f := dsd.DequeueFront()
		if f != 4 {
			t.Fatalf("Expected 4 got %d", f)
		}

		r := dsd.DequeueBack()
		if r != 5 {
			t.Fatalf("Expected 5 got %d", r)
		}

		require.ElementsMatch(t, dsd.Values(), []int{2, 1, 3}, "expected elements [2, 1, 3] got ", dsd.Values())

		dsd.DequeueBack()
		dsd.DequeueBack()
		dsd.DequeueBack()
		if !dsd.IsEmpty() {
			t.Fatalf("Expected empty Deque. Got %+v", dsd.Values())
		}

		underflow := dsd.DequeueFront()
		if underflow != -1 {
			t.Fatalf("Expected Deque to underflow. Got %+v", underflow)
		}
	})

}
