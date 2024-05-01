package main

import (
	"errors"
	"fmt"
)

func main() {

	/* Задача 11.1
	Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1»
	*/
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2: %w", err1)
	err3 := fmt.Errorf("ошибка3: %w", err2)
	fmt.Println(err3)

	/* Задача 11.2
	Нужно создать, используя оборачивание, ошибку
	«ошибка3:ошибка2:ошибка1». Из созданной цепочки ошибок нужно получить ошибку «ошибка2:ошибка1» и вывести в stdout
	*/
	err := errors.Unwrap(err3)
	fmt.Println(err)

	/* Задача 11.3
	Нужно создать, используя оборачивание, ошибку
	«ошибка3:ошибка2:ошибка1». Не используя unwrap, нужно определить была ли ошибка «ошибка1»
	*/
	fmt.Println("Ошибка1 есть в цепочке:", errors.Is(err3, err1))

	/* Задача 11.4
	Нужно создать, используя оборачивание, ошибку
	«ошибка3:ошибка2:ошибка1». Также нужно создать свою ошибку в виде структуры myFirstError,
	которая обязательно должна иметь метод Error() string. Необходимо убедиться, что в
	созданной цепочке ошибок не было ошибки типа myFirstError
	*/
	myErr := myFirstError{}
	if errors.As(err3, &myErr) {
		fmt.Println("Ошибка myFirstError есть в цепочке")
	} else {
		fmt.Println("Ошибки myFirstError нет в цепочке")
	}

}

type myFirstError struct {
	message string
}

func (error myFirstError) Error() string {
	return error.message
}
