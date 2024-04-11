package main

import "fmt"

type Square int

func main() {

	/* Задача 5.1
	Необходимо создать указатель на строковое значение
	*/
	strValue := "Your ADV could be here"
	var pointer *string = &strValue
	fmt.Println("Pointer", pointer)

	/* Задача 5.2
	Необходимо создать переменную. В консоль вывести её значение и адрес
	*/
	variable1 := 777
	fmt.Println("Value of variable1:", variable1)
	fmt.Println("Address of variable1:", &variable1)

	/* Задача 5.3
	Необходимо создать переменную и указатель на неё. Необходимо через указатель изменить значение переменной
	*/
	variable2 := 777
	pointerVariable2 := &variable2
	*pointerVariable2 = 666
	fmt.Println("Value of variable2:", variable2)

	/* Задача 5.4
	«Создать группу переменных, в консоль вывести их значения и адреса.
	Проанализировать, как отличаются адреса и почему»
	*/
	variable3, variable4, variable5 := 1, 1, "2"

	fmt.Println("Value of variable3:", variable3)
	fmt.Println("Address of variable3:", &variable3)
	fmt.Println("Value of variable4:", variable4)
	fmt.Println("Address of variable4:", &variable4)
	fmt.Println("Value of variable5:", variable5)
	fmt.Println("Address of variable5:", &variable5)
	fmt.Println()

	/* Задача 5.5
	Необходимо создать функцию change, которая принимает
	параметр и изменяет его значение. В функции main необходимо
	создать локальную переменную и вызвать change таким образом,
	чтобы она изменила значение локальной переменной.
	*/
	variable6 := 1
	change(&variable6)
	fmt.Println("Value of variable6:", variable6)

	/* Задача 5.6
	Необходимо создать пользовательский тип square на базе int.
	Далее необходимо создать переменную типа square и значением 25.
	Необходимо вывести значение переменной в консоль
	*/
	var variable7 Square = 25
	fmt.Println("Value of variable7:", variable7)

	/* Задача 5.7
	Необходимо создать пользовательский тип square на базе int.
	Далее необходимо создать переменную типа square и значением 30.
	Значение переменной s необходимо увеличить на 15 и вывести
	результат в консоль.
	*/
	var variable8 Square = 30
	variable8 += 15
	fmt.Println("Value of variable8:", variable8)

	/* Задача 5.8
	Необходимо создать пользовательский тип square на базе int. Далее необходимо создать переменную s типа square и значением 34.
	Значение переменной s необходимо увеличить на 10 и вывести
	результат в консоль. Тип square при выводе в консоль должен
	автоматически дописывать м², то есть результат должен
	выглядеть: 44 м².
	*/
	var variable9 Square = 34
	variable9 += 10
	fmt.Println("Value of variable9:", variable9)
}

func change(value *int) {
	*value = 666
}

func (s Square) String() string {
	return fmt.Sprintf("%d м²", s)
}
