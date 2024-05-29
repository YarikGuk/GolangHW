package main

import (
	"fmt"
	"time"
)

func main() {

	/* Задача 14.1
	Необходимо создать дочернюю горутину, которая выведет в stdout
	«Привет из дочерней горутины!»

	p.s. Последние кавычки елочкой некорректые, заменил на обычные
	*/
	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()
	time.Sleep(100 * time.Duration(time.Millisecond))

	/* Задача 14.2
	Необходимо создать небуферизованный канал, который будет
	принимать строковое значение. Записать в канал «Привет, строковый канал!»,
	далее прочитать это значение и вывести в stdout
	*/
	channel := make(chan string)

	go func() {
		channel <- "Привет, строковый канал!"
	}()

	msg := <-channel
	fmt.Println(msg)

	/* Задача 14.3
	Необходимо создать канал с буфером 4, записать в него 2 значения:
	«Привет», «буферизованный канал!».
	Далее необходимо прочитать все значения из канала и вывести в stdout
	*/
	channel2 := make(chan string, 4)

	// Запись двух значений в канал
	channel2 <- "Привет"
	channel2 <- "буферизованный канал!"

	// Чтение значений из канала и вывод их в stdout
	for i := 0; i < 2; i++ {
		msg := <-channel2
		fmt.Println(msg)
	}

	/* Задача 14.4
	Если запустить следующий код программы:
	package main
	import "fmt"
	func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
	<-ch
	stop <- struct{}{}
	}()
	<-stop
	fmt.Println("happy end")
	}
	возникнет ошибка блокировки:
	fatal error: all goroutines are asleep - deadlock!
	goroutine 1 [chan receive]:
	main.main()
	/go/src/course/main.go:12 +0x9c
	goroutine 18 [chan receive]:
	main.main.func1()
	/go/src/course/main.go:9 +0x2c
	created by main.main in goroutine 1
	/go/src/course/main.go:8 +0x90
	exit status 2

	Не меняя логику дочерней горутины, нужно дописать логику главной горутины так,
	чтобы ошибки блокировки не возникало.
	Вместо ошибки в консоль должна выводится фраза «happy end».
	Так как канал ch является «особым» , то в него ЗАПРЕЩАЕТСЯ писать значения.
	*/
	ch := make(chan int)
	stop := make(chan struct{})

	go func() {
		<-ch
		stop <- struct{}{}
	}()
	close(ch)

	<-stop
	fmt.Println("happy end")

	/* Задача 14.5
	Следующий код программы запускает излишне «болтливую» горутину:

	package main
	import (
	"fmt"
	"time"
	)
	func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
	for {
	select {
	case <-stop:
	break OUT
	case v, ok := <-ch:
	if !ok {
	break OUT
	}
	fmt.Println(v)
	default:
	continue
	}
	}
	fmt.Println("завершение работы горутины_1")
	}()
	go func() {
	var i int
	OUT:
	for {
	i++
	select {
	case <-stop:
	break OUT
	default:
	time.Sleep(time.Second)
	if ch == nil {
	continue
	}
	ch <- i
	}
	}
	fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
	}
	которая в консоль выводит счёт. Горутина настолько утомительна, что в коде поставлено ограничение по времени выполнения программы. Не меняя логику дочерних горутин, необходимо дописать логику так, чтобы «болтушка» не смогла вывести в консоль свой утомительный счёт. Вместо пустой болтовни должны увидеть:
	завершение работы горутины_1
	завершение работы горутины_2
	завершение работы главной горутины

	Причём логика главной горутины должна отработать в полном объёме,
	без изменения длительности временных задержек.
	*/

	ch1 := make(chan int)
	stop1 := make(chan struct{}, 2)

	go func() {
	OUT:
		for {
			select {
			case <-stop1:
				break OUT
			case v, ok := <-ch1:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()

	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop1:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch1 == nil {
					continue
				}
				select {
				case ch1 <- i:

				default:

				}
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	time.Sleep(5 * time.Second)
	stop1 <- struct{}{}
	stop1 <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")

	/* Задача 14.6
	Посмотрите внимательно на код:
	package main
	import (
	"fmt"
	"time"
	)
	func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
	for {
	select {
	case <-stop:
	break OUT
	case v, ok := <-ch:
	if !ok {
	break OUT
	}
	fmt.Println(v)

	default:
	continue
	}
	}
	fmt.Println("завершение работы горутины_1")
	}()
	go func() {
	var i int
	OUT:
	for {
	i++
	select {
	case <-stop:
	break OUT
	default:
	time.Sleep(time.Second)
	if ch == nil {
	continue
	}
	ch <- i
	}
	}
	fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	37
	fmt.Println("завершение работы главной горутины")
	}
	Запустите код. Затем посмотрите внимательно на звёзды и скажите: сколько горутин запускалось, а также расскажите о дальнейшей их судьбе.
	*/

	ch2 := make(chan int)
	stop2 := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop2:
				break OUT
			case v, ok := <-ch2:
				if !ok {
					break OUT
				}
				fmt.Println(v)

			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop2:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch2 == nil {
					continue
				}
				ch2 <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop2 <- struct{}{}
	stop2 <- struct{}{}
	time.Sleep(time.Second)
	//37
	fmt.Println("завершение работы главной горутины")

}
