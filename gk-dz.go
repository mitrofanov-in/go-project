package main

import "fmt"

var s int = 10000
var v1 int = 1
var v2 int = 2
var v3 int = 5
var count int = 0
var flagR int = 1
var flagL int = 1
var tsum int = 0

func main() {

	for s, tsum := 10000, 0; s > 10; s = s - tsum {

		if flagR == 1 {
			t1c := s / (v2 + v3)
			tsum = s - ((v1 + v2) * t1c)
			flagR = 0
			flagL = 1
			fmt.Println("Ко второму", t1c, tsum, s)
			count++
			continue
		}

		if flagL == 1 {
			t2c := s / (v1 + v3)
			tsum = s - ((v1 + v2) * t2c)
			flagL = 0
			flagR = 1
			fmt.Println("К первому", t2c, tsum, s)
			count++
			continue
		}
	}
	fmt.Println(count)
}
