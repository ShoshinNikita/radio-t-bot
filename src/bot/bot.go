package bot

import (
	"net/http"
	"strings"

	"dialogs"
	"logging"
)

// Init return function for serving of requests
func Init() func(http.ResponseWriter, *http.Request) {
	api := dialogs.DialogsAPI{DistributeFunc: distribute}
	return api.StartSevring()
}

var commands = []struct {
	keyWords []string
	handler  func() (string, string, []dialogs.Button, bool, error)
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

func distribute(req dialogs.Request) (text, tts string, buttons []dialogs.Button, endSession bool) {
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

	var (
		rightCommand bool // default == false
		err          error
	)
	for i := 0; i < len(commands) && !rightCommand; i++ {
		for j := 0; j < len(commands[i].keyWords) && !rightCommand; j++ {
			keyWord := commands[i].keyWords[j]
			if strings.Contains(command, keyWord) {
				text, tts, buttons, endSession, err = commands[i].handler()
				rightCommand = true
			}
		}
	}

	if err != nil {
		logging.LogError(err)
		text, tts, buttons, endSession = serveError()
	} else if !rightCommand {
		text, tts, buttons, endSession = wrongCommand()
	}

	return text, tts, buttons, endSession
}
