package main

import (
	"fmt"
	"container/heap"
)


type node struct{
	value string
	parent int
	left int
	right int
	encoded string
}


type tree struct {
	n []node
}



func main() {
// Example from Pg. No. 50
	arr := []int{60, 25, 10, 5}
// Example from Pg. No. 51
	arr = []int{3, 2, 6, 8, 2, 6}
// Test Case 1: From http://www.algorithmsilluminated.org/datasets/problem14.6test1.txt
	arr = []int{37, 59, 43, 27, 30, 96, 96, 71, 8, 76}
// Test Case 2: From http://www.algorithmsilluminated.org/datasets/problem14.6test2.txt
	arr = []int{895, 121, 188, 953, 378, 849, 153, 579, 144, 727, 589, 301, 442, 327, 930}

	binaryCodes := huffman(arr)
	fmt.Println("binaryCodes =", binaryCodes)
}

// Huffman Algorithm for finding Prefix-free binary codes (From Pg. No.49) (Using Heaps)
// [ 𝑂(𝑛∗𝑙𝑜𝑔(𝑛)) ]
func huffman(arr []int) [][]string {
	heapSlice := &heapFreq{}
	mapTrees := map[string]tree{}
	var huffTree tree
	var code string
	var prefixFreeCodes func(nodeIdx int)
	prefixFreeCodes = func(nodeIdx int) {
		i := nodeIdx
		leftNode := huffTree.n[i].left
		rightNode := huffTree.n[i].right
		if leftNode < 0 && rightNode < 0 {
			huffTree.n[i].encoded = code
			return
		}
		if leftNode >= 0 {
			code += "0"
			prefixFreeCodes(leftNode)
			code = code[:len(code)-1]
		}
		if rightNode >= 0 {
			code += "1"
			prefixFreeCodes(rightNode)
			code = code[:len(code)-1]
		}
		return
	}
	for i:=0; i<len(arr); i++ {
		temp := fmt.Sprintf("t%v", i+1)
		heap.Push(heapSlice, &freq{s:temp, f:arr[i]})
		mapTrees[temp] = tree{[]node{node{temp, -1, -1, -1, ""}}}
	}
	subscript := len(mapTrees)+1
	for len(mapTrees) >= 2 {
		temp := fmt.Sprintf("t%v", subscript)
		temp1 := (heap.Pop(heapSlice)).(*freq)
		temp2 := (heap.Pop(heapSlice)).(*freq)
		heap.Push(heapSlice, &freq{s:temp, f:temp1.f+temp2.f})
		tree1 := mapTrees[temp1.s]
		tree2 := mapTrees[temp2.s]
		delete(mapTrees, temp1.s)
		delete(mapTrees, temp2.s)
		mapTrees[temp] = mergeTrees(tree1, tree2)
		subscript++
	}
	huffTree = mapTrees[fmt.Sprintf("t%v", subscript-1)]
	prefixFreeCodes(0)
	var binaryCodes [][]string
	for i:=0; i<len(huffTree.n); i++ {
		if huffTree.n[i].value != "" {
			binaryCodes = append(binaryCodes, []string{huffTree.n[i].value, huffTree.n[i].encoded})
		}
	}
	return binaryCodes
}


// This function is dependency for huffman()
func mergeTrees(tree1, tree2 tree) tree {
	if len(tree1.n) < len(tree2.n) {
		tree1.n, tree2.n = tree2.n, tree1.n
 	}
 	for i:=0; i<len(tree1.n); i++ {
 		tree1.n[i].parent++
 		if tree1.n[i].left >= 0 {
 			tree1.n[i].left++
 		}
 		if tree1.n[i].right >= 0 {
 			tree1.n[i].right++
 		}
 	}
 	increment := len(tree1.n)+1
 	tree2.n[0].parent = -increment
 	for i:=0; i<len(tree2.n); i++ {
 		tree2.n[i].parent += increment
 		if tree2.n[i].left >= 0 {
 			tree2.n[i].left += increment
 		}
 		if tree2.n[i].right >= 0 {
 			tree2.n[i].right += increment
 		}
 	}
 	tree3 := tree{[]node{node{"", -1, 1, increment, ""}}}
 	tree3.n = append(tree3.n, tree1.n...)
 	tree3.n = append(tree3.n, tree2.n...)
 	return tree3
}



/* 
Heap Implementation - Dependency for huffman()
*/
type freq struct {
	s string
	f int
}

type heapFreq []*freq

func (h heapFreq) Len() int { return len(h) }
func (h heapFreq) Less(i, j int) bool { return h[i].f < h[j].f }
func (h heapFreq) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapFreq) Push(x interface{}) {
	*h = append(*h, x.(*freq))
}

func (h *heapFreq) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return popped
}