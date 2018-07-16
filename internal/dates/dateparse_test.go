package dates_test

import (
	"testing"
	"time"

	. "dates"
)

func TestParse(t *testing.T) {
	tests := []struct {
		test   string
		answer string
	}{
		{"Sat, 26 May 2018 18:11:11 EST", "26 Мая 2018"},
		{"Sat, 21 Apr 2018 17:59:40 EST", "21 Апреля 2018"},
		{"Sat, 17 Mar 2018 18:48:06 EST", "17 Марта 2018"},
	}
	for i, test := range tests {
		if ParseDate(test.test) != test.answer {
			t.Errorf("Test #%d Want: %s Got: %s", i, test.answer, test.test)
		}
	}
}

func TestNextSaturday(t *testing.T) {
	tests := []struct {
		time  int64
		days  int
		hours int
	}{
		{
			1527605489, // 2018/5/29 17.51.29
			4, 6,
		},
		{
			1525186289, // 2018/5/1 17.51.29
			4, 6,
		},
		{
			1527951089, // 2018/6/2 17.51.29
			0, 6,
		},
	}

	for i, test := range tests {
		d, h := NextSaturday(time.Unix(test.time, 0))
		if d != test.days || h != test.hours {
			t.Errorf("Test #%d Want: %d %d Got: %d %d", i, test.days, test.hours, d, h)
		}
	}
}

func TestNextGeekSaturday(t *testing.T) {
	tests := []struct {
		time  int64
		days  int
		hours int
	}{
		{
			1527605489, // 2018/05/29 17.51.29
			4, 6,
		},
		{
			1525186289, // 2018/05/01 17.51.29
			4, 6,
		},
		{
			1527951089, // 2018/06/02 17.51.29
			0, 6,
		},
		{
			1526840659, // 2018/05/20 21.24.19
			13, 2,
		},
	}

	for i, test := range tests {
		d, h := NextGeekSaturday(time.Unix(test.time, 0))
		if d != test.days || h != test.hours {
			t.Errorf("Test #%d Want: %d %d Got: %d %d", i, test.days, test.hours, d, h)
		}
	}
}
