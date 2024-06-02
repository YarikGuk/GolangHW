package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const numGoroutines = 5
const numIncrements = 10

func main() {

	/* Задача 15.1
	Необходимо запустить 5 горутин. Не используя time.Sleep нужно
	обеспечить вывод в консоль каждой горутиной своего уникального
	сообщения. Например:
	горутина: 1
	горутина: 2
	…
	*/
	wg := sync.WaitGroup{}
	for i := 1; i <= numGoroutines; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("горутина:", i)
		}()
	}
	wg.Wait()

	/* Задача 15.2
	Используя пакет atomic, необходимо реализовать счётчик,
	с которым параллельно могут работать несколько горутин.
	*/
	var counter int64
	var wg1 sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for j := 0; j < numIncrements; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg1.Wait()
	fmt.Println(counter)

	/* Задача 15.3
	Используя Mutex, необходимо реализовать счётчик, с которым
	параллельно могут работать несколько горутин
	*/

	var counter2 int
	var mutex sync.Mutex
	var wg2 sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < numIncrements; j++ {
				mutex.Lock()
				counter2++
				mutex.Unlock()
			}
		}()
	}
	wg2.Wait()
	fmt.Println(counter2)

	/* Задача 15.4
	Необходимо создать функцию start, которая в консоль выводит
	некоторое сообщение. Необходимо запустить 10 горутин, которые
	будут запускать функцию start и выводить в консоль факт своего запуска,
	причём необходимо обеспечить однократный запуск функции start
	*/

	var once sync.Once
	var wg3 sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			once.Do(start)
			fmt.Printf("Горутина %d\n", id)
		}(i)
	}

	wg3.Wait()

	/* Задача 15.5
	Необходимо реализовать интерфейс

	type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
	}

	Речь про температуру окружающей среды. ReadTemp читает
	температуру, ChangeTemp изменяет температуру. Код должен быть потокобезопасным,
	т.е. при работе с температурой нескольких параллельных
	горутин не должно возникать состояние гонки.
	*/

	ws := &MeteoState{temp: "10°C"}

	var wg4 sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			fmt.Printf("Температура = %s\n", ws.ReadTemp())
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			newTemp := fmt.Sprintf("%d°C", 10+id)
			ws.ChangeTemp(newTemp)
			fmt.Printf("Температура изменена на %s\n", newTemp)
		}(i)
	}

	wg4.Wait()
}

func start() {
	fmt.Println("Некоторое сообщение")
}

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type MeteoState struct {
	temp  string
	mutex sync.Mutex
}

func (ws *MeteoState) ReadTemp() string {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	return ws.temp
}

func (ws *MeteoState) ChangeTemp(v string) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	ws.temp = v
}
