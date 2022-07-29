package bcdt

import (
	"bytes"
	"encoding/binary"
	"time"
)

const yearZero = 2022

func Encode(epoch int64) uint32 {
	t := time.Unix(epoch, 0).UTC()
	var i uint32

	// byte 1
	i = uint32(t.Year() - yearZero)
	i <<= 4

	i |= uint32(t.Month())

	// byte 2
	i <<= 8
	i |= uint32(t.Day())

	// byte 3
	i <<= 8
	i |= uint32(t.Hour())

	// byte 4
	i <<= 8
	i |= uint32(t.Minute())

	return i
}

func DecodeByteArray(be []byte) int64 {
	var encoded uint32
	buf := bytes.NewReader(be)
	binary.Read(buf, binary.BigEndian, &encoded)

	return Decode(encoded)
}

func Decode(datetime uint32) int64 {
	var offset uint32 = 0

	minute := datetime & (255 << offset)
	minute >>= offset
	offset += 8

	hour := datetime & (255 << offset)
	hour >>= offset
	offset += 8

	day := datetime & (255 << offset)
	day >>= offset
	offset += 8

	month := datetime & (15 << offset)
	month >>= offset
	offset += 4

	year := datetime & (15 << offset)
	year >>= offset

	return time.Date(int(year+yearZero), time.Month(month), int(day), int(hour), int(minute), 0, 0, time.UTC).Unix()
}
