/*
Задача на распаковку.

Создать Go функцию, осуществляющую примитивную распаковку строки,
содержащую повторяющиеся символы / руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительное задание: поддержка escape - последовательностей
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку.
Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
 */
package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(input string) (string, error) {
	runes := []rune(input)

	if len(runes) == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", errors.New("incorrect input")
	}

	var outputBuilder strings.Builder
	numberChar := ""
	var currentChar string
	var prevChar string
	var isEscapeChar bool

	for i, runeCode := range runes {

		currentChar = string(runeCode)

		switch {

		case isEscapeChar:

			if prevChar == "\\" {
				if currentChar != "\\" {
					outputBuilder.WriteString(currentChar)
					prevChar = string(runeCode)
				} else {
					outputBuilder.WriteString(currentChar)
					prevChar = ""
					continue
				}
			} else if _, err := strconv.Atoi(prevChar); err == nil {
				numberChar += string(runeCode)
				writeChar(prevChar, numberChar, &outputBuilder)
			}

			if currentChar == "\\" {
				prevChar = "\\"
			}

			if prevChar == "" {
				numberChar += string(runeCode)
				writeChar("\\", numberChar, &outputBuilder)
			}

		case unicode.IsLetter(runeCode):

			if numberChar != "" {
				writeChar(prevChar, numberChar, &outputBuilder)
				numberChar = ""
			}

			outputBuilder.WriteString(currentChar)

		case unicode.IsDigit(runeCode):
			prevChar = currentChar
			numberChar += string(runeCode)

			if i == len(runes)-1 {
				writeChar(prevChar, numberChar, &outputBuilder)
			}

		default:
			isEscapeChar = true
			prevChar = string(runeCode)
		}
	}

	return outputBuilder.String(), nil
}

func writeChar(prevChar string, numberChar string, outputBuilder *strings.Builder) {
	num, _ := strconv.Atoi(numberChar)
	for i := 0; i < num-1; i++ {
		outputBuilder.WriteString(prevChar)
	}
}
