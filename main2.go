package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	numChan := make(chan int, 10)
	strChan := make(chan string, 10)
	var wg sync.WaitGroup

	// Записываем числа в канал
	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
		}
		close(numChan)
	}()

	// Запускаем 10 горутин для обработки чисел
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range numChan {
				strChan <- strconv.Itoa(num)
			}
		}()
	}

	// Ожидаем завершения всех горутин и закрываем strChan
	go func() {
		wg.Wait()
		close(strChan)
	}()

	// Читаем и выводим строки
	for str := range strChan {
		fmt.Println(str)
	}
}
