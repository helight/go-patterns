package mediator

import (
	"fmt"
)

type User struct {
	name string
}
func (u *User)SendMessage(message string){
	GetChatRoomInstance().ShowMsg(*u, message)
}

func (u *User)GetName() string{
	return u.name
}

type ChatRoom struct {
}

func (c *ChatRoom) ShowMsg(user User, msg string) {
	fmt.Println("[ ", user.GetName() , " ] : ", msg)
}

var chatroom *ChatRoom

func GetChatRoomInstance() *ChatRoom {
	if chatroom == nil {
		chatroom = &ChatRoom{}
	}
	return chatroom
}

func main() {
}
