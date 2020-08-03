package main

import (
	"fmt"
	"mypackages/graph"
	"mypackages/unionds"
	"math/rand"
	"time"
	"sort"
)



func main() {
	rand.Seed(time.Now().UnixNano())
// The below functions works only on undirected graphs
// Undirected Graph Example from Page No. 90
	edges := []graph.Edge{graph.Edge{K:"a", N:"b", W:4}, graph.Edge{K:"b", N:"a", W:4}, 
	 					  graph.Edge{K:"b", N:"d", W:5}, graph.Edge{K:"d", N:"b", W:5}, 
	 					  graph.Edge{K:"b", N:"c", W:3}, graph.Edge{K:"c", N:"b", W:3}, 
	 					  graph.Edge{K:"a", N:"c", W:2}, graph.Edge{K:"c", N:"a", W:2}, 
	 					  graph.Edge{K:"c", N:"d", W:6}, graph.Edge{K:"d", N:"c", W:6}, 
	 					  graph.Edge{K:"b", N:"e", W:1}, graph.Edge{K:"e", N:"b", W:1}, 
	 					  graph.Edge{K:"e", N:"d", W:7}, graph.Edge{K:"d", N:"e", W:7}}
	fmt.Println("-----Kruskal Algorithm for finding Minimum Spanning Tree-----")
	mstEdges := kruskal(edges)
	fmt.Println("mstEdges =", mstEdges)
	
// Undirected Graph Example from Page No. 90 (Same Above Example)
	edges = []graph.Edge{graph.Edge{K:"a", N:"b", W:4}, graph.Edge{K:"b", N:"a", W:4}, 
	 					  graph.Edge{K:"b", N:"d", W:5}, graph.Edge{K:"d", N:"b", W:5}, 
	 					  graph.Edge{K:"b", N:"c", W:3}, graph.Edge{K:"c", N:"b", W:3}, 
	 					  graph.Edge{K:"a", N:"c", W:2}, graph.Edge{K:"c", N:"a", W:2}, 
	 					  graph.Edge{K:"c", N:"d", W:6}, graph.Edge{K:"d", N:"c", W:6}, 
	 					  graph.Edge{K:"b", N:"e", W:1}, graph.Edge{K:"e", N:"b", W:1}, 
	 					  graph.Edge{K:"e", N:"d", W:7}, graph.Edge{K:"d", N:"e", W:7}}
	fmt.Println("-----Kruskal Algorithm for finding Minimum Spanning Tree using Union-Find Data Structure-----")
	mstEdges = kruskalUnionBased(edges)
	fmt.Println("mstEdges =", mstEdges)

}


// Kruskal Algorithm for finding Minimum Spanning Tree (Explanation From Pg. No. 92)
// [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠∗𝐸𝑑𝑔𝑒𝑠) ]
// The below function works only on undirected graphs
func kruskal(edges []graph.Edge) []graph.Edge {
	var mstEdges []graph.Edge
	g1 := graph.CreateGraph()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})
	for i:=0; i<len(edges); i+=2 {
		g1.MakeUnvisited()
		g1.AddEdges(edges[i], edges[i+1])
		if !isCyclic(&g1) {
			mstEdges = append(mstEdges, edges[i])			
		} else { // If the added edges made the graph cyclic
			g1.DeleteEdges(edges[i], edges[i+1])
		}
	}
	return mstEdges
}

// This function is dependency for kruskal()
// This function checks if there is a cycle in an undirected graph
func isCyclic(g *graph.Graph, pars ...string) bool {
	var key string
	var parent string
	var cycle bool
	if len(pars) == 0 {
		var keys []string
		for item := range (*g).M {
			keys = append(keys, item)
		}
		key = keys[rand.Intn(len(keys))]
		parent = "" 
	} else if len(pars) == 2 {
		key = pars[0]
		parent = pars[1]
	} else {
		panic("Passesd wrong no.of parameters")
	}
	g.SetE(key, true)
	for neighbour := range (*g).M[key].N {
		if neighbour == parent {
			continue
		}
	// If the neighbour is not parent and if it is already visited means there is a atleast one cycle
		if (*g).M[neighbour].E { 
			cycle = true
			break
		} else { // If the neighbour is unvisited, call the function recursively on that
			cycle = isCyclic(g, neighbour, key)
			if cycle {
				break
			}
		}
	}
	return cycle
}


// Kruskal Algorithm using Union-Find Based or Disjoint-Set Data Structure (Explanation From Pg. No. 97)
// [ 𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠+𝐸𝑑𝑔𝑒𝑠)∗𝑙𝑜𝑔(𝑛)) ]
// The below function works only on undirected graphs
func kruskalUnionBased(edges []graph.Edge) []graph.Edge {
	var mstEdges []graph.Edge
	var vertices []string
	keys := map[string]bool{}
	for _, edge := range edges {
		keys[edge.K] = true
		keys[edge.N] = true
	}
	for key := range keys {
		vertices = append(vertices, key)
	}
	uds := unionds.Initialize(vertices...)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})
	for _, edge := range edges {
		if uds.Find(edge.K) != uds.Find(edge.N) {
			mstEdges = append(mstEdges, edge)
			uds.Union(edge.K, edge.N)
		}
	}
	return mstEdges
}