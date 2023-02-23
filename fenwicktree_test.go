package godsa

import "testing"

func TestFenwickBuild(t *testing.T) {
	ft := NewFenwickTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	//[1 3 3 10 5 11 7 36 9]
	t.Logf("%v", ft)

}

func TestSum(t *testing.T) {
	arr := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	ft := NewFenwickTree(arr)
	if ft.getSum(8) != 27 {
		t.Error()
	}

	if ft.getSum(9) != 34 {
		t.Error()
	}

	arr[3] += 6
	ft.Update(3, 6)

	if ft.getSum(5) != 18 {
		t.Error(ft.getSum(5))
	}
}
