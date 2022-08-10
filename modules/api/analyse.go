package api

import (
	"encoding/json"
	"fmt"
)

// 上报信息
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

func Analyse(post []byte, done chan []byte) {

	var report Report

	err := json.Unmarshal(post, &report)

	if err != nil {
		fmt.Println("解析失败报文:" + string(post))
		return
	}

	// 消息类型
	switch report.PostType {
	case "message":
		var message Message
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		fmt.Println("消息类型:" + message.SubType)
		fmt.Println("消息内容:" + message.Message)
		fmt.Println("消息id:" + fmt.Sprintf("%d", message.MessageId))
		fmt.Println("消息发送人id:" + fmt.Sprintf("%d", message.UserId))
		fmt.Println("消息发送人昵称:" + message.Sender.Nickname)

		msg, _ := json.Marshal(SendGroupMsg(message.GroupId, message.Message))

		done <- msg

	case "request":
		var message Request
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		fmt.Println("消息类型:" + message.RequestType)
	case "notice":
		var message Notice
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		fmt.Println("消息类型:" + message.RequestType)
	case "meta_event":
		var message MetaEvent
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		fmt.Println("消息类型:" + message.MetaEventType)
	default:
		fmt.Println("未知消息类型:" + report.PostType)

	}

}
