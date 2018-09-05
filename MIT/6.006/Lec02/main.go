package main

import (
	"fmt"
	"sort"
)

var ints = []int{8, 2, 4, 9, 3, 6}

func main() {
	a := clone(ints)
	insertSort(a)
	fmt.Println(a)
	a = clone(ints)
	binaryInsertSort(a)
	fmt.Println(a)
	a = clone(ints)
	fmt.Println(mergeSort(a))
}

func clone(a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}

// 从第二个元素开始依次和前面已经排好序的序列比较，将该元素插入该序列中，形成新的有序序列
// Θ(n^2)
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j-1 >= 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// 分而治之 Θ(nlogn)，相比 insertSort，mergeSort 需要额外的空间
func mergeSort(a []int) (b []int) {
	if len(a) == 1 {
		return a
	}
	mid := len(a) / 2
	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	out := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}
	out = append(out, left[i:]...)
	out = append(out, right[j:]...)
	return out
}

// 比较 Θ(nlogn)，插入 Θ(n^2)
func binaryInsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		x := a[i]
		insertIndex := sort.Search(len(a[:i]), func(j int) bool { return a[j] >= x })
		copy(a[insertIndex+1:i+1], a[insertIndex:i])
		a[insertIndex] = x
	}
}
