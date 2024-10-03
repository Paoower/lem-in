package room

// Room datatype
type Room struct {
	Name     string
	X        int
	Y        int
	Occupied bool
	Links    []*Room
}

// Room constructor
func NewRoom(name string, x int, y int) *Room {
	return &Room{
		Name:     name,
		X:        x,
		Y:        y,
		Occupied: false,
		Links:    []*Room{},
	}
}
