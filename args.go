package main

import (
	//"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	e_string := ""
	arguments := os.Args

	if len(arguments) == 1 {
		e_string = "good"
	} else {
		e_string = arguments[6]
	}
	fmt.Println(len(arguments), e_string)

	var res float64 = 0
	var average float64 = 0

	for i := 1; i < len(arguments); i++ {
		if arguments[i] != "END" {
			n, _ := strconv.ParseFloat(arguments[i], 64)
			fmt.Println(n)
			res = res + n
		} else {
			continue
		}
	}

	average = res / 2

	fmt.Println(average)

	//	io.WriteString(os.Stdout, e_string)

	/*
		var f *os.File
		f = os.Stdin
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	*/
}
