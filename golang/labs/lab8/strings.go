package lab8

import (
	"io"
	"os"
	"strings"
)

func CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	file.Close()
}

func WriteFile(path, text string) {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(text)
}

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	var text string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		text = string(data[:n])
	}
	return text
}

func SearchText(path, searchText string) bool {
	text := ReadFile(path)

	search := strings.Contains(text, searchText)
	return search
}
