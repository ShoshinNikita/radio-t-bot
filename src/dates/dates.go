package dates

import (
	"math"
	"time"
)

func roundDate(t time.Time) time.Time {
	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	r := time.Duration(h)*time.Hour +
		time.Duration(m)*time.Minute +
		time.Duration(s)*time.Second

	return t.Add(-r)
}

func toInt(f float64) int {
	return int(math.Ceil(f))
}

func NextSaturday(now time.Time) (days, hours int) {
	sat := now
	for sat.Weekday() != time.Saturday {
		sat = sat.Add(24 * time.Hour)
	}
	sat = roundDate(sat)
	sat = sat.Add(23 * time.Hour)
	left := sat.Sub(now)

	days = toInt(left.Hours()) / 24
	hours = toInt(left.Hours()) % 24

	return days, hours
}

func NextGeekSaturday(now time.Time) (days, hours int) {
	if now.Day() < 7 && now.Weekday() <= time.Saturday {
		return NextSaturday(now)
	}

	geekSat := now.AddDate(0, 1, -now.Day())
	for geekSat.Weekday() != time.Saturday {
		geekSat = geekSat.Add(24 * time.Hour)
	}
	geekSat = roundDate(geekSat)
	geekSat = geekSat.Add(23 * time.Hour)
	left := geekSat.Sub(now)

	days = toInt(left.Hours()) / 24
	hours = toInt(left.Hours()) % 24

	return days, hours
}
