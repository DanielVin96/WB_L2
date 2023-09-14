package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	// создание временного файла
	// fileContent := "The quick brown fox \njumps over the lazy dog \nThe lazy dog barks\nThe quick brown fox jumps\n"
	// file, err := createTempFile(fileContent)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer file.Close()
	file, err := os.Open("grep.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// тесты
	tests := []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{"-i", "-B", "1", "-A", "1", "fox", file.Name()},
			expected: "The quick brown fox\njumps over the lazy dog\nThe quick brown fox jumps\n",
		},
		{
			args:     []string{"-i", "-C", "1", "dog", file.Name()},
			expected: "jumps over the lazy dog\nThe lazy dog barks\nThe quick brown fox jumps\n",
		},
		{
			args:     []string{"-F", "quick brown", file.Name()},
			expected: "The quick brown fox\n",
		},
		{
			args:     []string{"-n", "-v", "lazy", file.Name()},
			expected: "1:The quick brown fox\n4:The quick brown fox jumps\n",
		},
		{
			args:     []string{"-c", "lazy", file.Name()},
			expected: "1\n",
		},
	}

	for _, test := range tests {
		outStream := new(bytes.Buffer)
		errStream := new(bytes.Buffer)

		// запуск команды с флагами и аргументами
		args := append([]string{"go run task5.go"}, test.args...)
		cmd := newCommand(args, outStream, errStream)
		err = cmd.Run()
		if err != nil {
			t.Errorf("Command %v failed with error: %v", args, err)
		}

		// проверка результата
		if outStream.String() != test.expected {
			t.Errorf("Expected output %q, but got %q", test.expected, outStream.String())
		}
	}
}

// // вспомогательная функция для создания временного файла
// func createTempFile(content string) (*os.File, error) {
// 	file, err := os.CreateTemp("", "testfile")
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = file.WriteString(content)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = file.Sync()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }

// вспомогательная функция для создания команды с заданными потоками вывода/ошибок
func newCommand(args []string, outStream, errStream *bytes.Buffer) *exec.Cmd {
	cmd := exec.Command("sh", "-c", strings.Join(args, " "))
	cmd.Stdout = outStream
	cmd.Stderr = errStream
	return cmd
}
