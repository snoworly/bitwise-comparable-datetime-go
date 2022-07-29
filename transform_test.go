package bcdt_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	bcdt "github.com/snoworly/bitwise-comparable-datetime-go"
)

func TestFollowRealtime(t *testing.T) {
	t0 := time.Now()
	epoch := t0.Unix()
	fmt.Println("date epoch    ", epoch)

	encoded := bcdt.Encode(epoch)
	decoded := bcdt.Decode(encoded)

	fmt.Println("output dec    ", encoded)
	fmt.Println("output bin    ", strconv.FormatInt(int64(decoded), 2))
	fmt.Println("output decoded", time.Unix(decoded, 0).UTC())

	datetime := time.Unix(decoded, 0).UTC()
	trailing := t0.Sub(datetime)
	if trailing > 1*time.Minute {
		t.Error("encoding didn't work as we're trailing for more than one minute")
	} else {
		fmt.Println("trailing", trailing, "OK")
	}
}

func TestEncode(t *testing.T) {
	date := time.Date(2022, 7, 29, 0, 0, 0, 0, time.UTC)
	fmt.Println("date epoch    ", date.Unix())
	fmt.Println("date lookup   ", date)
	encoded := bcdt.Encode(date.Unix())
	fmt.Println("date dec      ", encoded)
	fmt.Println("output bin    ", strconv.FormatInt(int64(encoded>>16), 2))
	fmt.Println("output bin  y ", strconv.FormatInt(int64((encoded&(15<<28))>>28), 2))
	fmt.Println("output bin  m ", strconv.FormatInt(int64((encoded&(15<<24))>>24), 2))
	fmt.Println("output bin  d ", strconv.FormatInt(int64((encoded&(63<<16))>>16), 2))

	var expected uint32 = 119341056
	if encoded != expected {
		t.Errorf("expected %v but got %v", expected, encoded)
	}
}

func TestSplit(t *testing.T) {
	date := time.Date(2022, time.December, 31, 23, 59, 0, 0, time.UTC)
	fmt.Println("date epoch    ", date.Unix())
	fmt.Println("date lookup   ", date)
	encoded := bcdt.Encode(date.Unix())
	fmt.Println("date dec      ", encoded)

	year := int64((encoded & (15 << 28)) >> 28)
	month := int64((encoded & (15 << 24)) >> 24)
	day := int64((encoded & (255 << 16)) >> 16)
	hour := int64((encoded & (255 << 8)) >> 8)
	min := int64((encoded & (255 << 0)) >> 0)

	fmt.Println("output bin    ", strconv.FormatInt(int64(encoded), 2))
	fmt.Println("output dec  y ", year)
	fmt.Println("output bin  y ", strconv.FormatInt(year, 2))
	fmt.Println("output dec  m ", month)
	fmt.Println("output bin  m ", strconv.FormatInt(month, 2))
	fmt.Println("output dec  d ", day)
	fmt.Println("output bin  d ", strconv.FormatInt(day, 2))
	fmt.Println("output dec  h ", hour)
	fmt.Println("output bin  h ", strconv.FormatInt(hour, 2))
	fmt.Println("output dec  m ", min)
	fmt.Println("output bin  m ", strconv.FormatInt(min, 2))
}
