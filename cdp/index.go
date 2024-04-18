package cdp

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有跨域请求
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		// 检测是否为Debugger.enable命令
		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("error unmarshalling message:", err)
			continue
		}

		if msg["method"] == "Debugger.enable" {
			// 构造一个简单的响应
			response := map[string]interface{}{
				"id":     msg["id"],                // 回应相同的ID
				"result": map[string]interface{}{}, // 通常这里会有更复杂的数据
			}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Println("error marshalling response:", err)
				continue
			}
			if err := ws.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func Start() {
	http.HandleFunc("/", handleConnections)
	log.Println("http server started on :9229")
	err := http.ListenAndServe(":9229", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
