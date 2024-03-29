package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
// с использованием конкурентных вычислений.

func main() {
	// Инициализация заданной последовательности
	array := [5]int{2, 4, 6, 8, 10}

	// Инициализация канала для передачи квадратов элементов
	cur := make(chan int, len(array))

	var result int

	// Объявление переменной - счетчика типа sync.WaitGroup
	var wg sync.WaitGroup

	// Цикл по элементам последовательности
	for _, num := range array {

		// Увеличение счетчика на 1
		wg.Add(1)
		// Запуск горутины с анонимной функцией
		go func(num int, cur chan<- int, wg *sync.WaitGroup) {
			// Передача квадрата элемента в канал
			cur <- num * num
			// Отложенное уменьшение счетчика
			defer wg.Done()
		}(num, cur, &wg)
	}

	// Главный поток дожидается завершения всех горутин
	wg.Wait()

	// Закрытие канала
	close(cur)

	// Цикл по значениям канала cur
	for val := range cur {
		result += val
	}

	// Вывод результата в консоль
	fmt.Println(result)

}
