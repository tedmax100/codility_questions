package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution([]int{3, 1, 2, 4, 3}))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var length = len(A)

	var after []int = make([]int, len(A))

	after[length-1] = A[length-1]

	for i := length - 2; i >= 0; i-- {
		after[i] = after[i+1] + A[i]
	}

	var min int = int(math.Abs(float64((after[0] - after[1]*2))))

	for P := 2; P < length; P++ {
		temp := int(math.Abs(float64(after[0] - after[P]*2)))

		if min > temp {
			min = temp
		}
	}
	return min
}
