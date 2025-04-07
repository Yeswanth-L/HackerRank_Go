package main

import (
	"fmt"
	"sort"
)

func flatlandSpaceStations(n int, c []int) int {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	maxDist := c[0] // Distance from city 0 to first station

	for i := 0; i < len(c)-1; i++ {
		// Midpoint between two stations
		dist := (c[i+1] - c[i]) / 2
		if dist > maxDist {
			maxDist = dist
		}
	}

	lastDist := n - 1 - c[len(c)-1] // Distance from last station to last city
	if lastDist > maxDist {
		maxDist = lastDist
	}

	return maxDist
}
func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	cities := make([]int,m)
	for i:=0;i<m;i++{
		fmt.Scan(&cities[i])
	}
	fmt.Println(flatlandSpaceStations(n,cities))
}