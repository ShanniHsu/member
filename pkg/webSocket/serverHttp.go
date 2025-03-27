package webSocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func ServerHTTP() {
	go func() {
		g := gin.New()
		g.Use(gin.Recovery())
		err := g.SetTrustedProxies(nil)
		if err != nil {
			panic(err)
		}

		public := g.Group("/socket")
		public.GET("", SocketHandler)
		if err = g.Run(":8080"); err != nil {
			panic(err)
		}
	}()
}

// 客戶端發送訊息，Server自動回覆
func SocketHandler(c *gin.Context) {
	// CheckOrigin: 會檢查是否跨域
	// Buffer: 單位是Bytes，依需求設定(設為0，則為不限制大小)
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// 將HTTP連線轉換成WebSocket
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// 這邊開一個goroutine，如果有接收到斷線資訊則關閉此連線
	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			panic(err)
		}
	}()

	for {
		// 監聽客戶端傳來的訊息
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))
		// Server自動傳訊息給客戶端
		err = ws.WriteJSON(struct {
			Reply string `json:"reply"`
		}{
			Reply: "Echo...",
		})
		if err != nil {
			panic(err)
		}
	}
}

// https://tocandraw.com/coding/golang/111/
// 使用ws://localhost:8080/socket
// 可使用Google擴充工具 PieSocket模仿client
