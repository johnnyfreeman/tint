package tui

// Split represents a split pane layout with constraints
type Split struct {
	first      Component
	second     Component
	constraint ConstraintSet
	vertical   bool
	width      int
	height     int
}

// NewSplit creates a new split layout
func NewSplit(vertical bool) *Split {
	return &Split{
		vertical:   vertical,
		constraint: NewConstraintSet(NewPercentage(0.5)), // Default 50/50 split
	}
}

// NewHSplit creates a horizontal split (top/bottom)
func NewHSplit() *Split {
	return NewSplit(false)
}

// NewVSplit creates a vertical split (left/right)
func NewVSplit() *Split {
	return NewSplit(true)
}

// SetFirst sets the first pane component
func (s *Split) SetFirst(component Component) {
	s.first = component
}

// SetSecond sets the second pane component
func (s *Split) SetSecond(component Component) {
	s.second = component
}

// SetConstraint sets the constraint for the first pane
// The second pane gets the remaining space
func (s *Split) SetConstraint(constraint ConstraintSet) {
	s.constraint = constraint
}

// SetRatio sets a ratio constraint for the first pane
func (s *Split) SetRatio(ratio float64) {
	s.constraint = NewConstraintSet(NewRatio(ratio))
}

// SetPercentage sets a percentage constraint for the first pane
func (s *Split) SetPercentage(pct float64) {
	s.constraint = NewConstraintSet(NewPercentage(pct))
}

// SetFixed sets a fixed size constraint for the first pane
func (s *Split) SetFixed(size int) {
	s.constraint = NewConstraintSet(NewLength(size))
}

// SetSize sets the width and height of the split
func (s *Split) SetSize(width, height int) {
	s.width = width
	s.height = height
}

// GetSize returns the current width and height
func (s *Split) GetSize() (width, height int) {
	return s.width, s.height
}

// Draw renders the split layout to the screen
func (s *Split) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	// Split decides to use available space for layout
	splitWidth := availableWidth
	splitHeight := availableHeight
	
	s.drawWithBounds(screen, x, y, splitWidth, splitHeight, theme)
}

// drawWithBounds draws the split with specific bounds
func (s *Split) drawWithBounds(screen *Screen, x, y, width, height int, theme *Theme) {
	// Clear the split area
	ClearComponentArea(screen, x, y, width, height, theme)

	var firstSize int
	if s.vertical {
		// Vertical split - calculate width for first pane
		firstSize = s.constraint.Calculate(width, 1.0)

		// Draw first pane
		if s.first != nil && firstSize > 0 {
			s.first.Draw(screen, x, y, firstSize, height, theme)
		}

		// Draw second pane
		secondX := x + firstSize
		secondWidth := width - firstSize
		if s.second != nil && secondWidth > 0 {
			s.second.Draw(screen, secondX, y, secondWidth, height, theme)
		}
	} else {
		// Horizontal split - calculate height for first pane
		firstSize = s.constraint.Calculate(height, 1.0)

		// Draw first pane
		if s.first != nil && firstSize > 0 {
			s.first.Draw(screen, x, y, width, firstSize, theme)
		}

		// Draw second pane
		secondY := y + firstSize
		secondHeight := height - firstSize
		if s.second != nil && secondHeight > 0 {
			s.second.Draw(screen, x, secondY, width, secondHeight, theme)
		}
	}
}

// GetFirst returns the first pane component
func (s *Split) GetFirst() Component {
	return s.first
}

// GetSecond returns the second pane component
func (s *Split) GetSecond() Component {
	return s.second
}

// HandleInput processes keyboard input
func (s *Split) HandleInput(key string) {
	// Split doesn't handle input itself
}

// IsVertical returns whether this is a vertical split
func (s *Split) IsVertical() bool {
	return s.vertical
}

// Component interface methods

// Focus gives keyboard focus to this component
func (s *Split) Focus() {
	// Focus first pane
	if s.first != nil {
		if focusable, ok := s.first.(Focusable); ok {
			focusable.Focus()
		}
	}
}

// Blur removes keyboard focus from this component
func (s *Split) Blur() {
	// Blur both panes
	if s.first != nil {
		if focusable, ok := s.first.(Focusable); ok {
			focusable.Blur()
		}
	}
	if s.second != nil {
		if focusable, ok := s.second.(Focusable); ok {
			focusable.Blur()
		}
	}
}

// IsFocused returns whether this component currently has focus
func (s *Split) IsFocused() bool {
	// Check if either pane is focused
	if s.first != nil {
		if focusable, ok := s.first.(Focusable); ok && focusable.IsFocused() {
			return true
		}
	}
	if s.second != nil {
		if focusable, ok := s.second.(Focusable); ok && focusable.IsFocused() {
			return true
		}
	}
	return false
}

// HandleKey processes keyboard input when focused
func (s *Split) HandleKey(key string) bool {
	// Pass to focused pane
	if s.first != nil {
		if handler, ok := s.first.(interface{ HandleKey(string) bool }); ok {
			if focusable, ok := s.first.(Focusable); ok && focusable.IsFocused() {
				return handler.HandleKey(key)
			}
		}
	}
	if s.second != nil {
		if handler, ok := s.second.(interface{ HandleKey(string) bool }); ok {
			if focusable, ok := s.second.(Focusable); ok && focusable.IsFocused() {
				return handler.HandleKey(key)
			}
		}
	}
	return false
}
