package types

import (
	"time"

	"github.com/google/uuid"
)

type Payload[T any] struct {
	Type      string    `json:"type"`
	Client    Client    `json:"client"`
	Data      T         `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

type MessageData struct {
	Message string `json:"message"`
}

type TypingData struct {
	Status bool `json:"status"`
}

type Client struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
