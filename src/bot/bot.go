package bot

import (
	"encoding/json"
	"net/http"
	"strings"

	"logging"
)

// Index serves POST requests
func Index(w http.ResponseWriter, r *http.Request) {
	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	var res Response
	res.Response.Text, res.Response.TTS, res.Response.Buttons, res.Response.EndSession = distribute(req)

	res.Session.SessionID = req.Session.SessionID
	res.Session.MessageID = req.Session.MessageID
	res.Session.UserID = req.Session.UserID
	res.Version = req.Version

	json.NewEncoder(w).Encode(res)
}

var commands = []struct {
	keyWords []string
	handler  func() (string, string, []Button, bool, error)
}{
	{[]string{
		"подкаст выходного дня",
		"радио-т",
		"радиот",
		"радио т",
		"что такое подкаст выходного дня",
		"что такое радио-т",
		"что такое радиот",
		"что такое радио т"},
		mainInfo}, // main info
	{[]string{
		"сайт"},
		siteURL}, // url of the site
	{[]string{
		"последний выпуск",
		"новый выпуск",
		"информация о последнем выпуске",
		"информация о новом выпуске"},
		lastRelease}, // info about the last release
	{[]string{"следующий выпуск"},
		nextRelease}, // date of the next release
	{[]string{
		"следующий гиковский выпуск",
		"гиковский выпуск"},
		nextGeekRelease}, // date of the next geek release
	{[]string{
		"помощь",
		"что умеет",
		"что умеешь",
		"что ты умеешь",
		"что этот бот умеет",
		"что этот бот может",
		"что можешь",
		"что ты можешь"},
		botInfo}, // send info about the bot
	{[]string{
		"проиграй",
		"запусти",
		"включи",
		"послушать",
		"слушать"},
		playRelease},
	{[]string{
		"закончить",
		"всё",
		"все",
		"конец",
		"до свидания"},
		endConverseation}, // stop dialogue
}

func distribute(req Request) (text, tts string, buttons []Button, endSession bool) {
	command := strings.ToLower(req.Request.Command)
	// For Yandex ping
	if command == "ping" {
		return "pong", "pong", buttons, true
	}

	logging.LogRequest(command, req.Session.SessionID)

	// If phrase == "Запусти навык подкаст выходного дня"
	if command == "" {
		text, tts, buttons, endSession, _ = botInfo()
		return text, tts, buttons, endSession
	}

	rightCommand := false
	var err error
	for i := 0; i < len(commands) && !rightCommand; i++ {
		for j := 0; j < len(commands[i].keyWords) && !rightCommand; j++ {
			keyWord := commands[i].keyWords[j]
			if strings.Contains(command, keyWord) {
				text, tts, buttons, endSession, err = commands[i].handler()
				rightCommand = true
			}
		}
	}

	if err != nil || !rightCommand {
		buttons = []Button{
			Button{Title: "Помощь"},
			Button{Title: "Сайт подкаста", URL: "https://radio-t.com/", Hide: false},
			Button{Title: "Последний выпуск", Hide: false},
			Button{Title: "Следующий выпуск", Hide: false},
			Button{Title: "Следующий гиковский выпуск", Hide: false},
			Button{Title: "Закончить ❌"},
		}
		endSession = false

		if err != nil {
			logging.LogError(err)
			text = errorText
			tts = errorTTS
		} else if !rightCommand {
			text = wrongCommandText
			tts = wrongCommandTTS
		}
	}

	return text, tts, buttons, endSession
}
