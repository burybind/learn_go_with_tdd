package propertyBasedTests

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"testing/quick"
)

var cases = []struct {
	desc   string
	arabic uint16
	roman  string
}{
	{desc: "1 gets converted to I", arabic: 1, roman: "I"},
	{desc: "2 gets converted to II", arabic: 2, roman: "II"},
	{desc: "3 gets converted to III", arabic: 3, roman: "III"},
	{desc: "4 gets converted to IV (can't repeat more than 3 times)", arabic: 4, roman: "IV"},
	{desc: "5 gets converted to V", arabic: 5, roman: "V"},
	{desc: "6 gets converted to VI", arabic: 6, roman: "VI"},
	{desc: "7 gets converted to VII", arabic: 7, roman: "VII"},
	{desc: "8 gets converted to VIII", arabic: 8, roman: "VIII"},
	{desc: "9 gets converted to IX", arabic: 9, roman: "IX"},
	{desc: "10 gets converted to X", arabic: 10, roman: "X"},
	{desc: "14 gets converted to XIV", arabic: 14, roman: "XIV"},
	{desc: "18 gets converted to XVIII", arabic: 18, roman: "XVIII"},
	{desc: "20 gets converted to XX", arabic: 20, roman: "XX"},
	{desc: "39 gets converted to XXXIX", arabic: 39, roman: "XXXIX"},
	{desc: "40 gets converted to XL", arabic: 40, roman: "XL"},
	{desc: "47 gets converted to XLVII", arabic: 47, roman: "XLVII"},
	{desc: "49 gets converted to XLIX", arabic: 49, roman: "XLIX"},
	{desc: "50 gets converted to L", arabic: 50, roman: "L"},
	{"90 gets converted to XC", 90, "XC"},
	{"100 gets converted to C", 100, "C"},
	{"", 400, "CD"},
	{"", 500, "D"},
	{"", 900, "CM"},
	{"", 1000, "M"},
	{"", 1984, "MCMLXXXIV"},
	{"", 3999, "MMMCMXCIX"},
	{"", 2014, "MMXIV"},
	{"", 1006, "MVI"},
	{"", 798, "DCCXCVIII"},
}

// these are "example"-based tests where we provide the tooling some examples around our code to verify
func TestRomanNumerals(t *testing.T) {

	for _, tt := range cases {
		if tt.desc == "" {
			tt.desc = strconv.Itoa(int(tt.arabic))
		}
		t.Run(tt.desc, func(t *testing.T) {
			got := ConvertToRoman(tt.arabic)

			if got != tt.roman {
				t.Errorf("got: %q, roman: %q", got, tt.roman)
			}
		})
	}
}

// these are "example"-based tests where we provide the tooling some examples around our code to verify
func TestConvertToArabic(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", tt.roman, tt.arabic), func(t *testing.T) {
			got := ConvertToArabic(tt.roman)
			if got != tt.arabic {
				t.Errorf("got %d, want %d", got, tt.arabic)
			}
		})
	}
}


/*
property-based tests help test rules we know about our domain by throwing random data at our code and verifying the
rules you describe always hold true
 */
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic // converting to roman then back to int should yield the same number as was started with
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount:      1000,
		Values: generateValidValuesForRoman,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

// generateValidValuesForRoman will lock the tested numbers down to only uints from 0 to 3999 as that is all that is supported
func generateValidValuesForRoman(values []reflect.Value, rand *rand.Rand) {
	// this function needs to fill in the values slice. The values slice has a length equal to quick check's `MaxCount` config setting
	for i := 0; i < len(values); i++ {
		values[i] = reflect.ValueOf(uint16(rand.Intn(3999)))
	}
	return
}
