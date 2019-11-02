package game

import (
	"sync"
)

// Game ... default entity to structure rooms
type Game struct {
	running bool

	/*OnMessageReceived chan *Message
	OnUserJoined      chan *UserJoined
	OnUserQuit        chan *UserQuit

	OnAvatarJoinedRoom chan *AvatarJoinedRoom
	OnAvatarLeftRoom   chan *AvatarLeftRoom

	Receivers []Receiver
	*/
	CommandProcessor *CommandProcessor

	//	Avatars map[string]*Avatar
}

var instance *Game
var once sync.Once

// GetInstance ... returns the usermanager instance
func GetInstance() *Game {
	once.Do(func() {
		instance = &Game{
			running:          false,
			CommandProcessor: NewCommandProcessor(),
		}
		instance.run()
	})
	return instance
}

// Load ... Init game from id
func (game *Game) Init(id string) {
	//TODO: remove from init and add into a specific loader?

}

// main game loop
func (game *Game) run() {

	go func() {
		for {
			/*
				select {
				case userJoined := <-game.OnUserJoined:

					game.loadAvatarForUser(userJoined.User)

					//TODO: remove? only inform avatars in same room?
				//	game.sendMessage(NewMessage(game.SystemUser, userJoined.User.ID+" joined."))

				case userQuit := <-game.OnUserQuit:

					_ = userQuit

				//	game.sendMessage(NewMessage(game.SystemUser, userQuit.User.ID+" quitted."))

				case avatarJoinedRoom := <-game.OnAvatarJoinedRoom:

					//	var user = avatarJoinedRoom.Avatar.CurrentUser
					//	var msg = NewMessage(nil, "=== "+avatarJoinedRoom.Room.Title+" ===\n"+avatarJoinedRoom.Room.Description)

					//	msg.ToUser = user

					game.sendMessage(avatarJoinedRoom)

				case message := <-game.OnMessageReceived:

					// only broadcast if commandprocessor didnt process it
					if !game.CommandProcessor.Process(game, message) {
						game.sendMessage(message)
					}
				}
			*/
		}
	}()
}
