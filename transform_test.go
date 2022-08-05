package bcdt_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
	"time"

	bcdt "github.com/snoworly/bitwise-comparable-datetime-go"
)

func TestZero(t *testing.T) {
	got := bcdt.DecodeByteArray([4]byte{0, 0, 0, 0})

	var expected int64
	if expected != got {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestOracleExample(t *testing.T) {
	epoch := 1659079415
	encoded := bcdt.Encode(int64(epoch))

	got := make([]byte, 4)
	binary.BigEndian.PutUint32(got, encoded)

	expected := []byte{7, 29, 7, 23}
	if bytes.Compare(got, expected) != 0 {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestDecodeByteArray(t *testing.T) {
	v := bcdt.DecodeByteArray([4]byte{7, 29, 7, 23})
	expected := 1659079380

	if v != int64(expected) {
		t.Errorf("expected %v but got %v", expected, v)
	}
}

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

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, encoded)
	fmt.Printf("date bytes     %x\n", bs)
	fmt.Println("date bytes    ", bs)

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

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, encoded)
	fmt.Printf("date bytes     %x\n", bs)
	fmt.Println("date bytes    ", bs)
}
