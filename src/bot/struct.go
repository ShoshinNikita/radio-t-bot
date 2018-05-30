package bot

type Request struct {
	Meta struct {
		Locale   string `json:"locale,omitempty"`
		Timezone string `json:"timezone,omitempty"`
		ClientID string `json:"client_id,omitempty"`
	} `json:"meta,omitempty"`
	Request struct {
		Command           string `json:"command,omitempty"`
		OriginalUtterance string `json:"original_utterance,omitempty"`
		Type              string `json:"type,omitempty"`
		Markup            struct {
			DangerousContext bool `json:"dangerous_context,omitempty"`
		} `json:"markup,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
	} `json:"request,omitempty"`
	Session struct {
		New       bool   `json:"new,omitempty"`
		MessageID int64  `json:"message_id,omitempty"`
		SessionID string `json:"session_id,omitempty"`
		SkillID   string `json:"skill_id,omitempty"`
		UserID    string `json:"user_id,omitempty"`
	} `json:"session,omitempty"`
	Version string `json:"version,omitempty"`
}

type Button struct {
	Title   string      `json:"title,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	URL     string      `json:"url,omitempty"`
	Hide    bool        `json:"hide,omitempty"`
}

type Response struct {
	Response struct {
		Text       string   `json:"text,omitempty"`
		TTS        string   `json:"tts,omitempty"`
		Buttons    []Button `json:"buttons,omitempty"`
		EndSession bool     `json:"end_session,omitempty"`
	} `json:"response,omitempty"`
	Session struct {
		MessageID int64  `json:"message_id,omitempty"`
		SessionID string `json:"session_id,omitempty"`
		UserID    string `json:"user_id,omitempty"`
	} `json:"session,omitempty"`
	Version string `json:"version,omitempty"`
}
