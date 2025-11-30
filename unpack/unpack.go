package unpack

import (
	"fmt"
	"strings"
	"unicode"
)

// String осуществляет примитивную распаковку строки.
func String(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var builder strings.Builder
	runes := []rune(s)
	n := len(runes)

	lastWasMultiplier := false // был ли предыдущий неэкранированный символ — управляющей цифрой?

	for i := 0; i < n; i++ {
		r := runes[i]

		// Обработка экранирования
		if r == '\\' {
			if i+1 >= n {
				return "", fmt.Errorf("некорректная строка: неполная escape-последовательность в конце")
			}
			i++                         // Переходим к следующему символу после '\'
			builder.WriteRune(runes[i]) // Записываем экранированный символ
			lastWasMultiplier = false
			continue
		}

		// Обработка неэкранированной цифры
		if unicode.IsDigit(r) {
			// Цифра должна следовать за чем-то, иначе это ошибка
			if builder.Len() == 0 {
				return "", fmt.Errorf("некорректная строка: цифра в начале строки")
			}

			// Цифра не может идти сразу после другой неэкранированной цифры
			if lastWasMultiplier {
				return "", fmt.Errorf("некорректная строка: цифра следует за цифрой")
			}

			// Преобразуем руну-цифру в число.
			count := int(r - '0')
			// Получаем последний символ из builder
			lastRune := []rune(builder.String())
			lastChar := lastRune[len(lastRune)-1]

			// Повторяем этот символ нужное количество раз.
			for j := 0; j < count-1; j++ {
				builder.WriteRune(lastChar)
			}

			lastWasMultiplier = true
		} else {
			// Обычный символ
			builder.WriteRune(r)
			lastWasMultiplier = false
		}
	}

	return builder.String(), nil
}
