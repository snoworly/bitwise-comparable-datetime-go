package bcdt

import (
	"encoding/binary"
	"time"
)

func QueryYearMonth(year, month int) [1]byte {
	q := QueryDatetime(year, month, 0, 0, 0)

	return [1]byte{q[0]}
}

func QueryDate(year, month, day int) [2]byte {
	q := QueryDatetime(year, month, day, 0, 0)

	return [2]byte{q[0], q[1]}
}

func QueryDatetime(year, month, day, hour, minute int) [4]byte {
	date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	encoded := Encode(date.Unix())

	var bcdtquery [4]byte
	binary.BigEndian.PutUint32(bcdtquery[:], encoded)

	return bcdtquery
}
