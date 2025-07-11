package tui

// LinearLayout arranges components in a line (horizontal or vertical)
type LinearLayout struct {
	config LayoutConfig
	items  []LinearItem
	width  int
	height int
}

// LinearItem represents an item in a linear layout
type LinearItem struct {
	Component  Component
	Constraint ConstraintSet
}

// NewLinearLayout creates a new linear layout
func NewLinearLayout(direction Direction) *LinearLayout {
	return &LinearLayout{
		config: NewLayoutConfig(direction),
		items:  []LinearItem{},
	}
}

// SetSpacing sets the spacing between items
func (l *LinearLayout) SetSpacing(spacing int) {
	l.config.Spacing = spacing
}

// SetPadding sets the padding around the layout
func (l *LinearLayout) SetPadding(padding Margin) {
	l.config.Padding = padding
}

// SetAlignment sets how items are aligned
func (l *LinearLayout) SetAlignment(alignment Alignment) {
	l.config.Alignment = alignment
}

// Add adds a component with a constraint
func (l *LinearLayout) Add(component Component, constraint ConstraintSet) {
	l.items = append(l.items, LinearItem{
		Component:  component,
		Constraint: constraint,
	})
}

// AddFixed adds a component with a fixed size
func (l *LinearLayout) AddFixed(component Component, size int) {
	l.Add(component, NewConstraintSet(NewLength(size)))
}

// AddFlex adds a component with a ratio constraint
func (l *LinearLayout) AddFlex(component Component, flex float64) {
	l.Add(component, NewConstraintSet(NewRatio(flex)))
}

// AddPercentage adds a component with a percentage constraint
func (l *LinearLayout) AddPercentage(component Component, pct float64) {
	l.Add(component, NewConstraintSet(NewPercentage(pct)))
}

// SetSize sets the width and height of the layout
func (l *LinearLayout) SetSize(width, height int) {
	l.width = width
	l.height = height
}

// GetSize returns the current width and height
func (l *LinearLayout) GetSize() (width, height int) {
	return l.width, l.height
}

// Draw renders the linear layout to the screen
func (l *LinearLayout) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	// LinearLayout decides to use available space for layout
	layoutWidth := availableWidth
	layoutHeight := availableHeight
	
	l.drawWithBounds(screen, x, y, layoutWidth, layoutHeight, theme)
}

// drawWithBounds draws the layout with specific bounds
func (l *LinearLayout) drawWithBounds(screen *Screen, x, y, width, height int, theme *Theme) {
	if len(l.items) == 0 {
		return
	}

	// Clear the layout area
	ClearComponentArea(screen, x, y, width, height, theme)

	// Apply padding
	contentRect := ApplyMargin(Rectangle{X: x, Y: y, Width: width, Height: height}, l.config.Padding)

	// Calculate item sizes
	constraints := make([]ConstraintSet, len(l.items))
	for i, item := range l.items {
		constraints[i] = item.Constraint
	}

	var sizes []int
	if l.config.Direction == Horizontal {
		// Account for spacing
		availableSize := contentRect.Width - (len(l.items)-1)*l.config.Spacing
		sizes = CalculateConstraints(constraints, availableSize)
	} else {
		// Account for spacing
		availableSize := contentRect.Height - (len(l.items)-1)*l.config.Spacing
		sizes = CalculateConstraints(constraints, availableSize)
	}

	// Draw items
	currentX := contentRect.X
	currentY := contentRect.Y

	for i, item := range l.items {
		var itemX, itemY, itemWidth, itemHeight int

		if l.config.Direction == Horizontal {
			itemX = currentX
			itemWidth = sizes[i]
			itemHeight = contentRect.Height

			// Apply alignment
			switch l.config.Alignment {
			case AlignStart:
				itemY = contentRect.Y
			case AlignCenter:
				itemY = contentRect.Y + (contentRect.Height-itemHeight)/2
			case AlignEnd:
				itemY = contentRect.Y + contentRect.Height - itemHeight
			default: // AlignStretch
				itemY = contentRect.Y
			}

			currentX += itemWidth + l.config.Spacing
		} else {
			itemY = currentY
			itemHeight = sizes[i]
			itemWidth = contentRect.Width

			// Apply alignment
			switch l.config.Alignment {
			case AlignStart:
				itemX = contentRect.X
			case AlignCenter:
				itemX = contentRect.X + (contentRect.Width-itemWidth)/2
			case AlignEnd:
				itemX = contentRect.X + contentRect.Width - itemWidth
			default: // AlignStretch
				itemX = contentRect.X
			}

			currentY += itemHeight + l.config.Spacing
		}

		// Draw the component
		item.Component.Draw(screen, itemX, itemY, itemWidth, itemHeight, theme)
	}
}

// Clear removes all items from the layout
func (l *LinearLayout) Clear() {
	l.items = []LinearItem{}
}

// HandleInput processes keyboard input
func (l *LinearLayout) HandleInput(key string) {
	// LinearLayout doesn't handle input itself
}

// Count returns the number of items in the layout
func (l *LinearLayout) Count() int {
	return len(l.items)
}

// GetItem returns the component at the specified index
func (l *LinearLayout) GetItem(index int) Component {
	if index >= 0 && index < len(l.items) {
		return l.items[index].Component
	}
	return nil
}

// Component interface methods

// Focus gives keyboard focus to this component
func (l *LinearLayout) Focus() {
	// Focus first focusable item
	for _, item := range l.items {
		if focusable, ok := item.Component.(Focusable); ok {
			focusable.Focus()
			break
		}
	}
}

// Blur removes keyboard focus from this component
func (l *LinearLayout) Blur() {
	// Blur all items
	for _, item := range l.items {
		if focusable, ok := item.Component.(Focusable); ok {
			focusable.Blur()
		}
	}
}

// IsFocused returns whether this component currently has focus
func (l *LinearLayout) IsFocused() bool {
	// Check if any item is focused
	for _, item := range l.items {
		if focusable, ok := item.Component.(Focusable); ok {
			if focusable.IsFocused() {
				return true
			}
		}
	}
	return false
}

// HandleKey processes keyboard input when focused
func (l *LinearLayout) HandleKey(key string) bool {
	// Pass to focused item
	for _, item := range l.items {
		if handler, ok := item.Component.(interface{ HandleKey(string) bool }); ok {
			if focusable, ok := item.Component.(Focusable); ok && focusable.IsFocused() {
				return handler.HandleKey(key)
			}
		}
	}
	return false
}

// HBox creates a horizontal linear layout
func HBox() *LinearLayout {
	return NewLinearLayout(Horizontal)
}

// VBox creates a vertical linear layout
func VBox() *LinearLayout {
	return NewLinearLayout(Vertical)
}
