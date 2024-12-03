package main

import (
	"fmt"
	"live-chat-be/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/ws", handlers.SocketHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5010"
	}
	fmt.Println("server running on port:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
