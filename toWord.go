// Package numnam converts integers into words.
package numnam

import "math"

// Exported these variables so other languages can be translated, or changed to Long Scale and other integer naming systems.
var (
	Minus = "Minus"
	Zero  = "Zero" // Zero is hardcoded to reduce processing.
	To19  = []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
	}
	Tens    = []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	Hundred = "Hundred"
	Large   = []string{"Thousand", "Million", "Billion", "Trillion", "Quadrillion", "Quintillion"}
)

// ToWordU converts a uint into a snake-cased string.
// For example, `1234` returns `OneThousandTwoHundredThirtyFour`.
func ToWordU(num uint) string {
	if num == 0 {
		return Zero
	}
	return toWord(num)
}

// ToWord converts an int into a snake-cased string.
// For example, `-1234` returns `MinusOneThousandTwoHundredThirtyFour`.
func ToWord(num int) string {
	switch {
	case num == 0:
		return Zero
	case num == math.MinInt:
		return Minus + toWord(uint(num))
	case num < 0:
		return Minus + toWord(uint(-num))
	}
	return toWord(uint(num))
}

func toWord(num uint) string {
	switch {
	case num == 0:
		return ""
	case num < 20:
		return To19[num]
	case num < 100:
		return Tens[num/10-2] + toWord(num%10)
	case num < 1000:
		return To19[num/100] + Hundred + toWord(num%100)
	}
	for i, word := range Large {
		po := pow1000(i + 2)
		if num < po || po == num && num == math.MaxUint {
			po = pow1000(i + 1)
			return toWord(num/po) + word + toWord(num%po)
		}
	}
	return "" // Unhandled number.
}

func pow1000(p int) uint {
	if p >= 7 {
		return math.MaxUint
	}
	return uint(math.Pow(1000, float64(p)))
}
