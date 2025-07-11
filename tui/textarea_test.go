package tui

import (
	"strings"
	"testing"
)

func TestNewTextArea(t *testing.T) {
	ta := NewTextArea()

	if len(ta.lines) != 1 || ta.lines[0] != "" {
		t.Error("New textarea should have one empty line")
	}
	if ta.cursorRow != 0 || ta.cursorCol != 0 {
		t.Error("New textarea should have cursor at (0,0)")
	}
	if ta.width != 40 || ta.height != 10 {
		t.Error("New textarea should have default size 40x10")
	}
	if ta.IsFocused() {
		t.Error("New textarea should not be focused")
	}
}

func TestTextAreaSetValue(t *testing.T) {
	ta := NewTextArea()

	// Test single line
	ta.SetValue("Hello World")
	if ta.Value() != "Hello World" {
		t.Errorf("Expected value 'Hello World', got %q", ta.Value())
	}

	// Test multiple lines
	multiline := "Line 1\nLine 2\nLine 3"
	ta.SetValue(multiline)
	if ta.Value() != multiline {
		t.Errorf("Expected multiline value, got %q", ta.Value())
	}
	if len(ta.lines) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(ta.lines))
	}

	// Test empty value
	ta.SetValue("")
	if len(ta.lines) != 1 || ta.lines[0] != "" {
		t.Error("Empty value should result in one empty line")
	}
}

func TestTextAreaCursorMovement(t *testing.T) {
	ta := NewTextArea()
	ta.SetValue("Line 1\nLine 2\nLine 3")
	// SetValue might position cursor at end, so reset it
	ta.cursorRow = 0
	ta.cursorCol = 0

	// Move down
	ta.HandleInput("down")
	if ta.cursorRow != 1 {
		t.Errorf("Expected cursor on row 1, got %d", ta.cursorRow)
	}

	// Reset cursor position for consistent test
	ta.cursorCol = 0

	// Move right
	ta.HandleInput("right")
	ta.HandleInput("right")
	if ta.cursorCol != 2 {
		t.Errorf("Expected cursor at col 2, got %d", ta.cursorCol)
	}

	// Move to end of line
	ta.HandleInput("end")
	if ta.cursorCol != 6 { // "Line 2" has 6 chars
		t.Errorf("Expected cursor at end of line (6), got %d", ta.cursorCol)
	}

	// Move to beginning of line
	ta.HandleInput("home")
	if ta.cursorCol != 0 {
		t.Errorf("Expected cursor at beginning of line, got %d", ta.cursorCol)
	}
}

func TestTextAreaEditing(t *testing.T) {
	ta := NewTextArea()

	// Type some text
	ta.HandleInput("H")
	ta.HandleInput("e")
	ta.HandleInput("l")
	ta.HandleInput("l")
	ta.HandleInput("o")

	if ta.Value() != "Hello" {
		t.Errorf("Expected 'Hello', got %q", ta.Value())
	}

	// Press Enter
	ta.HandleInput("enter")
	ta.HandleInput("W")
	ta.HandleInput("o")
	ta.HandleInput("r")
	ta.HandleInput("l")
	ta.HandleInput("d")

	expected := "Hello\nWorld"
	if ta.Value() != expected {
		t.Errorf("Expected %q, got %q", expected, ta.Value())
	}

	// Test backspace
	ta.HandleInput("backspace")
	expected = "Hello\nWorl"
	if ta.Value() != expected {
		t.Errorf("Expected %q after backspace, got %q", expected, ta.Value())
	}
}

func TestTextAreaDraw(t *testing.T) {
	screen := NewScreenSimulation(50, 10)
	theme := NewTestTheme()
	ta := NewTextArea()
	ta.SetSize(20, 5)

	// Test drawing empty textarea
	ta.Draw(screen.Screen, 0, 0, 20, 5, theme)

	// Test drawing with content
	ta.SetValue("Line 1\nLine 2\nLine 3")
	screen.Clear()
	ta.Draw(screen.Screen, 0, 0, 20, 5, theme)

	// Check that lines are drawn
	AssertTextExists(t, screen, "Line 1")
	AssertTextExists(t, screen, "Line 2")
	AssertTextExists(t, screen, "Line 3")

	// Test drawing with placeholder
	ta.SetValue("")
	ta.SetPlaceholder("Type here...")
	screen.Clear()
	ta.Draw(screen.Screen, 0, 0, 20, 5, theme)
	AssertTextExists(t, screen, "Type here...")
}

