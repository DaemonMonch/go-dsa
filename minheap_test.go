package godsa

import (
	"sort"
	"testing"
)

func TestHeapAdd(t *testing.T) {
	h := NewHeap()
	h.Add(5)
	h.Add(10)
	h.Add(3)
	h.Add(2)
	h.Add(8)
	h.Add(1)
	h.Add(13)
	h.Add(17)
	t.Log(h)
	h.Pop()
	t.Log(h)
	h.Pop()
	t.Log(h)
	h.Add(1)
	t.Log(h)
	h.Add(4)
	t.Log(h)
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func TestHeapSort(t *testing.T) {
	a := []int{6, 4, 33, 111, 2, 6, 2, 5, 78, 8, 1, -1}
	h := NewHeap()
	for i := 0; i < len(a); i++ {
		h.Add(a[i])

	}

	for i := 0; i < len(a); i++ {
		a[i] = h.Pop()

	}

	if !sort.IsSorted(SortBy(a)) {
		t.Error(a)
	}

}

func TestHeapPop(t *testing.T) {
	a := []int{6}
	h := NewHeap()
	for i := 0; i < len(a); i++ {
		h.Add(a[i])

	}
	if h.Pop() != 6 {
		t.Error()
	}

	if h.Pop() != 0 {
		t.Error()
	}
}
