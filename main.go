package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	rome   = "Rome"
	arabic = "Arabic"
)

type Arg struct {
	val       int
	typeOfVal string
}

func main() {
	for {
		fmt.Println("Input")

		writer := bufio.NewScanner(os.Stdin)
		writer.Scan()
		input := writer.Text()

		fields := strings.Fields(input)

		validateLenInput(fields)

		firstArg := parseArg(fields[0])
		secondArg := parseArg(fields[2])

		validateArgs(firstArg.val, secondArg.val)

		if isSameTypes(firstArg, secondArg) {
			log.Println("ошибка:так как используются одновременно разные системы счисления.")
			return
		}

		operator := fields[1]
		result := calculate(firstArg.val, secondArg.val, operator)

		if firstArg.typeOfVal == rome && secondArg.typeOfVal == rome {
			if result < 1 {
				fmt.Println("Output")
				panic("ошибка: так как в римской системе нет 0 и отрицательных чисел.")
			}
			fmt.Println("Output")
			fmt.Println(arabicToRoman(result))
			continue
		}

		fmt.Println("Output")
		fmt.Println(result)
	}
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("деление на 0 запрещено")
		}
		return a / b
	default:
		panic("неправильный операнд")
	}
}

func parseArg(s string) *Arg {
	var result *Arg
	var err error

	result, err = parseArgsRome(s)
	if err != nil {
		result, err = parseArgsArabic(s)
		if err != nil {
			panic("не подходящее число")
		}
	}
	return result
}

func parseArgsRome(s string) (*Arg, error) {
	toArabic, err := romanToArabic(s)
	if err != nil {
		return nil, err
	}
	return &Arg{
		val:       toArabic,
		typeOfVal: rome,
	}, nil
}

func parseArgsArabic(s string) (*Arg, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &Arg{
		val:       n,
		typeOfVal: arabic,
	}, nil
}
