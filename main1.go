package main

import (
	"fmt"
	"sync"
)

var globalVar int       // глобальная переменная
var mu sync.Mutex       // мьютекс для синхронизации

func increment() {
	mu.Lock()             // захват мьютекса
	defer mu.Unlock()     // освобождение мьютекса
	globalVar++           // увеличение переменной
}

func main() {
	const numGoroutines = 5
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()              // ожидание завершения горутин
	fmt.Println(globalVar) // вывод результата
}
