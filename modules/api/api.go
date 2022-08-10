package api

import "time"

type WebSocket struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

func SendPrivateMsg(userId int64, message string) *WebSocket {
	//TODO send private message
	var ws *WebSocket

	ws.Action = "send_private_msg"
	ws.Params = map[string]interface{}{
		"user_id":     userId,
		"message":     message,
		"auto_escape": true,
		"time":        time.Now().Unix(),
	}
	ws.Echo = "send_private_msg"

	return ws

}

func SendGroupMsg(gruopId int64, message string) WebSocket {
	//TODO send private message
	var ws WebSocket

	ws.Action = "send_group_msg"
	ws.Params = map[string]interface{}{
		"group_id":    gruopId,
		"message":     message,
		"auto_escape": true,
		"time":        time.Now().Unix(),
	}
	ws.Echo = "send_group_msg"

	return ws

}
