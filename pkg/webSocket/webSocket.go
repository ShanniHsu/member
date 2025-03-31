package webSocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

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
		log.Println("WebSocket連線失敗: ", err)
		return
	}

	defer func() {
		fmt.Println("用戶斷線，發送關閉訊息")
		// 這邊是伺服器端主動優雅關閉連線
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Bye"))
		conn.Close()
	}()

	// 紀錄這個客戶端
	clients[conn] = true
	fmt.Println("clients: ", clients)

	for {
		var msg Message
		// 監聽客戶端傳來的訊息
		err = conn.ReadJSON(&msg)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				log.Println("訊息格式錯誤,請確認格式")
				continue
			}

			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("客戶端正常關閉連線")
				delete(clients, conn)
				break
			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseInternalServerErr) {
				log.Println("客戶端異常斷線")
				delete(clients, conn)
				break
			} else {
				log.Println("其他錯誤: ", err)
				delete(clients, conn)
				break
			}

		}
		// 把收到的訊息發送到廣播
		broadcast <- msg
	}
}

func Broadcast() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("傳送訊息錯誤:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// https://tocandraw.com/coding/golang/111/
// 使用ws://localhost:8080/socket
// 可使用Google擴充工具 PieSocket模仿client
