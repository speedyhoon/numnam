package numnam_test

import (
	"fmt"
	"github.com/speedyhoon/numnam"
	"math"
	"testing"
)

var positiveInts2 = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",

	math.MaxInt8:   "one hundred twenty-seven",
	math.MaxUint8:  "two hundred fifty-five",
	math.MaxInt16:  "thirty-two thousand seven hundred sixty-seven",
	math.MaxUint16: "sixty-five thousand five hundred thirty-five",
	math.MaxInt32:  "two billion one hundred forty-seven million four hundred eighty-three thousand six hundred forty-seven",                                                                                               // 2,147,483,647
	math.MaxUint32: "four billion two hundred ninety-four million nine hundred sixty-seven thousand two hundred ninety-five",                                                                                               // 4,294,967,295
	math.MaxInt64:  "nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven", // 9,223,372,036,854,775,807

	77:                  "seventy-seven",
	99:                  "ninety-nine",
	100:                 "one hundred",
	777:                 "seven hundred seventy-seven",
	999:                 "nine hundred ninety-nine",
	1000:                "one thousand",
	7777:                "seven thousand seven hundred seventy-seven",
	9999:                "nine thousand nine hundred ninety-nine",
	10000:               "ten thousand",
	77777:               "seventy-seven thousand seven hundred seventy-seven",
	99999:               "ninety-nine thousand nine hundred ninety-nine",
	100000:              "one hundred thousand",
	777777:              "seven hundred seventy-seven thousand seven hundred seventy-seven",
	999999:              "nine hundred ninety-nine thousand nine hundred ninety-nine",
	1000000:             "one million",
	7777777:             "seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	9999999:             "nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	10000000:            "ten million",
	77777777:            "seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	99999999:            "ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	100000000:           "one hundred million",
	777777777:           "seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	999999999:           "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	1000000000:          "one billion",
	7777777777:          "seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	9999999999:          "nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	10000000000:         "ten billion",
	77777777777:         "seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	99999999999:         "ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	100000000000:        "one hundred billion",
	777777777777:        "seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	999999999999:        "nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	1000000000000:       "one trillion",
	7777777777777:       "seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	9999999999999:       "nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	10000000000000:      "ten trillion",
	77777777777777:      "seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	99999999999999:      "ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	100000000000000:     "one hundred trillion",
	777777777777777:     "seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	999999999999999:     "nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	1000000000000000:    "one quadrillion",
	7777777777777777:    "seven quadrillion seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	9999999999999999:    "nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	10000000000000000:   "ten quadrillion",
	77777777777777777:   "seventy-seven quadrillion seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	99999999999999999:   "ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	100000000000000000:  "one hundred quadrillion",
	777777777777777777:  "seven hundred seventy-seven quadrillion seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	999999999999999999:  "nine hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
	1000000000000000000: "one quintillion",
	1777777777777777777: "one quintillion seven hundred seventy-seven quadrillion seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
	1999999999999999999: "one quintillion nine hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
}

func TestToWordsU(t *testing.T) {
	tests := map[uint]string{
		0:                    "zero",
		7777777777777777777:  "seven quintillion seven hundred seventy-seven quadrillion seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
		9999999999999999999:  "nine quintillion nine hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		10000000000000000000: "ten quintillion",
		16999999999999999999: "sixteen quintillion nine hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		17777777777777777777: "seventeen quintillion seven hundred seventy-seven quadrillion seven hundred seventy-seven trillion seven hundred seventy-seven billion seven hundred seventy-seven million seven hundred seventy-seven thousand seven hundred seventy-seven",
		17999999999999999999: "seventeen quintillion nine hundred ninety-nine quadrillion nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		math.MaxUint64:       "eighteen quintillion four hundred forty-six quadrillion seven hundred forty-four trillion seventy-three billion seven hundred nine million five hundred fifty-one thousand six hundred fifteen", // 18,446,744,073,709,551,615
	}
	for num, expected := range tests {
		t.Run(fmt.Sprintf("uint=%d", num), func(t *testing.T) {
			if got := numnam.ToWordsU(num); got != expected {
				t.Errorf("ToWordsU():\n%v, want:\n %v", got, expected)
			}
		})
	}

	for num, expected := range positiveInts2 {
		t.Run(fmt.Sprintf("positiveInt=%d", num), func(t *testing.T) {
			if got := numnam.ToWordsU(uint(num)); got != expected {
				t.Errorf("ToWordU():\n%v, want:\n %v", got, expected)
			}
		})
	}
}

func TestToWords(t *testing.T) {
	for num, expected := range positiveInts2 {
		t.Run(fmt.Sprintf("positiveInts=%d", num), func(t *testing.T) {
			if got := numnam.ToWords(num); got != expected {
				t.Errorf("ToWordsU():\n%v, want:\n %v", got, expected)
			}
		})

		t.Run(fmt.Sprintf("negativeInts=%d", num), func(t *testing.T) {
			expected = "minus " + expected
			if got := numnam.ToWords(-num); got != expected {
				t.Errorf("ToWordsU():\n%v, want:\n %v", got, expected)
			}
		})
	}

	others := map[int]string{
		0:             "zero",
		math.MinInt8:  "minus one hundred twenty-eight",
		math.MinInt16: "minus thirty-two thousand seven hundred sixty-eight",
		math.MinInt32: "minus two billion one hundred forty-seven million four hundred eighty-three thousand six hundred forty-eight",
		math.MinInt64: "minus nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred eight",
	}
	for num, expected := range others {
		t.Run(fmt.Sprintf("others=%d", num), func(t *testing.T) {
			if got := numnam.ToWords(num); got != expected {
				t.Errorf("ToWordsU():\n%v, want:\n %v", got, expected)
			}
		})
	}
}
