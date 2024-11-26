package types

import (
	"time"

	"github.com/google/uuid"
)

type Payload[T any] struct {
	Type      string    `json:"type"`
	Data      T         `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

type MessageData struct {
	Message string `json:"message"`
}

type Client struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
