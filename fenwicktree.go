package godsa

type fenwickTree []int

func NewFenwickTree(arr []int) fenwickTree {
	l := len(arr)
	ft := fenwickTree(make([]int, l+1))
	for i := 0; i < len(arr); i++ {
		ft.Update(i, arr[i])
	}

	return ft
}

func (ft fenwickTree) Update(i, v int) {
	ftIdx := i + 1

	for ftIdx < len(ft) {
		ft[ftIdx] += v
		ftIdx += ftIdx & (-ftIdx)
	}
}

func (ft fenwickTree) getSum(i int) int {
	sum := 0
	i += 1
	for i > 0 {

		sum += ft[i]
		i -= i & (-i)
	}
	return sum
}
