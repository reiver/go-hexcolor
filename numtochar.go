package hexcolor

func numtochar(n uint8) (byte, bool) {
	switch {
	case 0 <= n && n <= 9:
		return n + '0', true
	case 10 <= n && n <= 15:
		return n - 10 + 'A', true
	default:
		return 0, false
	}
}
