package webSocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//func ServerHTTP() {
//	go func() {
//		g := gin.New()
//		g.Use(gin.Recovery())
//		err := g.SetTrustedProxies(nil)
//		if err != nil {
//			panic(err)
//		}
//
//		public := g.Group("/socket")
//		public.GET("", SocketHandler)
//		if err = g.Run(":8080"); err != nil {
//			panic(err)
//		}
//	}()
//}

type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool) // 線上連線的客戶端
var broadcast = make(chan Message)           // 訊息廣播 Channel

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 多個客戶端發送訊息
func SocketHandler(c *gin.Context) {
	// 將HTTP連線轉換成WebSocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// 這邊開一個goroutine，如果有接收到斷線資訊則關閉此連線
	defer func() {
		closeSocketErr := conn.Close()
		if closeSocketErr != nil {
			panic(err)
		}
	}()

	fmt.Println("conn: ", &conn)
	// 紀錄這個客戶端
	clients[conn] = true

	for {
		var msg Message
		// 監聽客戶端傳來的訊息
		err = conn.ReadJSON(&msg)
		if err != nil {
			log.Panicln("讀取訊息錯誤: ", err)
			delete(clients, conn)
			break
		}
		fmt.Println("msg: ", msg)
		// 把收到的訊息發送到廣播
		broadcast <- msg

		// 以下需再檢查廣播可能有問題要調整

		// 廣播訊息
		broadcastMsg := <-broadcast
		fmt.Println("broadcastMsg: ", broadcastMsg)
		err = conn.WriteJSON(&broadcastMsg)
		if err != nil {
			log.Println("讀取錯誤: ", err)
			conn.Close()
			delete(clients, conn)
		}

	}
}

// https://tocandraw.com/coding/golang/111/
// 使用ws://localhost:8080/socket
// 可使用Google擴充工具 PieSocket模仿client
