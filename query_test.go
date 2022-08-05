package bcdt_test

import (
	"bytes"
	"testing"
	"time"

	bcdt "github.com/snoworly/bitwise-comparable-datetime-go"
)

func TestQueryDateTime(t *testing.T) {
	var epoch int64 = 1659079415
	datetime := time.Unix(epoch, 0).UTC()

	got := bcdt.QueryDatetime(datetime.Year(), int(datetime.Month()), datetime.Day(), datetime.Hour(), datetime.Minute())

	expected := []byte{7, 29, 7, 23}
	if bytes.Compare(got[:], expected) != 0 {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestQueryDate(t *testing.T) {
	var epoch int64 = 1659079415
	datetime := time.Unix(epoch, 0).UTC()

	got := bcdt.QueryDate(datetime.Year(), int(datetime.Month()), datetime.Day())

	expected := []byte{7, 29}
	if bytes.Compare(got[:], expected) != 0 {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestQueryYearMonth(t *testing.T) {
	var epoch int64 = 1674930387
	datetime := time.Unix(epoch, 0).UTC()

	got := bcdt.QueryYearMonth(datetime.Year(), int(datetime.Month()))

	expected := []byte{12}
	if bytes.Compare(got[:], expected) != 0 {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
