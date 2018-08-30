// #You Will All Conform
// #Input is a vector of F's and B's, in terms of forwards and backwards caps
// #Output is a set of commands (printed out) to get either all F's or all B's
// #Fewest commands are the goal

package main

import "fmt"

var caps = []byte{'F', 'F', 'B', 'B', 'B', 'F', 'B', 'B', 'B', 'F', 'F', 'B', 'F'}
var cap2 = []byte{'F', 'F', 'B', 'B', 'B', 'F', 'B', 'B', 'B', 'F', 'F', 'F', 'F'}
var cap3 = []byte{'F', 'F', 'B', 'H', 'B', 'F', 'B', 'B', 'B', 'F', 'H', 'F', 'F'}

func main() {
	fmt.Println("naive method:")
	naiveConform(caps)
	naiveConform(cap2)
	fmt.Println("optimized naive method:")
	optConform(caps)
	optConform(cap2)
	fmt.Println("one Pass method:")
	onePassConform(caps)
	onePassConform(cap2)
	fmt.Println("one Pass with pretty printing method:")
	onePassPrettyPrintConform(caps)
	onePassPrettyPrintConform(cap2)
	fmt.Println("bypass H method:")
	bypassHConform(cap3)
}

type interval struct {
	begin, end int
	cap        byte
}

func naiveConform(caps []byte) {
	start := 0
	fcnt, bcnt := 0, 0
	var intervals []interval
	for i := 1; i < len(caps); i++ {
		if caps[start] != caps[i] {
			intervals = append(intervals, interval{start, i - 1, caps[start]})
			if caps[start] == 'F' {
				fcnt++
			} else {
				bcnt++
			}
			start = i
		}
	}
	intervals = append(intervals, interval{start, len(caps) - 1, caps[start]})
	if caps[start] == 'F' {
		fcnt++
	} else {
		bcnt++
	}
	var flip byte
	if fcnt < bcnt {
		flip = 'F'
	} else {
		flip = 'B'
	}
	for _, interval := range intervals {
		if interval.cap == flip {
			fmt.Printf("People in positions %d thru %d flip your caps!\n", interval.begin, interval.end)
		}
	}
}

func optConform(caps []byte) {
	start := 0
	fcnt, bcnt := 0, 0
	var intervals []interval
	// 加一个哨兵，从而可以直接处理最后一个 interval
	caps = append(caps, 'x')
	for i := 1; i < len(caps); i++ {
		if caps[start] != caps[i] {
			intervals = append(intervals, interval{start, i - 1, caps[start]})
			if caps[start] == 'F' {
				fcnt++
			} else {
				bcnt++
			}
			start = i
		}
	}
	var flip byte
	if fcnt < bcnt {
		flip = 'F'
	} else {
		flip = 'B'
	}
	for _, interval := range intervals {
		if interval.cap == flip {
			fmt.Printf("People in positions %d thru %d flip your caps!\n", interval.begin, interval.end)
		}
	}
}

// F,B 的 interval 肯定是交错的，所以第一个 interval 的类型的个数肯定等于第二个 interval 的类型的个数或者+1
// 所以我们只要翻转与第二个 interval 的类型相同的 interval 即可
func onePassConform(caps []byte) {
	caps = append(caps, caps[0])
	for i := 0; i < len(caps)-1; i++ {
		if caps[i] != caps[i+1] {
			if caps[i+1] != caps[0] {
				fmt.Printf("People in positions %d ", i+1)
			} else {
				fmt.Printf("thru %d flip your caps!\n", i)
			}
		}
	}
}

// exercise 2
func onePassPrettyPrintConform(caps []byte) {
	if len(caps) == 0 {
		fmt.Println("empty list,no command requires!")
		return
	}
	caps = append(caps, caps[0])
	var start int
	for i := 0; i < len(caps)-1; i++ {
		if caps[i] != caps[i+1] {
			if caps[i+1] != caps[0] {
				start = i + 1
			} else {
				if i == start {
					fmt.Printf("People in position %d flip your cap!\n", i)
				} else {
					fmt.Printf("People in positions %d thru %d flip your caps!\n", start, i)
				}
			}
		}
	}
}

// exercise 3
func bypassHConform(caps []byte) {
	if len(caps) == 0 {
		fmt.Println("empty list,no command requires!")
		return
	}
	var i int
	// 首先获得不需要 flip 的类型，添加至末尾作为哨兵
	for i = 0; i < len(caps); i++ {
		if caps[i] != 'H' {
			break
		}
	}
	noAction := caps[i]
	caps = append(caps, noAction)
	var start int
	for ; i < len(caps)-1; i++ {
		if caps[i] != caps[i+1] {
			if caps[i] == 'H' && caps[i+1] == noAction ||
				caps[i] == noAction && caps[i+1] == 'H' {
				continue
			}
			if caps[i+1] != noAction && caps[i+1] != 'H' {
				start = i + 1
			} else {
				if i == start {
					fmt.Printf("People in position %d flip your cap!\n", i)
				} else {
					fmt.Printf("People in positions %d thru %d flip your caps!\n", start, i)
				}
			}
		}
	}
}
