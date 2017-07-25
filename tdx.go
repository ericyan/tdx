package tdx

import (
	"bytes"
	"encoding/binary"
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

func (day *dayBar) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.LittleEndian, day)
	if err != nil {
		return err
	}

	return nil
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

func (five *fiveBar) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.LittleEndian, five)
	if err != nil {
		return err
	}

	return nil
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

func (lcn *lcnBar) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.LittleEndian, lcn)
	if err != nil {
		return err
	}

	return nil
}
