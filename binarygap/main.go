package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// write your code in Go 1.4
	// 將數字轉成二進制文字
	binStr := strconv.FormatInt(int64(N), 2)
	// 從左至右, 依序拿出來判斷, =0 則curMaxGap++, =1 則curMaxGap與maxGap比大小, 且curMaxGap=0
	var curMaxGap = 0
	var maxGap = 0

	// 第一個字母一定是1, 不用判斷
	for idx := 1; idx < len(binStr); idx++ {
		if (string([]rune(binStr)[idx])) == "0" {
			curMaxGap++
			// 如果最後一個字母是0 , 這整段不用納入
			if idx == len(binStr)-1 {
				curMaxGap = 0
			}
		}
		if (string([]rune(binStr)[idx])) == "1" {
			if curMaxGap > maxGap {
				maxGap = curMaxGap
			}
			curMaxGap = 0
		}
	}
	if curMaxGap > maxGap {
		maxGap = curMaxGap
	}
	return maxGap
}

func main() {
	fmt.Println(Solution(32))
}
