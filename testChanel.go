package main

import (
	"fmt"
	"time"
)

var data = []int{1, 2, 3, 4, 5, 6, 7}
var arr = []int{1, 12, 5, 67, 99}
var txt string = "test"

func wrText(tx chan string, x string) {
	tx <- x
	fmt.Println("Write string in channel")
}

func cRead(c chan []int, array []int) {
	c <- array
	//res := <-c
	//close(c)
	fmt.Println("From channel 1", array)
}
func cWrite(c chan []int, array []int) {
	c <- array
	//res := <-c
	//close(c)
	fmt.Println("From channel 2", array)
}

func main() {

	c := make(chan []int, len(data))
	tx := make(chan string)

	go cRead(c, data)

	tmp := <-c
	fmt.Println("Read tmp ", tmp)
	go cWrite(c, arr)
	msg := <-c
	//fmt.Println("Read xyz", xyz)

	go func() {
		for {
			select {
			case msg = <-c:
				fmt.Scan(&msg)
				fmt.Println("Умка сосиска", msg)
			}
		}
	}()

	num := 0
	go func() {
		for {

			fmt.Scan(&num)
			z := []int{num}
			c <- z
			fmt.Println("Вывести", num, c)
			time.Sleep(time.Second * 2)

		}
	}()

	fmt.Scan(&txt)

	go wrText(tx, txt)

	go func() {
		for {
			fmt.Scan(&txt)
			tx <- txt
			fmt.Println("Показать", tx)
			time.Sleep(time.Second * 5)
		}
	}()

}
