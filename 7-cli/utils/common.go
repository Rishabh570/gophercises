package utils

import "encoding/binary"

func ConvertIntToByteArray(num int) []byte {
	byteSlice := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteSlice, uint64(num))
	return byteSlice
}

func ConvertByteArrayToInt(byteSlice []byte) int {
	return int(binary.LittleEndian.Uint64(byteSlice))
}
