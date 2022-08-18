package api

import (
	"encoding/json"
	"fmt"
)

// 上报信息

func Receive(post []byte, done chan []byte) {

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
		fmt.Println("消息:" + string(post))
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

		// 判断为正常消息或者CQ消息
		Analyse(message.Message)

		//done <- msg
	case "request":
		var message Request
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		//fmt.Println("消息类型:" + message.RequestType)
	case "notice":
		var message Notice
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		//fmt.Println("消息类型:" + message.RequestType)
	case "meta_event":
		var message MetaEvent
		err := json.Unmarshal(post, &message)
		if err != nil {
			fmt.Println("解析失败报文:" + string(post))
			return
		}
		//fmt.Println("消息类型:" + message.MetaEventType)

	}

}

func Analyse(message string) {

}
