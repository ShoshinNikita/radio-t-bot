package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"

	"dates"
	"dialogs"
)

// Returns main information about the podcast
func mainInfo() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	text = mainInformationText
	tts = mainInformationTTS
	buttons = defaultButtons

	return text, tts, buttons, false, nil
}

// Returns URL of the site
func siteURL() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	buttons = append(buttons, dialogs.Button{Title: "Сайт подкаста", URL: site})
	return "Сайт подкаста Радио-Т", "", buttons, false, nil
}

// Returns number and URL of the last release
func lastRelease() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(rssURL)
	if err != nil {
		return "", "", buttons, true, err
	}

	number := getNumber(feed.Items[0].Title)
	url := feed.Items[0].GUID
	themes, err := parseNews(feed.Items[0].Description)
	if err != nil {
		return "", "", buttons, false, err
	}
	date := dates.ParseDate(feed.Items[0].Published)

	// Clear // in url
	// (https://radio-t.com/p/2018/05/26//podcast-599/ -> https://radio-t.com/p/2018/05/26/podcast-599/)
	url = strings.Replace(url, "//podcast", "/podcast", -1)

	buttons = []dialogs.Button{
		dialogs.Button{Title: "Перейти к выпуску", URL: url},
		dialogs.Button{Title: "Сайт подкаста", URL: site},
		dialogs.Button{Title: "Следующий выпуск"},
		dialogs.Button{Title: "Следующий гиковский выпуск"},
		dialogs.Button{Title: "Закончить ❌"},
	}

	text = fmt.Sprintf(lastReleaseText, number, date, themes)
	tts = fmt.Sprintf(lastReleaseTTS, number, date)

	return text, tts, buttons, false, nil
}

// Returns left time (id days) until the next release
func nextRelease() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	days, hours := dates.NextSaturday(time.Now())
	day := dates.ParseDays(days)
	hour := dates.ParseHours(hours)

	text = fmt.Sprintf(nextReleaseText, days, day, hours, hour)
	tts = fmt.Sprintf(nextReleaseTTS, days, day, hours, hour)
	buttons = []dialogs.Button{
		dialogs.Button{Title: "Сайт подкаста", URL: site},
		dialogs.Button{Title: "Последний выпуск"},
		dialogs.Button{Title: "Следующий гиковский выпуск"},
		dialogs.Button{Title: "Закончить ❌"},
	}

	return text, tts, buttons, false, nil
}

// Returns left time (id days) until next geek release
func nextGeekRelease() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	days, hours := dates.NextGeekSaturday(time.Now())
	day := dates.ParseDays(days)
	hour := dates.ParseHours(hours)

	text = fmt.Sprintf(nextGeekReleaseText, days, day, hours, hour)
	tts = fmt.Sprintf(nextGeekReleaseTTS, days, day, hours, hour)
	buttons = []dialogs.Button{
		dialogs.Button{Title: "Сайт подкаста", URL: site},
		dialogs.Button{Title: "Последний выпуск"},
		dialogs.Button{Title: "Следующий выпуск"},
		dialogs.Button{Title: "Закончить ❌"},
	}

	return text, tts, buttons, false, nil
}

func playRelease() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	buttons = defaultButtons
	return playReleaseText, playReleaseTTS, buttons, false, nil
}

func botInfo() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	buttons = defaultButtons
	return botInfoText, botInfoTTS, buttons, false, nil
}

func endConverseation() (text, tts string, buttons []dialogs.Button, endSession bool, err error) {
	return endConverseationText, endConverseationTTS, buttons, true, nil
}
