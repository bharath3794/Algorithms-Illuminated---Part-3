package main

import (
	"fmt"
	"sort"
)

type job struct {
	weight int
	length int
	name string
}



func main() {
	// Example from Pg. No. 22
	jobs := []job{job{weight:3, length:5, name:"j1"}, job{weight:1, length:2, name:"j2"}}
	// Example with 12 Jobs (Link: http://www.algorithmsilluminated.org/datasets/problem13.4test.rtf)
	jobs = []job{job{weight:8, length:50, name:"j1"}, job{weight:74, length:59, name:"j2"}, 
				 job{weight:31, length:73, name:"j3"}, job{weight:45, length:79, name:"j4"}, 
				 job{weight:24, length:10, name:"j5"}, job{weight:41, length:66, name:"j6"}, 
				 job{weight:93, length:43, name:"j7"}, job{weight:88, length:4, name:"j8"}, 
				 job{weight:28, length:30, name:"j9"}, job{weight:41, length:13, name:"j10"}, 
				 job{weight:4, length:70, name:"j11"}, job{weight:10, length:58, name:"j12"}}
	fmt.Println("-----Method 1-----")
	fmt.Println("-----GreedyDiff-----")
	weightedSum, schedule := greedyDiff(jobs)
	fmt.Println("weightedSum =", weightedSum, "\nschedule =", schedule)
	fmt.Println("-----GreedyRatio-----")
	weightedSum, schedule = greedyRatio(jobs)
	fmt.Println("weightedSum =", weightedSum, "\nschedule =", schedule)

	fmt.Println("-----Method 2 (Other Way of doing it)-----")
	fmt.Println("-----GreedyDiff2-----")
	weightedSum, schedule = greedyDiff2(jobs)
	fmt.Println("weightedSum =", weightedSum, "\nschedule =", schedule)
	fmt.Println("-----GreedyRatio2-----")
	weightedSum, schedule = greedyRatio2(jobs)
	fmt.Println("weightedSum =", weightedSum, "\nschedule =", schedule)
}


/*
GreedyDiff and GreedyRatio (Programming Problems: Problem 13.4)
Based on explanation from Pg. No. 22 (GreedyDiff and GreedyRatio), 
Pg. No. 23 (Quiz 13.3), Pg.No. 25 (Solution to Quiz 13.3)
*/
func greedyDiff(jobs []job) (int, []job) {
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].weight-jobs[i].length != jobs[j].weight-jobs[j].length {
			return jobs[i].weight-jobs[i].length > jobs[j].weight-jobs[j].length
		} else {
			return jobs[i].weight > jobs[j].weight
		}
	})
	var weightedSum int
	completionTime := []int{jobs[0].length}
	for i:=0; i<len(jobs); i++ {
		if i < len(jobs)-1 {
			completionTime = append(completionTime, completionTime[i]+jobs[i+1].length)
		}
		weightedSum += completionTime[i]*jobs[i].weight
	}
	return weightedSum, jobs
}


func greedyRatio(jobs []job) (int, []job) {
	sort.Slice(jobs, func(i, j int) bool {
		ithTerm := float64(jobs[i].weight)/float64(jobs[i].length)
		jthTerm := float64(jobs[j].weight)/float64(jobs[j].length)
		if  ithTerm != jthTerm {
			return ithTerm > jthTerm
		} else {
			return jobs[i].weight > jobs[j].weight
		}
	})
	var weightedSum int
	completionTime := []int{jobs[0].length}
	for i:=0; i<len(jobs); i++ {
		if i < len(jobs)-1 {
			completionTime = append(completionTime, completionTime[i]+jobs[i+1].length)
		}
		weightedSum += completionTime[i]*jobs[i].weight
	}
	return weightedSum, jobs
}


/*
************** Method 2 (Other Way of doing it) *******************
*/
func greedyDiff2(jobs []job) (int, []job) {
	var diff [][]int
	for i:=0; i<len(jobs); i++ {
		diff = append(diff, []int{jobs[i].weight-jobs[i].length, i})
	}
	sort.Slice(diff, func(i, j int) bool {
		if diff[i][0] != diff[j][0] {
			return diff[i][0] > diff[j][0]
		} 
		return jobs[diff[i][1]].weight > jobs[diff[j][1]].weight
	})
	schedule := []job{}
	for _, item := range diff {
		schedule = append(schedule, jobs[item[1]])
	}

	var weightedSum int
	completionTime := []int{schedule[0].length}
	for i:=0; i<len(schedule); i++ {
		if i < len(schedule)-1 {
			completionTime = append(completionTime, completionTime[i]+schedule[i+1].length)
		}
		weightedSum += completionTime[i]*schedule[i].weight
	}
	return weightedSum, schedule
}
/*
************** Method 2 (Other Way of doing it) *******************
*/
func greedyRatio2(jobs []job) (int, []job) {
	var ratio [][]float64
	for i:=0; i<len(jobs); i++ {
		ratio = append(ratio, []float64{float64(jobs[i].weight)/float64(jobs[i].length), float64(i)})
	}
	sort.Slice(ratio, func(i, j int) bool {
		if ratio[i][0] != ratio[j][0] {
			return ratio[i][0] > ratio[j][0]
		} 
		return jobs[int(ratio[i][1])].weight > jobs[int(ratio[j][1])].weight
	})
	schedule := []job{}
	for _, item := range ratio {
		schedule = append(schedule, jobs[int(item[1])])
	}

	var weightedSum int
	completionTime := []int{schedule[0].length}
	for i:=0; i<len(schedule); i++ {
		if i < len(schedule)-1 {
			completionTime = append(completionTime, completionTime[i]+schedule[i+1].length)
		}
		weightedSum += completionTime[i]*schedule[i].weight
	}
	return weightedSum, schedule
}