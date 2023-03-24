package main

import "fmt"

var msg string

func init() { //функция дополнения пакетов, всегда выволняется перед main
	msg = "LOH"
}

func main() {
	var p *int
	number := 2
	p = &number
	fmt.Println(msg)
	change(&msg)
	fmt.Println(msg)
	fmt.Println(*p)
	fmt.Print(p)
	fmt.Println(" ", &number)
	*p = 12 //перезаписали значение ячейки
	fmt.Print(number)
}

func change(msg *string) { //& - обращение к ячейке(номеру) памяти, * - преобразования номера в значение которое в ней хранится
	*msg += " (вызвано из функции print)"
	fmt.Println("Измененние переменной")
}
