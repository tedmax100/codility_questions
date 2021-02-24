package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3))
}

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	var length = len(A)
	if len(A) == 0 {
		return A
	}

	var temp []int = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		temp[(i+K)%length] = A[i]
	}
	return temp
}
