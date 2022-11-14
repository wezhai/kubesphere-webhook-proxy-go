package qyapi

type MarkdownMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}
