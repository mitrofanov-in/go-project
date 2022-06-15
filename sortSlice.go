//go:build windows && amd64

package main

import (
	"fmt"
	"sort"
)

func main() {

	type jStruct struct {
		name string
		age  int
	}

	test_slice := make([]jStruct, 0)
	names := []string{"frenk", "jon", "shon", "svon", "clon", "ton", "ozon"}

	fmt.Println(names)
	ages := []int{20, 21, 21, 19, 18, 16, 62}

	a := len(names)
	b := len(ages)

	if a == b {
		for i := 0; i < len(names); i++ {
			test_slice = append(test_slice, jStruct{names[i], ages[i]})
		}

		for _, y := range test_slice {
			fmt.Println(y)
		}
		sort.Slice(test_slice, func(i, j int) bool {
			return test_slice[i].age < test_slice[j].age
		})
		fmt.Println(test_slice)
	} else {
		fmt.Println("Error slice not success")
	}

}
