package handlers

import (
	"encoding/json"
	"fmt"
	"live-chat-be/services"
	"live-chat-be/types"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var clients = map[*websocket.Conn]types.Client{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketHandler(writer http.ResponseWriter, request *http.Request) {
	// get name
	queryParams := request.URL.Query()
	name := queryParams.Get("name")

	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	currentTime := time.Now()
	conn.SetCloseHandler(func(code int, text string) error {
		currentTime = time.Now()
		fmt.Printf("%s: %s disconnected\n", currentTime.Format("2006-01-02 15:04:05"), clients[conn].Name)
		services.BroadcastMessage(&types.Payload[types.Client]{
			Type:      "left",
			Data:      clients[conn],
			Timestamp: currentTime,
		}, conn, &clients)
		conn.Close()
		delete(clients, conn)
		return nil
	})
	clients[conn] = types.Client{
		Id:   uuid.New(),
		Name: name,
	}

	fmt.Printf("%s: %s connected\n", currentTime.Format("2006-01-02 15:04:05"), clients[conn].Name)
	services.BroadcastMessage(&types.Payload[types.Client]{
		Type:      "joined",
		Data:      clients[conn],
		Timestamp: currentTime,
	}, conn, &clients)

	for {
		var JSONMessage map[string]any
		if err := conn.ReadJSON(&JSONMessage); err != nil {
			fmt.Println(err)
			break
		}
		switch JSONMessage["type"] {
		case "message":
			jsonData, _ := json.Marshal(JSONMessage)
			var message types.Payload[types.MessageData]
			_ = json.Unmarshal(jsonData, &message)
			message.Timestamp = time.Now()
			message.Client = clients[conn]
			fmt.Printf("%s: %s send \"%s\"\n", message.Timestamp.Format("2006-01-02 15:04:05"), clients[conn].Name, message.Data.Message)
			services.BroadcastMessage(&message, conn, &clients)
		default:
			fmt.Printf("type \"%s\" not found\n", JSONMessage["type"])
		}
	}
}
