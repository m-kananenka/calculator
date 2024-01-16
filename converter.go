package main

import "fmt"

var romanNumerals = map[string]int{
	"I": 1, "V": 5, "X": 10, "C": 100,
}

func romanToArabic(s string) (int, error) {
	conversion := 0
	for i := 0; i < len(s); i++ {
		romanChar := string(s[i])
		if _, ok := romanNumerals[romanChar]; !ok {
			return -1, fmt.Errorf("ошибка: не верная римская цифра")
		}
		number := romanNumerals[romanChar]
		if i < len(s)-1 && number < romanNumerals[string(s[i+1])] {
			conversion -= number
		} else {
			conversion += number
		}
	}
	return conversion, nil
}

func arabicToRoman(num int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for num >= conversion.value {
			roman += conversion.digit
			num -= conversion.value
		}
	}
	return roman
}
