package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-----Needleman–Wunsch Algorithm-----")
// Example from Pg. No. 151 (Problem Definition)
	similarity := nwString("AGGGCT", "AGGCA", 2, 1)
	fmt.Println("Similarity Distance of given two strings is =", similarity)
// Example from Pg. No. 152 (Quiz 17.1)
	similarity = nwString("AGTACG", "ACATAG", 2, 1)
	fmt.Println("Similarity Distance of given two strings is =", similarity)
// test case from http://www.algorithmsilluminated.com/datasets/problem17.8nw.txt
	x := "ACACATGCATCATGACTATGCATGCATGACTGACTGCATGCATGCATCCATCATGCATGCATCGATGCATGCATGACCACCTGTGTGACACATGCATGCGTGTGACATGCGAGACTCACTAGCGATGCATGCATGCATGCATGCATGC"
	y := "ATGATCATGCATGCATGCATCACACTGTGCATCAGAGAGAGCTCTCAGCAGACCACACACACGTGTGCAGAGAGCATGCATGCATGCATGCATGCATGGTAGCTGCATGCTATGAGCATGCAG"
	similarity = nwString(x, y, 5, 4)
	fmt.Println("Similarity Distance of given two strings is =", similarity)

	fmt.Println("-----Needleman–Wunsch Reconstruction Algorithm-----")
// Example from Pg. No. 151 (Problem Definition)
	valuesS1, valuesS2 := nwReconstruction("AGGGCT", "AGGCA", 2, 1)
	fmt.Printf("valuesS1 = %v\nvaluesS2 = %v\n", valuesS1, valuesS2)
// Example from Pg. No. 152 (Quiz 17.1)
	valuesS1, valuesS2 = nwReconstruction("AGTACG", "ACATAG", 2, 1)
	fmt.Printf("valuesS1 = %v\nvaluesS2 = %v\n", valuesS1, valuesS2)
}


/*
Needleman–Wunsch algorithm for String Sequence Problem
(Based on Explanation from Pg. No. 158)
Complexity - [ 𝑂(𝑚∗𝑛) ] (m: Length of first string, n: Length of second string)
*/
func nwString(s1, s2 string, alpXY, alpGap int) int {
	cache := make([][]int, len(s1)+1)
	for i := range cache {
		cache[i] = make([]int, len(s2)+1)
		cache[i][0] = i*alpGap
	}
	for j:=0; j<len(s2)+1; j++ {
		cache[0][j] = j*alpGap
	}
	var (
		t1 int
		t2 int
		t3 int
	)
	for i:=1; i<len(s1)+1; i++ {
		for j:=1; j<len(s2)+1; j++ {
			var curalpXY int
			if s1[i-1] != s2[j-1] {
				curalpXY = alpXY
			}
			t1 = cache[i-1][j-1]+curalpXY
			t2 = cache[i-1][j]+alpGap
			t3 = cache[i][j-1]+alpGap
			cache[i][j] = int(math.Min(float64(t1), math.Min(float64(t2), float64(t3))))
		}
	}
	return cache[len(s1)][len(s2)]
}


/*
Needleman–Wunsch Reconstruction algorithm for String Sequence Problem
(Own Implementation)
Complexity - [ 𝑂(𝑚∗𝑛) ] (m: Length of first string, n: Length of second string)
*/
func nwReconstruction(s1, s2 string, alpXY, alpGap int) ([]string, []string) {
	cache := make([][]int, len(s1)+1)
	for i := range cache {
		cache[i] = make([]int, len(s2)+1)
		cache[i][0] = i*alpGap
	}
	for j:=0; j<len(s2)+1; j++ {
		cache[0][j] = j*alpGap
	}
	var (
			t1 int
			t2 int
			t3 int
			temp int
		)
	nwStringAlgo := func() {
		for i:=1; i<len(s1)+1; i++ {
			for j:=1; j<len(s2)+1; j++ {
				var curalpXY int
				if s1[i-1] != s2[j-1] {
					curalpXY = alpXY
				}
				t1 = cache[i-1][j-1]+curalpXY
				t2 = cache[i-1][j]+alpGap
				t3 = cache[i][j-1]+alpGap
				cache[i][j] = int(math.Min(float64(t1), math.Min(float64(t2), float64(t3))))
			}
		}
	}
	
	nwStringAlgo()

	var valuesS1 []string
	var valuesS2 []string

	i := len(s1)
	j := len(s2)
	for i>=1 && j>=1 {
		var curalpXY int
		if s1[i-1] != s2[j-1] {
			curalpXY = alpXY
		}
		t1 = cache[i-1][j-1]+curalpXY
		t2 = cache[i-1][j]+alpGap
		t3 = cache[i][j-1]+alpGap
		temp = int(math.Min(float64(t1), math.Min(float64(t2), float64(t3))))
		if temp == t1 {
			valuesS1 = append(valuesS1, string(s1[i-1]))
			valuesS2 = append(valuesS2, string(s2[j-1]))
			i = i-1
			j = j-1
		} else if temp == t2 {
			valuesS1 = append(valuesS1, string(s1[i-1]))
			valuesS2 = append(valuesS2, "_")
			i = i-1
		} else { // temp == t3
			valuesS1 = append(valuesS1, "_")
			valuesS2 = append(valuesS2, string(s2[j-1]))
			j = j-1
		}
	}

	for i, j := 0, len(valuesS1)-1; i<j; i, j = i+1, j-1 {
		valuesS1[i], valuesS1[j] = valuesS1[j], valuesS1[i]
	}

	for i, j := 0, len(valuesS2)-1; i<j; i, j = i+1, j-1 {
		valuesS2[i], valuesS2[j] = valuesS2[j], valuesS2[i]
	}
	return valuesS1, valuesS2
}