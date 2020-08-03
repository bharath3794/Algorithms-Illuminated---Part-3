package main

import (
	"fmt"
	"mypackages/graph"
	"math"
	"math/rand"
	"time"
	"container/heap"
)



func main() {
	rand.Seed(time.Now().UnixNano())
// Example from Page No. 70
	edges := []graph.Edge{graph.Edge{K:"a", N:"b", W:1}, graph.Edge{K:"b", N:"a", W:1}, 
						  graph.Edge{K:"b", N:"d", W:2}, graph.Edge{K:"d", N:"b", W:2}, 
						  graph.Edge{K:"a", N:"d", W:3}, graph.Edge{K:"d", N:"a", W:3}, 
						  graph.Edge{K:"a", N:"c", W:4}, graph.Edge{K:"c", N:"a", W:4}, 
						  graph.Edge{K:"c", N:"d", W:5}, graph.Edge{K:"d", N:"c", W:5}}
	g1 := graph.CreateGraph(edges...)
	g2 := g1.Copy()
	fmt.Println("-----Prim Algorithm for Minimum Spanning Tree-----")
	mstEdges := primMST(g1)
	fmt.Printf("Edges that are part of Minimum Spanning Tree are\n%v\n", mstEdges)

	
	fmt.Println("-----Prim Algorithm for Minimum Spanning Tree using Heaps-----")
	mstEdges = primHeapBased(g2)
	fmt.Printf("Edges that are part of Minimum Spanning Tree are\n%v\n", mstEdges)

}


// Prim's Algorithm for finding Minimum Spanning Tree (From Page No. 72)
// 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠∗𝐸𝑑𝑔𝑒𝑠)
// If you're familiar with Dijkstra's Algorithm for Shortest Path Finding, 
// Prim's Algorithm is more or less similar to it.
func primMST(g graph.Graph) []graph.Edge {
	var mstEdges []graph.Edge
	keys := []string{}
	for key := range g.M {
		keys = append(keys, key)
	}
	idx := rand.Intn(len(keys))
	s := keys[idx]
	xMap := map[string]bool{s:true}
	for len(xMap) < len(g.M) {
		minCost := math.MaxInt64
        var minEdge graph.Edge
        for key := range xMap {
        	for neighbour := range g.M[key].N {
        		if _, ok := xMap[neighbour]; !ok {
        			if g.M[key].N[neighbour] < minCost {
        				minEdge = graph.Edge{key, neighbour, g.M[key].N[neighbour]}
        				minCost = g.M[key].N[neighbour]
        			}
        		}
        	}
        }
        xMap[minEdge.N] = true
        mstEdges = append(mstEdges, minEdge)
	}
	return mstEdges
}

/*
* Prim's Algorithm Using Heaps (Explanation From Pg. No. 79)
* Based on Book: [ 𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠+𝐸𝑑𝑔𝑒𝑠)∗𝑙𝑜𝑔(𝑛)) ];
* Below Implementation: [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠∗𝐸𝑑𝑔𝑒𝑠∗𝑙𝑜𝑔(𝑛)) ];
* While we delete the element in Heap (see Pg. No. 79, Prim's (Heap-Based) Line 12-16), in the below code 
we always need to search for the element first in the heap to know its index and then we would delete 
that element based on the index. If we don't search for key each time deleting the key, the complexity 
would normally be the mentioned complexity of [𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠+𝐸𝑑𝑔𝑒𝑠)∗𝑙𝑜𝑔(𝑛))]. As we are searching each time 
when we want to delete the key, it adds up to the complexity and 
the complexity would be [𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠*𝐸𝑑𝑔𝑒𝑠∗𝑙𝑜𝑔(𝑛))]
*/
func primHeapBased(g graph.Graph) []graph.Edge {
	var mstEdges []graph.Edge
	keys := []string{}
	for key := range g.M {
		keys = append(keys, key)
	}
	idx := rand.Intn(len(keys))
	s := keys[idx]
	xMap := map[string]bool{s:true}
	refMap := map[string]graph.Edge{}
	heapSlice := &heapVertices{}
	for _, key := range keys {
		if key == s {
			continue
		}
		if _, ok := g.M[key].N[s]; ok {
			refMap[key] = graph.Edge{s, key, g.M[key].N[s]}
			heap.Push(heapSlice, &vertex{key, g.M[key].N[s]})
		} else {
			refMap[key] = graph.Edge{"", "", math.MaxInt64}
			heap.Push(heapSlice, &vertex{key, math.MaxInt64})
		}
	}

	for len(*heapSlice) > 0 {
		popped := heap.Pop(heapSlice)
		w := popped.(*vertex)
		xMap[w.s] = true
		mstEdges = append(mstEdges, refMap[w.s])
		for neighbour := range g.M[w.s].N {
			if _, ok := xMap[neighbour]; !ok && g.M[w.s].N[neighbour] < refMap[neighbour].W {
				refMap[neighbour] = graph.Edge{w.s, neighbour, g.M[w.s].N[neighbour]}
				heapSlice.update(neighbour, g.M[w.s].N[neighbour])
			}
		}
	}
	return mstEdges 
}


type vertex struct {
	s string
	i int
}

type heapVertices []*vertex

func (h heapVertices) Len() int { return len(h) }
func (h heapVertices) Less(i, j int) bool { return h[i].i < h[j].i }
func (h heapVertices) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapVertices) Push(x interface{}) {
	*h = append(*h, x.(*vertex))
}

func (h *heapVertices) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return popped
}

func (h *heapVertices) update(key string, length int) {
	var idx int
	for j, item := range *h {
		if (*item).s == key {
			(*item).i = length
			idx = j
			break
		}
 	}
 	heap.Fix(h, idx)
}