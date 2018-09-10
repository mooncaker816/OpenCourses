package main

import (
	"fmt"
	"time"
)

var ints = []int{3, 15, 17, 23, 11, 3, 4, 5, 17, 23, 34, 17, 18, 14, 12, 15}

func main() {
	start := time.Now()
	fmt.Println(maxSum(ints))
	fmt.Println(time.Since(start))
	lookup := make(map[int]int)
	start = time.Now()
	fmt.Println(maxSumWithLookup(ints, lookup))
	fmt.Println(time.Since(start))
}

func maxSum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	pickFirst := maxSum(a[2:]) + a[0]
	skipFirst := maxSum(a[1:])
	return max(pickFirst, skipFirst)
}

func maxSumWithLookup(a []int, lookup map[int]int) int {
	if len(a) == 0 {
		lookup[0] = 0
		return 0
	}
	if len(a) == 1 {
		lookup[1] = a[0]
		return a[0]
	}
	if result, ok := lookup[len(a)]; ok {
		return result
	}
	pickFirst := maxSumWithLookup(a[2:], lookup) + a[0]
	skipFirst := maxSumWithLookup(a[1:], lookup)
	lookup[len(a)] = max(pickFirst, skipFirst)
	return lookup[len(a)]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
