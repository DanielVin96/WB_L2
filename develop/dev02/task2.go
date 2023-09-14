package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Repeat(char rune, n int) []rune {
	res := make([]rune, 0)
	for i := 0; i < n; i++ {
		res = append(res, char)
	}
	return res
}

func RepeatLetter(s string) (string, error) {
	runes := []rune(s)        // Конверт в срез рун
	newStr := make([]rune, 0) // возвращаемая строка в дальнейшем
	last := rune(' ')         // последний символ - не цифра
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == `\` { // escape косой черты
			if i+1 < len(runes) {
				i++
				last = runes[i]
				newStr = append(newStr, last)
				continue
			}
		}
		num, err := strconv.Atoi(string(runes[i])) // проверка, если символ - цифра, то конверт в инт
		if err != nil {                            // если нет
			last = runes[i]               // сохраняем в переменную last
			newStr = append(newStr, last) // добавляем в NewStr
		} else {
			if last == ' ' { // строка начинается с цифры, то err
				return "", errors.New("Неверная строка")
			}
			newStr = append(newStr, Repeat(last, num-1)...) // повторяем символ и добавляем в newStr
		}
	}
	return string(newStr), nil // возвращаем новую строку
}

func main() {
	str := "a4bc2d5e"
	fmt.Println(RepeatLetter(str))
}
