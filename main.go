package main

import (
	"fmt"
	"live-chat-be/handlers"
	"live-chat-be/middlewares"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/clients", handlers.GetAllClientHandler)
	http.HandleFunc("/ws", handlers.SocketHandler)
	http.HandleFunc("/", handlers.RootHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5010"
	}
	fmt.Println("server running on port:", port)
	if err := http.ListenAndServe(":"+port, middlewares.CorsMiddleware(http.DefaultServeMux)); err != nil {
		panic(err)
	}
}
