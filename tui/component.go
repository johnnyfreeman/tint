package tui

// Component is the base interface for all UI components
type Component interface {
	// Draw renders the component to the screen at the specified position and size
	Draw(screen *Screen, x, y, width, height int, theme *Theme)

	// HandleInput processes keyboard input
	HandleInput(key string)
}

// Focusable is an interface for components that can receive focus
type Focusable interface {
	// Focus gives keyboard focus to this component
	Focus()

	// Blur removes keyboard focus from this component
	Blur()

	// IsFocused returns whether this component currently has focus
	IsFocused() bool
}
