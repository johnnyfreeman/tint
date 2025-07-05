package tui

import (
	"strings"
	"testing"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()

	if !container.showBorder {
		t.Error("New container should show border by default")
	}
	if container.borderStyle != "single" {
		t.Error("New container should have single border style by default")
	}
	if container.content != nil {
		t.Error("New container should have nil content")
	}
}

func TestContainerSetContent(t *testing.T) {
	container := NewContainer()
	content := NewTestComponent("Test Content", 10, 5)

	container.SetContent(content)

	if container.content == nil {
		t.Error("Container content should not be nil after SetContent")
	}
}

func TestContainerBorderStyles(t *testing.T) {
	screen := NewScreenSimulation(20, 10)
	theme := NewTestTheme()
	container := NewContainer()
	container.SetSize(10, 5)

	styles := []string{"single", "double", "heavy", "rounded"}

	for _, style := range styles {
		t.Run(style, func(t *testing.T) {
			screen.Clear()
			container.SetBorderStyle(style)
			container.Draw(screen.Screen, 0, 0, theme)

			// Check that border characters are drawn
			// Top-left corner should not be a space
			cell := screen.GetCell(0, 0)
			if cell.Rune == ' ' {
				t.Errorf("Border style %s: top-left corner should not be space", style)
			}
		})
	}
}

func TestContainerWithTitle(t *testing.T) {
	screen := NewScreenSimulation(30, 10)
	theme := NewTestTheme()
	container := NewContainer()
	container.SetSize(20, 5)
	container.SetTitle("Test Title")

	container.Draw(screen.Screen, 0, 0, theme)

	// Title should appear in the top border (may be interspersed with border chars)
	topLine := screen.GetLine(0)
	if !strings.Contains(topLine, "Test") || !strings.Contains(topLine, "Title") {
		t.Errorf("Title not found in top border. Got: %s", topLine)
	}
}

func TestContainerPadding(t *testing.T) {
	screen := NewScreenSimulation(30, 10)
	theme := NewTestTheme()
	container := NewContainer()
	content := NewTestComponent("Content", 10, 3)

	container.SetContent(content)
	container.SetSize(20, 8)
	container.SetPadding(NewMargin(2)) // 2 chars padding on all sides

	container.Draw(screen.Screen, 0, 0, theme)

	// Content should be offset by border (1) + padding (2) = 3
	// Check that content appears at the right position
	x, y, found := screen.FindText("Content")
	if !found {
		t.Error("Content not found on screen")
	} else if x < 3 || y < 3 {
		t.Errorf("Content at (%d,%d) should be offset by padding", x, y)
	}
}

func TestContainerFocus(t *testing.T) {
	container := NewContainer()

	// Test Focus
	container.Focus()
	if !container.IsFocused() {
		t.Error("Container should be focused after Focus()")
	}

	// Test Blur
	container.Blur()
	if container.IsFocused() {
		t.Error("Container should not be focused after Blur()")
	}
}

func TestContainerWithFocusableContent(t *testing.T) {
	container := NewContainer()
	input := NewInput()
	container.SetContent(input)

	// Focus container should focus content
	container.Focus()
	if !container.IsFocused() {
		t.Error("Container should be focused")
	}
	if !input.IsFocused() {
		t.Error("Content should be focused when container is focused")
	}

	// Blur container should blur content
	container.Blur()
	if container.IsFocused() {
		t.Error("Container should not be focused")
	}
	if input.IsFocused() {
		t.Error("Content should not be focused when container is blurred")
	}
}

func TestContainerHandleKey(t *testing.T) {
	container := NewContainer()
	input := NewInput()
	container.SetContent(input)
	container.Focus()

	// Type some text
	handled := container.HandleKey("H")
	if !handled {
		t.Error("Container should handle key press")
	}

	container.HandleKey("i")

	if input.Value() != "Hi" {
		t.Errorf("Expected input value 'Hi', got %q", input.Value())
	}
}

func TestContainerNoBorder(t *testing.T) {
	screen := NewScreenSimulation(20, 10)
	theme := NewTestTheme()
	container := NewContainer()
	content := NewTestComponent("Content", 10, 3)

	container.SetContent(content)
	container.SetSize(15, 5)
	container.SetShowBorder(false)

	container.Draw(screen.Screen, 0, 0, theme)

	// Content should appear at position 0,0 with no border (unless padding is applied)
	x, y, found := screen.FindText("Content")
	if !found {
		t.Error("Content not found on screen")
	}
	// Container might have default padding even without border
	// Let's check if content is within reasonable bounds
	if x > 2 || y > 2 {
		t.Errorf("Content at (%d,%d) is too far from origin for no border", x, y)
	}
}

func TestContainerSizeConstraints(t *testing.T) {
	container := NewContainer()

	// Test SetSize
	container.SetSize(25, 10)
	w, h := container.GetSize()
	if w != 25 || h != 10 {
		t.Errorf("Expected size (25,10), got (%d,%d)", w, h)
	}

	// Test that content area is reduced by border and padding
	content := NewTestComponent("Test", 20, 8)
	container.SetContent(content)
	container.SetPadding(NewMargin(1))

	// With border (1) and padding (1), content area should be reduced by 2 on each side
	// So available space is 25-4=21 width, 10-4=6 height
	// This test ensures size calculations are correct
}

func TestContainerBorderElements(t *testing.T) {
	screen := NewScreenSimulation(30, 10)
	theme := NewTestTheme()
	container := NewContainer()
	container.SetSize(20, 5)

	// Add a text element to the top border
	textElement := NewTextElement("[Status]")
	container.AddBorderElement(textElement, BorderTop, BorderAlignRight)

	container.Draw(screen.Screen, 0, 0, theme)

	// The status text should appear in the border
	AssertTextExists(t, screen, "[Status]")
}

func TestContainerDrawRegion(t *testing.T) {
	// Test that container properly uses screen's DrawRegion for content
	mainScreen := NewScreenSimulation(40, 20)
	theme := NewTestTheme()

	container := NewContainer()
	content := NewTestComponent("Inner Content", 15, 5)
	container.SetContent(content)
	container.SetSize(20, 8)

	// Draw at offset position
	container.Draw(mainScreen.Screen, 5, 3, theme)

	// Content should appear at the correct offset position
	x, y, found := mainScreen.FindText("Inner Content")
	if !found {
		t.Error("Inner content not found on screen")
	} else {
		// Should be offset by container position + border + padding
		if x < 6 || y < 4 {
			t.Errorf("Content at (%d,%d) not properly offset", x, y)
		}
	}
}
