package tdx

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
	"time"
)

var tz = time.FixedZone("UTC+8", int(8*time.Hour/time.Second))

// Bar provides a generic representation of bar data.
type Bar interface {
	Time() time.Time
	Open() float32
	High() float32
	Low() float32
	Close() float32
	Volume() uint32
	Turnover() float32
}

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

func (bar *dayBar) Time() time.Time {
	mmdd := int(bar.RawDate % 10000)

	return time.Date(int(bar.RawDate/10000), time.Month(mmdd/100), mmdd%100, 15, 0, 0, 0, tz)
}

func (bar *dayBar) Open() float32 {
	return float32(bar.RawOpen) / 100
}

func (bar *dayBar) High() float32 {
	return float32(bar.RawHigh) / 100
}

func (bar *dayBar) Low() float32 {
	return float32(bar.RawLow) / 100
}

func (bar *dayBar) Close() float32 {
	return float32(bar.RawClose) / 100
}

func (bar *dayBar) Volume() uint32 {
	return bar.RawVolume
}

func (bar *dayBar) Turnover() float32 {
	return bar.RawTurnover
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

func (bar *fiveBar) Time() time.Time {
	mmdd := int(bar.RawDate & (^uint16(0) >> (16 - 11)))

	return time.Date(2004+int(bar.RawDate>>11), time.Month(mmdd/100), mmdd%100, int(bar.RawTime/60), int(bar.RawTime%60), 0, 0, tz)
}

func (bar *fiveBar) Open() float32 {
	return float32(bar.RawOpen) / 100
}

func (bar *fiveBar) High() float32 {
	return float32(bar.RawHigh) / 100
}

func (bar *fiveBar) Low() float32 {
	return float32(bar.RawLow) / 100
}

func (bar *fiveBar) Close() float32 {
	return float32(bar.RawClose) / 100
}

func (bar *fiveBar) Volume() uint32 {
	return bar.RawVolume
}

func (bar *fiveBar) Turnover() float32 {
	return bar.RawTurnover
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

func (bar *lcnBar) Time() time.Time {
	mmdd := int(bar.RawDate & (^uint16(0) >> (16 - 11)))

	return time.Date(2004+int(bar.RawDate>>11), time.Month(mmdd/100), mmdd%100, int(bar.RawTime/60), int(bar.RawTime%60), 0, 0, tz)
}

func (bar *lcnBar) Open() float32 {
	return bar.RawOpen
}

func (bar *lcnBar) High() float32 {
	return bar.RawHigh
}

func (bar *lcnBar) Low() float32 {
	return bar.RawLow
}

func (bar *lcnBar) Close() float32 {
	return bar.RawClose
}

func (bar *lcnBar) Volume() uint32 {
	return bar.RawVolume
}

func (bar *lcnBar) Turnover() float32 {
	return bar.RawTurnover
}

// Dataset is a collection of bars with metadata.
type Dataset struct {
	Market  string // Market identification code (ISO 10383)
	Symbol  string // Local exchange ticker symbol
	BarSize uint   // Length of the bar defined in minutes
	Bars    []Bar
}

// DecodeFile decodes a bar data file that has been encoded in any of
// the supported formats. It detects file format by the file extension.
func DecodeFile(filepath string) (*Dataset, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	market, symbol, ext := path.Base(filepath)[0:2], path.Base(filepath)[2:8], path.Base(filepath)[8:]
	switch market {
	case "sh":
		market = "XSHG"
	case "sz":
		market = "XSHE"
	default:
		return nil, fmt.Errorf("unknown market: %s", market)
	}

	var barSize uint
	bars := make([]Bar, fi.Size()/32)
	for i := 0; i < len(bars); i++ {
		var bar Bar
		switch ext {
		case ".day":
			bar = new(dayBar)
			barSize = 1440
		case ".5":
			bar = new(fiveBar)
			barSize = 5
		case ".lc5":
			bar = new(lcnBar)
			barSize = 5
		case ".lc1":
			bar = new(lcnBar)
			barSize = 1
		default:
			return nil, fmt.Errorf("unsupported file: %s", filepath)
		}

		err = binary.Read(f, binary.LittleEndian, bar)
		if err != nil {
			return nil, err
		}

		bars[i] = bar
	}

	return &Dataset{market, symbol, barSize, bars}, nil
}
