package tdx

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
)

// A dayBar is a single record in end-of-day bar (.day) files.
type dayBar struct {
	RawDate     uint32  // yyyymmdd
	RawOpen     uint32  // in cents
	RawHigh     uint32  // in cents
	RawLow      uint32  // in cents
	RawClose    uint32  // in cents
	RawTurnover float32 // in yuan
	RawVolume   uint32  // in shares
	Reserved    [4]byte // unknown
}

// A fiveBar in a single record in five-minute bar (.5) files.
type fiveBar struct {
	RawDate     uint16  // higher 5 bits: years since 2004, higher 11 bits: mmdd
	RawTime     uint16  // minutes since 00:00:00 (UTC+8)
	RawOpen     uint32  // in cents
	RawHigh     uint32  // in cents
	RawLow      uint32  // in cents
	RawClose    uint32  // in cents
	RawTurnover float32 // in yuan
	RawVolume   uint32  // in shares
	Reserved    [4]byte // unknown
}

// A lcnBar is a single record in minute bar (.lc5/.lc1) files.
type lcnBar struct {
	RawDate     uint16  // higher 5 bits: years since 2004, higher 11 bits: mmdd
	RawTime     uint16  // minutes since 00:00:00 (UTC+8)
	RawOpen     float32 // in yuan
	RawHigh     float32 // in yuan
	RawLow      float32 // in yuan
	RawClose    float32 // in yuan
	RawTurnover float32 // in yuan
	RawVolume   uint32  // in shares
	Reserved    [4]byte // unknown
}

// DecodeFile decodes a bar data file that has been encoded in any of
// the supported formats. It detects file format by the file extension.
func DecodeFile(filepath string) ([]interface{}, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	bars := make([]interface{}, fi.Size()/32)
	for i := 0; i < len(bars); i++ {
		var bar interface{}
		switch path.Ext(filepath) {
		case ".day":
			bar = new(dayBar)
		case ".5":
			bar = new(fiveBar)
		case ".lc1", ".lc5":
			bar = new(lcnBar)
		default:
			return nil, fmt.Errorf("unsupported file: %s", filepath)
		}

		err = binary.Read(f, binary.LittleEndian, bar)
		if err != nil {
			return nil, err
		}

		bars[i] = bar
	}

	return bars, nil
}
