package main

import (
	"fmt"
	"time"
)

func main() {
// Example from Pg. No. 137
	input := [][]int{[]int{3, 4}, []int{2, 3}, []int{4, 2}, []int{4, 3}}
	cap := 6
// Own Example
	input = [][]int{[]int{1, 3}, []int{5, 7}, []int{8, 2}, []int{4, 1}, []int{3, 7}, 
					[]int{5, 3}, []int{1, 1}, []int{6, 4}}
	cap = 10


	fmt.Println("-----Naive Recursive Approach-----")
	summation, size := knapsackNaiveRecursive(input, cap)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)

	fmt.Println("-----Recursive Approach using Cache as Hash Map-----")
	summation, size = knapsackRecursiveCacheMap(input, cap)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)


	
	fmt.Println("-----Iterative Version-----")
	summation, size = knapsackIterative(input, cap)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)

		
	fmt.Println("-----Recursive Approach using Cache as 2D Array-----")
	summation, size = knapsackRecursiveCache2D(input, cap)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)


	fmt.Println("-----Knapsack Reconstruction using Iterative Version-----")
	values := knapsackReconstructionIterative(input, cap)
	fmt.Printf("Values = %v\n", values)

// Test Case from http://www.algorithmsilluminated.org/datasets/problem16.7test.txt
	testcase := [][]int{[]int{16808, 250}, []int{50074, 659}, []int{8931, 273}, []int{27545, 879}, 
					[]int{77924, 710}, []int{64441, 166}, []int{84493, 43}, []int{7988, 504}, 
					[]int{82328, 730}, []int{78841, 613}, []int{44304, 170}, []int{17710, 158}, 
					[]int{29561, 934}, []int{93100, 279}, []int{51817, 336}, []int{99098, 827}, 
					[]int{13513, 268}, []int{23811, 634}, []int{80980, 150}, []int{36580, 822}, 
					[]int{11968, 673}, []int{1394, 337}, []int{25486, 746}, []int{25229, 92}, 
					[]int{40195, 358}, []int{35002, 154}, []int{16709, 945}, []int{15669, 491}, 
					[]int{88125, 197}, []int{9531, 904}, []int{27723, 667}, []int{28550, 25}, 
					[]int{97802, 854}, []int{40978, 409}, []int{8229, 934}, []int{60299, 982}, 
					[]int{28636, 14}, []int{23866, 815}, []int{39064, 537}, []int{39426, 670}, 
					[]int{24116, 95}, []int{75630, 502}, []int{46518, 196}, []int{30106, 405}, 
					[]int{19452, 299}, []int{82189, 124}, []int{99506, 883}, []int{6753, 567}, 
					[]int{36717, 338}, []int{54439, 145}, []int{51502, 898}, []int{83872, 829}, 
					[]int{11138, 359}, []int{53178, 398}, []int{22295, 905}, []int{21610, 232}, 
					[]int{59746, 176}, []int{53636, 299}, []int{98143, 400}, []int{27969, 413}, 
					[]int{261, 558}, []int{41595, 9}, []int{16396, 969}, []int{19114, 531}, 
					[]int{71007, 963}, []int{97943, 366}, []int{42083, 853}, []int{30768, 822}, 
					[]int{85696, 713}, []int{73672, 902}, []int{48591, 832}, []int{14739, 58}, 
					[]int{31617, 791}, []int{55641, 680}, []int{37336, 7}, []int{97973, 99}, 
					[]int{49096, 320}, []int{83455, 224}, []int{12290, 761}, []int{48906, 127}, 
					[]int{36124, 507}, []int{45814, 771}, []int{35239, 95}, []int{96221, 845}, 
					[]int{12367, 535}, []int{25227, 395}, []int{41364, 739}, []int{7845, 591}, 
					[]int{36551, 160}, []int{8624, 948}, []int{97386, 218}, []int{95273, 540}, 
					[]int{99248, 386}, []int{13497, 886}, []int{40624, 421}, []int{28145, 969}, 
					[]int{35736, 916}, []int{61626, 535}, []int{46043, 12}, []int{54680, 153}}
	cap = 10000
	fmt.Println("----Performing knapsackRecursiveCacheMap() on test case-----")
	start := time.Now()
	summation, size = knapsackRecursiveCacheMap(testcase, cap)
	timeTaken := time.Since(start)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)
	fmt.Println("Time Taken for performing above computation =", timeTaken)

	fmt.Println("----Performing knapsackIterative() on test case-----")
	start = time.Now()
	summation, size = knapsackIterative(testcase, cap)
	timeTaken = time.Since(start)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)
	fmt.Println("Time Taken for performing above computation =", timeTaken)

	fmt.Println("----Performing knapsackRecursiveCache2D() on test case-----")
	start = time.Now()
	summation, size = knapsackRecursiveCache2D(testcase, cap)
	timeTaken = time.Since(start)
	fmt.Printf("Maximum Summation = %v, Total Size = %v\n", summation, size)
	fmt.Println("Time Taken for performing above computation =", timeTaken)

}


