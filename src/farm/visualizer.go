package farm

import (
	"fmt"
	"strings"
)

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

// Creates a visual representation of the ant farm
func (f *Farm) Visualize() string {
	minX, maxX, minY, maxY := f.getDimensions()
	width := maxX - minX + 3
	height := maxY - minY + 3

	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = "  "
		}
	}

	// Mark rooms and their names
	for _, room := range f.Rooms {
		x := room.X - minX + 1
		y := room.Y - minY + 1

		roomMarker := "[ ]"
		if room == f.Rooms[0] {
			roomMarker = "[S]" // Start room
		} else if room == f.Rooms[len(f.Rooms)-1] {
			roomMarker = "[E]" // End room
		}

		// Room name above its coordinates
		grid[y-1][x] = room.Name
		grid[y][x] = roomMarker
	}

	// Render the visualization
	var visualization strings.Builder
	visualization.WriteString("Ant Farm Visualization:\n")
	for _, row := range grid {
		visualization.WriteString(strings.Join(row, " ") + "\n")
	}

	return visualization.String()
}

// Adds ant visualization to the existing Solve method
func (f *Farm) VisualSolve() {
	// Print initial farm layout
	fmt.Println(f.Visualize())
	fmt.Println("Ant Movements:")
	f.Solve()
}
