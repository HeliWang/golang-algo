// A symbol table implemented with a binary search tree.
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BST.java.html
// https://algs4.cs.princeton.edu/code/javadoc/edu/princeton/cs/algs4/BST.html
package binarySearchTree

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	var tree *BST = new(BST)
	if tree.Size() != 0 {
		t.Fail()
	}
	tree.Put(1, 3)
	tree.Put(1, 4)
	tree.Put(2, 3)
	tree.Put(4, 5)
	tree.Put(10, 1)

	if tree.Size() != 4 {
		t.Error("Wrong Size")
	}

	if !tree.Contains(1) {
		t.Fail()
	}

	if tree.Get(1) != 4 {
		t.Fail()
	}

	if tree.Get(0) != 0 {
		t.Fail()
	}

	minkey, minval := tree.Min()
	if minkey != 1 {
		t.Fail()
	}
	if minval != 4 {
		t.Fail()
	}

	tree.DeleteMin()

	if tree.Size() != 3 {
		t.Error()
	}

	minkey, minval = tree.Min()
	if minkey != 2 {
		t.Fail()
	}
	if minval != 3 {
		t.Fail()
	}

	if tree.Contains(1) {
		t.Fail()
	}

	var n1 *Node = tree.Floor(3)
	if n1 == nil || n1.key != 2 {
		t.Fail()
	}

	var n2 *Node = tree.Ceiling(1)
	if n2 == nil || n2.key != 2 {
		t.Fail()
	}

	var n3 *Node = tree.Floor(2)
	if n3 == nil || n3.key != 2 {
		t.Fail()
	}

	var n4 *Node = tree.Ceiling(2)
	if n4 == nil || n4.key != 2 {
		t.Fail()
	}

	tree.DeleteMin()

	minkey, minval = tree.Min()
	if minkey != 4 {
		t.Fail()
	}
	if minval != 5 {
		t.Fail()
	}

	tree.DeleteMax()

	minkey, minval = tree.Min()
	if minkey != 4 {
		t.Fail()
	}
	if minval != 5 {
		t.Fail()
	}

	maxkey, maxval := tree.Max()
	if maxkey != 4 {
		t.Fail()
	}
	if maxval != 5 {
		t.Fail()
	}

	tree.Delete(4)
	if tree.Size() != 0 {
		t.Error()
	}
}

func Test2(t *testing.T) {
	var tree *BST = new(BST)
	if tree.Size() != 0 {
		t.Fail()
	}
	tree.Put(1, 3)
	tree.Put(1, 4)
	tree.Put(2, 3)
	tree.Put(4, 5)
	tree.Put(10, 1)
	tree.Put(11, 1)
	tree.Put(13, 1)
	tree.Put(12, 1)
	tree.Put(22, 1)
	tree.Put(-10, 1)
	tree.Put(-20, 1)
	tree.Put(-1, 1)
	tree.Delete(13)
	if tree.Size() != 10 {
		t.Error("Wrong Tree Size: " + string(tree.Size()) + "!")
	}
	tree.Delete(4)
	if tree.Size() != 9 {
		t.Error("Wrong Tree Size: " + string(tree.Size()) + "!")
	}
	tree.Delete(-10)
	if tree.Size() != 8 {
		t.Error("Wrong Tree Size: " + string(tree.Size()) + "!")
	}
	tree.Delete(1)
	if tree.Size() != 7 {
		t.Error("Wrong Tree Size: " + string(tree.Size()) + "!")
	}
	tree.DeleteMin()
	if tree.Size() != 6 {
		t.Error("Wrong Tree Size: " + string(tree.Size()) + "!")
	}
	minkey, minval := tree.Min()
	if minkey != -1 {
		t.Fail()
	}
	if minval != 1 {
		t.Fail()
	}

}

func Test3(t *testing.T) {
	var tree *BST = new(BST)
	if tree.Size() != 0 {
		t.Fail()
	}
	tree.Put(4, 3)
	tree.Put(2, 4)
	tree.Put(1, 3)
	tree.Put(3, 5)
	if tree.Select(0).key != 1 {
		t.Error("Select Wrong" + strconv.Itoa(tree.Rank(0)))
	}
	if tree.Select(1).key != 2 {
		t.Error("Select Wrong" + strconv.Itoa(tree.Rank(0)))
	}
	if tree.Select(2).key != 3 {
		t.Error("Select Wrong" + strconv.Itoa(tree.Rank(0)))
	}
	if tree.Select(3).key != 4 {
		t.Error("Select Wrong" + strconv.Itoa(tree.Rank(0)))
	}

	if tree.Rank(0) != 0 {
		t.Error("Rank Wrong" + strconv.Itoa(tree.Rank(0)))
	}

	if tree.Rank(1) != 0 {
		t.Error("Rank Wrong" + strconv.Itoa(tree.Rank(0)))
	}

	if tree.Rank(2) != 1 {
		t.Error("Rank Wrong" + strconv.Itoa(tree.Rank(0)))
	}

	if tree.Rank(6) != 4 {
		t.Error("Rank Wrong" + strconv.Itoa(tree.Rank(0)))
	}
}

func Test4(t *testing.T) {
	var tree *BST = new(BST)
	if tree.Size() != 0 {
		t.Fail()
	}
	tree.Put(4, 3)
	tree.Put(2, 4)
	tree.Put(1, 3)
	tree.Put(3, 5)
	var ltrRes []int = tree.Keys()
	if !reflect.DeepEqual(ltrRes, []int{1, 2, 3, 4}) {
		t.Error("Put Wrong")
	}
}

func Test5(t *testing.T) {
	var tree *BST = new(BST)
	if tree.Size() != 0 {
		t.Fail()
	}
	tree.Put(4, 3)
	tree.Put(2, 4)
	tree.Put(1, 3)
	tree.Put(3, 5)
	var ltrRes []int = tree.RangeKeys(2, 3)

	if !reflect.DeepEqual(ltrRes, []int{2, 3}) || tree.RangeSize(2, 3) != 2 {
		t.Error("Range Wrong")
	}
}

func Test6(t *testing.T) {
	var tree *BST = new(BST)
	if tree.Size() != 0 {
		t.Fail()
	}
	tree.Put(4, 3)
	tree.Put(2, 4)
	tree.Put(1, 3)
	tree.Put(3, 5)
	var ltrRes []int = tree.LevelOrder()
	fmt.Printf("Level Order:")
	for _, v := range ltrRes {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

/* An example of using the errors package
func (t* BST) get(n *Node, key int) (val int, err error) {
	if n == nil {
		err = errors.New(KeyNotExist)
	} else {
		if key == n.key {
			return n.val, nil
		} else if key < n.key {
			val1, err1 := t.get(n.left, key)
			if err1 != nil {
				return val1, nil
			}
		} else {
			val2, err2 := t.get(n.right, key)
			if err2 != nil {
				return val2, nil
			}
		}
		err = errors.New(KeyNotExist)
	}
	return
}
*/
