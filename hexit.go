package hexit

// Alphabet to use. Is initialised to use
// lower-case letters and digits: 0-9a-f
var Alphabet = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
}

var Table []string

func init() {
	BuildTable()
}

// BuildTable re-builds the uint8 lookup table.
func BuildTable() {
	Table = make([]string, 256)

	for i := 0; i < 256; i++ {
		Table[i] = string(HexUint8(uint8(i)))
	}
}

// HexUint8Str returns a hexadecimal representation
// of its argument as string.
func HexUint8Str(i uint8) string {
	return Table[i]
}

// HexUint8 returns a hexadecimal representation
// of its argument.
func HexUint8(i uint8) []byte {
	l := i & 0x0F
	h := (i >> 4) & 0x0F
	return []byte{Alphabet[h], Alphabet[l]}
}

// HexUint8To writes a hexadecimal representation
// of its argument to the destination slice. Slice must
// be at least two bytes long.
func HexUint8To(i uint8, dst []byte) {
	l := i & 0x0F
	h := (i >> 4) & 0x0F
	dst[0] = Alphabet[h]
	dst[1] = Alphabet[l]
}

// HexUint16Str returns a hexadecimal representation
// of its argument as string.
func HexUint16Str(i uint16) string {
	return string(HexUint16(i))
}

// HexUint16 retruns a hexadecimal representation
// of its argument.
func HexUint16(i uint16) []byte {
	ll := i & 0x0F
	lh := (i >> 4) & 0x0F
	hl := (i >> 8) & 0x0F
	hh := (i >> 12) & 0x0F
	return []byte{
		Alphabet[hh], Alphabet[hl],
		Alphabet[lh], Alphabet[ll],
	}
}

// HexUint16To writes a hexadecimal representation
// of its argument to the destination slice. Slice must
// be at least 4 bytes long.
func HexUint16To(i uint16, dst []byte) {
	ll := i & 0x0F
	lh := (i >> 4) & 0x0F
	hl := (i >> 8) & 0x0F
	hh := (i >> 12) & 0x0F

	dst[0] = Alphabet[hh]
	dst[1] = Alphabet[hl]
	dst[2] = Alphabet[lh]
	dst[3] = Alphabet[ll]
}

// HexUint32Str returns a hexadecimal representation
// of its argument as string.
func HexUint32Str(i uint32) string {
	return string(HexUint32(i))
}

// HexUint32 retruns a hexadecimal representation
// of its argument.
func HexUint32(i uint32) []byte {
	al := i & 0x0F
	ah := (i >> 4) & 0x0F
	bl := (i >> 8) & 0x0F
	bh := (i >> 12) & 0x0F
	cl := (i >> 16) & 0x0F
	ch := (i >> 20) & 0x0F
	dl := (i >> 24) & 0x0F
	dh := (i >> 28) & 0x0F
	return []byte{
		Alphabet[dh], Alphabet[dl],
		Alphabet[ch], Alphabet[cl],
		Alphabet[bh], Alphabet[bl],
		Alphabet[ah], Alphabet[al],
	}
}

// HexUint32To writes a hexadecimal representation
// of its argument to the destination slice. Slice must
// be at least 8 bytes long.
func HexUint32To(i uint32, dst []byte) {
	al := i & 0x0F
	ah := (i >> 4) & 0x0F
	bl := (i >> 8) & 0x0F
	bh := (i >> 12) & 0x0F
	cl := (i >> 16) & 0x0F
	ch := (i >> 20) & 0x0F
	dl := (i >> 24) & 0x0F
	dh := (i >> 28) & 0x0F

	dst[0] = Alphabet[dh]
	dst[1] = Alphabet[dl]
	dst[2] = Alphabet[ch]
	dst[3] = Alphabet[cl]
	dst[4] = Alphabet[bh]
	dst[5] = Alphabet[bl]
	dst[6] = Alphabet[ah]
	dst[7] = Alphabet[al]
}

// HexUint64Str returns a hexadecimal representation
// of its argument as string.
func HexUint64Str(i uint64) string {
	return string(HexUint64(i))
}

// HexUint64 returns a hexadecimal representation
// of its argument.
func HexUint64(i uint64) []byte {
	al := i & 0x0F
	ah := (i >> 4) & 0x0F
	bl := (i >> 8) & 0x0F
	bh := (i >> 12) & 0x0F
	cl := (i >> 16) & 0x0F
	ch := (i >> 20) & 0x0F
	dl := (i >> 24) & 0x0F
	dh := (i >> 28) & 0x0F
	el := (i >> 32) & 0x0F
	eh := (i >> 36) & 0x0F
	fl := (i >> 40) & 0x0F
	fh := (i >> 44) & 0x0F
	gl := (i >> 48) & 0x0F
	gh := (i >> 52) & 0x0F
	hl := (i >> 56) & 0x0F
	hh := (i >> 60) & 0x0F
	return []byte{
		Alphabet[hh], Alphabet[hl],
		Alphabet[gh], Alphabet[gl],
		Alphabet[fh], Alphabet[fl],
		Alphabet[eh], Alphabet[el],
		Alphabet[dh], Alphabet[dl],
		Alphabet[ch], Alphabet[cl],
		Alphabet[bh], Alphabet[bl],
		Alphabet[ah], Alphabet[al],
	}
}

// HexUint64To writes a hexadecimal representation
// of its argument to the destination slice. Slice
// must be at least 16 bytes long.
func HexUint64To(i uint64, dst []byte) {
	al := i & 0x0F
	ah := (i >> 4) & 0x0F
	bl := (i >> 8) & 0x0F
	bh := (i >> 12) & 0x0F
	cl := (i >> 16) & 0x0F
	ch := (i >> 20) & 0x0F
	dl := (i >> 24) & 0x0F
	dh := (i >> 28) & 0x0F
	el := (i >> 32) & 0x0F
	eh := (i >> 36) & 0x0F
	fl := (i >> 40) & 0x0F
	fh := (i >> 44) & 0x0F
	gl := (i >> 48) & 0x0F
	gh := (i >> 52) & 0x0F
	hl := (i >> 56) & 0x0F
	hh := (i >> 60) & 0x0F

	dst[0] = Alphabet[hh]
	dst[1] = Alphabet[hl]
	dst[2] = Alphabet[gh]
	dst[3] = Alphabet[gl]
	dst[4] = Alphabet[fh]
	dst[5] = Alphabet[fl]
	dst[6] = Alphabet[eh]
	dst[7] = Alphabet[el]
	dst[8] = Alphabet[dh]
	dst[9] = Alphabet[dl]
	dst[10] = Alphabet[ch]
	dst[11] = Alphabet[cl]
	dst[12] = Alphabet[bh]
	dst[13] = Alphabet[bl]
	dst[14] = Alphabet[ah]
	dst[15] = Alphabet[al]
}
