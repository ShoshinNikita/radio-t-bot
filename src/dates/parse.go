package dates

import (
	"strings"
)

var months = map[string]string{
	"Jan": "Января",
	"Feb": "Февраля",
	"Mar": "Марта",
	"Apr": "Апреля",
	"May": "Мая",
	"Jun": "Июня",
	"Jul": "Июля",
	"Aug": "Августа",
	"Sep": "Сентября",
	"Oct": "Октября",
	"Nov": "Ноября",
	"Dec": "Декабря",
}

// [start, end]
type interval struct {
	start int
	end   int
}

var dayForms = map[interval]string{
	interval{0, 0}:   "дней",
	interval{1, 1}:   "день",
	interval{2, 4}:   "дня",
	interval{5, 20}:  "дней",
	interval{21, 21}: "день",
	interval{22, 24}: "дня",
	interval{25, 30}: "дней",
	interval{31, 31}: "день",
	interval{32, 34}: "дня",
	interval{35, 40}: "дней",
	interval{41, 41}: "день",
	interval{42, 44}: "дня",
	interval{45, 50}: "дней",
}

var hourForms = map[interval]string{
	interval{0, 0}:   "часов",
	interval{1, 1}:   "час",
	interval{2, 4}:   "часа",
	interval{5, 20}:  "часов",
	interval{21, 21}: "час",
	interval{22, 23}: "часа",
}

// ParseDate transform string of time
// "Sat, 26 May 2018 18:11:11 EST" -> "26 Мая 2018"
func ParseDate(s string) (res string) {
	arr := strings.Split(s, " ")
	res += arr[1] + " "
	res += months[arr[2]] + " "
	res += arr[3]

	return res
}

// ParseDays returns correct form of the word "день"
func ParseDays(d int) (res string) {
	for k, v := range dayForms {
		if k.start <= d && d <= k.end {
			res = v
			break
		}
	}

	return res
}

// ParseHours returns correct form of the word "час"
func ParseHours(h int) (res string) {
	for k, v := range hourForms {
		if k.start <= h && h <= k.end {
			res = v
			break
		}
	}

	return res
}
