package tui

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
	"testing"
)

// NewTestTheme creates a simple theme for testing
func NewTestTheme() *Theme {
	// Use a predefined theme for testing
	theme := GetTheme("monochrome")
	return &theme
}

// AssertScreenContent is a helper function to assert screen content in tests
func AssertScreenContent(t *testing.T, screen *ScreenSimulation, expected string) {
	t.Helper()
	actual := screen.GetContent()
	if actual != expected {
		t.Errorf("Screen content mismatch\nExpected:\n%s\n\nActual:\n%s", expected, actual)
	}
}

// AssertLineContent is a helper function to assert a specific line's content
func AssertLineContent(t *testing.T, screen *ScreenSimulation, lineNum int, expected string) {
	t.Helper()
	actual := screen.GetLine(lineNum)
	if actual != expected {
		t.Errorf("Line %d content mismatch\nExpected: %q\nActual: %q", lineNum, expected, actual)
	}
}

// AssertCellRune is a helper function to assert a cell's rune value
func AssertCellRune(t *testing.T, screen *ScreenSimulation, x, y int, expected rune) {
	t.Helper()
	cell := screen.GetCell(x, y)
	if cell.Rune != expected {
		t.Errorf("Cell at (%d,%d) rune mismatch\nExpected: %q\nActual: %q", x, y, expected, cell.Rune)
	}
}

// AssertCellWidth is a helper function to assert a cell's width
func AssertCellWidth(t *testing.T, screen *ScreenSimulation, x, y int, expected int) {
	t.Helper()
	cell := screen.GetCell(x, y)
	if cell.Width != expected {
		t.Errorf("Cell at (%d,%d) width mismatch\nExpected: %d\nActual: %d", x, y, expected, cell.Width)
	}
}

// AssertCellStyle checks if a cell has specific style attributes
func AssertCellStyle(t *testing.T, screen *ScreenSimulation, x, y int, bold, italic, underline, dim bool) {
	t.Helper()
	cell := screen.GetCell(x, y)

	if cell.Bold != bold {
		t.Errorf("Cell at (%d,%d) bold mismatch\nExpected: %v\nActual: %v", x, y, bold, cell.Bold)
	}
	if cell.Italic != italic {
		t.Errorf("Cell at (%d,%d) italic mismatch\nExpected: %v\nActual: %v", x, y, italic, cell.Italic)
	}
	if cell.Underline != underline {
		t.Errorf("Cell at (%d,%d) underline mismatch\nExpected: %v\nActual: %v", x, y, underline, cell.Underline)
	}
	if cell.Dim != dim {
		t.Errorf("Cell at (%d,%d) dim mismatch\nExpected: %v\nActual: %v", x, y, dim, cell.Dim)
	}
}

// AssertTextExists checks if text exists anywhere on the screen
func AssertTextExists(t *testing.T, screen *ScreenSimulation, text string) {
	t.Helper()
	_, _, found := screen.FindText(text)
	if !found {
		t.Errorf("Text %q not found on screen\nScreen content:\n%s", text, screen.GetContent())
	}
}

// AssertTextNotExists checks if text does not exist on the screen
func AssertTextNotExists(t *testing.T, screen *ScreenSimulation, text string) {
	t.Helper()
	_, _, found := screen.FindText(text)
	if found {
		t.Errorf("Text %q found on screen but should not exist\nScreen content:\n%s", text, screen.GetContent())
	}
}

// AssertCursorPosition checks if cursor is at expected position
func AssertCursorPosition(t *testing.T, screen *ScreenSimulation, expectedX, expectedY int) {
	t.Helper()
	actualX, actualY := screen.GetCursor()
	if actualX != expectedX || actualY != expectedY {
		t.Errorf("Cursor position mismatch\nExpected: (%d,%d)\nActual: (%d,%d)",
			expectedX, expectedY, actualX, actualY)
	}
}

// MockKeyEvent simulates a key press event for testing
func MockKeyEvent(key string) string {
	// This returns the key string as components expect it
	// Can be expanded to handle special keys
	return key
}

// CreateTestComponent creates a simple test component that implements Component interface
type TestComponent struct {
	focused bool
	width   int
	height  int
	content string
}

func NewTestComponent(content string, width, height int) *TestComponent {
	return &TestComponent{
		content: content,
		width:   width,
		height:  height,
	}
}

func (tc *TestComponent) Draw(screen *Screen, x, y, width, height int, theme *Theme) {
	// Draw content with optional focus indicator
	var style lipgloss.Style
	if tc.focused {
		style = lipgloss.NewStyle().Foreground(theme.Palette.Primary)
	} else {
		style = lipgloss.NewStyle().Foreground(theme.Palette.Text)
	}

	// Handle multi-line content
	lines := strings.Split(tc.content, "\n")
	for i, line := range lines {
		if i >= height {
			break // Don't draw beyond available height
		}
		screen.DrawString(x, y+i, line, style)
	}
}

func (tc *TestComponent) Focus() {
	tc.focused = true
}

func (tc *TestComponent) Blur() {
	tc.focused = false
}

func (tc *TestComponent) IsFocused() bool {
	return tc.focused
}

func (tc *TestComponent) HandleInput(key string) {
	// Simple key handler for testing
	// This method doesn't need to return a value in the new interface
}

func (tc *TestComponent) SetSize(width, height int) {
	tc.width = width
	tc.height = height
}

func (tc *TestComponent) GetSize() (width, height int) {
	return tc.width, tc.height
}
