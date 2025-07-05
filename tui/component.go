package tui

// Component is the base interface for all UI components
type Component interface {
	// Draw renders the component to the screen at the specified position
	Draw(screen *Screen, x, y int, theme *Theme)

	// Focus gives keyboard focus to this component
	Focus()

	// Blur removes keyboard focus from this component
	Blur()

	// IsFocused returns whether this component currently has focus
	IsFocused() bool

	// HandleKey processes keyboard input when focused
	// Returns true if the key was handled, false otherwise
	HandleKey(key string) bool
}

// SizedComponent is a component that has explicit dimensions
type SizedComponent interface {
	Component

	// SetSize sets the width and height of the component
	SetSize(width, height int)

	// GetSize returns the current width and height
	GetSize() (width, height int)
}

// ContainerComponent is a component that can contain other components
type ContainerComponent interface {
	SizedComponent

	// DrawWithBorder draws the component with a border and optional title
	DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string)
}
