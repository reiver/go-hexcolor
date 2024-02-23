package hexcolor

// Format encode and returns a hex color code for the specified color.
func Format(red uint8, green uint8, blue uint8) string {
	var buffer [7]byte

	buffer[0] = '#'
	buffer[1], _ = numtochar((red   & 0xF0) >> 4)
	buffer[2], _ =  numtochar(red   & 0x0F)
	buffer[3], _ = numtochar((green & 0xF0) >> 4)
	buffer[4], _ =  numtochar(green & 0x0F)
	buffer[5], _ = numtochar((blue  & 0xF0) >> 4)
	buffer[6], _ =  numtochar(blue  & 0x0F)

	return string(buffer[:])
}
