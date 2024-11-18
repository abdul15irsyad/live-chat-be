package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
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
		BroadcastMessage(&Message{
			Type:    "badge",
			Content: clients[conn].Name + " left",
		}, conn, true)
		return nil
	})
	clients[conn] = Client{
		Name: name,
	}

	fmt.Printf("%s connected\n", clients[conn].Name)
	BroadcastMessage(&Message{
		Type:    "badge",
		Content: clients[conn].Name + " joined",
	}, conn, false)

	for {
		message := Message{}
		if err := conn.ReadJSON(&message); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(clients[conn].Name+" send: ", message.Content)
		BroadcastMessage(&message, conn, true)
	}
}

func BroadcastMessage(message *Message, conn *websocket.Conn, isExcludeSender bool) {
	for client := range clients {
		if client == conn && isExcludeSender {
			continue
		}
		if err := client.WriteJSON(*message); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
