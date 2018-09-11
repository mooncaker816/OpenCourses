package main

import "testing"

var funcs = []struct {
	name string
	f    func(m1, m2 [][]int) *pos
}{
	{"RabinKarp", searchSubMatrix},
	{"BruteForce", bruteForce},
}

func BenchmarkRabinKarp(b *testing.B) {
	for _, f := range funcs {
		b.Run(f.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f.f(m0, m1)
			}
		})
	}
}
