package godsa

import "testing"

func TestAdd(t *testing.T) {
	st := newNode()
	intervals := []int{3, 5, 2, 6, 2, 1, 2, 4, 6, 3, 1, 5, 7, 5}
	st.add(0, len(intervals)-1, intervals)
	print(st)
	// t.Log(st.find(0, 5), st.s)
}

func TestFind(t *testing.T) {
	st := newNode()
	intervals := []int{3, 5, 2, 6, 2, 1, 2, 4, 6, 3, 1, 5, 7, 5}
	st.add(0, len(intervals)-1, intervals)

	t.Log(st.querySum(3, 3))
}
