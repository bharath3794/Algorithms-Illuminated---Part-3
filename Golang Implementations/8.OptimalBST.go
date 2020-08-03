package main

import (
	"fmt"
	"math"
)

func main() {

// Example from Youtube Tutorial: https://youtu.be/Xt8kuygspOk
	floatSlice := [][]float64{[]float64{1, 0.213}, []float64{2, 0.02}, []float64{3, 0.547}, 
							  []float64{4, 0.1}, []float64{5, 0.12}}

	reconstructedKeys := optimalBST(floatSlice)
	fmt.Println("Pass these keys in-order to a Normal BST to generate optimal BST =\n", reconstructedKeys)

// Generating a BST from the above reconstructed keys which creates an Optimal BST
	st1 := nodes{}

	for i:=0; i<len(reconstructedKeys); i++ {
		st1.insert(int(reconstructedKeys[i][0]))
	}

	st1.show()

// Test Case from: http://www.algorithmsilluminated.com/datasets/problem17.8optbst.txt
	fmt.Println("-----Test Case-----")
	testCase := [][]float64{[]float64{1, 2}, []float64{2, 8}, []float64{3, 2}, []float64{4, 5}, 
							[]float64{5, 5}, []float64{6, 2}, []float64{7, 8}, []float64{8, 3}, 
							[]float64{9, 6}, []float64{10, 1}, []float64{11, 1}, []float64{12, 6}, 
							[]float64{13, 3}, []float64{14, 2}, []float64{15, 6}, []float64{16, 7}, 
							[]float64{17, 4}, []float64{18, 63}, []float64{19, 2}, []float64{20, 9}, 
							[]float64{21, 10}, []float64{22, 1}, []float64{23, 60}, []float64{24, 5}, 
							[]float64{25, 2}, []float64{26, 7}, []float64{27, 34}, []float64{28, 11}, 
							[]float64{29, 31}, []float64{30, 76}, []float64{31, 21}, []float64{32, 6}, 
							[]float64{33, 8}, []float64{34, 1}, []float64{35, 81}, []float64{36, 37}, 
							[]float64{37, 15}, []float64{38, 6}, []float64{39, 8}, []float64{40, 24}, 
							[]float64{41, 12}, []float64{42, 18}, []float64{43, 42}, []float64{44, 8}, 
							[]float64{45, 51}, []float64{46, 21}, []float64{47, 8}, []float64{48, 6}, 
							[]float64{49, 5}, []float64{50, 7}}

	reconstructedKeys = optimalBST(testCase)
	fmt.Println("Pass these keys in-order to a Normal BST to generate optimal BST =\n", reconstructedKeys)

// Generating a BST from the above reconstructed keys which creates an Optimal BST
	st2 := nodes{}

	for i:=0; i<len(reconstructedKeys); i++ {
		st2.insert(int(reconstructedKeys[i][0]))
	}

	st2.show()

}



/*
Optimal Binary Search Tree Implementation using Dynamic Programming
Based on Implementation from Youtube Tutorial: 
https://www.youtube.com/playlist?list=PL4IH6CVPpTZUwOOyNioJgPnDOnOYO0MO9

Complexity = [ 𝑂((𝑛)^3) ] (n: no.of keys or values in the passed array)
*/
func optimalBST(arr [][]float64) [][]float64 {
	cache := make([][]float64, len(arr)+2)
	indexRootTable := make([][]int, len(arr)+2)
	for i := range cache {
		cache[i] = make([]float64, len(arr)+1)
		indexRootTable[i] = make([]int, len(arr)+1)
	}
	for i:=1; i<len(cache)-1; i++ {
		cache[i][i] = arr[i-1][1]
// index_root_table has the key number from 1 - len(arr) == 0 - len(arr)-1 indexes of array arr
// To get the original key index in array we have to subtract 1 from the value in index_root_table
// So, Here index_root_table[i][i] has i as key number = arr[i-1] as key
		indexRootTable[i][i] = i
	}

	for d:=1; d<len(arr); d++ {
		for i:=1; i<len(arr)+1-d; i++ {
			j := d+i
			minVal := math.Inf(1)
			rootKey := -1
			additionalCost := 0.0
			for l:=i; l<=j; l++ {
				additionalCost += arr[l-1][1]
				curVal := cache[i][l-1]+cache[l+1][j]
				if curVal < minVal {
					minVal = curVal
					rootKey = l
				}
			}
			cache[i][j] = minVal + additionalCost
// index_root_table has the key number from 1 - len(arr) == 0 - len(arr)-1 indexes of array arr
// To get the original key index in array we have to subtract 1 from the value in index_root_table
// So, Here index_root_table[i][j] has rootKey as key number = arr[rootKey-1] as key
			indexRootTable[i][j] = rootKey
		}
	}

	// Reconstructing the keys in correct order such that we can pass it to normal BST 
	// in the reconstructed order to generate optimal BST
	reconstructedKeys := [][]float64{}
// While reconstruction, its important to note that index_root_table has the key number from 1 - len(arr) == 0 - len(arr)-1 indexes of array arr
// To get the original key index in array we have to subtract 1 from the value in index_root_table
	stack := [][]int{[]int{indexRootTable[1][len(arr)], 1, len(arr)}}
	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curRootIdx, i, j := popped[0], popped[1], popped[2]
		reconstructedKeys = append(reconstructedKeys, arr[curRootIdx-1])
		if curRootIdx > i {
			stack = append(stack, []int{indexRootTable[i][curRootIdx-1], i, curRootIdx-1})
		}
		if curRootIdx < j {
			stack = append(stack, []int{indexRootTable[curRootIdx+1][j], curRootIdx+1, j})
		}
	}
	return reconstructedKeys
}








/*
Dependency to create Optimal BST - Binary Search Trees (From Algorithms Illuminated Part 2)
*/
type node struct{
	value int
	parent int
	left int
	right int
	size int
}


type nodes struct {
	n []node
}

/* 
This method inserts an element into Search Tree as you enter the key.
The first entered element is considered as root node and 
all the next elements are placed appropriately to maintain the property of Search Tree
*/
func (v *nodes) insert(ele int) {
	if len(v.n) == 0 {
		v.n = append(v.n, node{ele, -1, -1, -1, 0})
		return
	}
	i := 0
	for true {
		if ele <= v.n[i].value {
			if v.n[i].left < 0 {
				v.n = append(v.n, node{ele, i, -1, -1, 0})
				v.n[i].left = len(v.n)-1
				break
			} else {
				i = v.n[i].left
				continue
			}
		} else {
			if v.n[i].right < 0 {
				v.n = append(v.n, node{ele, i, -1, -1, 0})
				v.n[i].right = len(v.n)-1
				break
			} else {
				i = v.n[i].right
				continue
			}
		}
	}

}

// This method prints the search tree slice
func (v *nodes) show(){
	fmt.Println(v.n)
}