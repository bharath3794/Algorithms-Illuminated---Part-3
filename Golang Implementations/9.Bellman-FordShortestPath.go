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
	fmt.Println("-----Bellman-Ford Algorithm for Shortest Path-----")
	distances, ok := bellmanFord(g1, "s")
	if ok {
		fmt.Println("Shortest Distances from Source to each vertex =", distances)
	} else {
		fmt.Println("Negative Cycle Found in the Graph")
	}
}



/*
* Bellman-Ford Algorithm for finding Shortest-Path of Graphs with Negative edge weights
* Based on Explanantion from Pg. No. 192; 
  Also Refer to MIT Tutorial on Bellman-Ford: https://youtu.be/ozsuci5pIso
* Complexity = [ 𝑂(𝐸𝑑𝑔𝑒𝑠∗𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠) ]
* As Edges for a fully connected graph can be of  (𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠∗(𝑉𝑒𝑡𝑖𝑐𝑒𝑠−1))/2 , leading to a 
  total Complexity of [ 𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠)^3) ] which is a polynomial complexity
*/
func bellmanFord(g graph.Graph, source string)  (map[string]int, bool) {
	cache := make([][]int, len(g.M)+1)
	for i := range cache {
		cache[i] = make([]int, len(g.M))
	}
	var keys []string
	for key := range g.M {
		keys = append(keys, key)
	}
	indexKeys := map[string]int{}
	for j := range keys {
		indexKeys[keys[j]] = j
		if keys[j] != source {
			cache[0][j] = math.MaxInt32
		} else {
			cache[0][j] = 0
		}
	}
	distances := map[string]int{}
	for i:=1; i<len(cache); i++ {
		stable := true
		for j:=0; j<len(keys); j++ {
			minVal := math.MaxInt32
			var curMin int
			for _, key := range keys {
				if val, ok := g.M[key].N[keys[j]]; ok {
					curMin = cache[i-1][indexKeys[key]] + val
					if curMin < minVal {
						minVal = curMin
					}
				}
			}
			if cache[i-1][j] <= minVal {
				cache[i][j] = cache[i-1][j]
			} else {
				cache[i][j] = minVal
			}
			if cache[i][j] != cache[i-1][j] {
				stable = false
			}
		}
		if stable == true {
			for k, key := range keys {
				distances[key] = cache[i-1][k]
			}
			return distances, true
		}
	}
	return distances, false
}
