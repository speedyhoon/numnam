package numnam_test

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/speedyhoon/numnam"
)

func TestOrdinal(t *testing.T) {
	tests := []struct {
		position uint64
		want     string
	}{
		{position: 0, want: "0th"},
		{position: 18446744073709551598, want: "18446744073709551598th"},
		{position: 18446744073709551599, want: "18446744073709551599th"},
		{position: 18446744073709551600, want: "18446744073709551600th"},
		{position: 18446744073709551601, want: "18446744073709551601st"},
		{position: 18446744073709551602, want: "18446744073709551602nd"},
		{position: 18446744073709551603, want: "18446744073709551603rd"},
		{position: 18446744073709551611, want: "18446744073709551611th"},
		{position: 18446744073709551612, want: "18446744073709551612th"},
		{position: 18446744073709551613, want: "18446744073709551613th"},
		{position: math.MaxUint64, want: "18446744073709551615th"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			if got := numnam.Ordinal(tt.position); got != tt.want {
				t.Errorf("numnam.Ordinal(%d), got %s, want %s", tt.position, got, tt.want)
			}
		})
	}

	for i := uint64(0); i < math.MaxUint16; i++ {
		t.Run(fmt.Sprintf("numnam.Ordinal(%d)", i), func(t *testing.T) {
			got := numnam.Ordinal(i)

			if !checkSuffix(got) {
				t.Errorf("%s is an invalid ordinal", got)
			}

			if strings.Contains(got, "=") {
				t.Errorf("%s should not contain `=`", got)
			}
		})
	}
}

func checkSuffix(ordinal string) bool {
	stuff := []string{"11th", "12th", "13th", "0th", "1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th"}
	for _, th := range stuff {
		if strings.HasSuffix(ordinal, th) {
			return true
		}
	}
	return false
}

func TestOrdinalEquals(t *testing.T) {
	tests := []struct {
		position uint64
		want     string
	}{
		{position: 0, want: "0th"},
		{position: 18446744073709551598, want: "18446744073709551598th"},
		{position: 18446744073709551599, want: "18446744073709551599th"},
		{position: 18446744073709551600, want: "18446744073709551600th"},
		{position: 18446744073709551601, want: "18446744073709551601st"},
		{position: 18446744073709551602, want: "18446744073709551602nd"},
		{position: 18446744073709551603, want: "18446744073709551603rd"},
		{position: 18446744073709551611, want: "18446744073709551611th"},
		{position: 18446744073709551612, want: "18446744073709551612th"},
		{position: 18446744073709551613, want: "18446744073709551613th"},
		{position: math.MaxUint64, want: "18446744073709551615th"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			got := numnam.OrdinalEqual(tt.position, false)
			if got != tt.want {
				t.Errorf("numnam.OrdinalEqual(%d), got %s, want %s", tt.position, got, tt.want)
			}

			got = numnam.OrdinalEqual(tt.position, true)
			if got != "="+tt.want {
				t.Errorf("numnam.OrdinalEqual(%d), got %s, want =%s", tt.position, got, tt.want)
			}
		})
	}

	for i := uint64(0); i < math.MaxUint16; i++ {
		t.Run(fmt.Sprintf("numnam.OrdinalEqual(%d)", i), func(t *testing.T) {
			got := numnam.OrdinalEqual(i, true)

			if !checkSuffix(got) {
				t.Errorf("%s is an invalid ordinal", got)
			}

			if !strings.HasPrefix(got, "=") {
				t.Errorf("%s should be prefixed with `=`", got)
			}

			if strings.Count(got, "=") != 1 {
				t.Errorf("%s should only have one `=`", got)
			}
		})
	}
}

func TestOrdinal0(t *testing.T) {
	tests := []struct {
		position uint64
		want     string
	}{
		{position: 0, want: "1st"},
		{position: 18446744073709551598, want: "18446744073709551599th"},
		{position: 18446744073709551599, want: "18446744073709551600th"},
		{position: 18446744073709551600, want: "18446744073709551601st"},
		{position: 18446744073709551601, want: "18446744073709551602nd"},
		{position: 18446744073709551602, want: "18446744073709551603rd"},
		{position: 18446744073709551603, want: "18446744073709551604th"},
		{position: 18446744073709551610, want: "18446744073709551611th"},
		{position: 18446744073709551611, want: "18446744073709551612th"},
		{position: 18446744073709551612, want: "18446744073709551613th"},
		{position: 18446744073709551613, want: "18446744073709551614th"},
		{position: math.MaxUint64, want: "18446744073709551616th"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			got := numnam.Ordinal0(tt.position)
			if got != tt.want {
				t.Errorf("numnam.Ordinal0(%d), got %s, want %s", tt.position, got, tt.want)
			}
		})
	}

	for i := uint64(0); i < math.MaxUint16; i++ {
		t.Run(fmt.Sprintf("numnam.Ordinal0(%d)", i), func(t *testing.T) {
			got := numnam.Ordinal0(i)

			if !checkSuffix(got) {
				t.Errorf("%s is an invalid ordinal", got)
			}

			if strings.Contains(got, "=") {
				t.Errorf("%s should not contain `=`", got)
			}
		})
	}
}

func TestOrdinal0Equals(t *testing.T) {
	tests := []struct {
		position uint64
		want     string
	}{
		{position: 0, want: "1st"},
		{position: 18446744073709551598, want: "18446744073709551599th"},
		{position: 18446744073709551599, want: "18446744073709551600th"},
		{position: 18446744073709551600, want: "18446744073709551601st"},
		{position: 18446744073709551601, want: "18446744073709551602nd"},
		{position: 18446744073709551602, want: "18446744073709551603rd"},
		{position: 18446744073709551603, want: "18446744073709551604th"},
		{position: 18446744073709551610, want: "18446744073709551611th"},
		{position: 18446744073709551611, want: "18446744073709551612th"},
		{position: 18446744073709551612, want: "18446744073709551613th"},
		{position: 18446744073709551613, want: "18446744073709551614th"},
		{position: math.MaxUint64, want: "18446744073709551616th"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			got := numnam.Ordinal0Equal(tt.position, false)
			if got != tt.want {
				t.Errorf("numnam.Ordinal0Equal(%d), got %s, want %s", tt.position, got, tt.want)
			}

			got = numnam.Ordinal0Equal(tt.position, true)
			if got != "="+tt.want {
				t.Errorf("numnam.Ordinal0Equal(%d), got %s, want =%s", tt.position, got, tt.want)
			}
		})
	}

	for i := uint64(0); i < math.MaxUint16; i++ {
		t.Run(fmt.Sprintf("numnam.Ordinal0Equal(%d)", i), func(t *testing.T) {
			got := numnam.Ordinal0Equal(i, true)

			if !checkSuffix(got) {
				t.Errorf("%s is an invalid ordinal", got)
			}

			if !strings.HasPrefix(got, "=") {
				t.Errorf("%s should be prefixed with `=`", got)
			}

			if strings.Count(got, "=") != 1 {
				t.Errorf("%s should only have one `=`", got)
			}
		})
	}
}
