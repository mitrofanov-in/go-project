package main

import (
	"fmt"
)

var size int
var maxN int
var minN int
var sum int
var moveMin int
var moveMax int

func main() {
	arr := []int{9, 1000, 52, 12, 2, 7, 6, 53, 54, 237}
	fmt.Println(arr)

	size = len(arr)

	//Ищем максимальный и минимальный элемент
	for i := 0; i < size; i++ {
		//Максимальный
		if arr[i] > arr[maxN] {
			maxN = i
		}
		//Минимальный
		if arr[i] < arr[minN] {
			minN = i
		}
	}

	moveMax = maxN + 1
	moveMin = minN + 1
	//Суммируем элементы
	// Иницируем цикл
	for i := 0; i < size; i++ {
		//Сдвигаем элемент от начального отсчтеа
		/*
			if i==minN{
				moveMax=maxN+1
			}
		*/

		/*
			if i==maxN{
				moveMin=minN+1
			}
		*/

		//Сумма между индексами
		// Когда максимальный индекс располагает справа
		if maxN > moveMin {

			sum = sum + arr[moveMin]
			fmt.Println("суммирование элементов", arr[moveMin], "прямой")
			moveMin++
		}

		//Сумма между индексами
		// Когда максимальный индекс располагает слева
		if minN > moveMax {

			sum = sum + arr[moveMax]
			fmt.Println("суммирование элементов", arr[moveMax], "обратный порядок")
			moveMax++
		}

	}
	fmt.Println("минимальный", minN, "максимальный", maxN, "итог", sum)
}
