// A symbol table implemented with a binary search tree.
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/BST.java.html
// https://algs4.cs.princeton.edu/code/javadoc/edu/princeton/cs/algs4/BST.html
package binarySearchTree

const KeyNotExist = "Key Not Exist"

type Node struct {
	key         int
	val         int
	size        int
	left, right *Node
}

type BinaryTree interface {
	IsEmpty() bool
	Size() int
	Contains(key int) bool
	Get(key int) int
	Put(key int, val int)
	DeleteMin()
	DeleteMax()
	Delete(key int)
	Min() int
	Max() int
	Floor() int
	Ceiling() int
	Select(k int)             // Return the key in the symbol table whose rank is k
	Rank(key int)             // Return the number of keys in the symbol table strictly less than `key`
	Keys()                    // Returns all keys in the symbol table as an Iterable
	RangeKeys(lo int, hi int) // Returns all keys in the symbol table in the given range.
	RangeSize(lo int, hi int) // Returns the number of keys in the symbol table in the given range.
	LevelOrder()
}

// The struct represents an ordered symbol table of int key-value pairs.
type BST struct {
	root *Node
}

// have such a helper function to avoid visiting nil node
func (t *BST) size(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.size
	}
}

// Returns the number of key-value pairs in this symbol table.
func (t *BST) Size() int {
	return t.size(t.root)
}

// Returns the node by key
func (t *BST) get(n *Node, key int) *Node {
	if n == nil {
		return nil
	} else {
		if key == n.key {
			return n
		} else if key < n.key {
			return t.get(n.left, key)
		} else {
			return t.get(n.right, key)
		}
	}
}

// Get value by key, return 0 if not exist
func (t *BST) Get(key int) int {
	n := t.get(t.root, key)
	if n != nil {
		return n.val
	} else {
		return 0
	}
}

// Return true if the key exists in the symbol table
func (t *BST) Contains(key int) bool {
	n := t.get(t.root, key)
	return n != nil
}

func (t *BST) put(node *Node, key int, val int) *Node {
	if node == nil {
		return &Node{key, val, 1, nil, nil}
	} else {
		if key < node.key {
			node.left = t.put(node.left, key, val)
			// have such a t.size helper function to avoid visiting nil node
			node.size = 1 + t.size(node.left) + t.size(node.right)
		} else if key == node.key {
			node.key = key
			node.val = val
		} else {
			node.right = t.put(node.right, key, val)
			// have such a t.size helper function to avoid visiting nil node
			node.size = 1 + t.size(node.left) + t.size(node.right)
		}
		return node
	}
}

// Inserts the specified key-value pair into the symbol table
func (t *BST) Put(key int, val int) {
	t.root = t.put(t.root, key, val)
}

func (t *BST) deleteMin(node *Node) *Node {
	/* Handle special case (node == nil) in major DeleteMin()
	if node == nil {
		return node
	}
	*/

	if node.left != nil {
		node.left = t.deleteMin(node.left)
		// don't forget to update size ---- review data structure
		node.size = 1 + t.size(node.left) + t.size(node.right)
		return node
	}
	return node.right
}

// Removes the smallest key and associated value from the symbol table.
func (t *BST) DeleteMin() {
	if t.Size() == 0 {
		return
	}
	t.root = t.deleteMin(t.root)
}

func (t *BST) deleteMax(node *Node) *Node {
	if node.right != nil {
		node.right = t.deleteMax(node.right)
		node.size = 1 + t.size(node.left) + t.size(node.right)
		return node
	}
	return node.left
}

// Removes the largest key and associated value from the symbol table
func (t *BST) DeleteMax() {
	if t.Size() == 0 {
		return
	}
	t.root = t.deleteMax(t.root)
}

func (t *BST) findMin(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.left != nil {
		return t.findMin(node.left)
	}
	return node
}

func (t *BST) findMax(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.right != nil {
		return t.findMax(node.right)
	}
	return node
}

func (t *BST) delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	} else if node.key == key {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			var minNode *Node = t.findMin(node.right)
			node.key = minNode.key
			node.val = minNode.val
			node.right = t.deleteMin(node.right)                   // dont forget node.right =
			node.size = 1 + t.size(node.left) + t.size(node.right) // dont forget update
			return node
		}
	} else if key < node.key {
		node.left = t.delete(node.left, key)
	} else {
		node.right = t.delete(node.right, key)
	}
	node.size = 1 + t.size(node.left) + t.size(node.right) // dont forget update
	return node
}

