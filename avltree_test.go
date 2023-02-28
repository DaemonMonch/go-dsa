package godsa

import "testing"

func TestAvlTree(t *testing.T) {
	tr := NewAvlTree(1)

	tr.Insert(2)
	tr.Insert(3)
	if tr.r.V != 2 {
		t.Error(tr.r.V)
	}

	tr.Insert(4)

	tr.Print()
}

func TestAvlTree2(t *testing.T) {
	tr := NewAvlTree(8)

	tr.Insert(2)
	tr.Insert(3)
	tr.Insert(4)
	tr.Print()
	tr.Insert(39)
	tr.Insert(9)
	tr.Insert(22)
	tr.Insert(6)
	tr.Insert(9)
	tr.Insert(29)
	tr.Print()
}

/* The constructed AVL Tree would be
            30
           /  \
         20   40
        /  \     \
       10  25    50

	   Preorder traversal of the constructed AVL tree is
		30 20 10 25 40 50
*/

func TestAvlTree3(t *testing.T) {
	tr := NewAvlTree(10)

	tr.Insert(20)
	tr.Insert(30)
	tr.Insert(40)
	tr.Insert(50)
	tr.Insert(25)
	tr.Print()
}
