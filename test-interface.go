package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type square struct {
	x float64
}

type circle struct {
	r float64
}

func (s square) Area() float64 {
	return s.x * s.x
}

func (s square) Perimeter() float64 {
	return 4 * s.x
}

func (s circle) Area() float64 {
	return s.r * s.r * math.Pi
}

func (s circle) Perimeter() float64 {
	return 2 * s.r * math.Pi
}

func Calculate(z shape) {
	_, ok := z.(circle)
	if ok {
		fmt.Println("Is a circle")
	}

	v, ok := z.(square)
	if ok {
		fmt.Println("Is a square", v)
	}
	fmt.Println(z.Area())
	fmt.Println(z.Perimeter())
}

func main() {
	X := square{10}
	fmt.Println("Perimeter:", X.Perimeter())
	Calculate(X)
	y := circle{5}
	Calculate(y)
}
