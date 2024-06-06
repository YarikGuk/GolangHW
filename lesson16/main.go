package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const numGoroutines = 5

func main() {

	/* Задача 16.1
	Нужно запустить 5 горутин и остановить, используя контекст с
	отменой
	*/
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < numGoroutines; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Cтоп горутина: ", i)
					return
				default:
					time.Sleep(time.Millisecond * 50)
					fmt.Println("Cложные вычисления горутины: ", i)
				}
			}
		}()
	}
	time.Sleep(time.Millisecond * 100)
	cancel()
	wg.Wait()

	/* Задача 16.2
	Нужно запустить 5 горутин и остановить через 2 секунды
	*/
	ctx2 := context.Background()
	ctx2, cancel2 := context.WithCancel(ctx2)
	defer cancel2()
	wg2 := sync.WaitGroup{}
	for i := 0; i < numGoroutines; i++ {
		i := i
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for {
				select {
				case <-ctx2.Done():
					fmt.Println("Cтоп горутина (2): ", i)
					return
				default:
					time.Sleep(time.Millisecond * 250)
					fmt.Println("Cложные вычисления горутины (2): ", i)
				}
			}
		}()
	}

	time.Sleep(2 * time.Second)
	cancel2()
	wg2.Wait()

	/* Задача 16.3
	Нужно запустить 5 горутин и остановить в некоторое время, которое рассчитывается по формуле: текущий момент + 2 секунды
	*/
	ctx3 := context.Background()
	ctx3, cancel3 := context.WithTimeout(ctx3, 2*time.Second)
	defer cancel3()
	var wg3 = sync.WaitGroup{}

	for i := 0; i < numGoroutines; i++ {
		i := i
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			for {
				select {
				case <-ctx3.Done():
					fmt.Println("Cтоп горутина (3): ", i)
					return
				default:
					time.Sleep(time.Millisecond * 250)
					fmt.Println("Cложные вычисления горутины (3): ", i)
				}
			}
		}()
	}

	wg3.Wait()

	/* Задача 16.4
	Нужно создать контекст со значениями

	some key1: some value1
	some key2: some value2

	Контекст следует передать в функцию, которая выведет значения some key1 и some key2 в stdout.	*/

	ctx4 := context.Background()
	ctx4 = context.WithValue(ctx4, key1, "some value1")
	ctx4 = context.WithValue(ctx4, key2, "some value2")

	printValues(ctx4)

	/* Задача 16.5
	Не используя context и буферизованные каналы необходимо написать программу, которая будет запускать 10 рабочих горутин и одну капризную управляющую горутину. Каждая рабочая горутина с задержкой в 1 секунду должна выводить в stdout сообщение «сложные вычисления горутины: 1», где 1 - порядковый номер горутины. Управляющая горутина через 3 секунды после своего запуска должна в stdout вывести «ой, всё!», после чего рабочие горутины должны в stdout вывести «stop горутина: 1», где 1 - порядковый номер горутины, и завершить своё выполнение. В консоли должны увидеть что-то подобное:

	сложные вычисления горутины: 3
	сложные вычисления горутины: 8
	сложные вычисления горутины: 4
	сложные вычисления горутины: 9
	сложные вычисления горутины: 7
	сложные вычисления горутины: 2
	сложные вычисления горутины: 0
	сложные вычисления горутины: 1
	сложные вычисления горутины: 5

	сложные вычисления горутины: 6
	ой, всё!
	сложные вычисления горутины: 3
	stop горутина: 3
	сложные вычисления горутины: 9
	stop горутина: 9
	сложные вычисления горутины: 7
	stop горутина: 7
	сложные вычисления горутины: 4
	stop горутина: 4
	сложные вычисления горутины: 8
	stop горутина: 8
	сложные вычисления горутины: 6
	stop горутина: 6
	сложные вычисления горутины: 2
	stop горутина: 2
	сложные вычисления горутины: 0
	stop горутина: 0
	сложные вычисления горутины: 1
	stop горутина: 1
	сложные вычисления горутины: 5
	stop горутина: 5
	*/

	var wg5 sync.WaitGroup
	stopCh := make(chan struct{})
	doneCh := make(chan int)

	for i := 0; i < 10; i++ {
		wg5.Add(1)
		i := i
		go func() {
			defer wg5.Done()
			for {
				select {
				case <-stopCh:
					fmt.Printf("сложные вычисления горутины: %d\n", i)
					doneCh <- i
					return
				default:
					time.Sleep(1 * time.Second)
					fmt.Printf("сложные вычисления горутины: %d\n", i)
				}
			}
		}()
	}

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		close(stopCh)
	}()

	go func() {
		wg5.Wait()
		close(doneCh)
	}()

	for id := range doneCh {
		fmt.Printf("stop горутина: %d\n", id)
	}

}

const (
	key1 = "some key1"
	key2 = "some key2"
)

func printValues(ctx context.Context) {
	value1 := ctx.Value(key1)
	value2 := ctx.Value(key2)

	fmt.Printf("%s: %s \n", key1, value1)
	fmt.Printf("%s: %s \n", key2, value2)

}
