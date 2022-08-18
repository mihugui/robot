package api

import (
	"encoding/json"
	"fmt"
	"strings"
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
		fmt.Println("消息类型:" + message.MessageType)
		fmt.Println("消息内容:" + message.Message)
		fmt.Println("消息id:" + fmt.Sprintf("%d", message.MessageId))
		fmt.Println("消息发送人id:" + fmt.Sprintf("%d", message.UserId))
		fmt.Println("消息发送人昵称:" + message.Sender.Nickname)

		// 判断为正常消息或者CQ消息
		msg := Analyse(message.Message)

		if message.MessageType == "group" {
			msgType, _ := json.Marshal(SendGroupMsg(message.GroupId, msg))
			done <- msgType
		}

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

func Analyse(message string) []CQMessage {

	var cqMsg []CQMessage

	// 判断是否符合规范
	msgs := strings.Split(message, " ")
	switch msgs[0] {
	case "测试":
		if len(msgs) == 1 {
			cqMsg = append(cqMsg, CQMessage{
				Type: "text",
				Data: CQData{
					Text: "缺少后续指令",
				},
			})
		} else {
			for index, msg := range msgs {
				if index == 0 {
					continue
				}
				cqMsg = append(cqMsg, CQMessage{
					Type: "text",
					Data: CQData{
						Text: msg,
					},
				})

				cqMsg = append(cqMsg, CQMessage{
					Type: "text",
					Data: CQData{
						Text: " ",
					},
				})
			}
		}
	case "对联":
		break

	}

	return cqMsg
}
