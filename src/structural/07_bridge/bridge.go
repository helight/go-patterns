package bridge

import "fmt"

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageSender interface {
	Send(text, to string)
}

type MessageSMS struct {
}

func (s *MessageSMS) Send(text,to string) {
	fmt.Println("send: %s to %s via SMS \r\n", text, to)
}

type MessageEmail struct {
}

func (s *MessageEmail) Send(text,to string) {
	fmt.Printf("send: %s to %s via Email \r\n", text, to)
}

type CommonMessage struct {
	method MessageSender
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

func NewCommonMessage(method MessageSender) *CommonMessage {
	return &CommonMessage {
		method: method,
	}
}

type UrgencyMessage struct {
	method MessageSender
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

func NewUrgencyMessage(method MessageSender) *UrgencyMessage{
	return &UrgencyMessage{
		method: method,
	}
}


func main() {

}
