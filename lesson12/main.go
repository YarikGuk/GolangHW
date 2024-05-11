package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	/* Задача 12.1
	Здесь нужно увеличить значение a на v.
	В случае невозможности приведения к int
	необходимо сообщить об этом и немедленно
	завершить полнение программы.
	*/

	a := 1
	do(a)

	/* Задача 12.2
	Следующий код программы должен вывести в консоль сообщения:
	Утка - Умею летать!
	Утка - Умею плавать!
	Воробей - Умею летать!
	*/
	var d, s Bird
	d = Duck{}
	Do(d)
	s = Sparrow{}
	Do(s)

	/* Задача 12.3
	Не изменяя структуры Xml, Csv и функцию main необходимо
	доработать следующий код так, чтобы в консоли увидели:
	Данные в формате xml
	Данные в формате csv
	*/
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)

	/* Задача 12.4
	В представленном ниже коде возникает паника.
	Нужно понять где и объяснить почему.
	Также нужно предложить не менее 3-х вариантов решения проблемы
	с объяснением причины устранения ошибки
	*/
	var mcDuck *McDuck = &McDuck{voice: "Whats up?"}
	song, err := Sing(mcDuck)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)

}

/* Задача 12.1 */
func do(v any) {
	a := 10

	value, ok := v.(int)
	if !ok {
		log.Fatal("Невозможно привести значение к типу int")
	}
	a += value

	fmt.Println(a)
}

/* Задача 12.2 */

type Bird interface {
	Fly()
}
type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}
func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}

func Do(b Bird) {
	b.Fly()

	duck, ok := b.(Duck)
	if ok {
		duck.Swim()
	}
}

/* Задача 12.3 */

type OutputFormat interface {
	Format()
}

type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}
func Report(x OutputFormat) {
	x.Format()
}

/* Задача 12.4 */

type McBird interface {
	Sing() (string, error)
}
type McDuck struct {
	voice string
}

func (d *McDuck) Sing() (string, error) {
	if d != nil {
		return d.voice, nil
	}
	return "", errors.New("Duck is not exist")
}
func Sing(b McBird) (string, error) {
	if b != nil && b.(*McDuck) != nil {
		return b.Sing()
	}
	return "", errors.New("Ошибка пения!")
}
