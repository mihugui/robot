package main

import "encoding/json"

type message struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int32  `json:"message_id"`
	UserId      int64  `json:"user_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        string `json:"font"`
	Sender      string `json:"sender"`
	TempSource  string `json:"temp_source"`
}

type WebSocket struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

func main() {
	jsonStr := `{
		"action": "send_group_msg",
		"params": {
			"message_type":"group",
			"time": 1659679907,
			"group_id":"708705387",
			"message": "2"
			},
		"echo": "'回声', 如"
	}`

	var ws WebSocket
	json.Unmarshal([]byte(jsonStr), &ws)
}
