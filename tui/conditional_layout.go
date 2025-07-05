package tui

// ConditionFunc is a function that returns whether a condition is met
type ConditionFunc func(width, height int) bool

// ConditionalItem represents an item that's shown based on a condition
type ConditionalItem struct {
	Component Component
	Condition ConditionFunc
}

// Conditional shows different components based on conditions
type Conditional struct {
	items      []ConditionalItem
	fallback   Component
	width      int
	height     int
}

// NewConditional creates a new conditional layout
func NewConditional() *Conditional {
	return &Conditional{
		items: []ConditionalItem{},
	}
}

// Add adds a component with a condition
func (c *Conditional) Add(component Component, condition ConditionFunc) {
	c.items = append(c.items, ConditionalItem{
		Component: component,
		Condition: condition,
	})
}

// AddMinSize adds a component that shows when size is at least the specified dimensions
func (c *Conditional) AddMinSize(component Component, minWidth, minHeight int) {
	c.Add(component, func(w, h int) bool {
		return w >= minWidth && h >= minHeight
	})
}

// AddMaxSize adds a component that shows when size is at most the specified dimensions
func (c *Conditional) AddMaxSize(component Component, maxWidth, maxHeight int) {
	c.Add(component, func(w, h int) bool {
		return w <= maxWidth && h <= maxHeight
	})
}

// AddWidthRange adds a component that shows when width is in the specified range
func (c *Conditional) AddWidthRange(component Component, minWidth, maxWidth int) {
	c.Add(component, func(w, h int) bool {
		return w >= minWidth && w <= maxWidth
	})
}

// AddHeightRange adds a component that shows when height is in the specified range
func (c *Conditional) AddHeightRange(component Component, minHeight, maxHeight int) {
	c.Add(component, func(w, h int) bool {
		return h >= minHeight && h <= maxHeight
	})
}

// AddAspectRatio adds a component that shows when aspect ratio matches
func (c *Conditional) AddAspectRatio(component Component, minRatio, maxRatio float64) {
	c.Add(component, func(w, h int) bool {
		if h == 0 {
			return false
		}
		ratio := float64(w) / float64(h)
		return ratio >= minRatio && ratio <= maxRatio
	})
}

// SetFallback sets the fallback component when no conditions match
func (c *Conditional) SetFallback(component Component) {
	c.fallback = component
}

// SetSize sets the width and height of the layout
func (c *Conditional) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// GetSize returns the current width and height
func (c *Conditional) GetSize() (width, height int) {
	return c.width, c.height
}

// Draw renders the conditional layout to the screen
func (c *Conditional) Draw(screen *Screen, x, y int, theme *Theme) {
	c.DrawWithBounds(screen, x, y, c.width, c.height, theme)
}

// DrawWithBounds draws the conditional with specific bounds
func (c *Conditional) DrawWithBounds(screen *Screen, x, y, width, height int, theme *Theme) {
	// Find the first component whose condition is met
	var componentToDraw Component
	
	for _, item := range c.items {
		if item.Condition(width, height) {
			componentToDraw = item.Component
			break
		}
	}
	
	// Use fallback if no conditions matched
	if componentToDraw == nil && c.fallback != nil {
		componentToDraw = c.fallback
	}
	
	// Draw the selected component
	if componentToDraw != nil {
		if drawer, ok := componentToDraw.(interface {
			DrawWithBounds(*Screen, int, int, int, int, *Theme)
		}); ok {
			drawer.DrawWithBounds(screen, x, y, width, height, theme)
		} else if sizable, ok := componentToDraw.(interface {
			SetSize(int, int)
			Draw(*Screen, int, int, *Theme)
		}); ok {
			sizable.SetSize(width, height)
			sizable.Draw(screen, x, y, theme)
		} else {
			componentToDraw.Draw(screen, x, y, theme)
		}
	}
}

// GetActiveComponent returns the currently active component based on size
func (c *Conditional) GetActiveComponent() Component {
	for _, item := range c.items {
		if item.Condition(c.width, c.height) {
			return item.Component
		}
	}
	return c.fallback
}

// ResponsiveLayout is a helper for creating responsive layouts
type ResponsiveLayout struct {
	*Conditional
}

// NewResponsiveLayout creates a new responsive layout
func NewResponsiveLayout() *ResponsiveLayout {
	return &ResponsiveLayout{
		Conditional: NewConditional(),
	}
}

// AddMobile adds a component for mobile-sized screens (< 80 cols)
func (r *ResponsiveLayout) AddMobile(component Component) {
	r.AddMaxSize(component, 79, 9999)
}

// AddTablet adds a component for tablet-sized screens (80-120 cols)
func (r *ResponsiveLayout) AddTablet(component Component) {
	r.AddWidthRange(component, 80, 120)
}

// AddDesktop adds a component for desktop-sized screens (> 120 cols)
func (r *ResponsiveLayout) AddDesktop(component Component) {
	r.AddMinSize(component, 121, 0)
}

// AddCompact adds a component for compact height (< 30 rows)
func (r *ResponsiveLayout) AddCompact(component Component) {
	r.AddMaxSize(component, 9999, 29)
}

// AddTall adds a component for tall screens (>= 30 rows)
func (r *ResponsiveLayout) AddTall(component Component) {
	r.AddMinSize(component, 0, 30)
}

// Component interface methods

// Focus gives keyboard focus to this component
func (c *Conditional) Focus() {
	if active := c.GetActiveComponent(); active != nil {
		if focusable, ok := active.(Focusable); ok {
			focusable.Focus()
		}
	}
}

// Blur removes keyboard focus from this component
func (c *Conditional) Blur() {
	if active := c.GetActiveComponent(); active != nil {
		if focusable, ok := active.(Focusable); ok {
			focusable.Blur()
		}
	}
}

// IsFocused returns whether this component currently has focus
func (c *Conditional) IsFocused() bool {
	if active := c.GetActiveComponent(); active != nil {
		if focusable, ok := active.(Focusable); ok {
			return focusable.IsFocused()
		}
	}
	return false
}

// HandleKey processes keyboard input when focused
func (c *Conditional) HandleKey(key string) bool {
	if active := c.GetActiveComponent(); active != nil {
		if handler, ok := active.(interface{ HandleKey(string) bool }); ok {
			return handler.HandleKey(key)
		}
	}
	return false
}