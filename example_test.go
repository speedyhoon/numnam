package numnam_test

import (
	"fmt"
	"github.com/speedyhoon/numnam"
)

func ExampleToWord() {
	fmt.Println(numnam.ToWord(-999))
	// Output: MinusNineHundredNinetyNine
}

func ExampleToWordU() {
	fmt.Println(numnam.ToWordU(1234))
	// Output: OneThousandTwoHundredThirtyFour
}

func ExampleToWords() {
	fmt.Println(numnam.ToWords(-865))
	// Output: minus eight hundred sixty-five
}

func ExampleToWordsU() {
	fmt.Println(numnam.ToWordsU(103582467))
	// Output: one hundred three million five hundred eighty-two thousand four hundred sixty-seven
}
