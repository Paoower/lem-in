package entities

type AntStatus int

const (
	AntMoved AntStatus = iota
	AntNotMoved
	AntDeleted
	AntNotValid
)

type Ant struct {
	Id			int
	Path		*Path // do not include start and end
	IndexRoom	int
}

func	NewAnt(id int, path *Path) *Ant {
	return &Ant{
		Id:  id,
		Path: path,
		IndexRoom: 0,
	}
}

func (status AntStatus) Message() string {
	switch status {
	case AntMoved:
		return "The ant has moved"
	case AntNotMoved:
		return "The ant has not moved"
	case AntDeleted:
		return "The ant has been deleted"
	case AntNotValid:
		return "The ant is not valid"
	default:
		return "Unknown status"
	}
}

func	(ant *Ant) Move() AntStatus {
	var newIndex	int
	var roomsLen	int
	var rooms		[]*Room
	var newRoom		*Room

	rooms = ant.Path.Rooms
	roomsLen = len(rooms)
	if ant.IndexRoom < 0 || ant.IndexRoom >= roomsLen {
		// index out of range
		return AntNotValid
	}
	newIndex = ant.IndexRoom + 1
	if newIndex >= roomsLen {
		// new index out of range
		rooms[ant.IndexRoom].Ants = []*Ant{}
		return AntDeleted
	}
	newRoom = rooms[newIndex]
	if len(newRoom.Ants) > 0 && newIndex != roomsLen - 1 {
		// room already used and is not the end
		return AntNotMoved
	}
	rooms[ant.IndexRoom].Ants = []*Ant{}
	newRoom.Ants = append(newRoom.Ants, ant)
	ant.IndexRoom = newIndex
	return AntMoved
}
