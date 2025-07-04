package tui

import (
	"strings"
)

// ScreenSimulation extends Screen with additional capabilities for testing
type ScreenSimulation struct {
	*Screen
	// Track cursor position for components that need it
	cursorX, cursorY int
	cursorVisible    bool
	// Track the last rendered output for assertions
	lastOutput string
}

// NewScreenSimulation creates a new simulated screen for testing
func NewScreenSimulation(width, height int) *ScreenSimulation {
	return &ScreenSimulation{
		Screen:        NewDefaultScreen(width, height),
		cursorX:       0,
		cursorY:       0,
		cursorVisible: false,
	}
}

// SetCursor sets the cursor position
func (s *ScreenSimulation) SetCursor(x, y int) {
	s.cursorX = x
	s.cursorY = y
}

// GetCursor returns the current cursor position
func (s *ScreenSimulation) GetCursor() (x, y int) {
	return s.cursorX, s.cursorY
}

// ShowCursor makes the cursor visible
func (s *ScreenSimulation) ShowCursor() {
	s.cursorVisible = true
}

// HideCursor makes the cursor invisible
func (s *ScreenSimulation) HideCursor() {
	s.cursorVisible = false
}

// IsCursorVisible returns whether the cursor is visible
func (s *ScreenSimulation) IsCursorVisible() bool {
	return s.cursorVisible
}

// Render renders the screen and stores the output
func (s *ScreenSimulation) Render() string {
	s.lastOutput = s.Screen.Render()
	return s.lastOutput
}

// GetLastOutput returns the last rendered output
func (s *ScreenSimulation) GetLastOutput() string {
	return s.lastOutput
}

// GetCell returns the cell at the specified position
func (s *ScreenSimulation) GetCell(x, y int) Cell {
	if x >= 0 && x < s.width && y >= 0 && y < s.height {
		return s.cells[y][x]
	}
	return NewCell(' ')
}

// GetLine returns the content of a specific line as a string
func (s *ScreenSimulation) GetLine(y int) string {
	if y < 0 || y >= s.height {
		return ""
	}
	
	var builder strings.Builder
	for x := 0; x < s.width; x++ {
		cell := s.cells[y][x]
		if !cell.IsContinuation() {
			builder.WriteRune(cell.Rune)
		}
	}
	return strings.TrimRight(builder.String(), " ")
}

// GetContent returns the visible content of the screen as a string
func (s *ScreenSimulation) GetContent() string {
	lines := make([]string, 0, s.height)
	for y := 0; y < s.height; y++ {
		line := s.GetLine(y)
		if line != "" || y == 0 { // Always include at least the first line
			lines = append(lines, line)
		} else {
			// Check if there's more content below
			hasMoreContent := false
			for checkY := y + 1; checkY < s.height; checkY++ {
				if s.GetLine(checkY) != "" {
					hasMoreContent = true
					break
				}
			}
			if hasMoreContent {
				lines = append(lines, line)
			}
		}
	}
	
	// Remove trailing empty lines
	for i := len(lines) - 1; i > 0; i-- {
		if lines[i] == "" {
			lines = lines[:i]
		} else {
			break
		}
	}
	
	return strings.Join(lines, "\n")
}

// AssertContent checks if the screen content matches the expected string
func (s *ScreenSimulation) AssertContent(expected string) bool {
	actual := s.GetContent()
	return actual == expected
}

// AssertLineContent checks if a specific line matches the expected string
func (s *ScreenSimulation) AssertLineContent(y int, expected string) bool {
	actual := s.GetLine(y)
	return actual == expected
}

// AssertCellContent checks if a specific cell contains the expected rune
func (s *ScreenSimulation) AssertCellContent(x, y int, expected rune) bool {
	cell := s.GetCell(x, y)
	return cell.Rune == expected
}

// FindText searches for text on the screen and returns its position
func (s *ScreenSimulation) FindText(text string) (x, y int, found bool) {
	for y := 0; y < s.height; y++ {
		line := s.GetLine(y)
		if idx := strings.Index(line, text); idx >= 0 {
			return idx, y, true
		}
	}
	return -1, -1, false
}

// CountOccurrences counts how many times a rune appears on the screen
func (s *ScreenSimulation) CountOccurrences(r rune) int {
	count := 0
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			if s.cells[y][x].Rune == r && !s.cells[y][x].IsContinuation() {
				count++
			}
		}
	}
	return count
}

// GetVisibleBounds returns the bounds of the visible content
func (s *ScreenSimulation) GetVisibleBounds() (minX, minY, maxX, maxY int) {
	minX, minY = s.width, s.height
	maxX, maxY = -1, -1
	
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			cell := s.cells[y][x]
			if !cell.IsDefault() && !cell.IsContinuation() {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	
	if maxX == -1 {
		// No visible content
		return 0, 0, 0, 0
	}
	
	return minX, minY, maxX, maxY
}

// Snapshot creates a copy of the current screen state
func (s *ScreenSimulation) Snapshot() *ScreenSimulation {
	snapshot := NewScreenSimulation(s.width, s.height)
	snapshot.cursorX = s.cursorX
	snapshot.cursorY = s.cursorY
	snapshot.cursorVisible = s.cursorVisible
	
	// Deep copy cells
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			snapshot.cells[y][x] = s.cells[y][x]
		}
	}
	
	return snapshot
}

// Diff compares this screen with another and returns the differences
func (s *ScreenSimulation) Diff(other *ScreenSimulation) []string {
	var differences []string
	
	if s.width != other.width || s.height != other.height {
		differences = append(differences, 
			"Screen dimensions differ")
		return differences
	}
	
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			if s.cells[y][x] != other.cells[y][x] {
				differences = append(differences,
					"Cell differs at ("+string(rune(x))+","+string(rune(y))+")")
			}
		}
	}
	
	return differences
}