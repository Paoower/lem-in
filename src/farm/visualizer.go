package farm

import (
	"fmt"
	"lem-in/src/objects"
	"os"
	"os/exec"
	"strings"
	"time"
)

var delay = (1500 * time.Millisecond)
var spacingFactor = 3

// ShowInitialState displays the initial farm state with all paths
func (f *Farm) ShowInitialState() {
	clearScreen()
	fmt.Println("\n=== Initial Ant Farm Configuration ===")
	fmt.Printf("Number of ants: %d\n", f.TotalAnts)
	fmt.Println("Room layout:")

	// Get dimensions for visualization
	minX, maxX, minY, maxY := f.getDimensions()

	width := (maxX - minX + 3) * spacingFactor
	height := (maxY - minY + 3) * spacingFactor

	// Initialize grid with more spacing
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = "  "
		}
	}

	// Draw all paths with adjusted coordinates
	for _, room := range f.Rooms {
		x1 := (room.X - minX + 1) * spacingFactor
		y1 := (room.Y - minY + 1) * spacingFactor
		for _, linkedRoom := range room.Links {
			x2 := (linkedRoom.X - minX + 1) * spacingFactor
			y2 := (linkedRoom.Y - minY + 1) * spacingFactor
			drawLine(grid, x1, y1, x2, y2)
		}
	}

	// Mark rooms with adjusted coordinates
	for _, room := range f.Rooms {
		x := (room.X - minX + 1) * spacingFactor
		y := (room.Y - minY + 1) * spacingFactor

		// Room name above coordinates
		grid[y-1][x] = room.Name

		// Special markers for start/end rooms
		if room == f.Rooms[0] {
			xEnd := (f.Rooms[len(f.Rooms)-1].X - minX + 1) * spacingFactor
			if x > xEnd {
				grid[y][x] = " [] E"
			} else {
				grid[y][x] = "E [] "
			}
		} else if room == f.Rooms[len(f.Rooms)-1] {
			xStart := (f.Rooms[0].X - minX + 1) * spacingFactor
			if x < xStart {
				grid[y][x] = "S [] "
			} else {
				grid[y][x] = " [] S"
			}
		} else {
			grid[y][x] = "[] "
		}
	}

	for _, row := range grid {
		fmt.Println(strings.Join(row, " "))
	}

	fmt.Println("\nMap Legend:")
	fmt.Println("E [] = Entrance (Start)")
	fmt.Println("S [] = Sortie/Exit (End)")
	fmt.Println("·  = Path/Tunnel")
	fmt.Println("[] = Empty Room")
	fmt.Println("[L1] = Room with Ant #1")
	fmt.Println("\nPress Enter to start ant movement simulation...")
	fmt.Scanln()
}

// Determines the maximum dimensions for the visualization
func (f *Farm) getDimensions() (int, int, int, int) {
	minX, minY := f.Rooms[0].X, f.Rooms[0].Y
	maxX, maxY := minX, minY

	for _, room := range f.Rooms {
		if room.X < minX {
			minX = room.X
		}
		if room.X > maxX {
			maxX = room.X
		}
		if room.Y < minY {
			minY = room.Y
		}
		if room.Y > maxY {
			maxY = room.Y
		}
	}

	return minX, maxX, minY, maxY
}

// Clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Function to draw a line between two points
func drawLine(grid [][]string, x1, y1, x2, y2 int) {
	dx := x2 - x1
	dy := y2 - y1
	steps := max(abs(dx), abs(dy))

	if steps == 0 {
		return
	}

	xIncrement := float64(dx) / float64(steps)
	yIncrement := float64(dy) / float64(steps)

	x := float64(x1)
	y := float64(y1)

	for i := 0; i <= steps; i++ {
		if int(y) >= 0 && int(y) < len(grid) && int(x) >= 0 && int(x) < len(grid[0]) {
			if grid[int(y)][int(x)] == "  " {
				grid[int(y)][int(x)] = "· "
			}
		}
		x += xIncrement
		y += yIncrement
	}
}

