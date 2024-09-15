# numnam
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/numnam)](https://goreportcard.com/report/github.com/speedyhoon/numnam)

A package to convert integers into words.

## Example
```go
package main

import (
	"github.com/speedyhoon/numnam"
	"math"
)

func main() {
	println(numnam.ToWord(-999))
	println(numnam.ToWordU(1234))
	println(numnam.ToWords(-865))
	println(numnam.ToWordsU(103582467))
	println(numnam.ToWordsU(math.MaxUint))
}
```
**Output:**
```text
MinusNineHundredNinetyNine
OneThousandTwoHundredThirtyFour
minus eight hundred sixty-five
one hundred three million five hundred eighty-two thousand four hundred sixty-seven
eighteen quintillion four hundred forty-six quadrillion seven hundred forty-four trillion seventy-three billion seven hundred nine million five hundred fifty-one thousand six hundred fifteen
```
