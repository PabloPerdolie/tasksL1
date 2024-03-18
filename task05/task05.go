package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {
	var N int    // Обявление переменной N
	fmt.Scan(&N) // Запись в переменную из консоли

	ch := make(chan int) // Создаем канал

	go func() {
		defer close(ch) // Закрываем канал по завершении функции
		for {
			ch <- rand.Intn(10)     // Отправляем значение в канал
			time.Sleep(time.Second) // Задержка 1 секунда
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println("Received:", value) // Выводим значение, полученное из канала
		}
	}()

	time.Sleep(time.Duration(N) * time.Second) // Задержка на N секунд
	fmt.Println("Program finished.")
}
