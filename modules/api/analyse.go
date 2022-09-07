package api

import (
	"encoding/json"
	"fmt"
	"robot/modules/utils/draw"
	"robot/modules/utils/joke"
	"robot/modules/utils/qiniu"
	"robot/modules/utils/weather"
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

		if len(msg) == 0 {
			return
		}

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
	case "文字转图片":
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

				// 转换成图片文字
				if draw.WordToPic(msg) {
					url, _ := qiniu.Upload(msg+".png", "out.png")
					cqMsg = append(cqMsg, CQMessage{
						Type: "image",
						Data: CQData{
							File: url,
						},
					})
				}

			}
		}
	case "天气":
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

				// 获取天气情况
				info := weather.GetWeather(msg)
				cqMsg = append(cqMsg, CQMessage{

					Type: "text",
					Data: CQData{
						Text: info,
					},
				})
			}
		}
	case "讲笑话":
		joke := joke.GetJoke()
		cqMsg = append(cqMsg, CQMessage{

			Type: "text",
			Data: CQData{
				Text: joke,
			},
		})

	}

	return cqMsg
}
