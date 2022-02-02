package main

import "fmt"

func main() {
	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4}))
}

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	var count = 0

	var res []int = make([]int, len(A))

	if X > len(A) {
		return -1
	}

	if len(A) == 1 && X == 1 && A[0] == 1 {
		return 0
	}

	if len(A) == 1 && X == 1 && A[0] > 1 {
		return -1
	}

	for i := 0; i < len(A); i++ {
		if A[i] <= X {
			if res[A[i]-1] == 0 {
				res[A[i]-1] = 1
				count++
			}
			if count == X {
				return i
			}
		}
	}
	return -1
}
