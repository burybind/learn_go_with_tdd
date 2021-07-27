package propertyBasedTests

import (
	"strings"
)

type RomanNumeral struct {
	Value uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals {
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
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

func ConvertToRoman(arabic uint16) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	total, prevVal, currentVal := uint16(0), uint16(0), uint16(0)

	for _, romanChar := range roman {
		currentVal = allRomanNumerals.ValueOf(string(romanChar))
		if currentVal > prevVal {
			total -= prevVal
		} else {
			total += prevVal
		}
		prevVal = currentVal
	}

	total += prevVal
	return total
}

func (r RomanNumerals) ValueOf(symbol string) uint16 {
	for _, numeral := range r {
		if numeral.Symbol == symbol {
			return numeral.Value
		}
	}
	return 0
}


