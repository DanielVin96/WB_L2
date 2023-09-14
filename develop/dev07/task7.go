package main

import (
	"fmt"
	"sync"
	"time"
)

func Union(channels ...<-chan interface{}) <-chan interface{} {
	// создаем канал, который будет возвращен
	single := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(1)
	// запускаем горутину для каждого done-канала
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			// читаем из done-канала и отправляем значение в single-канал
			for val := range ch {
				single <- val
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(single)
	return single // Возвращаем Single
}

func MakeChan(after time.Duration) <-chan interface{} { // метод создания каналов
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-Union(
		MakeChan(2*time.Hour),
		MakeChan(5*time.Minute),
		MakeChan(1*time.Second),
		MakeChan(1*time.Hour),
		MakeChan(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
