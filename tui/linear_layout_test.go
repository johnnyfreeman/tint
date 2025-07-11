package tui

import (
	"testing"
)

func TestLinearLayout(t *testing.T) {
	t.Run("Horizontal layout creation", func(t *testing.T) {
		layout := HBox()
		if layout.config.Direction != Horizontal {
			t.Errorf("Expected Horizontal direction, got %v", layout.config.Direction)
		}
	})

	t.Run("Vertical layout creation", func(t *testing.T) {
		layout := VBox()
		if layout.config.Direction != Vertical {
			t.Errorf("Expected Vertical direction, got %v", layout.config.Direction)
		}
	})

	t.Run("Add components", func(t *testing.T) {
		layout := NewLinearLayout(Horizontal)

		// Add fixed component
		layout.AddFixed(&mockComponent{}, 20)
		if layout.Count() != 1 {
			t.Errorf("Expected 1 component, got %d", layout.Count())
		}

		// Add flex component
		layout.AddFlex(&mockComponent{}, 1.5)
		if layout.Count() != 2 {
			t.Errorf("Expected 2 components, got %d", layout.Count())
		}

		// Add percentage component
		layout.AddPercentage(&mockComponent{}, 0.25)
		if layout.Count() != 3 {
			t.Errorf("Expected 3 components, got %d", layout.Count())
		}
	})

	t.Run("Clear layout", func(t *testing.T) {
		layout := NewLinearLayout(Vertical)
		layout.AddFixed(&mockComponent{}, 10)
		layout.AddFixed(&mockComponent{}, 20)

		if layout.Count() != 2 {
			t.Errorf("Expected 2 components before clear, got %d", layout.Count())
		}

		layout.Clear()
		if layout.Count() != 0 {
			t.Errorf("Expected 0 components after clear, got %d", layout.Count())
		}
	})

	t.Run("Get item", func(t *testing.T) {
		layout := NewLinearLayout(Horizontal)
		comp1 := &mockComponent{id: "comp1"}
		comp2 := &mockComponent{id: "comp2"}

		layout.AddFixed(comp1, 10)
		layout.AddFixed(comp2, 20)

		item := layout.GetItem(0)
		if item == nil {
			t.Error("Expected to get first component, got nil")
		}
		if mc, ok := item.(*mockComponent); ok {
			if mc.id != "comp1" {
				t.Errorf("Expected comp1, got %s", mc.id)
			}
		}

		item = layout.GetItem(1)
		if item == nil {
			t.Error("Expected to get second component, got nil")
		}
		if mc, ok := item.(*mockComponent); ok {
			if mc.id != "comp2" {
				t.Errorf("Expected comp2, got %s", mc.id)
			}
		}

		// Test out of bounds
		item = layout.GetItem(2)
		if item != nil {
			t.Error("Expected nil for out of bounds index")
		}

		item = layout.GetItem(-1)
		if item != nil {
			t.Error("Expected nil for negative index")
		}
	})

	t.Run("Set properties", func(t *testing.T) {
		layout := NewLinearLayout(Horizontal)

		// Test spacing
		layout.SetSpacing(5)
		if layout.config.Spacing != 5 {
			t.Errorf("Expected spacing 5, got %d", layout.config.Spacing)
		}

		// Test padding
		padding := NewMargin(10)
		layout.SetPadding(padding)
		if layout.config.Padding.Top != 10 {
			t.Errorf("Expected padding 10, got %d", layout.config.Padding.Top)
		}

		// Test alignment
		layout.SetAlignment(AlignCenter)
		if layout.config.Alignment != AlignCenter {
			t.Errorf("Expected AlignCenter, got %v", layout.config.Alignment)
		}
	})

	t.Run("Size management", func(t *testing.T) {
		layout := NewLinearLayout(Vertical)
		layout.SetSize(100, 50)

		width, height := layout.GetSize()
		if width != 100 || height != 50 {
			t.Errorf("Expected size 100x50, got %dx%d", width, height)
		}
	})
}

// Mock component for testing
type mockComponent struct {
	id      string
	width   int
	height  int
	drawX   int
	drawY   int
	focused bool
}

func (m *mockComponent) Draw(screen *Screen, x, y, width, height int, theme *Theme) {
	m.drawX = x
	m.drawY = y
	m.width = width
	m.height = height
}

func (m *mockComponent) Focus() {
	m.focused = true
}

func (m *mockComponent) Blur() {
	m.focused = false
}

func (m *mockComponent) IsFocused() bool {
	return m.focused
}

func (m *mockComponent) HandleInput(key string) {
	// No-op for mock component
}

func (m *mockComponent) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *mockComponent) DrawWithBounds(screen *Screen, x, y, width, height int, theme *Theme) {
	m.drawX = x
	m.drawY = y
	m.width = width
	m.height = height
}
