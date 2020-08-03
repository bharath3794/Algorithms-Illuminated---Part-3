package main

import (
	"fmt"
)

type vertex struct {
	name string
	weight int
}

func main() {
// Example from Pg. No. 122
	vertices := []vertex{vertex{"a", 1}, vertex{"b", 4}, vertex{"c", 5}, vertex{"d", 4}}

// Example from Pg. No. 131
	vertices = []vertex{vertex{"a", 3}, vertex{"b", 2}, vertex{"c", 1}, vertex{"d", 6}, vertex{"e", 4}, 
						vertex{"f", 5}}

	fmt.Println("-----Naive Recursive Approach-----")
	maxWIS := wisNaiveRecursive(vertices)
	fmt.Println("The maximum weighted independent set sum is", maxWIS)

	fmt.Println("-----Recursive Approach with Cache-----")
	maxWIS = wisRecursiveCache(vertices)
	fmt.Println("The maximum weighted independent set sum is", maxWIS)

	fmt.Println("-----Iterative Approach-----")
	maxWIS = wisIterative(vertices)
	fmt.Println("The maximum weighted independent set sum is", maxWIS)

	fmt.Println("-----WIS Reconstruction using Iterative Approach-----")
	sets := wisReconstructionIterative(vertices)
	fmt.Println("Sets =", sets)
	
	fmt.Println("-----WIS Reconstruction using Recursion with Cache-----")
	sets = wisReconstructionRecCache(vertices)
	fmt.Println("Sets =", sets)
}


/*
Weighted Independent Set Problem using Naive Recursive Approach
(Based on Explanation from Pg.No. 124)
Complexity - grows exponentially with increase in no.of vertices
*/
func wisNaiveRecursive(vertices []vertex) int {
	var maxWIS int
	if len(vertices) == 0 {
		return 0
	} else if len(vertices) == 1 {
		return vertices[0].weight
	}
	s1 := wisNaiveRecursive(vertices[:len(vertices)-1])
	s2 := wisNaiveRecursive(vertices[:len(vertices)-2])
	if s1 >= s2+vertices[len(vertices)-1].weight {
		maxWIS = s1
	} else {
		maxWIS = s2+vertices[len(vertices)-1].weight
	}
	return maxWIS
}

/*
Weighted Independent Set Problem using Recursion with Cache
Based on Explanation from Pg.No. 125 (Recursion with Cache)
Complexity - [ 𝑂(𝑛) ]
*/
func wisRecursiveCache(vertices []vertex) int {
	cache := map[string]int{}
	var wisRecursion func(ref []vertex) int
	wisRecursion = func(ref []vertex) int {
		if len(ref) == 0 {
			return 0
		} else if value, ok := cache[ref[len(ref)-1].name]; ok {
			return value
		} else if len(ref) == 1 {
			cache[ref[0].name] = ref[0].weight
			return ref[0].weight
		}
		s1 := wisRecursion(ref[:len(ref)-1])
		s2 := wisRecursion(ref[:len(ref)-2])
		if s1 >= s2+ref[len(ref)-1].weight {
			cache[ref[len(ref)-1].name] = s1
		} else {
			cache[ref[len(ref)-1].name] = s2+ref[len(ref)-1].weight
		}
		return cache[ref[len(ref)-1].name]
	}
	maxWIS := wisRecursion(vertices)
	fmt.Println("Cache =", cache)
	return maxWIS
}


/*
Weighted Independent Set Problem Iterative Version
Based on Explanation from Pg.No. 127
Complexity - [ 𝑂(𝑛) ]
*/

func wisIterative(vertices []vertex) int {
	maxWIS := []int{0, vertices[0].weight}
	for i:=2; i<len(vertices)+1; i++ {
		if maxWIS[i-1] >= maxWIS[i-2] + vertices[i-1].weight {
			maxWIS = append(maxWIS, maxWIS[i-1])
		} else {
			maxWIS = append(maxWIS, maxWIS[i-2] + vertices[i-1].weight)
		}
	}
	return maxWIS[len(maxWIS)-1]
}


/*
WIS Reconstrution using Iterative Version (Based on Explanation from Pg.No. 130)
(Giving the vertices that are part of the maximum weighted independent set)
Complexity - [ 𝑂(𝑛) ]
*/
func wisReconstructionIterative(vertices []vertex) []vertex {
	var maxWIS []int
	var wisIter func()
	wisIter = func() {
		maxWIS = append(maxWIS, 0, vertices[0].weight)
		for i:=2; i<len(vertices)+1; i++ {
			if maxWIS[i-1] >= maxWIS[i-2] + vertices[i-1].weight {
				maxWIS = append(maxWIS, maxWIS[i-1])
			} else {
				maxWIS = append(maxWIS, maxWIS[i-2] + vertices[i-1].weight)
			}
		}
	}
	wisIter()
	sets := []vertex{}
	i := len(maxWIS)-1
	for i >= 2 {
		if maxWIS[i-1] >= maxWIS[i-2] + vertices[i-1].weight {
			i = i-1
		} else {
			sets = append(sets, vertices[i-1])
			i = i-2
		}
	}
	if i == 1 {
		sets = append(sets, vertices[0])
	}
	return sets
}

/*
WIS Reconstrution using Recursive Version with Cache
(Giving the vertices that are part of the maximum weighted independent set)
(Based on Explanation from Pg.No. 130)
Complexity - [ 𝑂(𝑛) ]
*/
func wisReconstructionRecCache(vertices []vertex) []vertex {
	cache := map[string]int{}
	var wisRecCache func(ref []vertex) int
	wisRecCache = func(ref []vertex) int {
		if len(ref) == 0 {
			return 0
		} else if value, ok := cache[ref[len(ref)-1].name]; ok {
			return value
		} else if len(ref) == 1 {
			cache[ref[0].name] = ref[0].weight
			return ref[0].weight
		}
		s1 := wisRecCache(ref[:len(ref)-1])
		s2 := wisRecCache(ref[:len(ref)-2])
		if s1 >= s2+ref[len(ref)-1].weight {
			cache[ref[len(ref)-1].name] = s1
		} else {
			cache[ref[len(ref)-1].name] = s2+ref[len(ref)-1].weight
		}
		return cache[ref[len(ref)-1].name]
	}

	_ = wisRecCache(vertices)

	maxWIS := []int{0}

	for i:=0; i<len(vertices); i++ {
		maxWIS = append(maxWIS, cache[vertices[i].name])
	}
	sets := []vertex{}
	i := len(maxWIS)-1
	for i >= 2 {
		if maxWIS[i-1] >= maxWIS[i-2] + vertices[i-1].weight {
			i = i-1
		} else {
			sets = append(sets, vertices[i-1])
			i = i-2
		}
	}
	if i == 1 {
		sets = append(sets, vertices[0])
	}
	return sets
}