// Creates a visual representation of the ant farm with current ant positions and active paths
func (f *Farm) visualizeWithDelay(newAntsCount int) {
	minX, maxX, minY, maxY := f.getDimensions()
	width := (maxX - minX + 3) * spacingFactor
	height := (maxY - minY + 3) * spacingFactor
	grid := make([][]string, height)

	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = "  "
		}
	}

	// Create a map to track active paths
	activePaths := make(map[string]bool)

	// Find active paths based on ant positions
	for _, ant := range f.Ants {
		if ant.IndexRoom < len(ant.Path.Rooms)-1 {
			currentRoom := ant.Path.Rooms[ant.IndexRoom]
			nextRoom := ant.Path.Rooms[ant.IndexRoom+1]
			pathKey := fmt.Sprintf("%s-%s", currentRoom.Name, nextRoom.Name)
			activePaths[pathKey] = true
			reversePathKey := fmt.Sprintf("%s-%s", nextRoom.Name, currentRoom.Name)
			activePaths[reversePathKey] = true
		}
	}

	if newAntsCount >= 0 {
		solution := f.selectSolution()
		startRoom := f.Rooms[0]
		startX := (startRoom.X - minX + 1) * spacingFactor
		startY := (startRoom.Y - minY + 1) * spacingFactor

		for _, path := range solution.Paths {
			if len(path.Rooms) > 1 {
				firstRoom := path.Rooms[1] // First room after the E
				x2 := (firstRoom.X - minX + 1) * spacingFactor
				y2 := (firstRoom.Y - minY + 1) * spacingFactor
				drawLine(grid, startX, startY, x2, y2)
			}
		}
	}

	// Draw only active paths with adjusted coordinates
	for _, room := range f.Rooms {
		x1 := (room.X - minX + 1) * spacingFactor
		y1 := (room.Y - minY + 1) * spacingFactor

		for _, linkedRoom := range room.Links {
			pathKey := fmt.Sprintf("%s-%s", room.Name, linkedRoom.Name)
			if activePaths[pathKey] {
				x2 := (linkedRoom.X - minX + 1) * spacingFactor
				y2 := (linkedRoom.Y - minY + 1) * spacingFactor
				drawLine(grid, x1, y1, x2, y2)
			}
		}
	}

	// Mark rooms and their contents with adjusted coordinates
	for _, room := range f.Rooms {
		x := (room.X - minX + 1) * spacingFactor
		y := (room.Y - minY + 1) * spacingFactor
		grid[y-1][x] = room.Name

		// Add room marker and ant IDs
		var roomContent strings.Builder

		// Add start/end markers
		if room == f.Rooms[0] {
			roomContent.WriteString("E ")
		} else if room == f.Rooms[len(f.Rooms)-1] {
			roomContent.WriteString("S ")
		}

		if len(room.Ants) > 0 {
			antIDs := make([]string, 0)
			for _, ant := range room.Ants {
				antIDs = append(antIDs, fmt.Sprintf("L%d", ant.Id))
			}
			roomContent.WriteString(fmt.Sprintf("[%s]", strings.Join(antIDs, ",")))
		} else {
			roomContent.WriteString("[]")
		}

		grid[y][x] = roomContent.String()
	}

	// Render the visualization
	var visualization strings.Builder
	visualization.WriteString("\n=== Ant Movement Visualization ===\n\n")
	for _, row := range grid {
		visualization.WriteString(strings.Join(row, " ") + "\n")
	}

	clearScreen()
	fmt.Print(visualization.String())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (f *Farm) VisualSolve() {
	f.ShowInitialState()

	var solution *objects.Solution
	cpt := 0
	newAntsCount := 0

	// Premier visuel sans fourmis
	f.visualizeWithDelay(0)
	time.Sleep(delay)

	for {
		solution = f.selectSolution()
		f.moveCurrentsAnts()
		newAntsCount = f.addNewAnts(solution)
		if len(f.Ants) == 0 {
			break
		}
		cpt++
		clearScreen()
		fmt.Printf("\nMove #%d:\n", cpt)
		f.printAntsPositions()
		f.visualizeWithDelay(newAntsCount)
		time.Sleep(delay)
	}

	fmt.Printf("\nVisualization Complete!\nTotal Moves: %d\n", cpt)
}
