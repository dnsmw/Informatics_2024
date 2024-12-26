package lab8

import (
	"bufio"
	"fmt"
	"os"
)

const Path = "Lab8/text.txt"

func RunLab8() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)

	CreateFile(Path)

	fmt.Println("Введите текст который будет введён")
	TextForWrite, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	WriteFile(Path, TextForWrite)

	text := ReadFile(Path)
	fmt.Println(text)

	fmt.Println("Введите текст для поиска")
	TextForSearch, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	SearchText(Path, TextForSearch)

	Task1()
}
