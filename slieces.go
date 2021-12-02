package main

import (
	"fmt"
)

func main() {
	test_slice := []int{}

	test_slice = append(test_slice, 101)
	tmp := make([]int, 10)

	for _, i := range tmp {
		fmt.Println(tmp[i])
	}
	fmt.Println(test_slice[0])

	twoS := make([][]int, 6)

	for i := 0; i < len(twoS); i++ {
		for j := 0; j < 2; j++ {
			twoS[i] = append(twoS[i], i+j)
		}
	}

	for _, s := range twoS {
		for x, z := range s {
			fmt.Println(x, z)
		}
		fmt.Println()
	}
}
