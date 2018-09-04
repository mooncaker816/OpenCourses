package main

import (
	"fmt"
	"time"
)

var problemMatrix1 = [][]int{
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
var problemMatrix2 = [][]int{
	{0, 0, 9, 8, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 0, 0, 0, 0},
	{0, 2, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
}

func main() {
	var matrixes = [][][]int{problemMatrix1, problemMatrix2}
	for _, m := range matrixes {
		pb := newProblem(m)
		fmt.Println("alog1:")
		start := time.Now()
		e := peakFind1(pb)
		fmt.Printf("spend %d nanoseconds to get a peak: %v\n", time.Since(start).Nanoseconds(), e)
		fmt.Println("alog2:")
		start = time.Now()
		e = peakFind2(pb)
		fmt.Printf("spend %d nanoseconds to get a peak: %v\n", time.Since(start).Nanoseconds(), e)
		fmt.Println("alog4:")
		start = time.Now()
		e = peakFind4(pb)
		fmt.Printf("spend %d nanoseconds to get a peak: %v\n", time.Since(start).Nanoseconds(), e)
	}
}

type problem struct {
	matrix [][]int
	scope  scope
	cnt    int
}

func newProblem(m [][]int) *problem {
	var pb problem
	pb.matrix = m
	pb.scope = scope{0, len(m) - 1, 0, len(m[0]) - 1}
	pb.cnt = 0
	return &pb
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

func peakFind1(pb *problem) (e *element) {
	rNums, cNums := pb.scope.size()
	if rNums == 0 || cNums == 0 {
		return nil
	}
	midR, _ := pb.scope.mid()
	rowMax := findRowMax(pb.matrix, midR)
	if rNums == 1 {
		return rowMax
	}
	if pb.scope.rowInScope(midR-1) && rowMax.v < pb.matrix[midR-1][rowMax.c] {
		pb.scope.setBottom(midR - 1)
		return peakFind1(pb)
	} else if pb.scope.rowInScope(midR+1) && rowMax.v < pb.matrix[midR+1][rowMax.c] {
		pb.scope.setTop(midR + 1)
		return peakFind1(pb)
	} else {
		return rowMax
	}
}

func getSize(m [][]int) (rows, columns int) {
	return len(m), len(m[0])
}

func findRowMax(m [][]int, row int) *element {
	idx, val := 0, m[row][0]
	for i := 1; i < len(m[row]); i++ {
		if m[row][i] > val {
			val = m[row][i]
			idx = i
		}
	}
	return &element{row, idx, val}
}

func findColumnMax(m [][]int, column int) *element {
	idx, val := 0, m[0][column]
	for i := 1; i < len(m); i++ {
		if m[i][column] > val {
			val = m[i][column]
			idx = i
		}
	}
	return &element{idx, column, val}
}

func peakFind2(pb *problem) (e *element) {
	rNums, cNums := getSize(pb.matrix)
	if rNums == 0 || cNums == 0 {
		return nil
	}
	return peak2R(pb.matrix, &element{0, 0, pb.matrix[0][0]})
}

func peak2R(m [][]int, current *element) (e *element) {
	max := current
	changed := false
	if new := current.up(m); new != nil {
		if max.v < new.v {
			max = new
			changed = true
		}
	}
	if new := current.down(m); new != nil {
		if max.v < new.v {
			max = new
			changed = true
		}
	}
	if new := current.left(m); new != nil {
		if max.v < new.v {
			max = new
			changed = true
		}
	}
	if new := current.right(m); new != nil {
		if max.v < new.v {
			max = new
			changed = true
		}
	}
	if !changed {
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
	if e.c <= 0 {
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
	if e.r >= rNums-1 {
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
	if e.c >= cNums-1 {
		return nil
	}
	new.r = e.r
	new.c = e.c + 1
	new.v = m[new.r][new.c]
	return &new
}

func peakFind4(pb *problem) *element {
	rNums, cNums := pb.scope.size()
	if rNums == 0 || cNums == 0 {
		return nil
	}
	if rNums == 1 {
		return findRowMax(pb.matrix, pb.scope.top)
	}
	if cNums == 1 {
		return findColumnMax(pb.matrix, pb.scope.left)
	}
	midR, midC := pb.scope.mid()
	switch pb.cnt % 2 {
	case 0:
		rowMax := findRowMax(pb.matrix, midR)
		if pb.scope.rowInScope(midR-1) && rowMax.v < pb.matrix[midR-1][rowMax.c] {
			pb.scope.setBottom(midR - 1)
			pb.cnt++
			return peakFind4(pb)
		} else if pb.scope.rowInScope(midR+1) && rowMax.v < pb.matrix[midR+1][rowMax.c] {
			pb.scope.setTop(midR + 1)
			pb.cnt++
			return peakFind4(pb)
		} else {
			return rowMax
		}
	case 1:
		columnMax := findColumnMax(pb.matrix, midC)
		if pb.scope.columnInScope(midC-1) && columnMax.v < pb.matrix[columnMax.r][midC-1] {
			pb.scope.setRight(midC - 1)
			pb.cnt++
			return peakFind4(pb)
		} else if pb.scope.columnInScope(midC+1) && columnMax.v < pb.matrix[columnMax.r][midC+1] {
			pb.scope.setLeft(midC + 1)
			pb.cnt++
			return peakFind4(pb)
		} else {
			return columnMax
		}
	default:
		return nil
	}
}

type scope struct {
	top, bottom, left, right int
}

func (s *scope) rowInScope(r int) bool {
	return s.top <= r && s.bottom >= r
}

func (s *scope) columnInScope(c int) bool {
	return s.left <= c && s.right >= c
}

func (s *scope) setTop(top int) {
	s.top = top
}

func (s *scope) setBottom(bottom int) {
	s.bottom = bottom
}

func (s *scope) setLeft(left int) {
	s.left = left
}

func (s *scope) setRight(right int) {
	s.right = right
}

func (s *scope) mid() (r, c int) {
	rows, columns := s.size()
	return s.top + rows/2, s.left + columns/2
}

func (s *scope) size() (rows, columns int) {
	return s.bottom - s.top + 1, s.right - s.left + 1
}
