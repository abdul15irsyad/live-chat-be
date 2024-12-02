package services

import (
	"fmt"
	"live-chat-be/types"

	"github.com/gorilla/websocket"
)

func BroadcastMessage[T any](message *types.Payload[T], conn *websocket.Conn, clients *map[*websocket.Conn]types.Client) {
	for client := range *clients {
		if client == conn {
			continue
		}
		if err := client.WriteJSON(*message); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
