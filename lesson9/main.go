package main

import "fmt"

const (
	NONE = iota
	FRUITS
	VEGETABLES
)

func main() {

	/* Задача 9.1
	Необходимо написать функцию fruitMarket, которая будет принимать название фруктов,
	а возвращать их количество. Например, «апельсины» - 5. Сама функция должна создать карту:
	апельсин=5, яблоки=3, сливы=1, груши=0. Если запрашиваемых фруктов нет в карте, то в
	консоль должно выводится сообщение об отсутствии

	*/
	productCount := fruitMarket("яблоки")
	fmt.Println(productCount)

	/* Задача 9.2
	Необходимо создать срез целых чисел: 1, 2, 3. Далее создать 4
	вложенных цикла. Каждый цикл должен в консоль выводить текущее значение среза.
	Причём внутренние срезы должны содержать отступы для облегчения визуального восприятия.
	Внутренний срез на ключе 1 должен остановить все циклы, начиная со второго цикла.
	*/
	slice := []int{1, 2, 3}
	for _, v1 := range slice {
		fmt.Printf("v1: %d\n", v1)
		for _, v2 := range slice {
			fmt.Printf("\tv2: %d\n", v2)
			for _, v3 := range slice {
				fmt.Printf("\t\tv3: %d\n", v3)
				for _, v4 := range slice {
					fmt.Printf("\t\t\tv4: %d\n", v4)
					if v4 == 2 {
						break
					}
				}
				if v3 == 1 {
					break
				}
			}
			if v2 == 1 {
				break
			}
		}
	}

	/* Задача 9.3
	Необходимо написать функцию checkFood, которая принимает в качестве параметра название еды,
	а в консоль выводит «это фрукт», «это овощ». Для проверки необходимо использовать оператор выбора.
	Причём проверке подлежат: "груша", "яблоко", "апельсин", "тыква", "огурец", "помидор".
	Если checkFood получит название еды, которое не входит в этот список, то в консоли должно
	появиться сообщение: «что-то странное…». Для вывода информации в консоль необходимо использовать
	только fmt.Println, но не более трёх раз.
	*/
	checkFood("огурец")
	checkFood("яблоко")
	checkFood("тыблоко")
}

func fruitMarket(request string) int {
	products := map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0,
	}
	product, exist := products[request]

	if !exist || product == 0 {
		fmt.Printf("Продукт \"%s\" закончился!", request)
	}
	return product
}

func checkFood(foodName string) {
	foods := map[string]int{
		"груша":    FRUITS,
		"яблоко":   FRUITS,
		"апельсин": FRUITS,
		"тыква":    VEGETABLES,
		"огурец":   VEGETABLES,
		"помидор":  VEGETABLES,
	}

	foodType := foods[foodName]

	switch foodType {
	case FRUITS:
		fmt.Println("Это фрукт")
	case VEGETABLES:
		fmt.Println("Это овощ")
	case NONE:
		fmt.Println("Что-то странное…")
	}

}
