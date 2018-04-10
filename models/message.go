package models

// Message permite estructurar los mensajes de retornos de parte del servidor al usuario
type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
