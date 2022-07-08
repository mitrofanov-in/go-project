package main

import (
	"fmt"
)

var data = []int{1, 2, 3, 4, 5, 6, 7}
var arr = []int{1, 12, 5, 67, 99}

func tRout(c chan []int) {
	c <- data
	res := <-c
	fmt.Println(res)
}

func main() {

	c := make(chan []int, len(data))

	go tRout(c)

	c <- arr

	result := <-c

	fmt.Println(result)

}
