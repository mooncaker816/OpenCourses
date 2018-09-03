package main

import (
	"fmt"
	"strings"
	"time"
)

var grid1 = [][]int{
	{5, 1, 7, 6, 0, 0, 0, 3, 4},
	{2, 8, 9, 0, 0, 4, 0, 0, 0},
	{3, 4, 6, 2, 0, 5, 0, 9, 0},
	{6, 0, 2, 0, 0, 0, 0, 1, 0},
	{0, 3, 8, 0, 0, 6, 0, 4, 7},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 9, 0, 0, 0, 0, 0, 7, 8},
	{7, 0, 3, 4, 0, 0, 5, 6, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0}}

var grid2 = [][]int{
	{5, 1, 7, 6, 0, 0, 0, 3, 4},
	{0, 8, 9, 0, 0, 4, 0, 0, 0},
	{3, 0, 6, 2, 0, 5, 0, 9, 0},
	{6, 0, 0, 0, 0, 0, 0, 1, 0},
	{0, 3, 0, 0, 0, 6, 0, 4, 7},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 9, 0, 0, 0, 0, 0, 7, 8},
	{7, 0, 3, 4, 0, 0, 5, 6, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0}}

var grid3 = [][]int{
	{1, 0, 5, 7, 0, 2, 6, 3, 8},
	{2, 0, 0, 0, 0, 6, 0, 0, 5},
	{0, 6, 3, 8, 4, 0, 2, 1, 0},
	{0, 5, 9, 2, 0, 1, 3, 8, 0},
	{0, 0, 2, 0, 5, 8, 0, 0, 9},
	{7, 1, 0, 0, 3, 0, 5, 0, 2},
	{0, 0, 4, 5, 6, 0, 7, 2, 0},
	{5, 0, 0, 0, 0, 4, 0, 6, 3},
	{3, 2, 6, 1, 0, 7, 0, 0, 4}}

var grid4 = [][]int{
	{8, 5, 0, 0, 0, 2, 4, 0, 0},
	{7, 2, 0, 0, 0, 0, 0, 0, 9},
	{0, 0, 4, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 7, 0, 0, 2},
	{3, 0, 5, 0, 0, 0, 9, 0, 0},
	{0, 4, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 7, 0},
	{0, 1, 7, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 3, 6, 0, 4, 0}}

var grid5 = [][]int{
	{0, 0, 5, 3, 0, 0, 0, 0, 0},
	{8, 0, 0, 0, 0, 0, 0, 2, 0},
	{0, 7, 0, 0, 1, 0, 5, 0, 0},
	{4, 0, 0, 0, 0, 5, 3, 0, 0},
	{0, 1, 0, 0, 7, 0, 0, 0, 6},
	{0, 0, 3, 2, 0, 0, 0, 8, 0},
	{0, 6, 0, 5, 0, 0, 0, 0, 9},
	{0, 0, 4, 0, 0, 0, 0, 3, 0},
	{0, 0, 0, 0, 0, 9, 7, 0, 0}}

func main() {
	for _, grid := range [][][]int{grid1, grid2, grid3, grid4, grid5} {
		s := newSudoku(grid)
		start := time.Now()
		s.solve()
		fmt.Printf("spend %f seconds to solve the sodoku!\n", time.Since(start).Seconds())
		fmt.Println(s)
		s = newSudoku(grid, withStrategy(implication))
		start = time.Now()
		s.solve()
		fmt.Printf("spend %f seconds to solve the sodoku with implication!\n", time.Since(start).Seconds())
		fmt.Println(s)
	}
}

type sudoku struct {
	size       int
	grid       [][]int
	backtracks int
	strategy   strategy
}

type trace struct {
	rIdx, cIdx, value int
}

type strategy func(grid [][]int, root trace) []trace

type option func(s *sudoku)

func withStrategy(strategy strategy) option {
	return func(s *sudoku) {
		s.strategy = strategy
	}
}

func newSudoku(grid [][]int, opts ...option) *sudoku {
	s := new(sudoku)
	if len(grid) == 0 || len(grid) != len(grid[0]) {
		return nil
	}
	s.size = len(grid)
	s.grid = clone(grid)
	s.strategy = naive
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func clone(grid [][]int) [][]int {
	ret := make([][]int, len(grid))
	for i := range grid {
		ret[i] = make([]int, len(grid[i]))
		for j := range grid[i] {
			ret[i][j] = grid[i][j]
		}
	}
	return ret
}

func naive(grid [][]int, root trace) []trace {
	var trs []trace
	grid[root.rIdx][root.cIdx] = root.value
	trs = append(trs, root)
	return trs
}

func implication(grid [][]int, root trace) []trace {
	var trs []trace
	grid[root.rIdx][root.cIdx] = root.value
	trs = append(trs, root)

	seenR := make([]map[int]struct{}, len(grid))
	seenC := make([]map[int]struct{}, len(grid[0]))
	for i, row := range grid {
		seenR[i] = make(map[int]struct{})
		for _, v := range row {
			if v != 0 {
				seenR[i][v] = struct{}{}
			}
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		seenC[j] = make(map[int]struct{})
		for i := 0; i < len(grid); i++ {
			if grid[i][j] != 0 {
				seenC[j][grid[i][j]] = struct{}{}
			}
		}
	}

	// 找出九宫内已填数字
	gNum := (len(grid) / 3) * (len(grid[0]) / 3)
	seenG := make([]map[int]struct{}, gNum)
	implicated := true
	for implicated {
		implicated = false
		for i := 0; i < gNum; i++ {
			if seenG[i] == nil {
				seenG[i] = make(map[int]struct{})
			}
			minR := i / 3 * 3
			maxR := minR + 2
			minC := i % 3 * 3
			maxC := minC + 2
			for r := minR; r <= maxR; r++ {
				for c := minC; c <= maxC; c++ {
					if grid[r][c] != 0 {
						seenG[i][grid[r][c]] = struct{}{}
					}
				}
			}
			for r := minR; r <= maxR; r++ {
				for c := minC; c <= maxC; c++ {
					if grid[r][c] == 0 {
						numbers := findPossibleNumbers(seenR[r], seenC[c], seenG[i])
						if len(numbers) == 1 {
							for num := range numbers {
								grid[r][c] = num
								seenR[r][num] = struct{}{}
								seenC[c][num] = struct{}{}
								seenG[i][num] = struct{}{}
								trs = append(trs, trace{r, c, num})
								implicated = true
							}
						}
					}
				}
			}
		}
	}
	return trs
}

func findPossibleNumbers(rMap, cMap, gMap map[int]struct{}) map[int]struct{} {
	base := make(map[int]struct{})
	for i := 1; i <= 9; i++ {
		base[i] = struct{}{}
	}
	for k := range rMap {
		delete(base, k)
	}
	for k := range cMap {
		delete(base, k)
	}
	for k := range gMap {
		delete(base, k)
	}
	return base
}

func (s *sudoku) undo(trace []trace) {
	for _, tr := range trace {
		s.grid[tr.rIdx][tr.cIdx] = 0
	}
	s.backtracks++
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
			// s.grid[i][j] = v

			trs := s.strategy(s.grid, trace{i, j, v})
			if s.solve() {
				return true
			}
			// s.grid[i][j] = 0
			// s.backtracks++
			s.undo(trs)
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
