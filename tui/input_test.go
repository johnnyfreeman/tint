package tui

import (
	"testing"
)

func TestNewInput(t *testing.T) {
	input := NewInput()

	if input.Value() != "" {
		t.Error("New input should have empty value")
	}
	if input.cursor != 0 {
		t.Error("New input should have cursor at position 0")
	}
	if input.width != 20 {
		t.Error("New input should have default width of 20")
	}
	if input.IsFocused() {
		t.Error("New input should not be focused")
	}
}

func TestInputSetValue(t *testing.T) {
	input := NewInput()

	// Test SetValue
	input.SetValue("Hello World")
	if input.Value() != "Hello World" {
		t.Errorf("Expected value 'Hello World', got %q", input.Value())
	}
	if input.cursor != 11 {
		t.Errorf("Expected cursor at position 11, got %d", input.cursor)
	}

	// Test SetValueWithCursorAtStart
	input.SetValueWithCursorAtStart("Test")
	if input.Value() != "Test" {
		t.Errorf("Expected value 'Test', got %q", input.Value())
	}
	if input.cursor != 0 {
		t.Errorf("Expected cursor at position 0, got %d", input.cursor)
	}
}

func TestInputFocus(t *testing.T) {
	input := NewInput()

	// Test Focus
	input.Focus()
	if !input.IsFocused() {
		t.Error("Input should be focused after Focus()")
	}

	// Test Blur
	input.Blur()
	if input.IsFocused() {
		t.Error("Input should not be focused after Blur()")
	}
}

func TestInputHandleInput(t *testing.T) {
	tests := []struct {
		name           string
		initialValue   string
		initialCursor  int
		key            string
		expectedValue  string
		expectedCursor int
	}{
		// Character insertion
		{"Insert at start", "", 0, "H", "H", 1},
		{"Insert in middle", "Hllo", 1, "e", "Hello", 2},
		{"Insert at end", "Hello", 5, "!", "Hello!", 6},

		// Movement
		{"Move left", "Hello", 3, "left", "Hello", 2},
		{"Move right", "Hello", 2, "right", "Hello", 3},
		{"Move to home", "Hello", 3, "home", "Hello", 0},
		{"Move to end", "Hello", 2, "end", "Hello", 5},

		// Deletion
		{"Backspace", "Hello", 3, "backspace", "Helo", 2},
		{"Delete", "Hello", 2, "delete", "Helo", 2},
		{"Kill to end", "Hello World", 5, "ctrl+k", "Hello", 5},
		{"Kill to beginning", "Hello World", 6, "ctrl+u", "World", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := NewInput()
			input.SetValue(tt.initialValue)
			input.cursor = tt.initialCursor

			input.HandleInput(tt.key)

			if input.Value() != tt.expectedValue {
				t.Errorf("Expected value %q, got %q", tt.expectedValue, input.Value())
			}
			if input.cursor != tt.expectedCursor {
				t.Errorf("Expected cursor at %d, got %d", tt.expectedCursor, input.cursor)
			}
		})
	}
}

func TestInputDraw(t *testing.T) {
	screen := NewScreenSimulation(30, 5)
	theme := NewTestTheme()
	input := NewInput()

	// Test drawing empty input
	input.Draw(screen.Screen, 0, 0, theme)
	// Should be blank

	// Test drawing with value
	input.SetValue("Hello World")
	input.Draw(screen.Screen, 0, 1, theme)
	AssertTextExists(t, screen, "Hello World")

	// Test drawing with placeholder
	input.SetValue("")
	input.SetPlaceholder("Type here...")
	input.Draw(screen.Screen, 0, 2, theme)
	AssertTextExists(t, screen, "Type here...")

	// Test drawing focused input
	input.Focus()
	input.SetValue("Focused")
	input.Draw(screen.Screen, 0, 3, theme)
	AssertTextExists(t, screen, "Focused")
}

func TestInputUnicodeHandling(t *testing.T) {
	input := NewInput()

	// Test with Chinese characters
	input.SetValue("你好")
	if input.Value() != "你好" {
		t.Errorf("Expected value '你好', got %q", input.Value())
	}
	if input.cursor != 4 { // 2 chars * 2 width each
		t.Errorf("Expected cursor at position 4, got %d", input.cursor)
	}

	// Test cursor movement with wide characters
	input.HandleInput("left")
	if input.cursor != 2 {
		t.Errorf("Expected cursor at position 2 after left, got %d", input.cursor)
	}

	// Test deletion with wide characters
	input.HandleInput("backspace")
	if input.Value() != "好" {
		t.Errorf("Expected value '好' after backspace, got %q", input.Value())
	}
}

func TestInputScrolling(t *testing.T) {
	input := NewInput()
	input.SetWidth(10)

	// Add text longer than width
	longText := "This is a very long text that needs scrolling"
	input.SetValue(longText)

	// Cursor should be at the end
	expectedCursor := StringWidth(longText)
	if input.cursor != expectedCursor {
		t.Errorf("Expected cursor at %d, got %d", expectedCursor, input.cursor)
	}

	// Offset should have scrolled
	if input.offset == 0 {
		t.Error("Expected offset to be greater than 0 for scrolling")
	}

	// Move to beginning
	input.HandleInput("home")
	if input.cursor != 0 {
		t.Error("Cursor should be at 0 after home")
	}
	if input.offset != 0 {
		t.Error("Offset should be 0 when cursor is at beginning")
	}
}

func TestInputEdgeCases(t *testing.T) {
	input := NewInput()

	// Test empty input operations
	input.HandleInput("backspace") // Should not crash
	input.HandleInput("delete")    // Should not crash
	input.HandleInput("left")      // Should not crash

	// Test operations at boundaries
	input.SetValue("Test")
	input.cursor = 0
	input.HandleInput("left") // Should not go negative
	if input.cursor != 0 {
		t.Error("Cursor should stay at 0 when moving left at beginning")
	}

	input.cursor = 4
	input.HandleInput("right") // Should not go past end
	if input.cursor != 4 {
		t.Error("Cursor should stay at end when moving right at end")
	}
}

func TestInputWidthConstraints(t *testing.T) {
	screen := NewScreenSimulation(30, 5)
	theme := NewTestTheme()
	input := NewInput()
	input.SetWidth(5)

	// Test that long text is truncated in display
	input.SetValue("This is too long")
	input.Draw(screen.Screen, 0, 0, theme)

	// The visible area should be limited to width
	line := screen.GetLine(0)
	if StringWidth(line) > 5 {
		t.Errorf("Displayed text width %d exceeds input width 5", StringWidth(line))
	}
}
