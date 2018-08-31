// #Given a list of intervals when celebrities will be at the party
// #Output is the time that you want to go the party when the maximum number of
// #celebrities are still there.
package main

import (
	"fmt"
	"sort"
)

var sched = []schedule{{6, 8}, {6, 12}, {6, 7}, {7, 8}, {7, 10}, {8, 9}, {8, 10}, {9, 12},
	{9, 10}, {10, 11}, {10, 12}, {11, 12}}
var sched2 = []schedule{{6.0, 8.0}, {6.5, 12.0}, {6.5, 7.0}, {7.0, 8.0}, {7.5, 10.0}, {8.0, 9.0},
	{8.0, 10.0}, {9.0, 12.0}, {9.5, 10.0}, {10.0, 11.0}, {10.0, 12.0}, {11.0, 12.0}}

var sched3 = []schedulew{
	{6.0, 8.0, 2}, {6.5, 12.0, 1}, {6.5, 7.0, 2}, {7.0, 8.0, 2}, {7.5, 10.0, 3}, {8.0, 9.0, 2},
	{8.0, 10.0, 1}, {9.0, 12.0, 2}, {9.5, 10.0, 4}, {10.0, 11.0, 2}, {10.0, 12.0, 3}, {11.0, 12.0, 7},
}

type schedule struct {
	start, end float32
}

type schedulew struct {
	start, end, weight float32
}

func main() {
	bestTimeToPartyBF(sched)
	bestTimeToPartySort(sched)
	exercise2(sched)
	bestTimeToPartySort2(sched, 7, 9)
	bestTimeToPartySort2(sched, 10, 12)

	bestTimeToPartyBF(sched2)
	bestTimeToPartySort(sched2)
	exercise2(sched2)
	bestTimeToPartySort2(sched2, 7, 9)
	bestTimeToPartySort2(sched2, 10, 12)

	exercise3(sched3)
}

// Brute force algorithm
func bestTimeToPartyBF(scheds []schedule) {
	fmt.Println("Brute force alg:")
	start := scheds[0].start
	end := scheds[0].end
	for i := 1; i < len(scheds); i++ {
		if scheds[i].start < start {
			start = scheds[i].start
		}
		if scheds[i].end > end {
			end = scheds[i].end
		}
	}
	m := make(map[float32]int)
	seen := make(map[float32]struct{})
	for _, sc := range scheds {
		tm := sc.start
		if _, ok := seen[tm]; ok {
			continue
		}
		seen[tm] = struct{}{}
		for _, cp := range scheds {
			if tm >= cp.start && tm < cp.end {
				m[tm]++
			}
		}
	}
	maxcount := -1
	time := float32(-1.0)
	for i, v := range m {
		if v > maxcount {
			maxcount = v
			time = i
		}
	}
	fmt.Printf("Best time to attend the party is at %.1f o'clock: %d celebrities will be attending!\n", time, maxcount)
}

type halfsched struct {
	time  float32
	leave bool
}

type halfschedw struct {
	time   float32
	leave  bool
	weight float32
}

func bestTimeToPartySort(scheds []schedule) {
	fmt.Println("Sort alg:")
	times := buildHalfTimes(scheds)
	maxcount, time := chooseTime(times)
	fmt.Printf("Best time to attend the party is at %.1f o'clock: %d celebrities will be attending!\n", time, maxcount)
}

func bestTimeToPartySort2(scheds []schedule, ystart, yend float32) {
	fmt.Printf("Sort alg with own schedule [%.1f,%.1f):\n", ystart, yend)
	times := buildHalfTimes(scheds)
	maxcount, time := chooseTimeByOwn(times, ystart, yend)
	fmt.Printf("Best time to attend the party is at %.1f o'clock: %d celebrities will be attending!\n", time, maxcount)
}

func buildHalfTimes(scheds []schedule) []halfsched {
	var times []halfsched
	for _, sc := range scheds {
		times = append(times, halfsched{sc.start, false}, halfsched{sc.end, true})
	}
	sort.Slice(times, func(i, j int) bool { return times[i].time < times[j].time })
	return times
}

func chooseTime(times []halfsched) (maxcount int, time float32) {
	count := 0
	prev := times[0].time
	for _, tm := range times {
		if tm.time > prev {
			if count > maxcount {
				maxcount = count
				time = prev
			}
			prev = tm.time
		}
		if tm.leave {
			count--
		} else {
			count++
		}
	}
	return
}

func chooseTimeByOwn(times []halfsched, ystart, yend float32) (maxcount int, time float32) {
	count := 0
	prev := times[0].time
	for _, tm := range times {
		if tm.time > prev {
			if count > maxcount && ystart <= prev && prev < yend {
				maxcount = count
				time = prev
			}
			prev = tm.time
		}
		if tm.leave {
			count--
		} else {
			count++
		}
	}
	return
}

// exercise 2
func exercise2(scheds []schedule) {
	fmt.Println("exercise2:")
	m := make(map[schedule]int)
	for _, sc := range scheds {
		for _, cp := range scheds {
			if sc.start >= cp.start && sc.start < cp.end {
				m[sc]++
			}
		}
	}
	maxcount := -1
	time := float32(-1.0)
	for k, v := range m {
		if v > maxcount {
			maxcount = v
			time = k.start
		}
	}
	fmt.Printf("Best time to attend the party is at %.1f o'clock: %d celebrities will be attending!\n", time, maxcount)
}

// exercise 3
func exercise3(scheds []schedulew) {
	fmt.Println("exercise3:")
	times := buildWeightHalfTimes(scheds)
	maxweight, time := chooseTimeByWeight(times)
	fmt.Printf("Best time to attend the party is at %.1f o'clock where the weight of attending celebrities is %.1f and maximum!\n", time, maxweight)
}

func buildWeightHalfTimes(scheds []schedulew) []halfschedw {
	var times []halfschedw
	for _, sc := range scheds {
		times = append(times, halfschedw{sc.start, false, sc.weight}, halfschedw{sc.end, true, sc.weight})
	}
	sort.Slice(times, func(i, j int) bool { return times[i].time < times[j].time })
	return times
}

func chooseTimeByWeight(times []halfschedw) (maxweight, time float32) {
	var weight float32
	prev := times[0].time
	for _, tm := range times {
		if tm.time > prev {
			if weight > maxweight {
				maxweight = weight
				time = prev
			}
			prev = tm.time
		}
		if tm.leave {
			weight -= tm.weight
		} else {
			weight += tm.weight
		}
	}
	return
}
