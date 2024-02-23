package hexcolor

func chartonum(b byte) (uint8, bool) {
	switch {
	case '0' <= b && b <= '9':
		return b - '0', true
	case 'A' <= b && b <= 'F':
		return b - 'A' + 10, true
	case 'a' <= b && b <= 'f':
		return b - 'a' + 10, true
	default:
		return 0, false
	}
}