func (t *BST) Delete(key int) {
	t.root = t.delete(t.root, key)
}

func (t *BST) Min() (key int, val int) {
	minNode := t.findMin(t.root)
	return minNode.key, minNode.val
}

func (t *BST) Max() (key int, val int) {
	maxNode := t.findMax(t.root)
	return maxNode.key, maxNode.val
}

// Returns the node with the largest key in the symbol table less than or equal to key.
func (t *BST) floor(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.key == key {
		return node
	} else if node.key < key {
		// pay attention to this part!!
		// if node.right != nil {return t.floor(node.right, key)}
		//  maybe the right tree are all ndoes > key
		r := t.floor(node.right, key)
		if r != nil {
			return r
		} else {
			return node
		}
	} else {
		return t.floor(node.left, key)
	}
}

// Returns the node with the largest key in the symbol table less than or equal to key.
func (t *BST) Floor(key int) *Node {
	return t.floor(t.root, key)
}

// Returns the smallest key in the symbol table greater than or equal to key.
func (t *BST) ceiling(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.key == key {
		return node
	} else if node.key > key {
		// pay attention to this part!!
		// if node.right != nil {return t.floor(node.right, key)}
		//  maybe the right tree are all ndoes > key
		r := t.ceiling(node.left, key)
		if r != nil {
			return r
		} else {
			return node
		}
	} else {
		return t.ceiling(node.right, key)
	}
}

// Returns the smallest key in the symbol table greater than or equal to {@code key}.
func (t *BST) Ceiling(key int) *Node {
	return t.ceiling(t.root, key)
}

func (t *BST) selectHelper(node *Node, k int) *Node {
	if node == nil {
		return node
	}
	if t.size(node.left) == k { // say k = 0, should return the smallest
		return node
	} else if t.size(node.left) < k {
		return t.selectHelper(node.right, k-1-t.size(node.left))
	} else {
		return t.selectHelper(node.left, k)
	}
}

// Return the key in the symbol table whose rank is k
/* Rank Definition:
(1) If the target is found, then the index ( = how many keys < k) is returned.
(2) If the target is not found, then the index to be inserted
of k ( =  ( = how many keys < k)) is returned. */
func (t *BST) Select(k int) *Node {
	return t.selectHelper(t.root, k)
}

func (t *BST) rank(node *Node, key int) int {
	if node == nil {
		return 0
	}
	if node.key < key {
		return t.size(node.left) + 1 + t.rank(node.right, key)
	} else if node.key == key {
		return t.size(node.left)
	} else {
		return t.rank(node.left, key)
	}
}

// Return the number of keys in the symbol table strictly less than `key`
func (t *BST) Rank(key int) int {
	return t.rank(t.root, key)
}

func (t *BST) keys(node *Node, res *[]int) {
	if node == nil {
		return
	} else {
		t.keys(node.left, res)
		*res = append(*res, node.key)
		t.keys(node.right, res)
	}
}

// Returns all keys in the symbol table as an Iterable
func (t *BST) Keys() []int {
	var res []int
	t.keys(t.root, &res)
	return res
}

func (t *BST) rangekeys(node *Node, lo int, hi int, res *[]int) {
	if node == nil {
		return
	}
	if node.key < lo {
		t.rangekeys(node.right, lo, hi, res)
	} else if node.key > hi {
		t.rangekeys(node.left, lo, hi, res)
	} else {
		t.rangekeys(node.left, lo, hi, res)
		*res = append(*res, node.key)
		t.rangekeys(node.right, lo, hi, res)
	}
}

// Returns all keys in the symbol table in the given range.
func (t *BST) RangeKeys(lo int, hi int) []int {
	var res []int
	t.rangekeys(t.root, lo, hi, &res)
	return res
}

// Returns the number of keys in the symbol table in the given range.
func (t *BST) RangeSize(lo int, hi int) int {
	return len(t.RangeKeys(lo, hi))
}

// Returns the keys in the BST in level order
func (t *BST) LevelOrder() []int {
	queue := make([]*Node, 0)
	res := make([]int, 0)
	if t.root != nil {
		queue = append(queue, t.root)
	}
	for len(queue) != 0 {
		a := queue[0]
		res = append(res, a.key)
		queue = queue[1:]
		for _, element := range []*Node{a.left, a.right} {
			if element != nil {
				queue = append(queue, element)
			}
		}
	}
	return res
}

func main() {}
