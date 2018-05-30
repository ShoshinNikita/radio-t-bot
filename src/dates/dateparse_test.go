
package dates_test
/*
import (
	dp "dateparse"
	"testing"
	"time"
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
		if dp.ParseDate(test.test) != test.answer {
			t.Errorf("Test #%d Want: %s Got: %s", i, test.answer, test.test)
		}
	}
}

func TestToDays(t *testing.T) {
	tests := []struct {
		time   string
		answer int
	}{
		{"78h35m18s", 3},
		{"120h", 5},
		{"23h", 1},
		{"73h", 3},
	}
	for i, test := range tests {
		d, err := time.ParseDuration(test.time)
		if err != nil {
			t.Error(err)
			continue
		}

		res := dp.ToDays(d)
		if res != test.answer {
			t.Errorf("Test #%d Want: %d Got: %d", i, test.answer, res)
		}
	}
}

func TestNextSaturday(t *testing.T) {
	tests := []struct {
		time   int64
		answer int
	}{
		{
			1527605489, // 2018/5/29 14.51.29
			4,
		},
		{
			1525186289, // 2018/5/1 14.51.29
			4,
		},
		{
			1527951089, // 2018/6/2 14.51.29
			0,
		},
	}

	for i, test := range tests {
		d := time.Unix(test.time, 0)

		res := dp.NextSaturday(d)
		if res != test.answer {
			t.Errorf("Test #%d Want: %d Got: %d", i, test.answer, res)
		}
	}
}

func Test_roundDate(t *testing.T) {
	tests := []struct {
		test   time.Time
		answer time.Time
	}{}

}
*/