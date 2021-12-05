package main

import (
	"fmt"
	"strconv"
)

func main() {
	tMap := make(map[string]int)
	tMap["k1"] = 123
	tMap["k2"] = 321

	oMap := map[string]int{
		"k1": 124,
		"k2": 322,
	}

	_, ok := tMap["k3"]
	if ok {

	} else {
		tMap["k3"] = 100
	}

	testMap := map[string]interface{}{
		"juju": "jiji",
		"uhu":  123,
		"tiju": 123.011,
		"huhu": 54,
	}

	testString := fmt.Sprintf("%v", testMap["uhu"])

	fmt.Println("trnasfrom to string ", testString, len(testString))

	for _, value := range testMap {
		fmt.Println(value)
	}
	x := testMap["uhu"]
	x_int := x.(int)

	if tMap["k1"] == x_int {
		fmt.Println("super")
	}

	fmt.Println(tMap, oMap)

	for _, value := range testMap {
		switch value.(type) {
		case int:
			fmt.Println("int", value, value.(int))
		case string:
			fmt.Println("string", value, len(value.(string)))
		}
	}

	//t_int,_ := strconv.ParseInt(x_int, 10, 64)

	xar := [5]int64{1, 2, 6, 7, 10}
	xar_st := [5]string{"j", "k", "s", "say", "bay"}
	xMap := make(map[string]int)

	for j := 0; j < len(xar); j++ {
		xMap[xar_st[j]] = j
		fmt.Println(xMap)
	}
	fmt.Println(xMap)

	s_str := strconv.FormatInt(xar[3], 10)
	fmt.Println(s_str)
}
