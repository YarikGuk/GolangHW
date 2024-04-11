package main

import "fmt"

func main() {

	/* Задача 4.1
	Необходимо создать и вызвать функцию «hello», которая выводит
	«Hello, Go!» в stdout
	*/
	hello()

	/* Задача 4.2
	Необходимо создать и вызвать анонимную функцию, которая выводит «Hello, Go!» в stdout.
	*/
	var printHello = func() { fmt.Println("Hello, Go!") }
	printHello()

	/* Задача 4.3
	Необходимо создать и вызвать функцию «hello», которая в
	качестве параметра принимает и вызывает анонимную функцию.
	Анонимная функция должна выводить в stdout фразу «Hello, Go!».
	*/
	hello2(printHello)

	/* Задача 4.4
	Необходимо создать функцию «hello», которая возвращает анонимную функцию.
	Необходимо вызвать анонимную функцию, которая в stdout выводит «Hello, Go!».
	*/
	printHello = hello3()
	printHello()

	/* Задача 4.5
	Необходимо создать и вызвать функцию «hello», которая выводит
	«Hello, Go!» в stdout. Также, используя defer необходимо вывести фразу «завершение работы».
	*/
	defer func() {
		fmt.Println("Завершение работы")
	}()
	hello()

	/* Задача 4.6
	"Вывести первые 23 числа Фибоначчи, не используя циклы и максимум один оператор if"
	*/
	printFibonacci(23)
}

/* Задача 4.1*/
func hello() {
	fmt.Println("Hello, Go!")
}

/* Задача 4.3*/
func hello2(block func()) {
	block()
}

/* Задача 4.4*/
func hello3() func() {
	return func() {
		fmt.Println("Hello, Go!")
	}
}

/* Задача 4.6*/
func printFibonacci(amountOfNumbers int8) {
	firstNum := 0
	secondNum := 1
	var iterationNum int8 = 1

	fmt.Println("iteration:", iterationNum, ", number:", firstNum)
	iterationNum++

	printNextFibonacciNumber(amountOfNumbers, iterationNum, firstNum, firstNum+secondNum)

}
func getNextFibonacciNum(prevNum int, curNum int) (int, int) {
	return curNum, prevNum + curNum
}

func printNextFibonacciNumber(amountOfNumbers int8, iterationNum int8, prevNum int, curNum int) {
	newPrev, newCur := getNextFibonacciNum(prevNum, curNum)
	fmt.Println("iteration: ", iterationNum, ", number:", curNum)
	iterationNum++

	if iterationNum <= amountOfNumbers {
		printNextFibonacciNumber(amountOfNumbers, iterationNum, newPrev, newCur)
	} else {
		return
	}
}
