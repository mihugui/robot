package api

type CQMessage struct {
	Type string `json:"type"`
	Data CQData `json:"data"`
}

type CQData struct {
	// 消息
	Text string `json:"text"`
	// 图片
	File string `json:"file"`
	// @
	QQ string `json:"qq"`
	// 链接
	Url   string `json:"url"`
	Title string `json:"title"`
}
