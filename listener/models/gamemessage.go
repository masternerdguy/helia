package models

import "github.com/google/uuid"

//MessageRegistry Registry of game message types
type MessageRegistry struct {
	Join int
}

//NewMessageRegistry Returns a MessageRegistry with correct enum values
func NewMessageRegistry() *MessageRegistry {
	return &MessageRegistry{
		Join: 0,
	}
}

//GameMessage Message container exchanged between client and server
type GameMessage struct {
	MessageType int    `json:"type"`
	MessageBody string `json:"body"`
}

//ClientJoinBody Body for a server join request from the client
type ClientJoinBody struct {
	SessionID uuid.UUID `json:"sid"`
}

//ServerJoinBody Body for the response to a ClientJoinBody request from the client
type ServerJoinBody struct {
	UserID uuid.UUID `json:"uid"`
}
