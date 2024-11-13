package objects

type AntStatus int

const (
	AntStatusMoved AntStatus = iota
	AntStatusNotMoved
	AntStatusDeleted
	AntStatusNotValid
)

type Ant struct {
	Id        int
	Path      *Path // do not include start
	IndexRoom int
}

func NewAnt(id int, path *Path) *Ant {
	return &Ant{
		Id:        id,
		Path:      path,
		IndexRoom: 1,
	}
}

func (status AntStatus) Message() string {
	switch status {
	case AntStatusMoved:
		return "The ant has moved"
	case AntStatusNotMoved:
		return "The ant has not moved"
	case AntStatusDeleted:
		return "The ant has been deleted"
	case AntStatusNotValid:
		return "The ant is not valid"
	default:
		return "Unknown status"
	}
}

func (ant *Ant) Move() AntStatus {
	var newIndex int
	var roomsLen int
	var rooms []*Room
	var newRoom *Room

	rooms = ant.Path.Rooms
	roomsLen = len(rooms)
	if ant.IndexRoom < 0 || ant.IndexRoom >= roomsLen {
		// index out of range
		return AntStatusNotValid
	}
	newIndex = ant.IndexRoom + 1
	if newIndex >= roomsLen {
		// new index out of range
		rooms[ant.IndexRoom].Ants = []*Ant{}
		return AntStatusDeleted
	}
	newRoom = rooms[newIndex]
	if len(newRoom.Ants) > 0 && newIndex != roomsLen-1 {
		// room already used and is not the end
		return AntStatusNotMoved
	}
	rooms[ant.IndexRoom].Ants = []*Ant{}
	newRoom.Ants = append(newRoom.Ants, ant)
	// fmt.Println(len(newRoom.Ants))
	ant.IndexRoom = newIndex
	return AntStatusMoved
}
