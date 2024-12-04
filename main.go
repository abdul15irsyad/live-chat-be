package main

import (
	"fmt"
	"live-chat-be/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allows all origins, replace with specific domains if needed
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/clients", handlers.GetAllClientHandler)
	http.HandleFunc("/ws", handlers.SocketHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5010"
	}
	fmt.Println("server running on port:", port)
	if err := http.ListenAndServe(":"+port, corsHandler.Handler(http.DefaultServeMux)); err != nil {
		panic(err)
	}
}
