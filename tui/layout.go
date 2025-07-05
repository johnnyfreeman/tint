package tui

import "github.com/charmbracelet/lipgloss"

// Position represents a position on the screen
type Position struct {
	X, Y int
}

// Rectangle represents a rectangular area on the screen
type Rectangle struct {
	X, Y, Width, Height int
}

// LayoutDirection represents the direction of a layout
type LayoutDirection int

const (
	LayoutHorizontal LayoutDirection = iota
	LayoutVertical
)

// ClearArea clears a rectangular area on the screen with the given style
func ClearArea(screen *Screen, x, y, width, height int, style lipgloss.Style) {
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', style)
		}
	}
}

// ClearComponentArea clears a rectangular area using the theme's background color
// This should be called at the start of every component's Draw method to ensure
// the entire component area uses the theme's background color consistently
func ClearComponentArea(screen *Screen, x, y, width, height int, theme *Theme) {
	style := lipgloss.NewStyle().Background(theme.Palette.Background)
	ClearArea(screen, x, y, width, height, style)
}

// CenterComponent calculates the centered position for a component
func CenterComponent(containerWidth, containerHeight, componentWidth, componentHeight int) Position {
	x := (containerWidth - componentWidth) / 2
	y := (containerHeight - componentHeight) / 2
	
	// Ensure non-negative values
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	
	return Position{X: x, Y: y}
}

// AlignRight calculates the right-aligned position for a component
func AlignRight(containerWidth, componentWidth, margin int) int {
	x := containerWidth - componentWidth - margin
	if x < 0 {
		x = 0
	}
	return x
}

// AlignBottom calculates the bottom-aligned position for a component
func AlignBottom(containerHeight, componentHeight, margin int) int {
	y := containerHeight - componentHeight - margin
	if y < 0 {
		y = 0
	}
	return y
}

// GridLayout calculates positions for components in a grid
func GridLayout(items, cols, itemWidth, itemHeight, spacing int) []Rectangle {
	positions := make([]Rectangle, items)
	
	for i := 0; i < items; i++ {
		col := i % cols
		row := i / cols
		
		x := col * (itemWidth + spacing)
		y := row * (itemHeight + spacing)
		
		positions[i] = Rectangle{
			X:      x,
			Y:      y,
			Width:  itemWidth,
			Height: itemHeight,
		}
	}
	
	return positions
}

// Deprecated: Use LinearLayout from linear_layout.go instead
// FlexLayout provides flexible box layout calculations
type FlexLayout struct {
	Direction LayoutDirection
	Spacing   int
	Padding   int
}

// Deprecated: Use NewLinearLayout instead
func NewFlexLayout(direction LayoutDirection) *FlexLayout {
	return &FlexLayout{
		Direction: direction,
		Spacing:   1,
		Padding:   0,
	}
}

// Calculate computes positions for flex items
func (f *FlexLayout) Calculate(containerRect Rectangle, itemSizes []int) []Rectangle {
	positions := make([]Rectangle, len(itemSizes))
	
	// Starting position
	currentPos := f.Padding
	
	for i, size := range itemSizes {
		if f.Direction == LayoutHorizontal {
			positions[i] = Rectangle{
				X:      containerRect.X + currentPos,
				Y:      containerRect.Y + f.Padding,
				Width:  size,
				Height: containerRect.Height - 2*f.Padding,
			}
			currentPos += size + f.Spacing
		} else {
			positions[i] = Rectangle{
				X:      containerRect.X + f.Padding,
				Y:      containerRect.Y + currentPos,
				Width:  containerRect.Width - 2*f.Padding,
				Height: size,
			}
			currentPos += size + f.Spacing
		}
	}
	
	return positions
}

// Deprecated: Use Split from split_layout.go instead
// SplitLayout calculates positions for split pane layouts
type SplitLayout struct {
	Vertical bool
	Split    float64 // Percentage (0.0 to 1.0) or fixed pixels if > 1
}

// Deprecated: Use NewSplit instead
func NewSplitLayout(vertical bool, split float64) *SplitLayout {
	return &SplitLayout{
		Vertical: vertical,
		Split:    split,
	}
}

// Calculate computes positions for split panes
func (s *SplitLayout) Calculate(container Rectangle) (left, right Rectangle) {
	if s.Vertical {
		// Vertical split (left | right)
		splitX := int(s.Split)
		if s.Split <= 1.0 {
			// Percentage
			splitX = int(float64(container.Width) * s.Split)
		}
		
		left = Rectangle{
			X:      container.X,
			Y:      container.Y,
			Width:  splitX,
			Height: container.Height,
		}
		
		right = Rectangle{
			X:      container.X + splitX,
			Y:      container.Y,
			Width:  container.Width - splitX,
			Height: container.Height,
		}
	} else {
		// Horizontal split (top / bottom)
		splitY := int(s.Split)
		if s.Split <= 1.0 {
			// Percentage
			splitY = int(float64(container.Height) * s.Split)
		}
		
		left = Rectangle{
			X:      container.X,
			Y:      container.Y,
			Width:  container.Width,
			Height: splitY,
		}
		
		right = Rectangle{
			X:      container.X,
			Y:      container.Y + splitY,
			Width:  container.Width,
			Height: container.Height - splitY,
		}
	}
	
	return left, right
}

// Margin represents margins for all four sides
type Margin struct {
	Top, Right, Bottom, Left int
}

// NewMargin creates a margin with the same value for all sides
func NewMargin(all int) Margin {
	return Margin{Top: all, Right: all, Bottom: all, Left: all}
}

// NewMarginTB creates a margin with top/bottom and left/right values
func NewMarginTB(topBottom, leftRight int) Margin {
	return Margin{Top: topBottom, Right: leftRight, Bottom: topBottom, Left: leftRight}
}

// ApplyMargin returns a rectangle with margins applied
func ApplyMargin(rect Rectangle, margin Margin) Rectangle {
	return Rectangle{
		X:      rect.X + margin.Left,
		Y:      rect.Y + margin.Top,
		Width:  rect.Width - margin.Left - margin.Right,
		Height: rect.Height - margin.Top - margin.Bottom,
	}
}