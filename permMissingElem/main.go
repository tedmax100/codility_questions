package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{2, 3, 1, 5}))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	// 假設N是5, 表示1-6之間會少一個數字
	// 透過梯形公式算得1-6的總和, 來扣除A陣列的元素總和
	var length = len(A)
	var total = (1 + length + 1) * (length + 1) / 2

	for i := 0; i < len(A); i++ {
		total -= A[i]
	}
	return total
}
