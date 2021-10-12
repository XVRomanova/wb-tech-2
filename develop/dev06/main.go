/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Flag struct {
	fields       string
	delimiter    string
	withSeparate bool
}

func main() {
	fields := flag.String("f", "-1", "select only these fields")
	delimiter := flag.String("d", "\t", "use DELIM instead of TAB for field delimiter")
	withSeparate := flag.Bool("s", false, "do not print lines not containing delimiters")

	flag.Parse()
	if *fields == "-1" {
		fmt.Println("incorrect value of -f")
		return
	}

	param := Flag{
		fields:       *fields,
		delimiter:    *delimiter,
		withSeparate: *withSeparate,
	}

	lines, err := readDataFromStdin()
	if err != nil {
		fmt.Println(err)
	}

	output, err := cut(param, lines)
	if err != nil {
		fmt.Println(err)
	}

	for _, str := range output {
		fmt.Print(str)
	}
}

func cut(param Flag, input []string) ([]string, error) {

	var output []string

	for _, line := range input {

		if !strings.Contains(line, param.delimiter) && param.withSeparate {
			continue
		}

		if !strings.Contains(line, param.delimiter) && !param.withSeparate {
			output = append(output, line)
			output = append(output, "\n")
			//fmt.Println(line)
			continue
		}

		columns := strings.Split(line, param.delimiter)

		fields, err := parseFieldsToInt(param.fields)
		if err != nil {
			return nil, err
		}

		if len(fields) > 1 {
			for _, columnNumber := range fields {
				output = append(output, columns[columnNumber-1])
				output = append(output, param.delimiter)
				//fmt.Print(columns[columnNumber-1] + param.delimiter)
			}
			output = append(output, "\n")
			//fmt.Println()

		} else {
			columnNumber := fields[0] - 1
			output = append(output, columns[columnNumber])
			output = append(output, "\n")
			//fmt.Println(columns[columnNumber])
		}

	}

	return output, nil
}

func readDataFromStdin() ([]string, error) {
	var listOfStrings []string

	fmt.Print("Enter text:\n")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			listOfStrings = append(listOfStrings, text)
		} else {
			break
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return listOfStrings, nil
}

func parseFieldsToInt(fields string) ([]int, error) {
	var listOfFields []int

	if strings.ContainsAny(fields, "-") {
		numbers, err := parseStrToInt(fields, "-")
		if err != nil {
			return nil, err
		}

		firstField := numbers[0]
		lastField := numbers[1]
		for i := firstField; i < lastField+1; i++ {
			listOfFields = append(listOfFields, i)
		}
		return listOfFields, nil
	} else if strings.ContainsAny(fields, ",") {
		listOfFields, err := parseStrToInt(fields, ",")
		if err != nil {
			return nil, err
		}
		return listOfFields, nil
	} else {
		listOfFields, err := parseStrToInt(fields, "")
		if err != nil {
			return nil, err
		}
		return listOfFields, nil
	}

}

func parseStrToInt(fields string, sep string) ([]int, error) {
	var numbers []int

	listOfStrings := strings.Split(fields, sep)
	for _, currentString := range listOfStrings {
		num, err := strconv.Atoi(currentString)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
