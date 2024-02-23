package hexcolor

// Parse parses a hex color code and returns the red, green, and blue components separatly as uint8 numbers.
func Parse(value string) (red uint8, green uint8, blue uint8, err error) {

	if len(value) < 1 {
		return 0, 0, 0, errNotHexColorCode(value)
	}

	var s string = value

	{
		s0 := s[0]

		if '#' == s0 {
			s = s[1:]
		}
	}

	if 6 != len(s) {
		return 0, 0, 0, errNotHexColorCode(value)
	}

	r1char, r0char := s[0], s[1]
	g1char, g0char := s[2], s[3]
	b1char, b0char := s[4], s[5]

	var r1 byte
	{
		var ok bool

		r1, ok = chartonum(r1char)
		if !ok {
			return 0, 0, 0, errNotHexColorCode(value)
		}
	}

	var r0 byte
	{
		var ok bool

		r0, ok = chartonum(r0char)
		if !ok {
			return 0, 0, 0, errNotHexColorCode(value)
		}
	}

	var g1 byte
	{
		var ok bool

		g1, ok = chartonum(g1char)
		if !ok {
			return 0, 0, 0, errNotHexColorCode(value)
		}
	}

	var g0 byte
	{
		var ok bool

		g0, ok = chartonum(g0char)
		if !ok {
			return 0, 0, 0, errNotHexColorCode(value)
		}
	}

	var b1 byte
	{
		var ok bool

		b1, ok = chartonum(b1char)
		if !ok {
			return 0, 0, 0, errNotHexColorCode(value)
		}
	}

	var b0 byte
	{
		var ok bool

		b0, ok = chartonum(b0char)
		if !ok {
			return 0, 0, 0, errNotHexColorCode(value)
		}
	}

	red   = (r1 << 4) | r0
	green = (g1 << 4) | g0
	blue  = (b1 << 4) | b0

	return red, green, blue, nil
}
