package farm

import (
	"container/list"
	"fmt"
	"lem-in/src/path"
	"lem-in/src/room"
	"lem-in/src/tools"
)

func (f *Farm) BFS() {
	// Initializing first and last room
	start := f.Rooms[0]
	end := f.Rooms[len(f.Rooms)-1]

	// Store all paths and create BFS queue
	var allPaths []*path.Path
	queue := list.New()

	// Add the first path with the first room
	var firstPath *path.Path = path.NewPath([]*room.Room{start}, 1)
	queue.PushBack(firstPath)

	for queue.Len() > 0 {
		// Dequeue the current path
		p := queue.Remove(queue.Front()).(*path.Path)
		currentRoom := p.Route[len(p.Route)-1]

		// If we have reached the destination, store the path
		if currentRoom == end {
			allPaths = append(allPaths, p)
			continue
		}

		// Explore each linked room
		for _, linkedRoom := range currentRoom.Links {

			// Check if the room has already been visited in the current path
			if containsRoom(p.Route, linkedRoom) {
				continue
			}

			// Copy the current path to extend it
			var newPath *path.Path = path.NewPath(p.Route, len(p.Route)+1)

			// Append the new room to the path and add it to the queue
			newPath.Route = append(newPath.Route, linkedRoom)
			queue.PushBack(newPath)
		}
	}

	f.Paths = allPaths
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

func (f *Farm) getShortestPath(paths []*path.Path) *path.Path {
	if len(paths) == 1 {
		return paths[0]
	}

	firstPath := paths[0]
	restOfPaths := paths[1:]

	shortestFromRest := f.getShortestPath(restOfPaths)

	if calculatePathCost(firstPath) < calculatePathCost(shortestFromRest) {
		return firstPath
	}

	return shortestFromRest
}

func calculatePathCost(p *path.Path) int {
	unavailableRooms := len(p.Route) - p.Available
	return len(p.Route) + unavailableRooms
}

func (f *Farm) ants {
	for ants > 0 {
		
	}
}

func (f *Farm) SolveProblem() {
	f.BFS()
	if len(f.Paths) == 0 {
		fmt.Println("No path found")
		return
	}
	for _, p := range f.Paths {
		for _, r := range p.Route {
			fmt.Print(r.Name, " ")
		}
		tools.Ret()
	}

	p := f.getShortestPath(f.Paths)
	fmt.Println(p)
}
