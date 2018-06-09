// Package dialogs provide opportunities to work with Yandex.Dialogs
package dialogs

import (
	"encoding/json"
	"net/http"
)

type DialogsAPI struct {
	DistributeFunc func(Request) (text, tts string, buttons []Button, endSession bool)
}

func (api *DialogsAPI) StartSevring() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		json.NewDecoder(r.Body).Decode(&req)

		var res Response
		res.Response.Text, res.Response.TTS, res.Response.Buttons, res.Response.EndSession = api.DistributeFunc(req)

		res.Session.SessionID = req.Session.SessionID
		res.Session.MessageID = req.Session.MessageID
		res.Session.UserID = req.Session.UserID
		res.Version = req.Version

		json.NewEncoder(w).Encode(res)
	}
}
