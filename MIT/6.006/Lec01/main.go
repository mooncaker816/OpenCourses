package main

import (
	"fmt"
	"time"
)

var problemMatrix = [][]int{
	{4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2},
	{5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3},
	{6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4},
	{7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5},
	{8, 9, 10, 11, 12, 11, 10, 9, 8, 7, 6},
	{7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5},
	{6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4},
	{5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3},
	{4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2},
	{3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1},
	{2, 3, 4, 5, 6, 5, 4, 3, 2, 1, 0},
}

func main() {
	fmt.Println("alog1:")
	start := time.Now()
	e := peakFind1(problemMatrix)
	fmt.Printf("spend %d nanoseconds to get a peak: %v\n", time.Since(start).Nanoseconds(), e)
	fmt.Println("alog2:")
	start = time.Now()
	e = peakFind2(problemMatrix)
	fmt.Printf("spend %d nanoseconds to get a peak: %v\n", time.Since(start).Nanoseconds(), e)
}

type element struct {
	r, c, v int
}

func (e *element) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("row: %d,column: %d,value: %d", e.r, e.c, e.v)
}

func peakFind1(m [][]int) (e *element) {
	rNums, cNums := getSize(m)
	if rNums == 0 || cNums == 0 {
		return nil
	}
	midR := rNums / 2
	cIdx, max := find1DMax(m[midR])
	if rNums == 1 {
		return &element{0, cIdx, max}
	}
	if max < m[midR-1][cIdx] {
		return peakFind1(m[:midR])
	} else if midR+1 < rNums && max < m[midR+1][cIdx] {
		tmp := peakFind1(m[midR+1:])
		return &element{tmp.r + midR + 1, tmp.c, tmp.v}
	} else {
		return &element{midR, cIdx, max}
	}
}

func getSize(m [][]int) (rows, columns int) {
	return len(m), len(m[0])
}

func find1DMax(a []int) (idx, val int) {
	idx, val = 0, a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > val {
			val = a[i]
			idx = i
		}
	}
	return
}

func peakFind2(m [][]int) (e *element) {
	rNums, cNums := getSize(m)
	if rNums == 0 || cNums == 0 {
		return nil
	}
	return peak2R(m, &element{0, 0, m[0][0]})
}

func peak2R(m [][]int, current *element) (e *element) {
	max := current
	if new := current.up(m); new != nil {
		if max.v < new.v {
			max = new
		}
	}
	if new := current.down(m); new != nil {
		if max.v < new.v {
			max = new
		}
	}
	if new := current.left(m); new != nil {
		if max.v < new.v {
			max = new
		}
	}
	if new := current.right(m); new != nil {
		if max.v < new.v {
			max = new
		}
	}
	if *max == *current {
		return current
	}
	return peak2R(m, max)
}

func (e *element) up(m [][]int) *element {
	var new element
	if e.r <= 0 {
		return nil
	}
	new.r = e.r - 1
	new.c = e.c
	new.v = m[new.r][new.c]
	return &new
}

func (e *element) left(m [][]int) *element {
	var new element
	if e.c == 0 {
		return nil
	}
	new.r = e.r
	new.c = e.c - 1
	new.v = m[new.r][new.c]
	return &new
}

func (e *element) down(m [][]int) *element {
	var new element
	rNums, _ := getSize(m)
	if e.r == rNums-1 {
		return nil
	}
	new.r = e.r + 1
	new.c = e.c
	new.v = m[new.r][new.c]
	return &new
}

func (e *element) right(m [][]int) *element {
	var new element
	_, cNums := getSize(m)
	if e.c == cNums-1 {
		return nil
	}
	new.r = e.r
	new.c = e.c + 1
	new.v = m[new.r][new.c]
	return &new
}
