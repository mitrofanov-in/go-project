package main

import "fmt"

type Address struct {
	index       int
	houseNumber int
	street      string
}

func (addr Address) addInfo(inf chan struct{}) (int, int) {
	if addr.street == "Букашкино" {
		fmt.Println("Нашелся")
		return addr.index, addr.houseNumber
	} else {
		fmt.Println("Не ашелся")
	}
	return 0, 0
}

func main() {
	inf := make(chan struct{})
	//xStruct := Address{}
	//zStruct := Address{}
	qStruct := Address{69, 91, "Букашкино"}
	go qStruct.addInfo(inf)

	//fmt.Println(qStruct.addInfo(69, 91, "Букашкино"))
	/*
		fmt.Println(xStruct.addInfo(15, 67, "Бульбино"))
		fmt.Println(zStruct.addInfo(21, 23, "Барашкино"))
	*/

}
