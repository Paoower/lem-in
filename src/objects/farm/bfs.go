package farm

import (
	"container/list"
	"fmt"
	o "lem-in/src/objects"
	t "lem-in/src/tools"
)

func (f *Farm) BFS() {
	// Initializing first and last room
	start := f.Rooms[0]
	end := f.Rooms[len(f.Rooms)-1]

	// Store all paths and create BFS queue
	var allPaths []*o.Path
	queue := list.New()

	// Add the first path with the first room
	firstPath := o.NewPath([]*o.Room{start})
	queue.PushBack(firstPath)

	for queue.Len() > 0 {
		// Dequeue the current path
		p := queue.Remove(queue.Front()).(*o.Path)
		currentRoom := p.Rooms[len(p.Rooms)-1]

		// If we have reached the destination, store the path
		if currentRoom == end {
			allPaths = append(allPaths, p)
			// Don't continue here, allow exploration of other paths
		}

		// Explore each linked room
		for _, linkedRoom := range currentRoom.Links {
			// Check if the room has already been visited in the current path
			if containsRoom(p.Rooms, linkedRoom) {
				continue
			}

			// Create a new path by copying the current one and extending it
			newRoute := make([]*o.Room, len(p.Rooms))
			copy(newRoute, p.Rooms)
			newRoute = append(newRoute, linkedRoom)

			newPath := o.NewPath(newRoute)
			queue.PushBack(newPath)
		}
	}

	f.Paths = allPaths
}

// Helper function to check if a room is in a path
func containsRoom(route []*o.Room, room *o.Room) bool {
	for _, r := range route {
		if r == room {
			return true
		}
	}
	return false
}

func (f *Farm) GetAllPaths() {
	f.BFS()
	if len(f.Paths) == 0 {
		fmt.Println("No path found")
		return
	}
	for _, p := range f.Paths {
		for _, r := range p.Rooms {
			fmt.Print(r.Name, " ")
		}
		t.Ret()
	}
}
