package hexcolor

import (
	"testing"
)

func TestCharToNum(t *testing.T) {

	tests := []struct{
		Char byte
		ExpectedNum uint8
		ExpectedOK  bool
	}{
		{
			Char:        0,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        1,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        2,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        3,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        4,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        5,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        6,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        7,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        8,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        9,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        10,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        11,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        12,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        13,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        14,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        15,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        16,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        17,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        18,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        19,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        20,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        21,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        22,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        23,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        24,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        25,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        26,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        27,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        28,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        29,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        30,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:        31,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       ' ',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '!',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '"',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '#',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '$',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '%',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '&',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '\'',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '(',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       ')',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '*',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '+',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       ',',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '-',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '.',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '/',
			ExpectedNum: 0,
			ExpectedOK: false,
		},



		{
			Char:       '0',
			ExpectedNum: 0,
			ExpectedOK: true,
		},
		{
			Char:       '1',
			ExpectedNum: 1,
			ExpectedOK: true,
		},
		{
			Char:       '2',
			ExpectedNum: 2,
			ExpectedOK: true,
		},
		{
			Char:       '3',
			ExpectedNum: 3,
			ExpectedOK: true,
		},
		{
			Char:       '4',
			ExpectedNum: 4,
			ExpectedOK: true,
		},
		{
			Char:       '5',
			ExpectedNum: 5,
			ExpectedOK: true,
		},
		{
			Char:       '6',
			ExpectedNum: 6,
			ExpectedOK: true,
		},
		{
			Char:       '7',
			ExpectedNum: 7,
			ExpectedOK: true,
		},
		{
			Char:       '8',
			ExpectedNum: 8,
			ExpectedOK: true,
		},
		{
			Char:       '9',
			ExpectedNum: 9,
			ExpectedOK: true,
		},



		{
			Char:       ':',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       ';',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '<',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '=',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '>',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '?',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '@',
			ExpectedNum: 0,
			ExpectedOK: false,
		},



		{
			Char:       'A',
			ExpectedNum: 10,
			ExpectedOK: true,
		},
		{
			Char:       'B',
			ExpectedNum: 11,
			ExpectedOK: true,
		},
		{
			Char:       'C',
			ExpectedNum: 12,
			ExpectedOK: true,
		},
		{
			Char:       'D',
			ExpectedNum: 13,
			ExpectedOK: true,
		},
		{
			Char:       'E',
			ExpectedNum: 14,
			ExpectedOK: true,
		},
		{
			Char:       'F',
			ExpectedNum: 15,
			ExpectedOK: true,
		},



		{
			Char:       'G',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'H',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'I',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'J',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'K',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'L',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'M',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'N',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'O',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'P',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'Q',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'R',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'S',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'T',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'U',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'V',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'W',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'X',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'Y',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'Z',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '[',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '\\',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       ']',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '^',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '_',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '`',
			ExpectedNum: 0,
			ExpectedOK: false,
		},



		{
			Char:       'a',
			ExpectedNum: 10,
			ExpectedOK: true,
		},
		{
			Char:       'b',
			ExpectedNum: 11,
			ExpectedOK: true,
		},
		{
			Char:       'c',
			ExpectedNum: 12,
			ExpectedOK: true,
		},
		{
			Char:       'd',
			ExpectedNum: 13,
			ExpectedOK: true,
		},
		{
			Char:       'e',
			ExpectedNum: 14,
			ExpectedOK: true,
		},
		{
			Char:       'f',
			ExpectedNum: 15,
			ExpectedOK: true,
		},



		{
			Char:       'g',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'h',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'i',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'j',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'k',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'l',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'm',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'n',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'o',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'p',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'q',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'r',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       's',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       't',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'u',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'v',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'w',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'x',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'y',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       'z',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '{',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '|',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '}',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       '~',
			ExpectedNum: 0,
			ExpectedOK: false,
		},
		{
			Char:       127,
			ExpectedNum: 0,
			ExpectedOK: false,
		},
	}

	for testNumber, test := range tests {

		actualNum, actualOK := chartonum(test.Char)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("CHAR: %q (%U)", rune(test.Char), rune(test.Char))
				continue
			}
		}

		{
			expected := test.ExpectedNum
			actual   := actualNum

			if expected != actual {
				t.Errorf("For test #%d, the actual 'num' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("CHAR: %q (%U)", rune(test.Char), rune(test.Char))
				continue
			}
		}
	}
}
