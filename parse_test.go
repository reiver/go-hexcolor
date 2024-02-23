package hexcolor

import (
	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedRed  uint8
		ExpectedGreen uint8
		ExpectedBlue  uint8
	}{
		{
			Value:        "123456",
			ExpectedRed: 0x12,
			ExpectedGreen: 0x34,
			ExpectedBlue:    0x56,
		},
		{
			Value:       "#123456",
			ExpectedRed: 0x12,
			ExpectedGreen: 0x34,
			ExpectedBlue:    0x56,
		},



		{
			Value:        "ABCDEF",
			ExpectedRed: 0xAB,
			ExpectedGreen: 0xCD,
			ExpectedBlue:    0xEF,
		},
		{
			Value:       "#ABCDEF",
			ExpectedRed: 0xAB,
			ExpectedGreen: 0xCD,
			ExpectedBlue:    0xEF,
		},



		{
			Value:        "fedcba",
			ExpectedRed: 0xFE,
			ExpectedGreen: 0xDC,
			ExpectedBlue:    0xBA,
		},
		{
			Value:       "#fedcba",
			ExpectedRed: 0xFE,
			ExpectedGreen: 0xDC,
			ExpectedBlue:    0xBA,
		},



		{
			Value:        "FF1DCE",
			ExpectedRed: 0xFF,
			ExpectedGreen: 0x1D,
			ExpectedBlue:    0xCE,
		},
		{
			Value:       "#FF1DCE",
			ExpectedRed: 0xFF,
			ExpectedGreen: 0x1D,
			ExpectedBlue:    0xCE,
		},
		{
			Value:        "ff1dce",
			ExpectedRed: 0xFF,
			ExpectedGreen: 0x1D,
			ExpectedBlue:    0xCE,
		},
		{
			Value:       "#ff1dce",
			ExpectedRed: 0xFF,
			ExpectedGreen: 0x1D,
			ExpectedBlue:    0xCE,
		},



		{
			Value:        "1B2A34",
			ExpectedRed: 0x1B,
			ExpectedGreen: 0x2A,
			ExpectedBlue:    0x34,
		},
		{
			Value:       "#1B2A34",
			ExpectedRed: 0x1B,
			ExpectedGreen: 0x2A,
			ExpectedBlue:    0x34,
		},
		{
			Value:        "1b2a34",
			ExpectedRed: 0x1B,
			ExpectedGreen: 0x2A,
			ExpectedBlue:    0x34,
		},
		{
			Value:       "#1b2a34",
			ExpectedRed: 0x1B,
			ExpectedGreen: 0x2A,
			ExpectedBlue:    0x34,
		},



		{
			Value:        "8A928D",
			ExpectedRed: 0x8A,
			ExpectedGreen: 0x92,
			ExpectedBlue:    0x8D,
		},
		{
			Value:       "#8A928D",
			ExpectedRed: 0x8A,
			ExpectedGreen: 0x92,
			ExpectedBlue:    0x8D,
		},
		{
			Value:        "8a928d",
			ExpectedRed: 0x8A,
			ExpectedGreen: 0x92,
			ExpectedBlue:    0x8D,
		},
		{
			Value:       "#8a928d",
			ExpectedRed: 0x8A,
			ExpectedGreen: 0x92,
			ExpectedBlue:    0x8D,
		},
	}

	for testNumber, test := range tests {

		actualRed, actualGreen, actualBlue, err := Parse(test.Value)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.ExpectedRed
			actual := actualRed

			if expected != actual {
				t.Errorf("For test #%d, the actual 'red' is not what was expected.", testNumber)
				t.Logf("EXPECTED: 0x%02X", expected)
				t.Logf("ACTUAL:   0x%02X", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedGreen
			actual := actualGreen

			if expected != actual {
				t.Errorf("For test #%d, the actual 'green' is not what was expected.", testNumber)
				t.Logf("EXPECTED: 0x%02X", expected)
				t.Logf("ACTUAL:   0x%02X", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedBlue
			actual := actualBlue

			if expected != actual {
				t.Errorf("For test #%d, the actual 'blue' is not what was expected.", testNumber)
				t.Logf("EXPECTED: 0x%02X", expected)
				t.Logf("ACTUAL:   0x%02X", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
