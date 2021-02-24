package main

import "fmt"

func main() {
	fmt.Println(Solution(10, 85, 30))
}

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	var res = (Y - X) / D

	if (Y-X)%D == 0 {
		return res
	}
	return res + 1
}
