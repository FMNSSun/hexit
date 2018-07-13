package hexit

// Alphabet to use. Is initialised to use
// lower-case letters and digits: 0-9a-f.
// If you change this please make sure that
// you call BuildTable() to rebuild internal
// lookup tables!
var Alphabet = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
}

var table []string
var revTable []uint8
var revTable8 []uint8

func init() {
	BuildTable()
}

// BuildTable re-builds the uint8 lookup table.
func BuildTable() {
	table = make([]string, 256)

	for i := 0; i < 256; i++ {
		table[i] = string(HexUint8(uint8(i)))
	}

	revTable = make([]uint8, 256)

	for i, v := range Alphabet {
		revTable[v] = uint8(i)
	}
}

// UnhexUint8 converts from hexadecimal representation
// to uint8. Buf must be at least two bytes long.
func UnhexUint8(buf []byte) uint8 {
	h := revTable[buf[0]]
	l := revTable[buf[1]]
	return (h << 4) | l
}

// UnhexUint8Str converts from hexadecimal representation
// to uint8. Buf must be at least two bytes long.
func UnhexUint8Str(buf string) uint8 {
	h := revTable[buf[0]]
	l := revTable[buf[1]]
	return (h << 4) | l
}

// UnhexUint16 converts from hexadecimal representation
// to uint16. Buf must be at least 4 bytes long.
func UnhexUint16(buf []byte) uint16 {
	ah := uint16(revTable[buf[0]])
	al := uint16(revTable[buf[1]])
	bh := uint16(revTable[buf[2]])
	bl := uint16(revTable[buf[3]])
	return (ah << 12) | (al << 8) | (bh << 4) | bl
}

// UnhexUint16Str converts from hexadecimal representation
// to uint16. Buf must be at least 4 bytes long.
func UnhexUint16Str(buf string) uint16 {
	ah := uint16(revTable[buf[0]])
	al := uint16(revTable[buf[1]])
	bh := uint16(revTable[buf[2]])
	bl := uint16(revTable[buf[3]])
	return (ah << 12) | (al << 8) | (bh << 4) | bl
}

// UnhexUint32 converts from hexadecimal representation
// to uint32. Buf must be at least 8 bytes long.
func UnhexUint32(buf []byte) uint32 {
	ah := uint32(revTable[buf[0]])
	al := uint32(revTable[buf[1]])
	bh := uint32(revTable[buf[2]])
	bl := uint32(revTable[buf[3]])
	ch := uint32(revTable[buf[4]])
	cl := uint32(revTable[buf[5]])
	dh := uint32(revTable[buf[6]])
	dl := uint32(revTable[buf[7]])
	return (ah << 28) | (al << 24) | (bh << 20) | (bl << 16) | (ch << 12) | (cl << 8) | (dh << 4) | dl
}

// UnhexUint32Str converts from hexadecimal representation
// to uint32. Buf must be at least 8 bytes long.
func UnhexUint32Str(buf string) uint32 {
	ah := uint32(revTable[buf[0]])
	al := uint32(revTable[buf[1]])
	bh := uint32(revTable[buf[2]])
	bl := uint32(revTable[buf[3]])
	ch := uint32(revTable[buf[4]])
	cl := uint32(revTable[buf[5]])
	dh := uint32(revTable[buf[6]])
	dl := uint32(revTable[buf[7]])
	return (ah << 28) | (al << 24) | (bh << 20) | (bl << 16) | (ch << 12) | (cl << 8) | (dh << 4) | dl
}

// UnhexUint64 converts from hexadecimal representation
// to uint32. Buf must be at least 8 bytes long.
func UnhexUint64(buf []byte) uint64 {
	ah := uint64(revTable[buf[0]])
	al := uint64(revTable[buf[1]])
	bh := uint64(revTable[buf[2]])
	bl := uint64(revTable[buf[3]])
	ch := uint64(revTable[buf[4]])
	cl := uint64(revTable[buf[5]])
	dh := uint64(revTable[buf[6]])
	dl := uint64(revTable[buf[7]])
	eh := uint64(revTable[buf[8]])
	el := uint64(revTable[buf[9]])
	fh := uint64(revTable[buf[10]])
	fl := uint64(revTable[buf[11]])
	gh := uint64(revTable[buf[12]])
	gl := uint64(revTable[buf[13]])
	hh := uint64(revTable[buf[14]])
	hl := uint64(revTable[buf[15]])

	s := (ah << 60) | (al << 56) | (bh << 52) | (bl << 48) | (ch << 44) | (cl << 40) | (dh << 36) | (dl << 32)
	s |= (eh << 28) | (el << 24) | (fh << 20) | (fl << 16) | (gh << 12) | (gl << 8) | (hh << 4) | hl

	return s
}

// UnhexUint64Str converts from hexadecimal representation
// to uint32. Buf must be at least 16 bytes long.
func UnhexUint64Str(buf string) uint64 {
	ah := uint64(revTable[buf[0]])
	al := uint64(revTable[buf[1]])
	bh := uint64(revTable[buf[2]])
	bl := uint64(revTable[buf[3]])
	ch := uint64(revTable[buf[4]])
	cl := uint64(revTable[buf[5]])
	dh := uint64(revTable[buf[6]])
	dl := uint64(revTable[buf[7]])
	eh := uint64(revTable[buf[8]])
	el := uint64(revTable[buf[9]])
	fh := uint64(revTable[buf[10]])
	fl := uint64(revTable[buf[11]])
	gh := uint64(revTable[buf[12]])
	gl := uint64(revTable[buf[13]])
	hh := uint64(revTable[buf[14]])
	hl := uint64(revTable[buf[15]])

	s := (ah << 60) | (al << 56) | (bh << 52) | (bl << 48) | (ch << 44) | (cl << 40) | (dh << 36) | (dl << 32)
	s |= (eh << 28) | (el << 24) | (fh << 20) | (fl << 16) | (gh << 12) | (gl << 8) | (hh << 4) | hl

	return s
}

// HexUint8Str returns a hexadecimal representation
// of its argument as string.
func HexUint8Str(i uint8) string {
	return table[i]
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
