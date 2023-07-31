package poego

type GraphQLPostPayload struct {
	QueryName  string      `json:"queryName"`
	Variables  interface{} `json:"variables"`
	Extensions Extensions  `json:"extensions"`
}

type ChatInputMetadata struct {
	UseVoiceRecord bool `json:"useVoiceRecord"`
}

type Source struct {
	SourceType        SourceType        `json:"sourceType"`
	ChatInputMetadata ChatInputMetadata `json:"chatInputMetadata"`
}

type SendMessageArgs struct {
	ChatID        int           `json:"chatId"`
	Bot           string        `json:"bot"`
	Query         string        `json:"query"`
	Source        Source        `json:"source"`
	WithChatBreak bool          `json:"withChatBreak"`
	ClientNonce   string        `json:"clientNonce"`
	Sdid          string        `json:"sdid"`
	Attachments   []interface{} `json:"attachments"`
}

type Extensions struct {
	Hash string `json:"hash"`
}
