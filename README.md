# numnam
[![Om Nom in their superhero outfit sitting on a race car during cartoon episode Mechanic Rodeo at 3:13](/omnom.webp "Om Nom")](https://youtu.be/HrNzPNcoGiQ?t=188)

A Go package to convert integers into words.

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/numnam.svg)](https://pkg.go.dev/github.com/speedyhoon/numnam)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/numnam)](https://goreportcard.com/report/github.com/speedyhoon/numnam)

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

## Name
`numnam` short for "number name", is something you'd expect [Om Nom](https://cuttherope.fandom.com/wiki/Om_Nom) to say while eating any delicious 🍬, 🍭, 🍫, 🍕 or junk food.
