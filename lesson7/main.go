package main

import "fmt"

func main() {

	/* Задача 7.1
	Необходимо создать массив строк размером в 5 элементов. Результат вывести в консоль.
	*/
	var array = [5]string{"one", "two", "three", "four", "five"}
	fmt.Println(array)

	/* Задача 7.2
	Необходимо создать массив, содержащий значения: «яблоко», «груша», «слива», «абрикос». Результат вывести в консоль
	*/
	var fruits = [4]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(fruits)
	fmt.Println(array)

	/* Задача 7.3
	Необходимо создать массив, содержащий значения: «яблоко», «груша», «помидор», «абрикос». Далее «помидор» заменить на «персик».
	Результат вывести в консоль
	*/
	var fruitsWithImposter = [4]string{"яблоко", "груша", "помидор", "абрикос"}
	fruitsWithImposter[2] = "персик"
	fmt.Println(fruitsWithImposter)

	/* Задача 7.4
	Необходимо создать целочисленный срез со значениями: 5, 2, 8, 3, 1, 9.
	Результат вывести в консоль.
	*/
	slice := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(slice)

	/* Задача 7.5
	Необходимо создать пустой срез ёмкостью 10.
	Результат вывести в консоль
	*/
	var emptySlice = make([]int, 0, 10)
	fmt.Println(emptySlice)

	/* Задача 7.6
	Необходимо создать пустой срез ёмкостью 10.
	Результат вывести в консоль
	*/
	var notEmptySlice = make([]int, 0, 10)
	notEmptySlice = append(notEmptySlice, 4, 1, 8, 9)
	fmt.Println(notEmptySlice)

	/* Задача 7.7
	Необходимо создать целочисленные срезы со значениями:
	1, 2, 3 и 4, 5, 6.
	Далее объединить эти срезы в результирующий срез.
	Результирующий срез вывести в консоль
	*/
	firstSlice := []int{1, 2, 3}
	secondSlice := []int{4, 5, 6}

	resultSlice := append(firstSlice, secondSlice...)
	fmt.Println(resultSlice)
	fmt.Println(notEmptySlice)

	/* Задача 7.8
	Необходимо создать целочисленный срез: 1, 2, 3, 4, 5, 6.
	Далее удалить элемент «4» и вывести результат в консоль.
	*/
	intSlice := []int{1, 2, 3, 4, 5, 6}
	intSlice = append(intSlice[:3], intSlice[4:]...)
	fmt.Println(intSlice)

}
