package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution("50552"))
	fmt.Println(Solution("10101"))
	fmt.Println(Solution("88"))
}

func Solution(S string) int {
	// write your code in Go 1.4
	var max int64 = 0

	for i := 0; i < len(S)-1; i++ {
		strNum := string([]rune(S)[i : i+2])
		num, _ := strconv.ParseInt(strNum, 10, 64)

		if num > max {
			max = num
		}
	}

	return int(max)
}
