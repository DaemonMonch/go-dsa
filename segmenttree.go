package godsa

import (
	"container/list"
	"fmt"
)

func defaultMerge(a []int, b []int) interface{} {
	var ab []int
	ab = append(ab, a...)
	ab = append(ab, b...)
	return ab
}

type node struct {
	s, e  int
	ele   int
	left  *node
	right *node
}

func newNode() *node {
	return &node{}
}

func (n *node) String() string {
	return fmt.Sprintf("s %d e %d ele %v", n.s, n.e, n.ele)
}

func (n *node) add(s, e int, arr []int) int {

	if s >= e {
		n.ele = arr[e]
		n.s = s
		n.e = e
		return n.ele
	}
	n.s = s
	n.e = e
	mid := (s + e) / 2
	n.left = newNode()
	n.right = newNode()
	n.ele = n.left.add(s, mid, arr) + n.right.add(mid+1, e, arr)
	return n.ele
}

func (n *node) querySum(s, e int) int {
	return n.find(n.s, n.e, s, e)
}

func (n *node) find(ns, ne, s, e int) int {
	fmt.Println(ns, ne, "  ", s, e)
	if s <= ns && e >= ne {
		return n.ele
	}

	// if n.left == nil {
	// 	return n.ele
	// }

	if s > ne || e < ns {
		return 0
	}

	mid := (ns + ne) / 2
	// if mid > e {
	// 	return n.left.find(s, e)
	// } else if mid <= s {
	// 	return n.right.find(s, e)
	// } else {
	// 	return n.left.find(s, mid) + n.right.find(mid, e)
	// }
	return n.left.find(ns, mid, s, e) + n.right.find(mid+1, ne, s, e)
}

func print(root *node) {
	fmt.Println(root)
	queue := list.New()
	queue.PushBack(root.left)
	queue.PushBack(root.right)

	for {
		st := queue.Front()
		if st == nil {
			return
		}
		queue.Remove(st)
		n := st.Value.(*node)
		fmt.Println(n)
		if n.left != nil {
			queue.PushBack(n.left)
		}
		if n.right != nil {
			queue.PushBack(n.right)
		}
	}
}
