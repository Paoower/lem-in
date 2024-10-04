package farm

// import (
// 	"container/list"
// 	"fmt"
// 	"lem-in/src/path"
// 	"lem-in/src/room"
// 	"lem-in/src/tools"
// 	t "lem-in/src/tools"
// 	"time"
// )

// func (f *Farm) BFS() {
// 	// Initializing first and last room
// 	start := f.Rooms[0]
// 	end := f.Rooms[len(f.Rooms)-1]

// 	// Store all paths and create BFS queue
// 	var allPaths []*path.Path
// 	queue := list.New()

// 	// Add the first path with the first room
// 	var firstPath *path.Path = path.NewPath([]*room.Room{start}, 1)
// 	queue.PushBack(firstPath)

// 	for queue.Len() > 0 {
// 		// Dequeue the current path
// 		p := queue.Remove(queue.Front()).(*path.Path)
// 		currentRoom := p.Route[len(p.Route)-1]

// 		// If we have reached the destination, store the path
// 		if currentRoom == end {
// 			allPaths = append(allPaths, p)
// 			continue
// 		}

// 		// Explore each linked room
// 		for _, linkedRoom := range currentRoom.Links {

// 			// Check if the room has already been visited in the current path
// 			if containsRoom(p.Route, linkedRoom) {
// 				continue
// 			}

// 			// Copy the current path to extend it
// 			var newPath *path.Path = path.NewPath(p.Route, len(p.Route)+1)

// 			// Append the new room to the path and add it to the queue
// 			newPath.Route = append(newPath.Route, linkedRoom)
// 			queue.PushBack(newPath)
// 		}
// 	}

// 	f.Paths = allPaths
// }

// // Checks if a room is already in the path
// func containsRoom(path []*room.Room, room *room.Room) bool {
// 	for _, r := range path {
// 		if r == room {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (f *Farm) getShortestPath([]*path.Path) *path.Path {
// 	paths := f.Paths
// 	if len(paths) == 1 {
// 		return paths[0]
// 	}

// 	firstPath := paths[0]
// 	restOfPaths := paths[1:]

// 	shortestFromRest := f.getShortestPath(restOfPaths)

// 	if calculatePathCost(firstPath) < calculatePathCost(shortestFromRest) {
// 		return firstPath
// 	}

// 	return shortestFromRest
// }

// func calculatePathCost(p *path.Path) int {
// 	unavailableRooms := len(p.Route) - p.Available
// 	return len(p.Route) + unavailableRooms
// }

// type ant struct {
// 	cr  *room.Room
// 	pr  *room.Room
// 	cri int
// 	cp  *path.Path
// 	n   int
// }

// func deleteElement(slice []ant, index int) []ant {
// 	return append(slice[:index], slice[index+1:]...)
// }

// func (f *Farm) ants() {
// 	a := make([]ant, f.Ants)

// 	for i := range a {
// 		a[i].n = i + 1
// 	}

// 	i := 0
// 	for len(a) != 0 {
// 		ant := &a[i]

// 		if ant.cr == nil {
// 			ant.cp = f.getShortestPath(f.Paths)
// 			ant.cri = 0
// 			ant.cr = ant.cp.Route[0]
// 			fmt.Printf("ant %d starting\n", ant.n)
// 		}

// 		time.Sleep(1 * time.Second)

// 		if a[i].cr.Name == a[i].cp.Route[len(ant.cp.Route)-1].Name {
// 			fmt.Printf("removing ant %d from slice\n", ant.n)
// 			fr := a[i].cp.Route[len(ant.cp.Route)-1]
// 			fr.Occupied = false
// 			fmt.Println("freed space for room ", fr)
// 			a = deleteElement(a, i)
// 			continue
// 		}

// 		nr := a[i].cp.Route[ant.cri+1]
// 		fmt.Printf("next room for ant %d \n", ant.n)
// 		fmt.Println("\t", nr)
// 		if !nr.Occupied {
// 			fmt.Printf("moving ant %d to the next room\n", ant.n)
// 			time.Sleep(1 * time.Second)
// 			ant.cr = nr
// 			ant.cri++
// 			nr.Occupied = true

// 			pr := a[i].cp.Route[ant.cri-1]
// 			pr.Occupied = false
// 			fmt.Println("state of prev room ", pr)

// 			fmt.Println("state of next room ", nr)
// 		}

// 		if ant.cri != 0 {
// 			fmt.Printf("L%d-%s\n", ant.n, ant.cr.Name)
// 		}
// 		tools.Ret()

// 		i++
// 		if i >= len(a) {
// 			i = 0
// 		}
// 	}
// }

// func (f *Farm) SolveProblem() {
// 	f.BFS()
// 	if len(f.Paths) == 0 {
// 		fmt.Println("No path found")
// 		return
// 	}
// 	for _, p := range f.Paths {
// 		for _, r := range p.Route {
// 			fmt.Print(r.Name, " ")
// 		}
// 		t.Ret()
// 	}

// 	f.ants()
// }