/*
Knapsack Problem Naive Recursive Approach (Own Implementation)
Complexity - grows exponentially with increase in no.of vertices
*/
func knapsackNaiveRecursive(arr [][]int, cap int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	var (
		s1 int
		s2 int
		c1 int
		c2 int
		s int
		c int
	)

	if arr[0][1] <= cap {
		s1, c1 = knapsackNaiveRecursive(arr[1:], cap-arr[0][1])
		s1, c1 = s1+arr[0][0], c1+arr[0][1]
	}

	s2, c2 = knapsackNaiveRecursive(arr[1:], cap)
	if s1 > s2 {
		s, c = s1, c1
	} else if s1 < s2 {
		s, c = s2, c2
	} else { // when s1 == s2
		s = s1
		
		if c1 <=c2 {
			c = c1
		} else {
			c = c2
		}
	}
	return s, c
}



/*
Knapsack Problem Recursive Version with Cache as Hash Map (Own Implementation)
Complexity - [ 𝑂(𝑛∗𝐶) ] approximately (n: No.of Values, C: Knapsack Capacity)
*/
func knapsackRecursiveCacheMap(arr [][]int, cap int) (int, int) {
	cache := map[string][]int{}
	var knapsackRec func(arr [][]int, cap int) (int, int) 
	knapsackRec = func(arr [][]int, cap int) (int, int) {
		if len(arr) == 0 {
			return 0, 0
		}
		key := fmt.Sprintf("%v%v%v", arr[0][0], arr[0][1], cap)
		if _, ok := cache[key]; ok {
			return cache[key][0], cache[key][1]
		}
		var (
			s1 int
			s2 int
			c1 int
			c2 int
			s int
			c int
		)

		if arr[0][1] <= cap {
			s1, c1 = knapsackRec(arr[1:], cap-arr[0][1])
			s1, c1 = s1+arr[0][0], c1+arr[0][1]
		}

		s2, c2 = knapsackRec(arr[1:], cap)
		if s1 > s2 {
			s, c = s1, c1
		} else if s1 < s2 {
			s, c = s2, c2
		} else { // when s1 == s2
			s = s1
			
			if c1 <=c2 {
				c = c1
			} else {
				c = c2
			}
		}
		cache[key] = []int{s, c}
		return cache[key][0], cache[key][1]
	}
	summation, size := knapsackRec(arr, cap)
	return summation, size
}


/*
Knapsack Problem Iterative Version (Based on Explanation from Pg. No. 141)
Complexity - [ 𝑂(𝑛∗𝐶) ] (n: No.of Values, C: Knapsack Capacity)
*/
func knapsackIterative(arr [][]int, cap int) (int, int) {
	cacheSum := make([][]int, len(arr)+1)
	cacheSize := make([][]int, len(arr)+1)	
	for i := range cacheSum {
		cacheSum[i] = make([]int, cap+1)
		cacheSize[i] = make([]int, cap+1)
	}
	for i:=1; i<len(arr)+1; i++ {
		for j:=0; j<cap+1; j++ {
			if arr[i-1][1] > j {
				cacheSum[i][j] = cacheSum[i-1][j]
				cacheSize[i][j] = cacheSize[i-1][j]
			} else { // When arr[i-1][1] <= j
				t1 := cacheSum[i-1][j-arr[i-1][1]] + arr[i-1][0]
				t2 := cacheSum[i-1][j]
				if t1 < t2 {
					cacheSum[i][j] = t2
					cacheSize[i][j] = cacheSize[i-1][j]
				} else {
					cacheSum[i][j] = t1
					cacheSize[i][j] = cacheSize[i-1][j-arr[i-1][1]] + arr[i-1][1]
				}
			}
		}
	}
	return cacheSum[len(arr)][cap], cacheSize[len(arr)][cap]
}


