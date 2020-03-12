package server

//room is a kind of channel that allow users to make group conv
type room struct {
	NbChatter  uint8
	MessagesDB []string
}

//NewRoom is a room constructor
func newRoom(nbChatter uint8) room {
	return room{NbChatter: nbChatter}
}
