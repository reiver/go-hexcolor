package hexcolor

import (
	"testing"
)

func TestFormat(t *testing.T) {

	tests := []struct{
		Red   uint8
		Green uint8
		Blue  uint8
		Expected string
	}{
		{
			Red:         0,
			Green:         0,
			Blue:            0,
			Expected: "#000000",
		},

		{
			Red:         1,
			Green:         2,
			Blue:            3,
			Expected: "#010203",
		},

		{
			Red:       0x1A,
			Green:      0xB2,
			Blue:         0x3C,
			Expected: "#1AB23C",
		},
	}

	for testNumber, test := range tests {

		actual := Format(test.Red, test.Green, test.Blue)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual hex color code is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("(R,G,B):   (%02X,%02X,%02X)", test.Red, test.Green, test.Blue)
			continue
		}
	}
}
