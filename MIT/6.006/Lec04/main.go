package main

import (
	"fmt"
	"io"
	"strings"
)

var ints = []int{4, 1, 3, 21, 6, 9, 10, 14, 8, 7}

func main() {
	a := clone(ints)
	heapSort(a)
	fmt.Println(a)
}

func clone(a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}

type heap struct {
	data   []int
	lo, hi int
}

func (h heap) Fprint(w io.Writer) {
	w.Write([]byte(h.String()))
}

func (h heap) String() string {
	var b strings.Builder
	levelCnt := 1
	upLevelCnts := 0
	// totalLevels := math.Ceil(math.Log(float64(h.hi)))
	// spcaeCnt := int(math.Pow(2, totalLevels-1))
	// b.WriteString(strings.Repeat("  ", spcaeCnt))
	for i := 0; i < h.hi; i++ {
		leftSize, rightSize := h.size(i)
		b.WriteString(strings.Repeat("  ", leftSize))
		b.WriteString(fmt.Sprintf("%2d", h.data[i]))
		b.WriteString(strings.Repeat("  ", rightSize))
		if i == h.hi-1 {
			b.WriteString("\n")
			break
		}
		if i+1 >= upLevelCnts+levelCnt {
			upLevelCnts += levelCnt
			levelCnt *= 2
			b.WriteString("\n")
			continue
		}
		b.WriteString("  ")
	}
	return b.String()
}

func (h heap) root() int {
	return h.data[0]
}

func (h *heap) setLowIndex(lo int) {
	if lo >= 0 && lo < len(h.data) {
		h.lo = lo
	}
}

func (h *heap) setHighIndex(hi int) {
	if hi >= 0 && hi < len(h.data) {
		h.hi = hi
	}
}

func (h heap) parent(i int) (idx int) {
	if i < h.lo || i >= h.hi {
		return -1
	}
	return (i+1)/2 - 1
}

func (h heap) size(i int) (leftSize, rightSize int) {
	if left := h.leftChild(i); left != -1 {
		tmpLeft, tmpRight := h.size(left)
		leftSize = tmpLeft + tmpRight + 1
	}
	if right := h.rightChild(i); right != -1 {
		tmpLeft, tmpRight := h.size(right)
		rightSize = tmpLeft + tmpRight + 1
	}
	return
}

func (h heap) leftChild(i int) (idx int) {
	idx = 2*(i+1) - 1
	if idx < h.hi {
		return idx
	}
	return -1
}

func (h heap) rightChild(i int) (idx int) {
	idx = 2 * (i + 1)
	if idx < h.hi {
		return idx
	}
	return -1
}

// O(log n) from the root
func (h heap) maxHeapify(i int) {
	idx := h.getSwapNode(i)
	if idx != -1 {
		h.maxHeapify(idx)
	}
}

func (h heap) getSwapNode(i int) (idx int) {
	leftIdx := h.leftChild(i)
	rightIdx := h.rightChild(i)
	switch {
	case leftIdx == -1 && rightIdx == -1:
		return -1
	case leftIdx == -1 && rightIdx != -1:
		if h.data[rightIdx] > h.data[i] {
			h.data[i], h.data[rightIdx] = h.data[rightIdx], h.data[i]
			return rightIdx
		}
		return -1
	case leftIdx != -1 && rightIdx == -1:
		if h.data[leftIdx] > h.data[i] {
			h.data[i], h.data[leftIdx] = h.data[leftIdx], h.data[i]
			return leftIdx
		}
		return -1
	case leftIdx != -1 && rightIdx != -1:
		if h.data[leftIdx] <= h.data[i] && h.data[rightIdx] <= h.data[i] {
			return -1
		}
		if h.data[leftIdx] >= h.data[rightIdx] {
			h.data[i], h.data[leftIdx] = h.data[leftIdx], h.data[i]
			return leftIdx
		}
		h.data[i], h.data[rightIdx] = h.data[rightIdx], h.data[i]
		return rightIdx
	}
	return -1
}

func (h heap) buildMaxHeap() {
	for i := h.hi / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

// O(nlogn)
func heapSort(a []int) {
	// f, _ := os.Create("test")
	h := heap{a, 0, len(a)}
	for i := len(a) - 1; i >= 0; i-- {
		h.buildMaxHeap()
		// h.Fprint(f)
		// fmt.Println(h)
		h.data[0], h.data[i] = h.data[i], h.data[0]
		h.setHighIndex(i)
	}
	// f.Close()
}
