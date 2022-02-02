package main

import (
	"testing"

	"math/rand"

	"github.com/stretchr/testify/assert"
)

func Test_MaxProfit(t *testing.T) {
	var testCases = []struct {
		in  []int
		out int
	}{
		{
			in:  []int{7, 1, 5, 3, 6, 4},
			out: 7,
		}, {
			in:  []int{1, 2, 3, 4, 5},
			out: 4,
		}, {
			in:  []int{7, 6, 4, 3, 1},
			out: 0,
		}, {
			in:  []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			out: 0,
		}, {
			in:  []int{2, 2, 3, 2, 2, 2, 2, 2},
			out: 1,
		},
	}

	for idx := range testCases {
		assert.Equal(t, testCases[idx].out, maxProfit(testCases[idx].in))
	}
}

func Benchmark_MaxProfit(b *testing.B) {
	rand.Seed(int64(b.N))

	var inArr = make([][]int, 10)
	for i := 0; i < 10; i++ {
		inArr[i] = rand.Perm(rand.Intn(10))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var arr = inArr[rand.Intn(10)]
		maxProfit(arr)
	}
}
