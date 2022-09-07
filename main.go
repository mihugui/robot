package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"robot/modules/api"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "server.mihugui.cn:10080", "http service address")

func main() {

	var input string

	fmt.Print("请输入密钥:")
	fmt.Scanln(&input)

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, ForceQuery: true, RawQuery: "access_token=" + input}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan []byte)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			// 数据处理 放入管道
			api.Receive(message, done)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		// 管道消息
		case t := <-done:
			err := c.WriteMessage(websocket.TextMessage, t)
			if err != nil {
				log.Println("write:", err)
				return
			}
		// 定时发送消息
		case <-ticker.C:

		// 事件消息
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
