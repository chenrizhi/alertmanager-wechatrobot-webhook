package model

type WeChatMessage struct {
}

type WeChatMarkdown struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}
