package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Counter - структура счетчика
type Counter struct {
	mu    sync.Mutex // мьютекс для синхронизации доступа к счетчику
	value int        // значение счетчика
}

// Increment - функция для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()   // блокировка мьютекса
	c.value++     // инкрементирование счетчика
	c.mu.Unlock() // разблокировка мьютекса
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	numRoutines := 56

	wg.Add(numRoutines)

	// Цикл для запуска горутин
	for i := 0; i < numRoutines; i++ {
		go func() {
			counter.Increment() // Инкрементирование счетчика
			wg.Done()
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Вывод в консоль
	fmt.Println("Итоговое значение счетчика:", counter.value)
}
