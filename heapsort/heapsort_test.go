package heap_test

import (
	"reflect"
	"testing"

	heap "github.com/danielmbirochi/Golang-Algorithms-And-Data-Structures/heapsort"
)

func TestHeapsort(t *testing.T) {

	t.Run("TestBuildMaxHeap", func(t *testing.T) {
		input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
		h := heap.NewHeap(10).Init(input)

		heap.BuildMaxHeap(h)

		if !reflect.DeepEqual(h.Values(), []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}) {
			t.Fatal("Erro ao executar BuildMaxHeap")
		}
		t.Log("BuildMaxHeap executado com sucesso!")
	})

	t.Run("TestHeapsort", func(t *testing.T) {
		input := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
		h := heap.NewHeap(10).Init(input)

		heap.Heapsort(h)

		if !reflect.DeepEqual(h.Values(), []int{1, 2, 3, 4, 7, 8, 9, 10, 14, 16}) {
			t.Fatal("Error sorting Heap")
		}
		t.Log("Heap sorted successfully!")
	})

	t.Run("TestExtractMax", func(t *testing.T) {
		h := heap.NewHeap(10)
		input := []int{16, 14, 10, 9, 8, 7, 4, 3, 2, 1}
		for _, v := range input {
			h.Insert(v)
		}

		max := h.ExtractMax()

		if max != 16 {
			t.Fatal("Wrong max value!")
		}
		t.Logf("Heap max value extracted successfully : %d", max)
	})

	t.Run("TestIncreaseKey", func(t *testing.T) {
		h := heap.NewHeap(10)
		input := []int{16, 14, 10, 9, 8, 7, 4, 3, 2, 1}
		for _, v := range input {
			h.Insert(v)
		}

		h.IncreaseKey(3, 18)

		if h.Maximum() != 18 {
			t.Fatal("Wrong max value!")
		}
		t.Logf("Heap increased key value successfully : %d", h.Values())
	})

	t.Run("TestHeapInsert", func(t *testing.T) {
		h := heap.NewHeap(10)
		input := []int{16, 14, 10, 9, 8, 7, 4, 3, 2, 1}
		for _, v := range input {
			h.Insert(v)
		}

		h.Insert(12)
		t.Logf("Heap increased key value successfully : %d", h.Values())
	})

}
