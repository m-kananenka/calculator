package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArgsArabic_Positive(t *testing.T) {
	input := "123"
	expected := &Arg{
		val:       123,
		typeOfVal: arabic,
	}

	result, err := parseArgsArabic(input)
	assert.NoError(t, err)
	assert.Equal(t, result, expected)
}

func TestParseArgsArabic_Negative(t *testing.T) {
	input := "abc"

	_, err := parseArgsArabic(input)
	assert.Error(t, err)
}

func TestParseArgsRoman_Positive(t *testing.T) {
	input := "I"
	expected := &Arg{
		val:       1,
		typeOfVal: rome,
	}

	result, err := parseArgsRome(input)
	assert.NoError(t, err)
	assert.Equal(t, result, expected)
}
func TestParseArgsRoman_Negative(t *testing.T) {
	input := "abc"

	_, err := parseArgsRome(input)
	assert.Error(t, err)
}

func TestParseArgs_Positive(t *testing.T) {
	input := "I"
	expected := &Arg{
		val:       1,
		typeOfVal: rome,
	}

	result := parseArg(input)
	assert.Equal(t, result, expected)
}

func TestCalculate_Positive(t *testing.T) {
	a := 1
	b := 2
	operand := "+"
	expected := 3

	result := calculate(a, b, operand)
	assert.Equal(t, result, expected)
}

func TestCalculate_Positive2(t *testing.T) {
	a := 1
	b := 2
	operand := "-"
	expected := -1

	result := calculate(a, b, operand)
	assert.Equal(t, result, expected)
}

func TestCalculate_Positive3(t *testing.T) {
	a := 1
	b := 2
	operand := "*"
	expected := 2

	result := calculate(a, b, operand)
	assert.Equal(t, result, expected)
}

func TestCalculate_Positive4(t *testing.T) {
	a := 2
	b := 2
	operand := "/"
	expected := 1

	result := calculate(a, b, operand)
	assert.Equal(t, result, expected)
}
