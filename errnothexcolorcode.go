package hexcolor

import (
	"fmt"
)

func errNotHexColorCode(s string) error {
	return fmt.Errorf("hexcolor: %q is not a hex color code", s)
}
