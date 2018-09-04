package main

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func(pb *problem) *element
}{
	{"algorithm1", peakFind1},
	{"algorithm2", peakFind2},
	{"algorithm4", peakFind4},
}

var matrixes = [][][]int{problemMatrix1, problemMatrix2}

func BenchmarkPeakFinding(b *testing.B) {
	for i, m := range matrixes {
		pb := newProblem(m)
		for _, f := range funcs {
			b.Run(fmt.Sprintf("%s/matrix-%d:", f.name, i+1), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.f(pb)
				}
			})
		}
	}
}
