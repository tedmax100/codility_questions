package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{0, 1, 0, 1, 1}))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var lengthA = len(A)
	var count = 0
	var addBase = 0

	for i := 0; i < lengthA; i++ {
		if A[i] == 0 {
			addBase++
		} else {
			count += addBase
		}
	}
	if count > 1000000000 {
		return -1
	}
	return count
}
