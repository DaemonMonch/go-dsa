package godsa

type Heap struct {
	arr []int
	i   int
}

func NewHeap() *Heap {
	return &Heap{make([]int, 5), 0}
}

func (h *Heap) Add(a int) {
	if h.i == len(h.arr) {
		o := h.arr
		h.arr = make([]int, (h.i+1)<<1)
		for idx := 0; idx < h.i; idx++ {
			h.arr[idx] = o[idx]
		}
	}

	i := h.i
	h.arr[i] = a
	for i != 0 && h.arr[(i-1)>>1] > h.arr[i] {
		tmp := h.arr[(i-1)>>1]
		h.arr[(i-1)>>1] = h.arr[i]
		h.arr[i] = tmp
		i = (i - 1) >> 1
	}
	h.i++
}

func (h *Heap) heapify() {
	i := 0
	l := h.maxChild(0)
	for i != l {
		tmp := h.arr[l]
		h.arr[l] = h.arr[i]
		h.arr[i] = tmp
		i = l
		l = h.maxChild(i)
	}
}

func (h *Heap) Pop() int {
	if h.i == 0 {
		return 0
	}
	p := h.arr[0]
	h.arr[0] = h.arr[h.i-1]
	h.i--
	h.heapify()
	return p
}

func (h *Heap) maxChild(i int) int {
	l := i
	lc := i<<1 + 1
	rc := i<<1 + 2
	if lc < h.i && h.arr[l] > h.arr[lc] {
		l = lc
	}
	if rc < h.i && h.arr[rc] < h.arr[l] {
		l = rc
	}

	return l
}
