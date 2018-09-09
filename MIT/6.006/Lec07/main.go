package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ints = []int{4, 1, 3, 21, 6, 9, 10, 14, 8, 7}

func main() {
	fmt.Println(countingSort(clone(ints), 30))
	fmt.Println(countingSort2(clone(ints), 30))
	input := make([]int64, 10)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		input[i] = rand.Int63()
		// fmt.Println(input[i])
	}
	for _, v := range numberSortBasedCounting(input) {
		fmt.Println(v)
	}
}

// k 是 a 中元素的上界，0<=a[i]<k
func countingSort(a []int, k int) []int {
	counts := make([]int, k)
	b := make([]int, len(a))
	// counts[v] 表示 a 中值为 v 的元素有多少个
	for _, v := range a {
		counts[v]++
	}
	// counts[v] 表示值 <= v 的元素的个数，
	// 也就是说原始 a 中值为 v 的元素在输出 b 中的结束索引为 counts[v]-1（a 中可能有多个元素的值都为 v）
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	// 为了保证稳定性（相同值的元素顺序不变），在知道 a 中相同值的元素在 b 中的结束索引的条件下，我们需要从 a 的末尾开始循环
	for i := len(a) - 1; i >= 0; i-- {
		b[counts[a[i]]-1] = a[i]
		counts[a[i]]--
	}
	return b
}

func countingSort2(a []int, k int) []int {
	counts := make([]int, k)
	b := make([]int, len(a))
	// counts[v] 表示 a 中值为 v 的元素有多少个
	for _, v := range a {
		counts[v]++
	}
	// 等效地也可统计原始 a 中值 < v 的元素的个数到 counts[v]，这样我们就获得了在输出 b 中的值为 v 的元素的起始索引为 counts[v]
	total := len(a)
	for i := len(counts) - 1; i >= 0; i-- {
		total -= counts[i]
		counts[i] = total
	}
	// OR:
	// totalLess := 0
	// for i := 0; i < len(counts); i++ {
	// 	tmp := counts[i]
	// 	counts[i] = totalLess
	// 	totalLess += tmp
	// }

	for _, v := range a {
		b[counts[v]] = v
		counts[v]++
	}
	return b
}

func clone(a []int) []int {
	b := make([]int, len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}

func numberSortBasedCounting(a []int64) []int64 {
	// fmt.Println(tmp)
	for i := 19; i >= 0; i-- {
		tmp := decimalToDigits(a, 10)
		digits := make([]int, len(a))
		for j := 0; j < len(a); j++ {
			digits[j] = tmp[j][i]
		}
		a = radixSort(a, digits)
		// fmt.Println("start")
		// for _, v := range a {
		// 	fmt.Println(v)
		// }
		// fmt.Println("end")
	}
	return a
}

func decimalToDigits(a []int64, base int64) [][20]int {
	b := make([][20]int, len(a))
	for i, n := range a {
		var c [20]int
		for j := 19; j >= 0; j-- {
			c[j] = int(n % base)
			n /= base
		}
		b[i] = c
	}
	return b
}

func radixSort(a []int64, digits []int) []int64 {
	counts := make([]int, 10)
	b := make([]int64, len(a))
	for _, v := range digits {
		counts[v]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		b[counts[digits[i]]-1] = a[i]
		counts[digits[i]]--
	}
	return b
}
