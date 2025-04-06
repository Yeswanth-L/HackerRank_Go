package main

import "fmt"

/*
simpler way:
func strangeCounter(t int64) int64 {
	cycleStart := int64(3)
	startTime := int64(1)

	for t >= startTime + cycleStart {
		startTime += cycleStart
		cycleStart *= 2
	}
	return cycleStart - (t - startTime)
}

*/
func strangeCounter(t int) int {
	// Write your code here
	round := 3
	var arr []int
	breakk := true
	j:=0
	in :=0
	for breakk{
		for i:=round;i>=1;i--{
			arr = append(arr, i)
			if j==t{
				breakk = false
				in = j
			}
			j++
		}
		round = round*2	
	}
	return arr[in-1]
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(strangeCounter(a))
}