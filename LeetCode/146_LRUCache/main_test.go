package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input_Struct struct {
	Actions []string
	Values  [][]int
}

func createIntPtr(x int) *int {
	return &x
}

func Test_MaxProfit(t *testing.T) {
	var testCases = []struct {
		in  Input_Struct
		out []*int
	}{
		{
			in: Input_Struct{
				Actions: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				Values:  [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			out: []*int{nil, nil, nil, createIntPtr(1), nil, createIntPtr(-1), nil, createIntPtr(-1), createIntPtr(3), createIntPtr(4)},
		},
	}

	for idx := range testCases {
		excepted := testCases[idx].out[1:]
		testResult := Process(testCases[idx].in)
		for idx2, value := range excepted {
			assert.Equal(t, value, testResult[idx2])
		}
	}
}
