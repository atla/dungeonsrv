package game

// AvatarJoinedRoom ... asdd
type AvatarJoinedRoom struct {
	//	Avatar *Avatar
	//	Room   *Room
}

// AvatarLeftRoom ... asdd	
type AvatarLeftRoom struct {
	//Avatar *Avatar
	//Room   *Room
}

// Message ... main message container to pass data from users to server and back
type Message struct {
	//FromUser *users.User
	//ToUser   *users.User
	Data string
}

// NewMessage ... creates a new message
//func NewMessage(fromUser *users.User, data string) *Message {
func NewMessage(data string) *Message {
	return &Message{
		//FromUser: fromUser,
		Data: data,
	}
}
