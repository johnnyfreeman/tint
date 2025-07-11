package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Modal represents a modal dialog component that provides an elevated surface
// It should contain Container components for structure and content
type Modal struct {
	visible  bool
	width    int
	height   int
	focused  bool
	centered bool
	x, y     int // Position (used when not centered)
}

// NewModal creates a new modal dialog
func NewModal() *Modal {
	return &Modal{
		visible:  false,
		width:    40,
		height:   10,
		focused:  false,
		centered: true,
	}
}

// Show makes the modal visible
func (m *Modal) Show() {
	m.visible = true
}

// Hide makes the modal invisible
func (m *Modal) Hide() {
	m.visible = false
}

// Toggle switches the modal visibility
func (m *Modal) Toggle() {
	m.visible = !m.visible
}

// IsVisible returns whether the modal is visible
func (m *Modal) IsVisible() bool {
	return m.visible
}

// SetCentered sets whether the modal should be centered on screen
func (m *Modal) SetCentered(centered bool) {
	m.centered = centered
}

// SetPosition sets the modal position (when not centered)
func (m *Modal) SetPosition(x, y int) {
	m.x = x
	m.y = y
	m.centered = false
}

// Draw renders the modal to the screen at the specified position
func (m *Modal) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	if !m.visible {
		return
	}

	// Modal decides its own size (doesn't use all available space)
	// Use the modal's preferred dimensions
	modalWidth := m.width
	modalHeight := m.height
	
	// Constrain to available space if needed
	if modalWidth > availableWidth {
		modalWidth = availableWidth
	}
	if modalHeight > availableHeight {
		modalHeight = availableHeight
	}

	// Calculate actual position
	actualX, actualY := x, y
	if m.centered {
		// Center within available space
		actualX = x + (availableWidth - modalWidth) / 2
		actualY = y + (availableHeight - modalHeight) / 2
	} else if x == 0 && y == 0 {
		actualX, actualY = m.x, m.y
	}

	// Clear the modal area with surface color first
	surfaceStyle := lipgloss.NewStyle().Background(theme.Palette.Surface)
	ClearArea(screen, actualX, actualY, modalWidth, modalHeight, surfaceStyle)

	// Draw block shadow AFTER clearing (offset by 1 cell)
	shadowStyle := lipgloss.NewStyle().
		Background(theme.Palette.Shadow)
		// Shadow is created by background color on spaces
	shadowOffsetX := 1
	shadowOffsetY := 1

	// Use the new DrawBlockShadow method
	screen.DrawBlockShadow(actualX, actualY, modalWidth, modalHeight, shadowStyle, shadowOffsetX, shadowOffsetY)
}

// Focus gives keyboard focus to this component
func (m *Modal) Focus() {
	m.focused = true
}

// Blur removes keyboard focus from this component
func (m *Modal) Blur() {
	m.focused = false
}

// IsFocused returns whether this component currently has focus
func (m *Modal) IsFocused() bool {
	return m.focused
}

// HandleInput processes keyboard input
func (m *Modal) HandleInput(key string) {
	switch key {
	case "esc", "enter":
		m.Hide()
	}
}

// SetSize sets the width and height of the component
func (m *Modal) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// GetSize returns the current width and height
func (m *Modal) GetSize() (width, height int) {
	return m.width, m.height
}
