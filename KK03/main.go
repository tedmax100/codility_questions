package main

import "fmt"

func main() {
	fmt.Println(Solution("babba"))
	fmt.Println(Solution("baaabbabbb"))
	fmt.Println(Solution("abaaaa"))
}

func Solution(S string) int {
	// write your code in Go 1.4
	if len(S) == 0 {
		return 0
	}

	var maxCount = 0
	var startIdx = 0

	var preC = string([]rune(S)[0])
	var preCount = 1

	for i := 1; i < len(S); i++ {
		var c = string([]rune(S)[i])
		if c == preC {
			if preCount == 2 {
				var count = i - startIdx
				if count > maxCount {
					maxCount = count
				}

				startIdx = i - 1
				for i+1 < len(S) && string([]rune(S)[i+1]) == preC {
					i++
					startIdx++
				}
			} else {
				preCount++
			}
		} else {
			preC = c
			preCount = 1
		}
	}
	var count = len(S) - startIdx
	if count > maxCount {
		maxCount = count
	}
	return maxCount
	/* var maxCount = 3

	var length = len(S)

	if length < maxCount {
		return length
	}
	var count = 1
	var result = 1

	for i := 1; i < length-1; i++ {
		if string([]rune(S)[i-1]) == string([]rune(S)[i]) && string([]rune(S)[i+1]) == string([]rune(S)[i]) {
			if count+1 > result {
				result = count + 1
			}
			count = 1
		} else {
			count++
		}
	} */

}
