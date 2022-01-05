package entity

type Response struct {
	MessageType string      `json:"message_type,omitempty"`
	Message     string      `json:"message,omitempty"`
	Err         bool        `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
