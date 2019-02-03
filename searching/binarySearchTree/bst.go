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
func (self *BST) size(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.size
	}
}

// Returns the number of key-value pairs in this symbol table.
func (self *BST) Size() int {
	return self.size(self.root)
}

// Returns the node by key
func (self *BST) get(n *Node, key int) *Node {
	if n == nil {
		return nil
	} else {
		if key == n.key {
			return n
		} else if key < n.key {
			return self.get(n.left, key)
		} else {
			return self.get(n.right, key)
		}
	}
}

// Get value by key, return 0 if not exist
func (self *BST) Get(key int) int {
	n := self.get(self.root, key)
	if n != nil {
		return n.val
	} else {
		return 0
	}
}

// Return true if the key exists in the symbol table
func (self *BST) Contains(key int) bool {
	n := self.get(self.root, key)
	return n != nil
}

func (self *BST) put(node *Node, key int, val int) *Node {
	if node == nil {
		return &Node{key, val, 1, nil, nil}
	} else {
		if key < node.key {
			node.left = self.put(node.left, key, val)
			// have such a self.size helper function to avoid visiting nil node
			node.size = 1 + self.size(node.left) + self.size(node.right)
		} else if key == node.key {
			node.key = key
			node.val = val
		} else {
			node.right = self.put(node.right, key, val)
			// have such a self.size helper function to avoid visiting nil node
			node.size = 1 + self.size(node.left) + self.size(node.right)
		}
		return node
	}
}

// Inserts the specified key-value pair into the symbol table
func (self *BST) Put(key int, val int) {
	self.root = self.put(self.root, key, val)
}

func (self *BST) deleteMin(node *Node) *Node {
	/* Handle special case (node == nil) in major DeleteMin()
	if node == nil {
		return node
	}
	*/

	if node.left != nil {
		node.left = self.deleteMin(node.left)
		// don't forget to update size ---- review data structure
		node.size = 1 + self.size(node.left) + self.size(node.right)
		return node
	}
	return node.right
}

// Removes the smallest key and associated value from the symbol table.
func (self *BST) DeleteMin() {
	if self.Size() == 0 {
		return
	}
	self.root = self.deleteMin(self.root)
}

func (self *BST) deleteMax(node *Node) *Node {
	if node.right != nil {
		node.right = self.deleteMax(node.right)
		node.size = 1 + self.size(node.left) + self.size(node.right)
		return node
	}
	return node.left
}

// Removes the largest key and associated value from the symbol table
func (self *BST) DeleteMax() {
	if self.Size() == 0 {
		return
	}
	self.root = self.deleteMax(self.root)
}

func (self *BST) findMin(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.left != nil {
		return self.findMin(node.left)
	}
	return node
}

func (self *BST) findMax(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.right != nil {
		return self.findMax(node.right)
	}
	return node
}

func (self *BST) delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	} else if node.key == key {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			var minNode *Node = self.findMin(node.right)
			node.key = minNode.key
			node.val = minNode.val
			node.right = self.deleteMin(node.right)                      // dont forget node.right =
			node.size = 1 + self.size(node.left) + self.size(node.right) // dont forget update
			return node
		}
	} else if key < node.key {
		node.left = self.delete(node.left, key)
	} else {
		node.right = self.delete(node.right, key)
	}
	node.size = 1 + self.size(node.left) + self.size(node.right) // dont forget update
	return node
}

func (self *BST) Delete(key int) {
	self.root = self.delete(self.root, key)
}

func (self *BST) Min() (key int, val int) {
	minNode := self.findMin(self.root)
	return minNode.key, minNode.val
}

func (self *BST) Max() (key int, val int) {
	maxNode := self.findMax(self.root)
	return maxNode.key, maxNode.val
}

// Returns the node with the largest key in the symbol table less than or equal to key.
func (self *BST) floor(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.key == key {
		return node
	} else if node.key < key {
		// pay attention to this part!!
		// if node.right != nil {return self.floor(node.right, key)}
		//  maybe the right tree are all ndoes > key
		r := self.floor(node.right, key)
		if r != nil {
			return r
		} else {
			return node
		}
	} else {
		return self.floor(node.left, key)
	}
}

// Returns the node with the largest key in the symbol table less than or equal to key.
func (self *BST) Floor(key int) *Node {
	return self.floor(self.root, key)
}

// Returns the smallest key in the symbol table greater than or equal to key.
func (self *BST) ceiling(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.key == key {
		return node
	} else if node.key > key {
		// pay attention to this part!!
		// if node.right != nil {return self.floor(node.right, key)}
		//  maybe the right tree are all ndoes > key
		r := self.ceiling(node.left, key)
		if r != nil {
			return r
		} else {
			return node
		}
	} else {
		return self.ceiling(node.right, key)
	}
}

// Returns the smallest key in the symbol table greater than or equal to {@code key}.
func (self *BST) Ceiling(key int) *Node {
	return self.ceiling(self.root, key)
}

func (self *BST) selectHelper(node *Node, k int) *Node {
	if node == nil {
		return node
	}
	if self.size(node.left) == k { // say k = 0, should return the smallest
		return node
	} else if self.size(node.left) < k {
		return self.selectHelper(node.right, k-1-self.size(node.left))
	} else {
		return self.selectHelper(node.left, k)
	}
}

// Return the key in the symbol table whose rank is k
/* Rank Definition:
(1) If the target is found, then the index ( = how many keys < k) is returned.
(2) If the target is not found, then the index to be inserted
of k ( =  ( = how many keys < k)) is returned. */
func (self *BST) Select(k int) *Node {
	return self.selectHelper(self.root, k)
}

func (self *BST) rank(node *Node, key int) int {
	if node == nil {
		return 0
	}
	if node.key < key {
		return self.size(node.left) + 1 + self.rank(node.right, key)
	} else if node.key == key {
		return self.size(node.left)
	} else {
		return self.rank(node.left, key)
	}
}

// Return the number of keys in the symbol table strictly less than `key`
func (self *BST) Rank(key int) int {
	return self.rank(self.root, key)
}

func (self *BST) keys(node *Node, res *[]int) {
	if node == nil {
		return
	} else {
		self.keys(node.left, res)
		*res = append(*res, node.key)
		self.keys(node.right, res)
	}
}

// Returns all keys in the symbol table as an Iterable
func (self *BST) Keys() []int {
	var res []int
	self.keys(self.root, &res)
	return res
}

func (self *BST) rangekeys(node *Node, lo int, hi int, res *[]int) {
	if node == nil {
		return
	}
	if node.key < lo {
		self.rangekeys(node.right, lo, hi, res)
	} else if node.key > hi {
		self.rangekeys(node.left, lo, hi, res)
	} else {
		self.rangekeys(node.left, lo, hi, res)
		*res = append(*res, node.key)
		self.rangekeys(node.right, lo, hi, res)
	}
}

// Returns all keys in the symbol table in the given range.
func (self *BST) RangeKeys(lo int, hi int) []int {
	var res []int
	self.rangekeys(self.root, lo, hi, &res)
	return res
}

// Returns the number of keys in the symbol table in the given range.
func (self *BST) RangeSize(lo int, hi int) int {
	return len(self.RangeKeys(lo, hi))
}

// Returns the keys in the BST in level order
func (self *BST) LevelOrder() []int {
	queue := make([]*Node, 0)
	res := make([]int, 0)
	if self.root != nil {
		queue = append(queue, self.root)
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
