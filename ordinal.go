package numnam

import (
	"math"
	"strconv"
)

const first, second, third, ten, eleven, twelve, thirteen, hundred = 1, 2, 3, 10, 11, 12, 13, 100

// Ordinal returns the position integer as a rank/ordinal string.
// For example, Ordinal(3) outputs "3rd".
func Ordinal(position uint64) (ord string) {
	ord = strconv.FormatUint(position, ten)

	switch position % ten {
	case first:
		if position%hundred != eleven {
			return ord + "st"
		}
	case second:
		if position%hundred != twelve {
			return ord + "nd"
		}
	case third:
		if position%hundred != thirteen {
			return ord + "rd"
		}
	}

	return ord + "th"
}

// OrdinalEqual returns the position integer as a rank/ordinal string.
// For example, OrdinalEqual(3, true) outputs "=3rd".
// When isEqual is true, OrdinalEqual prefixes the output with an "=".
func OrdinalEqual(position uint64, isEqual bool) (ord string) {
	if isEqual {
		ord = "="
	}

	return ord + Ordinal(position)
}

// Ordinal0 returns the position integer as a rank/ordinal string where zero is the first position.
// For example, Ordinal0(3) outputs "4th".
// It is suitable for loops starting at zero, without having to manually increment or pass i+1 to Ordinal0.
func Ordinal0(position uint64) (ord string) {
	if position == math.MaxUint64 {
		return "18446744073709551616th"
	}
	return Ordinal(position + 1)
}

// Ordinal0Equal returns the position integer as a rank/ordinal string where zero is the first position.
// For example, Ordinal0Equal(3, true) outputs "=4th".
// When isEqual is true, Ordinal0Equal prefixes the output with an "=".
// It is suitable for loops starting at zero, without having to manually increment or pass i+1 to Ordinal0Equal.
func Ordinal0Equal(position uint64, isEqual bool) (ord string) {
	if isEqual {
		ord = "="
	}

	return ord + Ordinal0(position)
}
