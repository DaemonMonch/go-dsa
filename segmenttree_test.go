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

	t.Log(st.querySum(0, 8))
}

func TestUpdate(t *testing.T) {
	st := newNode()
	intervals := []int{3, 5, 2, 6, 2, 1, 2, 4, 6, 3, 1, 5, 7, 5}
	st.add(0, len(intervals)-1, intervals)
	st.update(0, 2)

	x := st.querySum(0, 1)
	t.Log(x)
	if x != 10 {
		t.Error("expect ", 10, "actual ", x)
	}

	st.update(3, 8)

	x = st.querySum(0, 1)
	t.Log(x)
	if x != 10 {
		t.Error("expect ", 10, "actual ", x)
	}

	x = st.querySum(0, 4)
	t.Log(x)
	if x != 28 {
		t.Error("expect ", 28, "actual ", x)
	}
}
