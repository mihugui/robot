package api

import (
	"time"
)

func SendPrivateMsg(userId int64, message []CQMessage) WebSocket {
	//TODO send private message
	var ws WebSocket

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

func SendGroupMsg(gruopId int64, message []CQMessage) WebSocket {
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
