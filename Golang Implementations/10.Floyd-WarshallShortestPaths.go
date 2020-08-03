package main

import (
	"fmt"
	"mypackages/graph"
	"math"
)

func main() {
// Example from Page No. 181
	edges := []graph.Edge{graph.Edge{K:"s", N:"v", W:1}, graph.Edge{K:"v", N:"t", W:-5}, 
						  graph.Edge{K:"s", N:"t", W:-2}}
// Example from Page No. 182
	edges = []graph.Edge{graph.Edge{K:"s", N:"v", W:10}, graph.Edge{K:"v", N:"u", W:-4}, 
						 graph.Edge{K:"u", N:"w", W:3}, graph.Edge{K:"w", N:"x", W:-5}, 
						 graph.Edge{K:"x", N:"v", W:4}}
// Example from Page No. 193
	edges = []graph.Edge{graph.Edge{K:"s", N:"v", W:4}, graph.Edge{K:"s", N:"u", W:2}, 
						 graph.Edge{K:"u", N:"v", W:-1}, graph.Edge{K:"v", N:"t", W:4}, 
						 graph.Edge{K:"u", N:"w", W:2}, graph.Edge{K:"w", N:"t", W:2}}
	
	g1 := graph.CreateGraph(edges...)
	// Normal Dijkstra Algorithm (From Page No. 92)
	fmt.Println("-----Floyd-Warshall Algorithm using 2D Array as Cache-----")
	distances, ok := floydWarshall(g1)
	if ok {
		fmt.Println("All-Pairs Shortest Paths =", distances)
	} else {
		fmt.Println("Negative Cycle Found in the Graph")
	}

	fmt.Println("-----Floyd-Warshall Algorithm Book Version using 3D Array as Cache-----")
	distances, ok = floydWarshallBookVersion(g1)
	if ok {
		fmt.Println("All-Pairs Shortest Paths =", distances)
	} else {
		fmt.Println("Negative Cycle Found in the Graph")
	}
	
// test case 1: http://www.algorithmsilluminated.com/datasets/problem18.8test1.txt
	edges = []graph.Edge{graph.Edge{K:"1", N:"2", W:2}, graph.Edge{K:"1", N:"5", W:3}, 
						 graph.Edge{K:"2", N:"4", W:-2}, graph.Edge{K:"3", N:"1", W:1}, 
						 graph.Edge{K:"4", N:"1", W:4}, graph.Edge{K:"4", N:"3", W:1}, 
						 graph.Edge{K:"4", N:"5", W:2}, graph.Edge{K:"5", N:"3", W:-1}}
// test case 2: http://www.algorithmsilluminated.com/datasets/problem18.8test2.txt
	edges = []graph.Edge{graph.Edge{K:"1", N:"2", W:2}, graph.Edge{K:"1", N:"5", W:3}, 
						 graph.Edge{K:"2", N:"4", W:-2}, graph.Edge{K:"3", N:"1", W:1}, 
						 graph.Edge{K:"4", N:"1", W:4}, graph.Edge{K:"4", N:"3", W:1}, 
						 graph.Edge{K:"4", N:"5", W:-1}, graph.Edge{K:"5", N:"3", W:-1}}

	g1 = graph.CreateGraph(edges...)
	fmt.Println("-----Floyd-Warshall Algorithm using 2D Array as Cache-----")
	distances, ok = floydWarshall(g1)
	if ok {
		fmt.Println("All-Pairs Shortest Paths =", distances)
	} else {
		fmt.Println("Negative Cycle Found in the Graph")
	}

	fmt.Println("-----Floyd-Warshall Algorithm Book Version using 3D Array as Cache-----")
	distances, ok = floydWarshallBookVersion(g1)
	if ok {
		fmt.Println("All-Pairs Shortest Paths =", distances)
	} else {
		fmt.Println("Negative Cycle Found in the Graph")
	}
}



/*
* Floyd-Warshall Algorithm for finding All-Pairs Shortest Paths with 2D Array as Cache 
  (i.e. Shortest Paths from each vertex to each other vertex)
* Implementation is based on Abdul Badari Youtube Video: https://youtu.be/oNI0rf2P9gE
* Complexity = [ 𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠)3) ] which is a polynomial complexity
*/
func floydWarshall(g graph.Graph) ([]graph.Edge, bool) {
	var keys []string
	for key := range g.M {
		keys = append(keys, key)
	}
	cache := make([][]int, len(keys))
	for i := range cache {
		cache[i] = make([]int, len(keys))
		for j:=0; j<len(keys); j++ {
			if i == j {
				cache[i][j] = 0
			} else if val, ok := g.M[keys[i]].N[keys[j]]; ok {
				cache[i][j] = val
			} else {
				cache[i][j] = math.MaxInt32
			}
		}
	}
	var distances []graph.Edge
	for k:=0; k<len(keys); k++ {
		for i:=0; i<len(keys); i++ {
			if i == k {
				continue
			}
			for j:=0; j<len(keys); j++ {
				if j == k {
					continue
				}
				if cache[i][j] > cache[i][k]+cache[k][j] {
					cache[i][j] = cache[i][k]+cache[k][j]
				}
			}
		}
	}
	for i := range(keys) {
		if cache[i][i] < 0 {
			return distances, false
		}
	}
	for i := range(keys) {
		for j := range(keys) {
			if i == j {
				continue
			}
			if float64(cache[i][j])/float64(math.MaxInt32) < 0.5 {
				distances = append(distances, graph.Edge{K:keys[i], N:keys[j], W:cache[i][j]})
			} 
		}
	}
	return distances, true
}


/*
Floyd-Warshall Algorithm Book Version with 3D Array as Cache
(Based on Implementation from Pg. No. 206)
Complexity = [ 𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠)3) ] which is a polynomial complexity
*/
func floydWarshallBookVersion(g graph.Graph) ([]graph.Edge, bool) {
	var keys []string
	for key := range g.M {
		keys = append(keys, key)
	}

	cache := make([][][]int, len(keys)+1)
	for i := range keys {
		cache[i] = make([][]int, len(keys))
		for k := range cache[i] {
			cache[i][k] = make([]int, len(keys))
		}
		for j := range keys {
			if i == j {
				cache[0][i][j] = 0
			} else if val, ok := g.M[keys[i]].N[keys[j]]; ok {
				cache[0][i][j] = val
			} else {
				cache[0][i][j] = math.MaxInt32
			}
		}
	}
	cache[len(keys)] = make([][]int, len(keys)) // Last item still not set
	for k := range cache[len(keys)] {
		cache[len(keys)][k] = make([]int, len(keys))
	}
	var distances []graph.Edge
	for k := range keys {
		for i := range keys {
			for j := range keys {
				if cache[k][i][j] <= cache[k][i][k]+cache[k][k][j] {
					cache[k+1][i][j] = cache[k][i][j]
				} else {
					cache[k+1][i][j] = cache[k][i][k]+cache[k][k][j]
				}
			}
		}
	}

	for i := range keys {
		if cache[len(keys)][i][i] < 0 {
			return distances, false
		}
	}
	
	for i := range cache[len(keys)] {
		for j := range cache[len(keys)][0] {
			if i == j {
				continue
			}
			if float64(cache[len(keys)][i][j])/float64(math.MaxInt32) < 0.5 {
				distances = append(distances, graph.Edge{K:keys[i], N:keys[j], W:cache[len(keys)][i][j]})
			} 
		}
	}
	return distances, true
}