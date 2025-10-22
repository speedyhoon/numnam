package numnam_test

import (
	"fmt"
	"github.com/speedyhoon/numnam"
	"math"
	"testing"
)

var positiveInts = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",

	math.MaxInt8:   "OneHundredTwentySeven",
	math.MaxUint8:  "TwoHundredFiftyFive",
	math.MaxInt16:  "ThirtyTwoThousandSevenHundredSixtySeven",
	math.MaxUint16: "SixtyFiveThousandFiveHundredThirtyFive",
	math.MaxInt32:  "TwoBillionOneHundredFortySevenMillionFourHundredEightyThreeThousandSixHundredFortySeven",                                                                                   // 2,147,483,647
	math.MaxUint32: "FourBillionTwoHundredNinetyFourMillionNineHundredSixtySevenThousandTwoHundredNinetyFive",                                                                                   // 4,294,967,295
	math.MaxInt64:  "NineQuintillionTwoHundredTwentyThreeQuadrillionThreeHundredSeventyTwoTrillionThirtySixBillionEightHundredFiftyFourMillionSevenHundredSeventyFiveThousandEightHundredSeven", // 9,223,372,036,854,775,807

	77:                  "SeventySeven",
	99:                  "NinetyNine",
	100:                 "OneHundred",
	777:                 "SevenHundredSeventySeven",
	999:                 "NineHundredNinetyNine",
	1000:                "OneThousand",
	7777:                "SevenThousandSevenHundredSeventySeven",
	9999:                "NineThousandNineHundredNinetyNine",
	10000:               "TenThousand",
	77777:               "SeventySevenThousandSevenHundredSeventySeven",
	99999:               "NinetyNineThousandNineHundredNinetyNine",
	100000:              "OneHundredThousand",
	777777:              "SevenHundredSeventySevenThousandSevenHundredSeventySeven",
	999999:              "NineHundredNinetyNineThousandNineHundredNinetyNine",
	1000000:             "OneMillion",
	7777777:             "SevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	9999999:             "NineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	10000000:            "TenMillion",
	77777777:            "SeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	99999999:            "NinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	100000000:           "OneHundredMillion",
	777777777:           "SevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	999999999:           "NineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	1000000000:          "OneBillion",
	7777777777:          "SevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	9999999999:          "NineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	10000000000:         "TenBillion",
	77777777777:         "SeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	99999999999:         "NinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	100000000000:        "OneHundredBillion",
	777777777777:        "SevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	999999999999:        "NineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	1000000000000:       "OneTrillion",
	7777777777777:       "SevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	9999999999999:       "NineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	10000000000000:      "TenTrillion",
	77777777777777:      "SeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	99999999999999:      "NinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	100000000000000:     "OneHundredTrillion",
	777777777777777:     "SevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	999999999999999:     "NineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	1000000000000000:    "OneQuadrillion",
	7777777777777777:    "SevenQuadrillionSevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	9999999999999999:    "NineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	10000000000000000:   "TenQuadrillion",
	77777777777777777:   "SeventySevenQuadrillionSevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	99999999999999999:   "NinetyNineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	100000000000000000:  "OneHundredQuadrillion",
	777777777777777777:  "SevenHundredSeventySevenQuadrillionSevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	999999999999999999:  "NineHundredNinetyNineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
	1000000000000000000: "OneQuintillion",
	1777777777777777777: "OneQuintillionSevenHundredSeventySevenQuadrillionSevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
	1999999999999999999: "OneQuintillionNineHundredNinetyNineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
}

func TestToWordU(t *testing.T) {
	tests := map[uint]string{
		0:                    "Zero",
		7777777777777777777:  "SevenQuintillionSevenHundredSeventySevenQuadrillionSevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
		9999999999999999999:  "NineQuintillionNineHundredNinetyNineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
		10000000000000000000: "TenQuintillion",
		16999999999999999999: "SixteenQuintillionNineHundredNinetyNineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
		17777777777777777777: "SeventeenQuintillionSevenHundredSeventySevenQuadrillionSevenHundredSeventySevenTrillionSevenHundredSeventySevenBillionSevenHundredSeventySevenMillionSevenHundredSeventySevenThousandSevenHundredSeventySeven",
		17999999999999999999: "SeventeenQuintillionNineHundredNinetyNineQuadrillionNineHundredNinetyNineTrillionNineHundredNinetyNineBillionNineHundredNinetyNineMillionNineHundredNinetyNineThousandNineHundredNinetyNine",
		math.MaxUint64:       "EighteenQuintillionFourHundredFortySixQuadrillionSevenHundredFortyFourTrillionSeventyThreeBillionSevenHundredNineMillionFiveHundredFiftyOneThousandSixHundredFifteen", // 18,446,744,073,709,551,615
	}
	for num, expected := range tests {
		t.Run(fmt.Sprintf("uint=%d", num), func(t *testing.T) {
			if got := numnam.ToWordU(num); got != expected {
				t.Errorf("ToWordU():\n%v, want:\n %v", got, expected)
			}
		})
	}

	for num, expected := range positiveInts {
		t.Run(fmt.Sprintf("positiveInt=%d", num), func(t *testing.T) {
			if got := numnam.ToWordU(uint(num)); got != expected {
				t.Errorf("ToWordU():\n%v, want:\n %v", got, expected)
			}
		})
	}
}

func TestToWord(t *testing.T) {
	for num, expected := range positiveInts {
		t.Run(fmt.Sprintf("positiveInts=%d", num), func(t *testing.T) {
			if got := numnam.ToWord(num); got != expected {
				t.Errorf("ToWord():\n%v, want:\n %v", got, expected)
			}
		})

		t.Run(fmt.Sprintf("negativeInts=%d", num), func(t *testing.T) {
			expected = "Minus" + expected
			if got := numnam.ToWord(-num); got != expected {
				t.Errorf("ToWord():\n%v, want:\n %v", got, expected)
			}
		})
	}

	others := map[int]string{
		0:             "Zero",
		math.MinInt8:  "MinusOneHundredTwentyEight",
		math.MinInt16: "MinusThirtyTwoThousandSevenHundredSixtyEight",
		math.MinInt32: "MinusTwoBillionOneHundredFortySevenMillionFourHundredEightyThreeThousandSixHundredFortyEight",
		math.MinInt64: "MinusNineQuintillionTwoHundredTwentyThreeQuadrillionThreeHundredSeventyTwoTrillionThirtySixBillionEightHundredFiftyFourMillionSevenHundredSeventyFiveThousandEightHundredEight",
	}
	for num, expected := range others {
		t.Run(fmt.Sprintf("others=%d", num), func(t *testing.T) {
			if got := numnam.ToWord(num); got != expected {
				t.Errorf("ToWord():\n%v, want:\n %v", got, expected)
			}
		})
	}
}
