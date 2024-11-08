package objects

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

func	(ant *Ant) Move() AntStatus {
	var newIndex	int
	var roomsLen	int
	var rooms		[]*Room
	var newRoomAnts	[]*Ant

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
	newRoomAnts = rooms[newIndex].Ants
	if len(newRoomAnts) > 0 && newIndex != roomsLen - 1 {
		// room already used and is not the end
		return AntNotMoved
	}
	rooms[ant.IndexRoom].Ants = []*Ant{}
	newRoomAnts = append(newRoomAnts, ant)
	ant.IndexRoom = newIndex
	return AntMoved
}
