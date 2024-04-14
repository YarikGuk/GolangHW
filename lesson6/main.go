package main

import (
	"fmt"
	"strings"
)

type contract struct {
	ID     int
	Number string
	Date   string
}
type anotherContract struct {
	ID     int
	Number string
	Date   string
}

type info struct {
	Addss string
	Phone string
}
type user struct {
	ID   int
	Name string
	Info info
}
type employee struct {
	ID   int
	Name string
	Info info
}

func main() {

	/* Задача 6.1
	Необходимо объявить глобальную структуру contract с полями: ID int, Number string, Date string.
	Далее создать экземпляр структуры со значениями полей:
	ID=1, Number=«#000A\n101», Date=«2024-01-31».
	В консоль нужно вывести структуру таким образом, чтобы данные отображались в виде:
	{ID:1Number:#000A\n101 Date:2024-01-31}
	*/
	var someContract = contract{
		ID:     1,
		Number: "#000A\n101",
		Date:   "2024-01-31",
	}
	fmt.Printf("{ID:%d Number:%s Date:%s}\n", someContract.ID, strings.ReplaceAll(someContract.Number, "\n", "\\n"), someContract.Date)
	fmt.Println()

	/* Задача 6.2
	Необходимо объявить локальную структуру contract с полями: ID int,
	Number string, Date string. Далее создать экземпляр структуры со значениями полей: ID=1,
	Number=«#000A101\t01», Date=«2024-01-31». В консоль нужно вывести структуру таким образом,
	чтобы данные отображались в виде: {ID:1 Number:#000A101 01 Date:2024-01-31}
	*/
	var someSecondContract = contract{
		ID:     1,
		Number: "#000A101\t01",
		Date:   "2024-01-31",
	}
	fmt.Printf("%+v", someSecondContract)
	fmt.Println()

	/* Задача 6.3
	Необходимо объявить глобальную структуру contract с полями: ID int, Number string, Date string.
	Далее создать экземпляр структуры со значениями полей: ID=1, Number=«#000A\n101», Date=«2024-01-31».
	При передачи экземпляра структуры в fmt.Println в консоли должно отображаться: Договор № #000A\n101 от 2024-01-31
	*/
	var someThirdContract = anotherContract{
		ID:     1,
		Number: "#000A\\n101",
		Date:   "2024-01-31",
	}
	fmt.Printf("%+v", someThirdContract)

	/* Задача 6.4
	Необходимо убрать повторяющийся код - поля Addss и Phone из структур:
	type user struct {
	ID int
	Name string
	Addss string
	Phone string
	}
	type employee struct {
	ID int
	Name string
	Addss string
	Phone string
	}

	После проведения рефакторинга строка fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
	должна выводить в консоль адрес и телефон пользователя и сотрудника соответственно
	*/

	u := user{
		ID:   1,
		Name: "User",
		Info: info{
			Addss: "Some street",
			Phone: "8-800-555-35-35",
		},
	}
	e := employee{
		ID:   1,
		Name: "Employee",
		Info: info{
			Addss: "employee street",
			Phone: "7777",
		},
	}

	fmt.Println(u.Info.Addss, u.Info.Phone, e.Info.Addss, e.Info.Phone)

}

func (c anotherContract) String() string {
	return fmt.Sprintf("Договор № %s от %s", strings.ReplaceAll(c.Number, "\n", "\\n"), c.Date)
}
