```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```
Программа выведет числа от 0 до 9 и закончится дедлоком, поскольку в главной горутине идет вывод значений из канала
Для коррекнтого завершения работы программы необходимо закрыть канал close(ch) после цикла внутри горутины.
Поскольку Цикл for n := range ch получает значения из канала до тех пор, пока он не закрыт.
