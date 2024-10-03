package farm

import (
	"container/list"
	"fmt"
	"lem-in/src/room"
	"lem-in/src/tools"
)

func (f *Farm) BFS() [][]*room.Room {
	// Initializing first and last room
	start := f.Rooms[0]
	end := f.Rooms[len(f.Rooms)-1]

	// Store all paths and create BFS queue
	var allPaths [][]*room.Room
	queue := list.New()

	// Add the first path with the first room
	queue.PushBack([]*room.Room{start})

	for queue.Len() > 0 {
		// Dequeue the current path
		path := queue.Remove(queue.Front()).([]*room.Room)
		currentRoom := path[len(path)-1]

		// If we have reached the destination, store the path
		if currentRoom == end {
			allPaths = append(allPaths, path)
			continue
		}

		// Explore each linked room
		for _, linkedRoom := range currentRoom.Links {
			// Check if the room has already been visited in the current path
			if containsRoom(path, linkedRoom) {
				continue
			}

			// Copy the current path to extend it
			newPath := make([]*room.Room, len(path))
			copy(newPath, path)

			// Append the new room to the path and add it to the queue
			newPath = append(newPath, linkedRoom)
			queue.PushBack(newPath)
		}
	}

	return allPaths
}

// Checks if a room is already in the path
func containsRoom(path []*room.Room, room *room.Room) bool {
	for _, r := range path {
		if r == room {
			return true
		}
	}
	return false
}

func (f *Farm) SolveProblem() {
	allPaths := f.BFS()
	if len(allPaths) == 0 {
		fmt.Println("No path found")
		return
	}
	for _, p := range allPaths {
		for _, r := range p {
			fmt.Print(r.Name, " ")
		}
		tools.Ret()
	}
}
