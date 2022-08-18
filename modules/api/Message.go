package api

type Report struct {
	PostType string `json:"post_type"`
	SelfId   int64  `json:"self_id"`
	Time     int64  `json:"time"`
}

// 上报消息
type Message struct {
	SubType    string        `json:"sub_type"`
	Message    string        `json:"message"`
	UserId     int64         `json:"user_id"`
	GroupId    int64         `json:"group_id"`
	MessageId  int32         `json:"message_id"`
	RawMessage string        `json:"raw_message"`
	Font       int           `json:"font"`
	Sender     MessageSender `json:"sender"`
}

// 消息发送人员
type MessageSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Card     string `json:"card"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}

type Request struct {
	RequestType string `json:"request_type"`
}

type Notice struct {
	RequestType string `json:"request_type"`
}

type MetaEvent struct {
	MetaEventType string `json:"meta_event_type"`
}
