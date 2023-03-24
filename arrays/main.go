package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	//var arr [3]int - простоая инициализация массива[3]
	arr := []int{0, 0, 0}
	//[] -  означает что это слайс(обёртка над массивомБ которая динамически изменяяется, а так же предеается с сылкой на массив)
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	fmt.Println(arr)
	er := check(arr)
	if er != nil {
		log.Fatal(er)
	}
	//более правильный способ объявления слайса
	mass := make([]string, 5) //2й параметр отвечает за размер, а 3й за стандарноое его наполнение(не обязателен)
	mass = append(mass, "6")
	mass = append(mass, "7")
	fmt.Println(mass)
	fmt.Println(len(mass))
	fmt.Println(cap(mass)) // сap реальный размер массива, при использовании append размер массива увеличивается 2х
	// до определённого значения затем, х уменьшается

	array := [6]int{5, 7, 3, 8, 1, 6}
	sum := 0
	for _, value := range array { //_ - ключ, value - значение (без ключа value будет принимать индексы от array)
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}

func check(arr []int) error {
	if len(arr) == 0 {
		return errors.New("нулевой массив")
	}
	arr[1] = 3
	fmt.Println(arr)
	return nil
}