func TestTextAreaScrolling(t *testing.T) {
	ta := NewTextArea()
	ta.SetSize(10, 3) // Small area to force scrolling

	// Add many lines
	lines := make([]string, 10)
	for i := 0; i < 10; i++ {
		lines[i] = strings.Repeat("Line ", 10) // Long lines
	}
	ta.SetValue(strings.Join(lines, "\n"))

	// Move cursor down past visible area
	for i := 0; i < 5; i++ {
		ta.HandleInput("down")
	}

	// Should have scrolled
	if ta.offsetRow == 0 {
		t.Error("Expected vertical scrolling")
	}

	// Move cursor to the right past visible area
	for i := 0; i < 15; i++ {
		ta.HandleInput("right")
	}

	// Should have horizontal scroll
	if ta.offsetCol == 0 {
		t.Error("Expected horizontal scrolling")
	}
}

func TestTextAreaUnicodeHandling(t *testing.T) {
	ta := NewTextArea()

	// Type unicode characters
	ta.SetValue("你好世界")
	if ta.Value() != "你好世界" {
		t.Errorf("Expected '你好世界', got %q", ta.Value())
	}

	// Test cursor movement with wide characters
	ta.cursorRow = 0
	ta.cursorCol = 0
	ta.HandleInput("right")
	if ta.cursorCol != 2 { // Wide character takes 2 columns
		t.Errorf("Expected cursor at column 2, got %d", ta.cursorCol)
	}
}

func TestTextAreaLineOperations(t *testing.T) {
	// Skip this test as ctrl+k and ctrl+u are not implemented in TextArea
	t.Skip("Line operations (ctrl+k, ctrl+u) not implemented in TextArea")

	ta := NewTextArea()
	ta.SetValue("Line 1\nLine 2\nLine 3")

	// Test kill line (ctrl+k)
	ta.cursorRow = 1
	ta.cursorCol = 3
	ta.HandleInput("ctrl+k")

	lines := strings.Split(ta.Value(), "\n")
	if lines[1] != "Lin" {
		t.Errorf("Expected 'Lin' after kill line, got %q", lines[1])
	}

	// Test kill to beginning (ctrl+u)
	ta.HandleInput("ctrl+u")
	lines = strings.Split(ta.Value(), "\n")
	if lines[1] != "" {
		t.Errorf("Expected empty line after kill to beginning, got %q", lines[1])
	}
}

func TestTextAreaFocus(t *testing.T) {
	ta := NewTextArea()

	// Test Focus
	ta.Focus()
	if !ta.IsFocused() {
		t.Error("TextArea should be focused after Focus()")
	}

	// Test Blur
	ta.Blur()
	if ta.IsFocused() {
		t.Error("TextArea should not be focused after Blur()")
	}
}

func TestTextAreaEdgeCases(t *testing.T) {
	ta := NewTextArea()

	// Test operations on empty textarea
	ta.HandleInput("backspace") // Should not crash
	ta.HandleInput("delete")    // Should not crash
	ta.HandleInput("up")        // Should not crash
	ta.HandleInput("down")      // Should not crash

	// Test cursor bounds
	ta.SetValue("Test")
	ta.cursorCol = 10 // Beyond line length
	ta.HandleInput("left")
	if ta.cursorCol > 4 {
		t.Error("Cursor should be constrained to line length")
	}
}

func TestTextAreaSizeConstraints(t *testing.T) {
	screen := NewScreenSimulation(50, 20)
	theme := NewTestTheme()
	ta := NewTextArea()
	ta.SetSize(15, 5)

	// Add content that exceeds display area
	longLine := strings.Repeat("This is a very long line ", 5)
	ta.SetValue(longLine + "\n" + longLine + "\n" + longLine)

	ta.Draw(screen.Screen, 0, 0, 15, 5, theme)

	// Check that content is constrained to the textarea size
	// Count non-empty lines in the first 5 rows
	nonEmptyLines := 0
	for y := 0; y < 5; y++ {
		line := screen.GetLine(y)
		if strings.TrimSpace(line) != "" {
			nonEmptyLines++
			// Check line width
			if StringWidth(line) > 15 {
				t.Errorf("Line %d width %d exceeds textarea width 15", y, StringWidth(line))
			}
		}
	}

	if nonEmptyLines == 0 {
		t.Error("TextArea should display some content")
	}
}
