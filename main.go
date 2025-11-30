package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"string-unpacker/unpack"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку для распаковки:")

	// Цикл чтения строк, пока пользователь не введет EOF или не закроет ввод
	for scanner.Scan() {
		input := scanner.Text()

		result, err := unpack.String(input)

		if err != nil {
			// В случае ошибки выводим её и переходим к следующей строке ввода
			fmt.Printf("Ошибка обработки строки \"%s\": %v\n", input, err)
		} else {
			// В случае успеха выводим результат
			fmt.Println(result)
		}

		fmt.Println("\nВведите следующую строку:")
	}

	if err := scanner.Err(); err != nil {
		// Обработка ошибок сканера
		log.Fatal("Ошибка чтения ввода:", err)
	}

}
