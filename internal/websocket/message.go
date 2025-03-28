package websocket

// import (
// 	"_/C_/Users/Алмат/Desktop/message_micro/internal/websocket"
// 	"net/http"
// )

// var upgrader = websocket.upgrader{}

// type WebsocketServer struct {
// 	RedisClient *redis.Client
// }

// func NewWebSocket(redisClient *redis.Client) *WebsocketServer {
// 	return &WebsocketServer{RedisClient: redisClient}
// }

// func (w *WebsocketServer) HandlerConn(wr http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(wr, r, nil)
// 	if err != nil {
// 		return
// 	}
// 	defer conn.Close()

// 	pubsub := w.RedisClient.Subscribe(r.Context(), "chat_messages")
// 	defer pubsub.Close()

// 	for msg := range pubsub.Channel() {
// 		err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
// 		if err != nil {
// 			return
// 		}
// 	}
// }