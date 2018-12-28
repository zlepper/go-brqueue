package brqueue

import "encoding/binary"

func byteArrayToInt(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(b))
}

func intToByteArray(i int32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(i))
	return b
}
