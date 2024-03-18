package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	// Обявление переменной N - количества воркеров
	var N int

	// Запись в переменную из консоли
	fmt.Scan(&N)

	// Инициализация глвного канала с данными
	numChan := make(chan int)

	// Канал для отслеживания сигналов ОС и регистрация сигнала CTRL+C или прерывание
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	// Канал для завершения работы горутин
	quit := make(chan bool)

	// Объявление переменной - счетчика типа sync.WaitGroup
	var wg sync.WaitGroup

	// Увеличение счетчика на 1
	wg.Add(1)

	// Запуск горутины с анонимной функцией
	go func(nums chan<- int, quit chan<- bool) {
		// Отложенное уменьшение счетчика
		defer wg.Done()
		// Бесконечный цикл
		for {
			select {
			// Поступил сигнал от системы
			case <-signalChan:
				// Отправление сигнала о заверешении
				quit <- true
				return
			default:
				// Отправление в главный канал целого числа
				nums <- rand.Intn(10)
			}
		}
	}(numChan, quit)

	// Цикл для создания N горутин
	for i := 0; i < N; i++ {
		// Увеличение счетчика на 1
		wg.Add(1)

		// Запуск горутины с анонимной функцией
		go func(nums <-chan int, i int, quit <-chan bool) {
			// Отложенное уменьшение счетчика
			defer wg.Done()

			// Бесконечный цикл
			for {
				select {
				// Поступил сигнал о заверешнии
				case <-quit:
					return
				default:
					fmt.Printf("Read worker %d: %d\n", i, <-nums)
				}
			}

		}(numChan, i, quit)
	}

	// Главный поток дожидается завершения всех горутин
	wg.Wait()
}
