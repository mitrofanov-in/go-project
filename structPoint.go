package main

import (
	"fmt"
)

type TestStruct struct {
	person string
	age    int
}

func CreateStructP(n string, i int) *TestStruct {
	if i > 300 {
		i = 0
	}
	return &TestStruct{n, i}
}

func CreateStructNoP(n string, i int) TestStruct {
	if i > 300 {
		i = 0
	}
	return TestStruct{n, i}
}

func main() {
	s1 := CreateStructP("Test", 17)
	s2 := CreateStructNoP("Tset", 71)
	fmt.Println("test", *s1, s2)
}
