package main

import "fmt"

var temp int
var arr = []int{58, 4, 6, 7, 15, 99, 21, 32, 2}
var puc = []int{4, 12, 51, 61, 42, 230}
var zzz = []int{5, 8, 12, 97, 76, 65, 54, 40}

func reverArr(a []int) []int {

	sizeAr := len(a)
	fmt.Println(sizeAr)

	for i, j := 0, sizeAr-1; j > i; i, j = i+1,
		j-1 {

		fmt.Println("ШАГ", j, "заменяем на", i)
		/*
			if j < i {
				fmt.Println(arr[i])
			}
		*/

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

	}
	return a
}

func main() {

	reverArr(arr)
	fmt.Println(arr)
	reverArr(puc)
	fmt.Println(puc)
	reverArr(zzz)
	fmt.Println(zzz)

}
