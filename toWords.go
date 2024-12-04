package numnam

import (
	"math"
	"strings"
)

// ToWordsU converts a uint into a string of space separated words.
// For example, `865` returns `eight hundred sixty-five`.
func ToWordsU(num uint) string {
	if num == 0 {
		return strings.ToLower(Zero)
	}
	return strings.ToLower(strings.Join(toWords(num), " "))
}

// ToWords converts an int into a string of space separated words.
// For example, `-865` returns `minus eight hundred sixty-five`.
func ToWords(num int) string {
	switch {
	case num == 0:
		return strings.ToLower(Zero)
	case num == math.MinInt:
		return strings.ToLower(Minus + " " + strings.Join(toWords(uint(num)), " "))
	case num < 0:
		return strings.ToLower(Minus + " " + strings.Join(toWords(uint(-num)), " "))
	}
	return strings.ToLower(strings.Join(toWords(uint(num)), " "))
}

func toWords(num uint) []string {
	switch {
	case num == 0:
		return nil
	case num < 20:
		return []string{To19[num]}
	case num < 100:
		s := []string{Tens[num/10-2]}
		if num%10 != 0 {
			s[0] += "-" + To19[num%10]
		}
		return s
	case num < 1000:
		return append([]string{To19[num/100], Hundred}, toWords(num%100)...)
	}
	for i, word := range Large {
		po := pow1000(i + 2)
		if num < po || po == num && num == math.MaxUint {
			po = pow1000(i + 1)
			return append(append(toWords(num/po), word), toWords(num%po)...)
		}
	}
	return nil // Unhandled number.
}
