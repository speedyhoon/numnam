# numnam
[![Om Nom in their superhero outfit sitting on a race car during cartoon episode Mechanic Rodeo at 3:13](/omnom.webp "Om Nom")](https://youtu.be/HrNzPNcoGiQ?t=188)

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/numnam.svg)](https://pkg.go.dev/github.com/speedyhoon/numnam)
[![Go Report Card](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/goReport.svg)](https://goreportcard.com/report/github.com/speedyhoon/numnam)
![license AGPL3](https://raw.githubusercontent.com/speedyhoon/speedyhoon/refs/heads/main/AGPL3.svg)

A Go package to convert integers into words and ordinal suffixes (`"st"`, `"nd"`, `"rd"`, `"th"`).

## Example
```go
package main

import (
	"math"

	"github.com/speedyhoon/numnam"
)

func main() {
	println(numnam.ToWord(-999))
	println(numnam.ToWordU(1234))
	println(numnam.ToWords(-865))
	println(numnam.ToWordsU(103582467))
	println(numnam.ToWordsU(math.MaxUint))
	println(numnam.Ordinal(2))
	println(numnam.OrdinalEqual(5, true))
	println(numnam.Ordinal0(0))
	println(numnam.Ordinal0Equal(7, false))
}
```
**Output:**
```text
MinusNineHundredNinetyNine
OneThousandTwoHundredThirtyFour
minus eight hundred sixty-five
one hundred three million five hundred eighty-two thousand four hundred sixty-seven
eighteen quintillion four hundred forty-six quadrillion seven hundred forty-four trillion seventy-three billion seven hundred nine million five hundred fifty-one thousand six hundred fifteen
2nd
=5th
1st
8th
```

## Name
`numnam` short for "number name", is something you'd expect [Om Nom](https://cuttherope.fandom.com/wiki/Om_Nom) to say while eating any delicious 🍬, 🍭, 🍫, 🍕 or junk food.
