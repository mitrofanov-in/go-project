package main

import (
	"fmt"
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

	//t_int,_ := strconv.ParseInt(x_int, 10, 64)

	fmt.Println(tMap, oMap)

}
