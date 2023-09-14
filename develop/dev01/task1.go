package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	requestTime, err := ntp.Query("pool.ntp.org") // Получаем точное время с сервера NTP
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при запросе точного времени: %s\n", err)
		os.Exit(1)
	}

	// Вывод текущего и точного времени
	currentTime := time.Now()                             // текущее
	exactTime := currentTime.Add(requestTime.ClockOffset) // Точное время
	fmt.Printf("Текущее время: %s\n", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Точное время: %s\n", exactTime.Format("2006-01-02 15:04:05"))
}
