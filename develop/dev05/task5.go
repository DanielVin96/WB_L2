package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// определение флагов
	after := flag.Int("A", 0, "print N lines after match")
	before := flag.Int("B", 0, "print N lines before match")
	context := flag.Int("C", 0, "print N lines around match")
	count := flag.Bool("c", false, "print count of matching lines")
	ignoreCase := flag.Bool("i", false, "ignore case")
	invert := flag.Bool("v", false, "invert match")
	fixed := flag.Bool("F", false, "fixed string match")
	lineNum := flag.Bool("n", false, "print line number")

	flag.Parse()

	// определение паттерна
	var pattern string
	if *fixed {
		pattern = flag.Arg(0)
	} else {
		pattern = strings.ToLower(flag.Arg(0))
	}

	// открытие файла
	//file, err := os.Open(flag.Arg(1))
	file, err := os.Open("grep.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	var lineNumber int
	for scanner.Scan() {
		line := scanner.Text()
		if *ignoreCase {
			line = strings.ToLower(line)
		}
		match := strings.Contains(line, pattern)
		if (*invert && !match) || (!*invert && match) {
			if *count {
				lineNumber++
			} else {
				if *lineNum {
					lineNumber++
					output = append(output, fmt.Sprintf("%d:%s", lineNumber, line))
				} else {
					output = append(output, line)
				}
			}
			if *after > 0 {
				for i := 1; i <= *after && scanner.Scan(); i++ {
					output = append(output, scanner.Text())
				}
			}
		} else if *before > 0 {
			if len(output) > *before {
				output = output[len(output)-*before:]
			}
			output = append(output, line)
		} else if *context > 0 {
			if len(output) > *context*2 {
				output = output[len(output)-*context*2:]
			}
			output = append(output, line)
		}
	}

	// вывод результата
	if *count {
		fmt.Println(lineNumber)
	} else {
		for _, line := range output {
			fmt.Println(line)
		}
	}
}
