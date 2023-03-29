package godsa

import (
	"fmt"
)

/*
	Comparison with Red Black Tree:
	The AVL tree and other self-balancing search trees like Red Black are useful to get all basic operations done in O(log n) time.
	The AVL trees are more balanced compared to Red-Black Trees, but they may cause more rotations during insertion and deletion.
	So if your application involves many frequent insertions and deletions, then Red Black trees should be preferred.
	And if the insertions and deletions are less frequent and search is the more frequent operation,
	then the AVL tree should be preferred over Red Black Tree.
*/

type Avlnode struct {
	V int
	l *Avlnode
	r *Avlnode
	h int
}

func (n *Avlnode) String() string {
	return fmt.Sprintf("V = %d ", n.V)
}

type AvlTree struct {
	r *Avlnode
}

func rotateRight(n *Avlnode) *Avlnode {
	nr := n.l
	n.l = nr.r
	nr.r = n

	return nr
}

func rotateLeft(n *Avlnode) *Avlnode {
	nr := n.r
	n.r = nr.l
	nr.l = n

	return nr
}

func NewAvlTree(v int) *AvlTree {
	return &AvlTree{&Avlnode{V: v}}
}

func (t *AvlTree) Insert(v int) {
	fmt.Println("insert ", v, "root ", t.r)
	cur := t.r
	stack := NewStack()
	for cur != nil {
		stack.Push(cur)

		if cur.V > v {
			cur = cur.l
		} else if cur.V < v {
			cur = cur.r
		} else {
			return
		}

	}

	cur = &Avlnode{V: v, h: stack.Len()}
	// if stack.Len() == 0 {
	// 	return
	// }

	_, pv := stack.Pop()
	p := pv.(*Avlnode)
	h := 1
	var lh, rh int
	if p.V < v {
		p.r = cur
		rh = h
	} else {
		p.l = cur
		lh = h
	}

	var gp *Avlnode
	for stack.Len() > 0 /* && math.Abs(float64(lh-rh)) < 2 */ {
		_, pv := stack.Pop()
		h++
		gp = pv.(*Avlnode)

		if p.V < gp.V {
			lh = h
		} else {
			rh = h
		}

		fmt.Println(lh-rh, p, gp)

		if lh-rh > 1 && p.V < gp.V {
			// ll
			fmt.Println("ll")
			if gp == t.r {
				t.r = rotateRight(gp)
			}

			if ok, pv := stack.Pop(); ok {
				ggp := (pv.(*Avlnode))
				if ggp.V > gp.V {
					ggp.l = rotateRight(gp)
				} else {
					ggp.r = rotateRight(gp)
				}
			}
			return
		}
		if lh-rh > 1 && p.V > gp.V {
			// lr
			fmt.Println("lr")
			gp.l = rotateLeft(gp.l)
			if gp == t.r {
				t.r = rotateRight(gp)
			}

			if ok, pv := stack.Pop(); ok {
				ggp := (pv.(*Avlnode))
				if ggp.V > gp.V {
					ggp.l = rotateRight(gp)
				} else {
					ggp.r = rotateRight(gp)
				}
			}
			return
		}

		if lh-rh < -1 && p.V > gp.V {
			// rr
			fmt.Println("rr")
			if gp == t.r {
				t.r = rotateLeft(gp)
			}

			if ok, pv := stack.Pop(); ok {
				ggp := (pv.(*Avlnode))
				if ggp.V > gp.V {
					ggp.l = rotateLeft(gp)
				} else {
					ggp.r = rotateLeft(gp)
				}
			}
			return
		}

		if lh-rh < -1 && p.V > gp.V {
			// rl
			fmt.Println("rl")
			gp.r = rotateRight(gp.r)
			if gp == t.r {
				t.r = rotateLeft(gp)
			}

			if ok, pv := stack.Pop(); ok {
				ggp := (pv.(*Avlnode))
				if ggp.V > gp.V {
					ggp.l = rotateLeft(gp)
				} else {
					ggp.r = rotateLeft(gp)
				}

			}
			return
		}

		p = gp

	}

}

func balance(stack *Stack, cur, pcur *Avlnode) {

}

func (t *AvlTree) Print() {
	printa(t.r)
	println()
}

func printa(n *Avlnode) {
	fmt.Print(n.V, ",")
	if n.l != nil {
		printa(n.l)
	}

	if n.r != nil {
		printa(n.r)
	}
}
