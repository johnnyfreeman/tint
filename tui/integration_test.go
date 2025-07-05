package tui

import (
	"testing"
)

func TestContainerWithInput(t *testing.T) {
	// Test a container with an input field
	screen := NewScreenSimulation(40, 10)
	theme := NewTestTheme()

	container := NewContainer()
	container.SetSize(30, 5)
	container.SetTitle("User Input")

	input := NewInput()
	input.SetPlaceholder("Enter name...")
	input.SetWidth(20)
	container.SetContent(input)

	// Draw and verify structure
	container.Draw(screen.Screen, 0, 0, theme)

	// Check that container border is drawn
	AssertTextExists(t, screen, "User Input")
	AssertTextExists(t, screen, "Enter name...")

	// Focus container (should focus input)
	container.Focus()

	// Type some text
	container.HandleKey("J")
	container.HandleKey("o")
	container.HandleKey("h")
	container.HandleKey("n")

	// Redraw
	screen.Clear()
	container.Draw(screen.Screen, 0, 0, theme)

	// Check input value is displayed
	AssertTextExists(t, screen, "John")
}

func TestNestedContainers(t *testing.T) {
	// Test containers within containers
	screen := NewScreenSimulation(50, 20)
	theme := NewTestTheme()

	// Outer container
	outer := NewContainer()
	outer.SetSize(40, 15)
	outer.SetTitle("Outer Container")

	// Inner container with content
	inner := NewContainer()
	inner.SetSize(30, 8)
	inner.SetTitle("Inner Container")
	inner.SetPadding(NewMargin(1))

	// Content
	input := NewInput()
	input.SetValue("Nested content")
	inner.SetContent(input)

	outer.SetContent(inner)

	// Draw
	outer.Draw(screen.Screen, 0, 0, theme)

	// Verify nested structure
	AssertTextExists(t, screen, "Outer Container")
	AssertTextExists(t, screen, "Inner Container")
	AssertTextExists(t, screen, "Nested content")
}

func TestFocusChain(t *testing.T) {
	// Test focus propagation through component hierarchy
	container1 := NewContainer()
	input1 := NewInput()
	container1.SetContent(input1)

	container2 := NewContainer()
	textarea := NewTextArea()
	container2.SetContent(textarea)

	// Focus first container
	container1.Focus()
	if !container1.IsFocused() || !input1.IsFocused() {
		t.Error("Focus should propagate to container and its content")
	}

	// Blur first, focus second
	container1.Blur()
	container2.Focus()

	if container1.IsFocused() || input1.IsFocused() {
		t.Error("First container and input should not be focused")
	}
	if !container2.IsFocused() || !textarea.IsFocused() {
		t.Error("Second container and textarea should be focused")
	}
}

func TestTableInContainer(t *testing.T) {
	screen := NewScreenSimulation(60, 20)
	theme := NewTestTheme()

	container := NewContainer()
	container.SetSize(50, 15)
	container.SetTitle("Data Table")

	table := NewTable()
	table.SetColumns([]TableColumn{
		{Title: "ID", Width: 5},
		{Title: "Name", Width: 15},
		{Title: "Status", Width: 10},
	})
	table.SetRows([]TableRow{
		{"1", "Alice", "Active"},
		{"2", "Bob", "Inactive"},
	})

	container.SetContent(table)
	container.Focus()

	// Draw
	container.Draw(screen.Screen, 0, 0, theme)

	// Verify table is drawn within container
	AssertTextExists(t, screen, "Data Table")
	AssertTextExists(t, screen, "ID")
	AssertTextExists(t, screen, "Name")
	AssertTextExists(t, screen, "Status")
	AssertTextExists(t, screen, "Alice")

	// Navigate in table
	container.HandleKey("down")
	container.HandleKey("right")

	// Edit cell
	container.HandleKey("enter")
	container.HandleKey("backspace")
	container.HandleKey("backspace")
	container.HandleKey("backspace")
	container.HandleKey("J")
	container.HandleKey("o")
	container.HandleKey("e")
	container.HandleKey("enter")

	// Verify edit
	if table.GetValue(1, 1) != "Joe" {
		t.Errorf("Expected 'Joe', got %q", table.GetValue(1, 1))
	}
}

func TestScreenComposition(t *testing.T) {
	// Test complex screen layout with multiple components
	mainScreen := NewScreenSimulation(80, 30)
	theme := NewTestTheme()

	// Header - using a simple component instead of borderless container
	headerText := NewTestComponent("=== Application Header ===", 78, 1)

	// Main content area with input
	content := NewContainer()
	content.SetSize(50, 20)
	content.SetTitle("Main Content")

	input := NewInput()
	input.SetValue("Main content area")
	content.SetContent(input)

	// Sidebar
	sidebar := NewContainer()
	sidebar.SetSize(25, 20)
	sidebar.SetTitle("Sidebar")

	sidebarText := NewTestComponent("Options\nSettings\nHelp", 20, 10)
	sidebar.SetContent(sidebarText)

	// Draw all components
	headerText.Draw(mainScreen.Screen, 1, 1, theme)
	content.Draw(mainScreen.Screen, 1, 5, theme)
	sidebar.Draw(mainScreen.Screen, 53, 5, theme)

	// Verify layout
	AssertTextExists(t, mainScreen, "=== Application Header ===")
	AssertTextExists(t, mainScreen, "Main Content")
	AssertTextExists(t, mainScreen, "Main content area")
	AssertTextExists(t, mainScreen, "Sidebar")
	AssertTextExists(t, mainScreen, "Options")
}

func TestUnicodeIntegration(t *testing.T) {
	// Test unicode handling across multiple components
	screen := NewScreenSimulation(50, 20)
	theme := NewTestTheme()

	container := NewContainer()
	container.SetSize(40, 10)
	container.SetTitle("日本語")

	ta := NewTextArea()
	ta.SetValue("こんにちは\n世界")
	container.SetContent(ta)

	container.Draw(screen.Screen, 0, 0, theme)

	// Verify unicode content is properly rendered
	AssertTextExists(t, screen, "日本語")
	AssertTextExists(t, screen, "こんにちは")
	AssertTextExists(t, screen, "世界")
}
