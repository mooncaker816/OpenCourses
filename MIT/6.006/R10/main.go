package main

import "fmt"

func main() {
	p := searchSubMatrix(m0, m1)
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("not found")
	}
	p = bruteForce(m0, m1)
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("not found")
	}
}

var primeRK = 23

var m0 = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 2, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 3, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 4, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

var m1 = [][]int{
	{1, 0, 0, 0},
	{0, 2, 0, 0},
	{0, 0, 3, 0},
	{0, 0, 0, 4},
}

type pos struct {
	r, c int
}

func searchSubMatrix(matrix, sub [][]int) *pos {
	M, N, subM, subN := len(matrix), len(matrix[0]), len(sub), len(sub[0])
	if subM > M || subN > N {
		return nil
	}
	hashSub, powC, powR := hashMatrix(sub)
	// colHash := initColumnHash(matrix, subM)
	colHash := make([]int, N)
	h := 0
	for i := 0; i < subN; i++ {
		cHash := 0
		for j := 0; j < subM; j++ {
			cHash = cHash*primeRK + matrix[j][i]
		}
		colHash[i] = cHash
		h = h*primeRK + colHash[i]
	}
	if h == hashSub && compare(matrix, sub, pos{0, 0}) {
		return &pos{0, 0}
	}
loop:
	for i := subM - 1; i < M; {
		for j := subN; j < N; j++ {
			if i == subM-1 {
				cHash := 0
				for k := 0; k < subM; k++ {
					cHash = cHash*primeRK + matrix[k][j]
				}
				colHash[j] = cHash
			} else {
				colHash[j] = colHash[j]*primeRK + matrix[i][j]
				colHash[j] -= powC * matrix[i-subM][j]
			}
			h = h*primeRK + colHash[j]
			h -= powR * colHash[j-subN]
			if h == hashSub && compare(matrix, sub, pos{i - subM + 1, j - subN + 1}) {
				return &pos{i - subM + 1, j - subN + 1}
			}
		}
		// shift down
		i++
		if i >= M {
			break loop
		}
		h = 0
		for k := 0; k < subN; k++ {
			colHash[k] = colHash[k]*primeRK + matrix[i][k]
			colHash[k] -= powC * matrix[i-subM][k]
			h = h*primeRK + colHash[k]
		}
		if h == hashSub && compare(matrix, sub, pos{i - subM + 1, 0}) {
			return &pos{i - subM + 1, 0}
		}
	}
	return nil
}

// func initColumnHash(a [][]int, size int) []int {
// 	colHash := make([]int, len(a[0]))
// 	for i := 0; i < len(a[0]); i++ {
// 		hash := 0
// 		for j := 0; j < size; j++ {
// 			hash = hash*primeRK + a[j][i]
// 		}
// 		colHash[i] = hash
// 	}
// 	return colHash
// }

func getPow(size int) int {
	pow, sq := 1, primeRK
	for i := size; i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return pow
}

func hashMatrix(a [][]int) (hash, powC, powR int) {
	for i := 0; i < len(a[0]); i++ {
		cHash := 0
		for j := 0; j < len(a); j++ {
			cHash = cHash*primeRK + a[j][i]
		}
		hash = hash*primeRK + cHash
	}
	return hash, getPow(len(a)), getPow(len(a[0]))
}

func compare(m, s [][]int, p pos) bool {
	if len(m)-p.r < len(s) || len(m[0])-p.c < len(s[0]) {
		return false
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[0]); j++ {
			if s[i][j] != m[p.r+i][p.c+j] {
				return false
			}
		}
	}
	return true
}

func bruteForce(matrix, sub [][]int) *pos {
	for i := 0; i < len(matrix)-len(sub); i++ {
		for j := 0; j < len(matrix[0])-len(sub[0]); j++ {
			p := pos{i, j}
			if compare(matrix, sub, p) {
				return &p
			}
		}
	}
	return nil
}
