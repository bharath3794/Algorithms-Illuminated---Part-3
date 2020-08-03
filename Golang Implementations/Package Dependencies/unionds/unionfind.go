/*
************Union-Find/ Disjoint-Set Data Structure************
Based on explanation from Page No. 98 to 103 and also from https://bit.ly/337fmv1 
particularly used path compression technique while implementing
Direct URL for above Shortened URL: https://opendsa-server.cs.vt.edu/ODSA/Books/Everything/html/UnionFind.html
For Complexities of these operations, refer to Pg. No. 96
***************************************************************

Usage:
	import this "unionds" package in Go main file

Initialize Union-Find Data Structure:
Syntax: Initialize(vertices ...string) Partition {...}
	Ex: uds := Initialize("1", "2", "3", "4", "5", "6") 
	Output: Partition{P: []Node{[{1 0 1} {2 1 1} {3 2 1} {4 3 1} {5 4 1} {6 5 1}]} 
					  Index: map[1:0 2:1 3:2 4:3 5:4 6:5]}
Finding to which set in the partition the given key belongs to:
Given a key, it returns the parent of the set this key belongs to
Syntax: uds.Find(key string) int {...}
	Ex: uds.Find("5") -> key "5" belongs to a set whose parent is 4, so it returns 4
Combine two sets by giving any of the keys in those sets using Union() method
Given two keys, it searches for the sets to which these two keys belongs to and merge these two sets
Syntax: uds.UnionUnion(key1, key2 string)
	Ex: 
	Step 1: uds := Initialize("1", "2", "3", "4", "5", "6")
	Output: Partition{P: []Node{[{1 0 1} {2 1 1} {3 2 1} {4 3 1} {5 4 1} {6 5 1}]} 
					  Index: map[1:0 2:1 3:2 4:3 5:4 6:5]}
	Step 2:	uds.Union("5", "6")
	Output: Partition{P: []Node{[{1 0 1} {2 1 1} {3 2 1} {4 3 1} {5 4 1} {6 4 1}]} 
					  Index: map[1:0 2:1 3:2 4:3 5:4 6:5]}
	The sets of those two keys "5", "6" are merged. Now key "6" parent is "5" thus they both are 
	linked together as a single set
*/




package unionds

import (
	"fmt"
)


type Node struct {
	N string // Name of the node N
	P int // Parent's Index P
	S int // Size S
}

type Partition struct {
	P []Node // Parition P
	Index map[string]int // Map for index of each node's name N
}

func Initialize(vertices ...string) Partition {
	slice := Partition{P:[]Node{}, Index:map[string]int{}}
	for _, vertex := range vertices {
		slice.P = append(slice.P, Node{vertex, len(slice.P), 1})
		slice.Index[vertex] = len(slice.P)-1
	}
	return slice
}


func (p *Partition) Find(key string) int {
	idx := p.Index[key]
	for p.P[idx].P != idx {
		idx = p.P[idx].P
	}
	return p.P[idx].P 
}

func (p *Partition) Union(key1, key2 string) {
	root1Idx := p.Find(key1)
	root2Idx := p.Find(key2)
	if root1Idx != root2Idx {
		if p.P[root1Idx].S >= p.P[root2Idx].S {
			p.P[root2Idx].P = root1Idx
			p.P[root1Idx].S += p.P[root2Idx].S
		} else {
			p.P[root1Idx].P = root2Idx
			p.P[root2Idx].S += p.P[root1Idx].S
		}
	}
}


func (p *Partition) Show() {
	fmt.Println(p.P)
}