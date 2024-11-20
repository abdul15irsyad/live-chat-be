package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Payload[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type JoinData struct {
	Name string `json:"name"`
}

type Client struct {
	Name string
}

var clients = map[*websocket.Conn]Client{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// get name
	queryParams := request.URL.Query()
	name := queryParams.Get("name")
	conn.SetCloseHandler(func(code int, text string) error {
		BroadcastMessage(&Payload[JoinData]{
			Type: "left",
			Data: JoinData{
				Name: clients[conn].Name,
			},
		}, conn)
		return nil
	})
	clients[conn] = Client{
		Name: name,
	}

	fmt.Printf("%s connected\n", clients[conn].Name)
	BroadcastMessage(&Payload[JoinData]{
		Type: "joined",
		Data: JoinData{
			Name: clients[conn].Name,
		},
	}, conn)

	for {
		var JSONMessage map[string]any
		if err := conn.ReadJSON(&JSONMessage); err != nil {
			fmt.Println(err)
			break
		}
		switch JSONMessage["type"] {
		case "message":
			jsonData, _ := json.Marshal(JSONMessage)
			var message Payload[MessageData]
			_ = json.Unmarshal(jsonData, &message)
			fmt.Printf("%s send: %s\n", clients[conn].Name, message.Data.Message)
			BroadcastMessage(&message, conn)
		default:
			fmt.Printf("type %s not found\n", JSONMessage["type"])
		}
	}
}

func BroadcastMessage[T any](message *Payload[T], conn *websocket.Conn) {
	for client := range clients {
		if client == conn {
			continue
		}
		if err := client.WriteJSON(*message); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
