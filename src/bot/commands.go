package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"

	"dates"
)

// Returns main information about the podcast
func mainInfo() (text, tts string, buttons []Button, endSession bool, err error) {
	text = mainInformationText
	tts = mainInformationTTS
	buttons = []Button{
		Button{Title: "Закончить ❌"},
		Button{Title: "Сайт подкаста", URL: site, Hide: false},
		Button{Title: "Последний выпуск", Hide: false},
		Button{Title: "Следующий выпуск", Hide: false},
		Button{Title: "Следующий гиковский выпуск", Hide: false},
	}

	return text, tts, buttons, false, nil
}

// Returns URL of the site
func siteURL() (text, tts string, buttons []Button, endSession bool, err error) {
	buttons = append(buttons, Button{Title: "Сайт подкаста", URL: site, Hide: false})
	return "", "", buttons, true, nil
}

// Returns number and URL of the last release
func lastRelease() (text, tts string, buttons []Button, endSession bool, err error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(rssURL)
	if err != nil {
		return "", "", buttons, true, err
	}

	number := getNumber(feed.Items[0].Title)
	url := feed.Items[0].GUID
	themes := parseNews(feed.Items[0].Description)
	date := dates.ParseDate(feed.Items[0].Published)

	// Clear // in url
	// (https://radio-t.com/p/2018/05/26//podcast-599/ -> https://radio-t.com/p/2018/05/26/podcast-599/)
	url = strings.Replace(url, "//podcast", "/podcast", -1)

	buttons = []Button{
		Button{Title: "Перейти к выпуску", URL: url},
		Button{Title: "Закончить ❌"},
		Button{Title: "Сайт подкаста", URL: site},
		Button{Title: "Следующий выпуск"},
		Button{Title: "Следующий гиковский выпуск"},
	}

	text = fmt.Sprintf(lastReleaseText, number, date, themes)
	tts = fmt.Sprintf(lastReleaseTTS, number, date)

	return text, tts, buttons, false, nil
}

// Returns left time (id days) until the next release
func nextRelease() (text, tts string, buttons []Button, endSession bool, err error) {
	days, hours := dates.NextSaturday(time.Now())
	day := dates.ParseDays(days)
	hour := dates.ParseHours(hours)

	text = fmt.Sprintf(nextReleaseText, days, day, hours, hour)
	tts = fmt.Sprintf(nextReleaseTTS, days, day, hours, hour)
	buttons = []Button{
		Button{Title: "Закончить ❌"},
		Button{Title: "Сайт подкаста", URL: site, Hide: false},
		Button{Title: "Последний выпуск", Hide: false},
		Button{Title: "Следующий гиковский выпуск", Hide: false},
	}

	return text, tts, buttons, false, nil
}

// Returns left time (id days) until next geek release
func nextGeekRelease() (text, tts string, buttons []Button, endSession bool, err error) {
	days, hours := dates.NextGeekSaturday(time.Now())
	day := dates.ParseDays(days)
	hour := dates.ParseHours(hours)

	text = fmt.Sprintf(nextGeekReleaseText, days, day, hours, hour)
	tts = fmt.Sprintf(nextGeekReleaseTTS, days, day, hours, hour)
	buttons = []Button{
		Button{Title: "Закончить ❌"},
		Button{Title: "Сайт подкаста", URL: site, Hide: false},
		Button{Title: "Последний выпуск", Hide: false},
		Button{Title: "Следующий выпуск", Hide: false},
	}

	return text, tts, buttons, false, nil
}

func endConverseation() (text, tts string, buttons []Button, endSession bool, err error) {
	return endConverseationText, endConverseationTTS, buttons, true, nil
}
