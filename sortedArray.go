package main

import "fmt"

/**
Delete(key int)
Min() int
Max() int
Floor(k int) int
Ceiling(k int) int
Select(k int)             // Return the key in the symbol table whose rank is k
Rank(key int)             // Return the number of keys in the symbol table strictly less than `key`
Keys()                    // Returns all keys in the symbol table as an Iterable
RangeKeys(lo int, hi int) // Returns all keys in the symbol table in the given range.
RangeSize(lo int, hi int) // Returns the number of keys in the symbol table in the given range.
**/

type Node struct {
	key int
	val int
}

type SortedArray struct {
	array []Node
}

func NewSortedArray() *SortedArray {
	arrObj := new(SortedArray)
	return arrObj
}

// If the target is found,
//   then the index ( = how many keys < k) is returned.
// If the target is not found, then the index to be
//   inserted of k ( =  ( = how many keys < k)) is returned.
// https://www.zhihu.com/question/27161493
func (self *SortedArray) BinarySearch(key int) int {
	left := 0
	right := len(self.array) - 1
	for left <= right {
		mid := (left + right) / 2
		if self.array[mid].key == key {
			return mid
		} else if self.array[mid].key < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func (self *SortedArray) Contains(key int) bool {
	idx := self.BinarySearch(key)
	return idx < len(self.array) && self.array[idx].key == key
}

func (self *SortedArray) Put(key int, val int) {
	idx := self.BinarySearch(key)
	// The append built-in function appends elements to the end of a slice.
	//    If it has sufficient capacity, the destination is resliced to accommodate the new elements.
	//    If it does not, a new underlying array will be allocated.
	// The following expression is WRONG, since the append(self.array[:idx], Node{key, val}) \
	//    will modify the underlying array
	// self.array = append(append(self.array[:idx], Node{key, val}), self.array[idx:]...)
	self.array = append(self.array[:idx], append([]Node{Node{key, val}}, self.array[idx:]...)...)
}

func (self *SortedArray) Print() {
	fmt.Printf("\n")
	for _, node := range self.array {
		fmt.Printf("%v ", node.key)
		// only printf can support placeholder
	}
}

func main() {
	arr := NewSortedArray()
	arr.Put(10, 5)
	arr.Print()
	arr.Put(4, 7)
	arr.Print()
	arr.Put(1, 4)
	arr.Print()
	arr.Put(3, 9)
	arr.Print()
}
