package numnam

import (
	"fmt"
	"math"
	"testing"
)

func Test_pow1000(t *testing.T) {
	tests := map[int]uint{
		0: 1,
		1: 1000,
		2: 1000000,
		3: 1000000000,
		4: 1000000000000,
		5: 1000000000000000,
		6: 1000000000000000000,
		7: math.MaxUint64,
		8: math.MaxUint64,
	}
	for input, expected := range tests {
		t.Run(fmt.Sprintf("[%d]", input), func(t *testing.T) {
			if got := pow1000(input); got != expected {
				t.Errorf("pow1000() = %v, want %v", got, expected)
			}
		})
	}
}
