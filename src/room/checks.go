package room

import (
	"errors"
)

// Checks if a room is reachable
func (room *Room) CheckRoomLink() error {
	if len(room.Links) == 0 {
		return errors.New("start or final room is unreachable")
	}
	return nil
}
