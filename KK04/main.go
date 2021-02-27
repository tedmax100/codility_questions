package main

import (
	"fmt"
	"math"
)

func main() {
	root := &Tree{X: 1,
		L: &Tree{
			X: 2,
			L: &Tree{
				X: 1,
			},
			R: &Tree{
				X: 2,
			},
		}, R: &Tree{
			X: 2,
			L: &Tree{
				X: 4,
			},
			R: &Tree{
				X: 1,
			},
		}}
	fmt.Println(Solution(root))

	root2 := &Tree{X: 1,
		R: &Tree{
			X: 2,
			L: &Tree{
				X: 1,
			},
			R: &Tree{
				X: 1,
				L: &Tree{
					X: 4,
				},
			},
		}}
	fmt.Println(Solution(root2))
}

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func Solution(T *Tree) int {
	// write your code in Go 1.4
	if T == nil {
		return 0
	}
	var hash map[int]int = make(map[int]int)
	return search(T, hash)

}

func search(node *Tree, hash map[int]int) int {
	if node == nil {
		return len(hash)
	}
	if _, exist := hash[node.X]; exist == true {
		return len(hash)

	} else {
		hash[node.X] = 1
	}

	var maxPath = math.Max(float64(search(node.L, hash)), float64(search(node.R, hash)))
	if value, exist := hash[node.X]; exist {
		hash[node.X] = value - 1
	}
	if hash[node.X] == 0 {
		delete(hash, node.X)
	}

	return int(maxPath)
}
