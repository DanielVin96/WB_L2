package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(word string) string {
	letters := []rune(word)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}

func FindAnagrams(words []string) *map[string][]string {
	// Словарь для избежания дубликатов
	dictionary := make(map[string]bool, len(words))
	// Мапа - предрезультат
	m := make(map[string][]string)
	// Мапа - результат
	r := make(map[string][]string)

	// Итерируемся по словам
	for _, upperWord := range words {

		// Переводим слово в нижний регистр
		word := strings.ToLower(upperWord)
		if len(word) < 2 {
			continue // Пропускаем слова из одной буквы
		}

		// Если слова уже нет в каком-либо множестве
		if !dictionary[word] {
			m[SortString(word)] = append(m[SortString(word)], word)
			dictionary[word] = true
		}
	}

	for _, v := range m {
		firstWord := v[0]
		sort.Strings(v)
		r[firstWord] = v
	}

	return &r
}

func main() {
	words := []string{"кулон", "клоун", "газон", "загон"}
	anag := *FindAnagrams(words)
	fmt.Println(anag)
}
