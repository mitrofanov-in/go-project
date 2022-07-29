package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Address struct {
	Index       int
	HouseNumber int
	Street      string
}

func (addr Address) addInfo(inf chan Address) {

	fmt.Println("Загруженная структура", inf)
	inf <- addr
	//sVal := <-inf
	if addr.Street == "Букашкино" {
		fmt.Println("Нашелся", addr)
		os.Exit(0)
	} else {
		fmt.Println("Не ашелся")
	}

}

func ReadCh(inf chan Address) {
	<-inf
	fmt.Println("Читаем из канала")
}

func main() {

	var stopChan = make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	inf := make(chan Address)

	for {
		/*
			reader := bufio.NewScanner(os.Stdin)
			reader.Scan()
			text := reader.Text()
		*/
		fmt.Println("Введите индекс, номер дома, улицу")
		var i, h int
		var st string
		fmt.Scan(&i, &h, &st)
		curStruct := Address{i, h, st}
		go curStruct.addInfo(inf)
		time.Sleep(time.Second * 3)
		ReadCh(inf)
		signal.Stop(stopChan)
	}

}
