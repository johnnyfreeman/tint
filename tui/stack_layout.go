package tui

// StackAlignment represents how items are aligned in a stack
type StackAlignment struct {
	Horizontal Alignment
	Vertical   Alignment
}

// StackItem represents an item in a stack layout
type StackItem struct {
	Component Component
	X         ConstraintSet // X position constraint
	Y         ConstraintSet // Y position constraint
	Width     ConstraintSet // Width constraint
	Height    ConstraintSet // Height constraint
	Alignment StackAlignment
}

// Stack arranges components in layers on top of each other
type Stack struct {
	items  []StackItem
	width  int
	height int
}

// NewStack creates a new stack layout
func NewStack() *Stack {
	return &Stack{
		items: []StackItem{},
	}
}

// Add adds a component to the stack with full constraints
func (s *Stack) Add(component Component, x, y, width, height ConstraintSet) {
	s.items = append(s.items, StackItem{
		Component: component,
		X:         x,
		Y:         y,
		Width:     width,
		Height:    height,
		Alignment: StackAlignment{
			Horizontal: AlignStart,
			Vertical:   AlignStart,
		},
	})
}

// AddCentered adds a component centered in the stack
func (s *Stack) AddCentered(component Component, width, height ConstraintSet) {
	s.items = append(s.items, StackItem{
		Component: component,
		X:         NewConstraintSet(NewPercentage(0.5)),
		Y:         NewConstraintSet(NewPercentage(0.5)),
		Width:     width,
		Height:    height,
		Alignment: StackAlignment{
			Horizontal: AlignCenter,
			Vertical:   AlignCenter,
		},
	})
}

// AddAnchored adds a component anchored to a specific position
func (s *Stack) AddAnchored(component Component, horizontal, vertical Alignment, width, height ConstraintSet) {
	var x, y ConstraintSet

	switch horizontal {
	case AlignStart:
		x = NewConstraintSet(NewLength(0))
	case AlignCenter:
		x = NewConstraintSet(NewPercentage(0.5))
	case AlignEnd:
		x = NewConstraintSet(NewPercentage(1.0))
	default:
		x = NewConstraintSet(NewLength(0))
	}

	switch vertical {
	case AlignStart:
		y = NewConstraintSet(NewLength(0))
	case AlignCenter:
		y = NewConstraintSet(NewPercentage(0.5))
	case AlignEnd:
		y = NewConstraintSet(NewPercentage(1.0))
	default:
		y = NewConstraintSet(NewLength(0))
	}

	s.items = append(s.items, StackItem{
		Component: component,
		X:         x,
		Y:         y,
		Width:     width,
		Height:    height,
		Alignment: StackAlignment{
			Horizontal: horizontal,
			Vertical:   vertical,
		},
	})
}

// AddFull adds a component that fills the entire stack
func (s *Stack) AddFull(component Component) {
	s.Add(
		component,
		NewConstraintSet(NewLength(0)),
		NewConstraintSet(NewLength(0)),
		NewConstraintSet(NewPercentage(1.0)),
		NewConstraintSet(NewPercentage(1.0)),
	)
}

// SetSize sets the width and height of the stack
func (s *Stack) SetSize(width, height int) {
	s.width = width
	s.height = height
}

// GetSize returns the current width and height
func (s *Stack) GetSize() (width, height int) {
	return s.width, s.height
}

// Draw renders the stack layout to the screen
func (s *Stack) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	// Stack decides to use available space for layout
	stackWidth := availableWidth
	stackHeight := availableHeight
	
	s.drawWithBounds(screen, x, y, stackWidth, stackHeight, theme)
}

// drawWithBounds draws the stack with specific bounds
func (s *Stack) drawWithBounds(screen *Screen, x, y, width, height int, theme *Theme) {
	// Draw items in order (first item is bottom, last is top)
	for _, item := range s.items {
		// Calculate position and size
		itemX := x + item.X.Calculate(width, 1.0)
		itemY := y + item.Y.Calculate(height, 1.0)
		itemWidth := item.Width.Calculate(width, 1.0)
		itemHeight := item.Height.Calculate(height, 1.0)

		// Apply alignment adjustments
		switch item.Alignment.Horizontal {
		case AlignCenter:
			itemX -= itemWidth / 2
		case AlignEnd:
			itemX -= itemWidth
		}

		switch item.Alignment.Vertical {
		case AlignCenter:
			itemY -= itemHeight / 2
		case AlignEnd:
			itemY -= itemHeight
		}

		// Ensure item stays within bounds
		if itemX < x {
			itemX = x
		}
		if itemY < y {
			itemY = y
		}
		if itemX+itemWidth > x+width {
			itemWidth = x + width - itemX
		}
		if itemY+itemHeight > y+height {
			itemHeight = y + height - itemY
		}

		// Skip if item is outside bounds or has no size
		if itemWidth <= 0 || itemHeight <= 0 {
			continue
		}

		// Draw the component
		item.Component.Draw(screen, itemX, itemY, itemWidth, itemHeight, theme)
	}
}

// Clear removes all items from the stack
func (s *Stack) Clear() {
	s.items = []StackItem{}
}

// HandleInput processes keyboard input
func (s *Stack) HandleInput(key string) {
	// Stack doesn't handle input itself
}

// Count returns the number of items in the stack
func (s *Stack) Count() int {
	return len(s.items)
}

// GetItem returns the component at the specified index
func (s *Stack) GetItem(index int) Component {
	if index >= 0 && index < len(s.items) {
		return s.items[index].Component
	}
	return nil
}

// BringToFront moves an item to the top of the stack
func (s *Stack) BringToFront(index int) {
	if index > 0 && index < len(s.items) {
		item := s.items[index]
		s.items = append(s.items[:index], s.items[index+1:]...)
		s.items = append(s.items, item)
	}
}

// SendToBack moves an item to the bottom of the stack
func (s *Stack) SendToBack(index int) {
	if index >= 0 && index < len(s.items)-1 {
		item := s.items[index]
		s.items = append(s.items[:index], s.items[index+1:]...)
		s.items = append([]StackItem{item}, s.items...)
	}
}

// Component interface methods

// Focus gives keyboard focus to this component
func (s *Stack) Focus() {
	// Focus the top-most focusable item
	for i := len(s.items) - 1; i >= 0; i-- {
		if focusable, ok := s.items[i].Component.(Focusable); ok {
			focusable.Focus()
			break
		}
	}
}

// Blur removes keyboard focus from this component
func (s *Stack) Blur() {
	// Blur all items
	for _, item := range s.items {
		if focusable, ok := item.Component.(Focusable); ok {
			focusable.Blur()
		}
	}
}

// IsFocused returns whether this component currently has focus
func (s *Stack) IsFocused() bool {
	// Check if any item is focused
	for _, item := range s.items {
		if focusable, ok := item.Component.(Focusable); ok {
			if focusable.IsFocused() {
				return true
			}
		}
	}
	return false
}

// HandleKey processes keyboard input when focused
func (s *Stack) HandleKey(key string) bool {
	// Pass to focused item (top-most first)
	for i := len(s.items) - 1; i >= 0; i-- {
		if handler, ok := s.items[i].Component.(interface{ HandleKey(string) bool }); ok {
			if focusable, ok := s.items[i].Component.(Focusable); ok && focusable.IsFocused() {
				return handler.HandleKey(key)
			}
		}
	}
	return false
}
