package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 9}))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	var length = len(A)
	var number = 0

	for i := 0; i < length; i++ {
		number ^= A[i]
	}
	return number
}
