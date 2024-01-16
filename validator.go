package main

func isSameTypes(firstArg *Arg, secondArg *Arg) bool {
	if firstArg.typeOfVal != secondArg.typeOfVal {
		return true
	}
	return false
}

func validateLenInput(fields []string) {
	if len(fields) > 3 {
		panic("ошибка: так как формат математической операции не удовлетворяет заданию — " +
			"два операнда и один оператор (+, -, /, *).")
	}

	if len(fields) < 3 {
		panic("ошибка:так как строка не является математической операцией.")
	}
}

func validateArgs(a, b int) {
	if a > 10 || a < 1 || b > 10 || b < 1 {
		panic("калькулятор должен принимать на вход числа от 1 до 10 включительно")
	}
}
