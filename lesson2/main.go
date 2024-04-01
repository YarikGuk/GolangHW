package main

import "fmt"

const someConst = 0
const constForShadowing = 666

const (
	const1 = 1
	const2 = 2
	const3 = 3
	const4 = 4
	const5 = 5
)

func main() {

	/* Задача 2.2
	Необходимо написать программу, которая в консоль выводит результат деления 16 на 3 в виде сообщения:
	«Результат: ..., остаток от деления: ..., тип результата: ...».
	*/
	result := 16 / 3
	remainder := 16 % 3

	fmt.Printf("\"Задача 2.2: Результат: %d, остаток от деления: %d, тип результата: %T\n", result, remainder, result)

	/* Задача 3.1
	Необходимо создать глобальную константу и вывести её значение в консоль.
	*/
	fmt.Printf("Задача 3.1: %d\n", someConst)

	/* Задача 3.2
	Необходимо создать локальную константу и вывести её значение в консоль.
	*/
	const localConst = 777
	fmt.Println("Задача 3.2:", localConst)

	/* Задача 3.3
	Необходимо создать глобальную константу. Далее затенить
	глобальную константу локальной и вывести результат в консоль
	*/
	const constForShadowing = 111
	fmt.Println("Задача 3.3:", constForShadowing)

	/* Задача 3.4
	Необходимо создать группу из пяти глобальных констант и вывести их значения в консоль.
	*/
	fmt.Printf("Задача 3.4: %d,%d,%d,%d,%d\n", const1, const2, const3, const4, const5)

	/* Задача 3.5
	Необходимо создать группу из пяти глобальных констант и вывести их значения в консоль.
	*/
	const (
		localConst1 = 1
		localConst2 = 2
		localConst3 = 3
		localConst4 = 4
		localConst5 = 5
	)
	fmt.Printf("Задача 3.5: %d,%d,%d,%d,%d\n", localConst1, localConst2, localConst3, localConst4, localConst5)

	/* Задача 3.6
	Необходимо создать типизированную локальную константу n типа int со значением 5.
	Также необходимо создать локальную переменную m, значение которой должно определяться выражением 3.4 + n.
	Также нужно вывести значение переменной в консоль (в консоли должно отобразиться значение 8.4).
	*/
	const n int = 5
	m := 3.4 + float64(n)
	fmt.Println("Задача 3.6:", m)

	/* Задача 3.7
	Использую генератор iota необходимо создать 5 констант. Их значения должны представлять собой: 1, 2, 4, 8, 16.
	*/
	const (
		iotaConst1 = 1 << iota
		iotaConst2
		iotaConst3
		iotaConst4
		iotaConst5
	)
	fmt.Printf("Задача 3.7: %d,%d,%d,%d,%d\n", iotaConst1, iotaConst2, iotaConst3, iotaConst4, iotaConst5)
}
