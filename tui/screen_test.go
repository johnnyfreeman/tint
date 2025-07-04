package tui

import (
	"testing"
	"github.com/charmbracelet/lipgloss"
)

func TestNewScreen(t *testing.T) {
	screen := NewScreen(10, 5)
	
	if screen.Width() != 10 {
		t.Errorf("Expected width 10, got %d", screen.Width())
	}
	if screen.Height() != 5 {
		t.Errorf("Expected height 5, got %d", screen.Height())
	}
	
	// Check all cells are initialized with spaces
	for y := 0; y < 5; y++ {
		for x := 0; x < 10; x++ {
			cell := screen.cells[y][x]
			if cell.Rune != ' ' {
				t.Errorf("Cell at (%d,%d) not initialized with space", x, y)
			}
		}
	}
}

func TestScreenDrawRune(t *testing.T) {
	screen := NewScreenSimulation(10, 5)
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	
	// Draw a simple ASCII character
	screen.DrawRune(2, 1, 'A', style)
	AssertCellRune(t, screen, 2, 1, 'A')
	
	// Test boundary conditions
	screen.DrawRune(-1, 0, 'B', style) // Should not panic
	screen.DrawRune(0, -1, 'C', style) // Should not panic
	screen.DrawRune(10, 0, 'D', style) // Should not panic
	screen.DrawRune(0, 5, 'E', style)  // Should not panic
}

func TestScreenDrawString(t *testing.T) {
	screen := NewScreenSimulation(20, 5)
	style := lipgloss.NewStyle()
	
	// Draw ASCII string
	screen.DrawString(0, 0, "Hello", style)
	AssertLineContent(t, screen, 0, "Hello")
	
	// Draw string with wide characters
	screen.DrawString(0, 1, "你好", style)
	AssertCellRune(t, screen, 0, 1, '你')
	AssertCellRune(t, screen, 2, 1, '好')
	AssertCellWidth(t, screen, 0, 1, 2)
	AssertCellWidth(t, screen, 1, 1, 0) // Continuation cell
	
	// Draw mixed content
	screen.DrawString(0, 2, "Test🎉", style)
	content := screen.GetLine(2)
	if content != "Test🎉" {
		t.Errorf("Mixed content not drawn correctly, got: %s", content)
	}
}

func TestScreenClear(t *testing.T) {
	screen := NewScreenSimulation(10, 5)
	style := lipgloss.NewStyle()
	
	// Fill screen with content
	screen.DrawString(0, 0, "Test", style)
	screen.DrawString(0, 1, "Content", style)
	
	// Clear screen
	screen.Clear()
	
	// Check all cells are spaces
	for y := 0; y < 5; y++ {
		for x := 0; x < 10; x++ {
			cell := screen.GetCell(x, y)
			if cell.Rune != ' ' {
				t.Errorf("Cell at (%d,%d) not cleared", x, y)
			}
		}
	}
}

func TestScreenClearWithStyle(t *testing.T) {
	screen := NewScreenSimulation(10, 5)
	bgStyle := lipgloss.NewStyle().Background(lipgloss.Color("#0000FF"))
	
	screen.ClearWithStyle(bgStyle)
	
	// Check all cells have the background color
	for y := 0; y < 5; y++ {
		for x := 0; x < 10; x++ {
			cell := screen.GetCell(x, y)
			if cell.Rune != ' ' {
				t.Errorf("Cell at (%d,%d) not cleared to space", x, y)
			}
			if _, isNoColor := cell.Background.(lipgloss.NoColor); isNoColor {
				t.Errorf("Cell at (%d,%d) doesn't have background color", x, y)
			}
		}
	}
}

func TestScreenDrawBox(t *testing.T) {
	screen := NewScreenSimulation(20, 10)
	style := lipgloss.NewStyle()
	
	screen.DrawBox(1, 1, 10, 5, style)
	
	// Check corners
	AssertCellRune(t, screen, 1, 1, '┌')
	AssertCellRune(t, screen, 10, 1, '┐')
	AssertCellRune(t, screen, 1, 5, '└')
	AssertCellRune(t, screen, 10, 5, '┘')
	
	// Check borders
	AssertCellRune(t, screen, 5, 1, '─') // Top
	AssertCellRune(t, screen, 5, 5, '─') // Bottom
	AssertCellRune(t, screen, 1, 3, '│') // Left
	AssertCellRune(t, screen, 10, 3, '│') // Right
}

func TestScreenDrawBoxWithTitle(t *testing.T) {
	screen := NewScreenSimulation(30, 10)
	borderStyle := lipgloss.NewStyle()
	titleStyle := lipgloss.NewStyle().Bold(true)
	
	screen.DrawBoxWithTitle(1, 1, 20, 5, "Test Title", borderStyle, titleStyle)
	
	// Check that title appears in the top border
	AssertTextExists(t, screen, "Test Title")
	
	// Check corners still exist
	AssertCellRune(t, screen, 1, 1, '┌')
	AssertCellRune(t, screen, 20, 1, '┐')
}

