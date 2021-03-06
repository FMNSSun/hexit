package hexit

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestUnhexUint8(t *testing.T) {
	for i := 0; i < 8192; i++ {
		exp := uint8(rand.Intn(256))
		s := []byte(fmt.Sprintf("%02x", exp))
		got := UnhexUint8(s)

		if exp != got {
			t.Errorf("Expected %x but got %x", exp, got)
			return
		}
	}
}

func TestUnhexUint16(t *testing.T) {
	for i := 0; i < 8192; i++ {
		exp := uint16(rand.Intn(65536))
		s := []byte(fmt.Sprintf("%04x", exp))
		got := UnhexUint16(s)

		if exp != got {
			t.Errorf("Expected %x but got %x", exp, got)
			return
		}
	}
}

func TestUnhexUint32(t *testing.T) {
	for i := 0; i < 8192; i++ {
		exp := rand.Uint32()
		s := []byte(fmt.Sprintf("%08x", exp))
		got := UnhexUint32(s)

		if exp != got {
			t.Errorf("Expected %x but got %x", exp, got)
			return
		}
	}
}

func TestUnhexUint64(t *testing.T) {
	for i := 0; i < 8192; i++ {
		exp := rand.Uint64()
		s := []byte(fmt.Sprintf("%016x", exp))
		got := UnhexUint64(s)

		if exp != got {
			t.Errorf("Expected %x but got %x", exp, got)
			return
		}
	}
}


func TestHexUint8(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := uint8(rand.Intn(256))
		exp := []byte(fmt.Sprintf("%02x", n))
		got := HexUint8(n)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func TestHexUint8Str(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := uint8(rand.Intn(256))
		exp := fmt.Sprintf("%02x", n)
		got := HexUint8Str(n)

		if exp != got {
			t.Errorf("Expected %q but got %q.", exp, got)
			return
		}
	}
}


func BenchmarkHexUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint8(rand.Intn(256))
		_ = HexUint8Str(n)
	}
}

func BenchmarkFmtUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint8(rand.Intn(256))
		_ = fmt.Sprintf("%02x", n)
	}
}

func BenchmarkStrconvAppendUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint8(rand.Intn(256))
		_ = strconv.AppendInt(nil, int64(n), 16)
	}
}

func BenchmarkUnhexUint8Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%02x", uint8(rand.Intn(256)))
		_ = UnhexUint8Str(s)
	}
}

func BenchmarkStrconvParseUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%02x", uint8(rand.Intn(256)))
		_, _ = strconv.ParseUint(s, 16, 8)
	}
}

func TestHexUint16(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := uint16(rand.Intn(65536))
		exp := []byte(fmt.Sprintf("%04x", n))
		got := HexUint16(n)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func BenchmarkHexUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint16(rand.Intn(65536))
		_ = HexUint16Str(n)
	}
}

func BenchmarkFmtUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint16(rand.Intn(65536))
		_ = fmt.Sprintf("%04x", n)
	}
}

func BenchmarkStrconvAppendUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint16(rand.Intn(65536))
		_ = strconv.AppendInt(nil, int64(n), 16)
	}
}

func BenchmarkUnhexUint16Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%04x", uint16(rand.Intn(65536)))
		_ = UnhexUint16Str(s)
	}
}

func BenchmarkStrconvParseUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%04x", uint16(rand.Intn(65536)))
		_, _ = strconv.ParseUint(s, 16, 16)
	}
}

func TestHexUint32(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := uint32(rand.Intn(4294967296))
		exp := []byte(fmt.Sprintf("%08x", n))
		got := HexUint32(n)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func BenchmarkHexUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint32(rand.Intn(4294967296))
		_ = HexUint32Str(n)
	}
}

func BenchmarkHexUint32To(b *testing.B) {
	dst := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		n := uint32(rand.Intn(4294967296))
		HexUint32To(n, dst)
	}
}

func BenchmarkFmtUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint32(rand.Intn(4294967296))
		_ = fmt.Sprintf("%08x", n)
	}
}

func BenchmarkStrconvAppendUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := uint32(rand.Intn(4294967296))
		_ = strconv.AppendInt(nil, int64(n), 16)
	}
}

func BenchmarkStrconvFormatUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Uint32()
		_ = strconv.FormatUint(uint64(n), 16)
	}
}

func BenchmarkUnhexUint32Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%08x", rand.Uint32())
		_ = UnhexUint32Str(s)
	}
}

func BenchmarkStrconvParseUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%08x", rand.Uint32())
		_, _ = strconv.ParseUint(s, 16, 32)
	}
}

func TestHexUint64(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := rand.Uint64()
		exp := []byte(fmt.Sprintf("%016x", n))
		got := HexUint64(n)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func BenchmarkHexUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Uint64()
		_ = HexUint64Str(n)
	}
}

func BenchmarkFmtUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Uint64()
		_ = fmt.Sprintf("%016x", n)
	}
}

func BenchmarkStrconvAppendUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Uint32()
		_ = strconv.AppendInt(nil, int64(n), 16)
	}
}

func BenchmarkStrconvFormatUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Uint64()
		_ = strconv.FormatUint(n, 16)
	}
}

func BenchmarkUnhexUint64Str(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%016x", rand.Uint64())
		_ = UnhexUint64Str(s)
	}
}

func BenchmarkStrconvParseUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%016x", rand.Uint64())
		_, _ = strconv.ParseUint(s, 16, 64)
	}
}

func TestHexUint8To(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := uint8(rand.Intn(256))
		exp := []byte(fmt.Sprintf("%02x", n))
		got := make([]byte, 2)
		HexUint8To(n, got)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func TestHexUint16To(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := uint16(rand.Intn(65536))
		exp := []byte(fmt.Sprintf("%04x", n))
		got := make([]byte, 4)
		HexUint16To(n, got)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func TestHexUint32To(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := rand.Uint32()
		exp := []byte(fmt.Sprintf("%08x", n))
		got := make([]byte, 8)
		HexUint32To(n, got)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}

func TestHexUint64To(t *testing.T) {
	for i := 0; i < 8192; i++ {
		n := rand.Uint64()
		exp := []byte(fmt.Sprintf("%016x", n))
		got := make([]byte, 16)
		HexUint64To(n, got)

		if !bytes.Equal(exp, got) {
			t.Errorf("Expected %d but got %d.", exp, got)
			return
		}
	}
}
