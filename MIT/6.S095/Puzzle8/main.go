package main

import (
	"fmt"
	"strings"
	"time"
)

var sudoku1 = sudoku{
	size: 9,
	grid: [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{2, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 4, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 3, 8, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}}

var sudoku2 = sudoku{
	size: 9,
	grid: [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{0, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 0, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 3, 0, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}}

var sudoku3 = sudoku{
	size: 9,
	grid: [][]int{
		{1, 0, 5, 7, 0, 2, 6, 3, 8},
		{2, 0, 0, 0, 0, 6, 0, 0, 5},
		{0, 6, 3, 8, 4, 0, 2, 1, 0},
		{0, 5, 9, 2, 0, 1, 3, 8, 0},
		{0, 0, 2, 0, 5, 8, 0, 0, 9},
		{7, 1, 0, 0, 3, 0, 5, 0, 2},
		{0, 0, 4, 5, 6, 0, 7, 2, 0},
		{5, 0, 0, 0, 0, 4, 0, 6, 3},
		{3, 2, 6, 1, 0, 7, 0, 0, 4}}}

var sudoku4 = sudoku{
	size: 9,
	grid: [][]int{
		{8, 5, 0, 0, 0, 2, 4, 0, 0},
		{7, 2, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 4, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 7, 0, 0, 2},
		{3, 0, 5, 0, 0, 0, 9, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 7, 0},
		{0, 1, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 6, 0, 4, 0}}}

var sudoku5 = sudoku{
	size: 9,
	grid: [][]int{{0, 0, 5, 3, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 7, 0, 0, 1, 0, 5, 0, 0},
		{4, 0, 0, 0, 0, 5, 3, 0, 0},
		{0, 1, 0, 0, 7, 0, 0, 0, 6},
		{0, 0, 3, 2, 0, 0, 0, 8, 0},
		{0, 6, 0, 5, 0, 0, 0, 0, 9},
		{0, 0, 4, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 9, 7, 0, 0}}}

func main() {
	for _, s := range []sudoku{sudoku1, sudoku2, sudoku3, sudoku4, sudoku5} {
		start := time.Now()
		s.solve()
		fmt.Printf("spend %f seconds to solve the sodoku!\n", time.Since(start).Seconds())
		fmt.Println(s)
	}
}

type sudoku struct {
	size       int
	grid       [][]int
	backtracks int
}

func (s sudoku) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("backtracks: %d\n", s.backtracks))
	for i := 0; i < s.size; i++ {
		for j := 0; j < s.size; j++ {
			b.WriteString(fmt.Sprintf("%d ", s.grid[i][j]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (s *sudoku) solve() bool {
	i, j := s.findNextCellToFill()
	if i == -1 {
		return true
	}

	for v := 1; v <= s.size; v++ {
		if s.isValid(i, j, v) {
			s.grid[i][j] = v
			if s.solve() {
				return true
			}
			s.grid[i][j] = 0
			s.backtracks++
		}
	}
	return false
}

func (s *sudoku) isValid(m, n int, v int) bool {
	// 检查行,列
	for i := 0; i < s.size; i++ {
		if s.grid[m][i] == v || s.grid[i][n] == v {
			return false
		}
	}
	// 检查9宫
	baseX, baseY := m/3*3, n/3*3
	for i := baseX; i < baseX+3; i++ {
		for j := baseY; j < baseY+3; j++ {
			if s.grid[i][j] == v {
				return false
			}
		}
	}
	return true
}

func (s *sudoku) findNextCellToFill() (i, j int) {
	for i, row := range s.grid {
		for j, cell := range row {
			if cell == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}
