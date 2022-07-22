package golukeutil

import "encoding/binary"

//
func Get16BitValNum(refByte0 byte, refByte1 byte) uint16 {
	var values []byte = []byte{refByte0, refByte1}
	return binary.BigEndian.Uint16(values)
}