/*
Knapsack Problem Recursive Version of Cache using 2D array
(Similar to above iterative version implementation of cache with 2D Array)
Complexity - [ 𝑂(𝑛∗𝐶) ] (n: No.of Values, C: Knapsack Capacity)
*/
func knapsackRecursiveCache2D(arr [][]int, cap int) (int, int) {
	cacheSum := make([][]int, len(arr)+1)
	cacheSize := make([][]int, len(arr)+1)
	cacheSum[0] = make([]int, cap+1)
	cacheSize[0] = make([]int, cap+1)
	for i:=1; i<len(cacheSum); i++ {
		cacheSize[i] = make([]int, cap+1)
		cacheSum[i] = make([]int, cap+1)
		for j:=0; j<cap+1; j++ {
			cacheSum[i][j] = -1
		}
	}
	var knapsackRec2D func(arr [][]int, cap int) (int, int) 
	knapsackRec2D = func(arr [][]int, cap int) (int, int) {
		if len(arr) == 0 {
			return 0, 0
		}
		if cacheSum[len(arr)][cap] != -1 {
			return cacheSum[len(arr)][cap], cacheSize[len(arr)][cap]
		}
		var (
			s1 int
			s2 int
			c1 int
			c2 int
			s int
			c int
		)
		if arr[0][1] <= cap {
			s1, c1 = knapsackRec2D(arr[1:], cap-arr[0][1])
			s1, c1 = s1+arr[0][0], c1+arr[0][1]
			s2, c2 = knapsackRec2D(arr[1:], cap)
			if s1 > s2 {
				s, c = s1, c1
			} else if s1 < s2 {
				s, c = s2, c2
			} else { // when s1 == s2
				s = s1
				
				if c1 <= c2 {
					c = c1
				} else {
					c = c2
				}
			}
			cacheSum[len(arr)][cap], cacheSize[len(arr)][cap] = s, c
		} else {
			s2, c2 = knapsackRec2D(arr[1:], cap)
			cacheSum[len(arr)][cap], cacheSize[len(arr)][cap] = s2, c2
		}
		return cacheSum[len(arr)][cap], cacheSize[len(arr)][cap]
	}
	summation, size := knapsackRec2D(arr, cap)
	return summation, size
} 


/*
Knapsack Reconstruction (Own Implementation + Explanantion from Pg. No. 144)
Complexity for generating cache 2D array using knapsack_iterative() version = [ 𝑂(𝑛∗𝐶) ]
Complexity for reconstruction = [ 𝑂(𝑛) ]
Total Complexity = [ 𝑂(𝑛∗𝐶) ] (n: No.of Values, C: Knapsack Capacity)
( Ignored [ 𝑂(𝑛) ] as [ 𝑂(𝑛) ] < [ 𝑂(𝑛∗𝐶) ] )
*/
func knapsackReconstructionIterative(arr [][]int, cap int) [][]int {
	cacheSum := make([][]int, len(arr)+1)
	for i := range cacheSum {
		cacheSum[i] = make([]int, cap+1)
	}
	knapsackIter := func() {
		for i:=1; i<len(arr)+1; i++ {
			for j:=0; j<cap+1; j++ {
				if arr[i-1][1] > j {
					cacheSum[i][j] = cacheSum[i-1][j]
				} else { // When arr[i-1][1] <= j
					t1 := cacheSum[i-1][j-arr[i-1][1]] + arr[i-1][0]
					t2 := cacheSum[i-1][j]
					if t1 < t2 {
						cacheSum[i][j] = t2
					} else {
						cacheSum[i][j] = t1
					}
				}
			}
		}
	}
	knapsackIter()
	var values [][]int
	j := cap
	for i:=len(cacheSum)-1; i>=1; i-- {
		if arr[i-1][1] <= j && cacheSum[i][j] > cacheSum[i-1][j] {
			values = append(values, arr[i-1])
			j = j-arr[i-1][1]
		}
	}
	return values
}
