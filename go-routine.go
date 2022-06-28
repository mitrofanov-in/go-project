package main

import (
	"fmt"
	"time"
)

var result int = 0

func j(int) {
	for i := 0; i < 10; i++ {
		result += i
	}
	fmt.Print(result)
}

func main() {
	start := time.Now()
	duration := time.Since(start)

	go j(1)
	fmt.Println(result)

	go j(101)
	fmt.Println(result)

	fmt.Println(duration)
}
