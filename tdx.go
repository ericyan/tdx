package tdx

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
)

// A dayBar is a single record in end-of-day bar (.day) files.
type dayBar struct {
	Date     uint32  // yyyymmdd
	Open     uint32  // in cents
	High     uint32  // in cents
	Low      uint32  // in cents
	Close    uint32  // in cents
	Turnover float32 // in yuan
	Volume   uint32  // in shares
	Reserved [4]byte // unknown
}

// A fiveBar in a single record in five-minute bar (.5) files.
type fiveBar struct {
	Date     uint16  // higher 5 bits: years since 2004, higher 11 bits: mmdd
	Time     uint16  // minutes since 00:00:00 (UTC+8)
	Open     uint32  // in cents
	High     uint32  // in cents
	Low      uint32  // in cents
	Close    uint32  // in cents
	Turnover float32 // in yuan
	Volume   uint32  // in shares
	Reserved [4]byte // unknown
}

// A lcnBar is a single record in minute bar (.lc5/.lc1) files.
type lcnBar struct {
	Date     uint16  // higher 5 bits: years since 2004, higher 11 bits: mmdd
	Time     uint16  // minutes since 00:00:00 (UTC+8)
	Open     float32 // in yuan
	High     float32 // in yuan
	Low      float32 // in yuan
	Close    float32 // in yuan
	Turnover float32 // in yuan
	Volume   uint32  // in shares
	Reserved [4]byte // unknown
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
