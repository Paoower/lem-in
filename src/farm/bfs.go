package farm

import (
	"container/list"
	"fmt"
	"lem-in/src/path"
	"lem-in/src/room"
	t "lem-in/src/tools"
	"strings"
)

func (f *Farm) BFS() {
	// Initializing first and last room
	start := f.Rooms[0]
	end := f.Rooms[len(f.Rooms)-1]

	// Store all paths and create BFS queue
	var allPaths []*path.Path
	queue := list.New()

	// Add the first path with the first room
	firstPath := path.NewPath([]*room.Room{start}, 1)
	queue.PushBack(firstPath)

	for queue.Len() > 0 {
		// Dequeue the current path
		p := queue.Remove(queue.Front()).(*path.Path)
		currentRoom := p.Route[len(p.Route)-1]

		// If we have reached the destination, store the path
		if currentRoom == end {
			allPaths = append(allPaths, p)
			// Don't continue here, allow exploration of other paths
		}

		// Explore each linked room
		for _, linkedRoom := range currentRoom.Links {
			// Check if the room has already been visited in the current path
			if containsRoom(p.Route, linkedRoom) {
				continue
			}

			// Create a new path by copying the current one and extending it
			newRoute := make([]*room.Room, len(p.Route))
			copy(newRoute, p.Route)
			newRoute = append(newRoute, linkedRoom)

			// Calculate the number of available rooms in the new path
			availableRooms := 0
			for _, room := range newRoute {
				if !room.Occupied {
					availableRooms++
				}
			}

			newPath := path.NewPath(newRoute, availableRooms)
			queue.PushBack(newPath)
		}
	}

	f.Paths = allPaths
}

// Helper function to check if a room is in a path
func containsRoom(route []*room.Room, room *room.Room) bool {
	for _, r := range route {
		if r == room {
			return true
		}
	}
	return false
}

func (f *Farm) getShortestPath(paths []*path.Path) *path.Path {
	if len(paths) == 0 {
		return nil
	}
	if len(paths) == 1 {
		return paths[0]
	}

	shortestPath := paths[0]
	shortestCost := calculatePathCost(shortestPath)

	for _, path := range paths[1:] {
		currentCost := calculatePathCost(path)
		if currentCost < shortestCost {
			shortestPath = path
			shortestCost = currentCost
		}
	}

	return shortestPath
}

func calculatePathCost(p *path.Path) int {
	unavailableRooms := len(p.Route) - p.Available
	return len(p.Route) + unavailableRooms
}

func deleteElement(slice []ant, index int) []ant {
	return append(slice[:index], slice[index+1:]...)
}

type ant struct {
	cr  *room.Room
	cri int
	cp  *path.Path
	n   int
}

func (f *Farm) ants() {
	a := make([]ant, f.Ants)
	for i := range a {
		a[i].n = i + 1
	}

	for len(a) > 0 {
		moved := false
		var turnMoves []string

		for i := 0; i < len(a); i++ {
			ant := &a[i]
			if ant.cr == nil {
				p := f.getShortestPath(f.Paths)
				if p.Route[0].Occupied {
					fmt.Println("prev ant already in the first room")
					p = f.getShortestPath(f.Paths[1:])
				}
				ant.cp = p
				fmt.Printf("ant %d choosing path: ", ant.n)
				fmt.Println(ant.cp)
				ant.cri = 0
				ant.cr = ant.cp.Route[0]
			}

			if ant.cr.Name == ant.cp.Route[len(ant.cp.Route)-1].Name {
				fr := ant.cp.Route[len(ant.cp.Route)-1]
				fr.Occupied = false
				a = deleteElement(a, i)
				i--
				continue
			}

			nr := ant.cp.Route[ant.cri+1]
			if !nr.Occupied {
				ant.cr = nr
				ant.cri++
				nr.Occupied = true
				if ant.cri > 0 {
					pr := ant.cp.Route[ant.cri-1]
					pr.Occupied = false
				}
				turnMoves = append(turnMoves, fmt.Sprintf("L%d-%s", ant.n, ant.cr.Name))
				moved = true
			}
		}

		if moved {
			fmt.Println(strings.Join(turnMoves, " "))
		}
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
		t.Ret()
	}

	f.ants()
}