func TestScreenDimArea(t *testing.T) {
	screen := NewScreenSimulation(10, 5)
	style := lipgloss.NewStyle()
	
	// Draw some content
	screen.DrawString(0, 0, "Normal", style)
	screen.DrawString(0, 1, "Dimmed", style)
	
	// Dim an area
	screen.DimArea(0, 1, 6, 1)
	
	// Check that cells in the dimmed area have dim flag
	for x := 0; x < 6; x++ {
		cell := screen.GetCell(x, 1)
		if !cell.Dim {
			t.Errorf("Cell at (%d,1) should be dimmed", x)
		}
	}
	
	// Check that other cells are not dimmed
	for x := 0; x < 6; x++ {
		cell := screen.GetCell(x, 0)
		if cell.Dim {
			t.Errorf("Cell at (%d,0) should not be dimmed", x)
		}
	}
}

func TestScreenDrawRegion(t *testing.T) {
	// Create source screen with content
	src := NewScreenSimulation(10, 5)
	style := lipgloss.NewStyle()
	src.DrawString(0, 0, "Source", style)
	src.DrawString(0, 1, "Content", style)
	
	// Create destination screen
	dst := NewScreenSimulation(20, 10)
	
	// Draw region from source to destination
	dst.DrawRegion(5, 3, src.Screen, 0, 0, 7, 2)
	
	// Check content was copied
	AssertLineContent(t, dst, 3, "     Source")
	AssertLineContent(t, dst, 4, "     Content")
}

func TestScreenWideCharacterHandling(t *testing.T) {
	screen := NewScreenSimulation(10, 5)
	style := lipgloss.NewStyle()
	
	// Test drawing wide character
	screen.DrawString(0, 0, "你", style)
	AssertCellRune(t, screen, 0, 0, '你')
	AssertCellWidth(t, screen, 0, 0, 2)
	
	// Check continuation cell
	cell := screen.GetCell(1, 0)
	if !cell.IsContinuation() {
		t.Error("Expected continuation cell at (1,0)")
	}
	
	// Test overwriting continuation cell
	screen.DrawString(1, 0, "X", style)
	AssertCellRune(t, screen, 0, 0, ' ') // Wide char should be cleared
	AssertCellRune(t, screen, 1, 0, 'X')
	
	// Test overwriting start of wide character
	screen.DrawString(2, 0, "好", style)
	screen.DrawString(2, 0, "Y", style)
	AssertCellRune(t, screen, 2, 0, 'Y')
	AssertCellRune(t, screen, 3, 0, ' ') // Continuation should be cleared
}

func TestScreenBrutalistBox(t *testing.T) {
	screen := NewScreenSimulation(20, 10)
	style := lipgloss.NewStyle()
	
	screen.DrawBrutalistBox(1, 1, 10, 5, style)
	
	// Check heavy border characters
	AssertCellRune(t, screen, 1, 1, '┏')  // Top-left
	AssertCellRune(t, screen, 10, 1, '┓') // Top-right
	AssertCellRune(t, screen, 1, 5, '┗')  // Bottom-left
	AssertCellRune(t, screen, 10, 5, '┛') // Bottom-right
	AssertCellRune(t, screen, 5, 1, '━')  // Top border
	AssertCellRune(t, screen, 1, 3, '┃')  // Left border
}

func TestScreenBlockShadow(t *testing.T) {
	screen := NewScreenSimulation(20, 10)
	shadowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	
	screen.DrawBlockShadow(2, 2, 5, 3, shadowStyle, 2, 1)
	
	// Check shadow blocks
	AssertCellRune(t, screen, 4, 5, '█') // Bottom shadow
	AssertCellRune(t, screen, 7, 3, '█') // Right shadow
	AssertCellRune(t, screen, 7, 5, '█') // Corner shadow
}

func TestScreenSimulationFeatures(t *testing.T) {
	screen := NewScreenSimulation(20, 10)
	
	// Test cursor operations
	screen.SetCursor(5, 3)
	x, y := screen.GetCursor()
	if x != 5 || y != 3 {
		t.Errorf("Cursor position incorrect: got (%d,%d), want (5,3)", x, y)
	}
	
	screen.ShowCursor()
	if !screen.IsCursorVisible() {
		t.Error("Cursor should be visible")
	}
	
	screen.HideCursor()
	if screen.IsCursorVisible() {
		t.Error("Cursor should be hidden")
	}
	
	// Test snapshot
	style := lipgloss.NewStyle()
	screen.DrawString(0, 0, "Original", style)
	snapshot := screen.Snapshot()
	
	screen.DrawString(0, 0, "Modified", style)
	
	// Snapshot should still have original content
	AssertLineContent(t, snapshot, 0, "Original")
	AssertLineContent(t, screen, 0, "Modified")
	
	// Test visible bounds
	screen.Clear()
	screen.DrawString(5, 2, "Text", style)
	minX, minY, maxX, maxY := screen.GetVisibleBounds()
	
	if minX != 5 || minY != 2 || maxX != 8 || maxY != 2 {
		t.Errorf("Visible bounds incorrect: got (%d,%d)-(%d,%d), want (5,2)-(8,2)",
			minX, minY, maxX, maxY)
	}
	
	// Test occurrence counting
	screen.Clear()
	screen.DrawString(0, 0, "Hello", style)
	screen.DrawString(0, 1, "World", style)
	
	count := screen.CountOccurrences('l')
	if count != 3 {
		t.Errorf("Expected 3 occurrences of 'l', got %d", count)
	}
}