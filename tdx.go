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
