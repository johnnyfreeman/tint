package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Container represents a component that can contain another component
// with optional borders, padding, and title
type Container struct {
	content        Component
	title          string
	showBorder     bool
	padding        Margin
	width          int
	height         int
	focused        bool
	borderStyle    string // "single", "double", "heavy", "rounded"
	borderElements []BorderElement
	useSurface     bool   // Whether to use Surface color instead of Background
}

// NewContainer creates a new container
func NewContainer() *Container {
	return &Container{
		showBorder:     true,
		padding:        NewMargin(1),
		borderStyle:    "single",
		borderElements: []BorderElement{},
	}
}

// SetContent sets the content component
func (c *Container) SetContent(component Component) {
	c.content = component
}

// SetShowBorder sets whether to show the border
func (c *Container) SetShowBorder(show bool) {
	c.showBorder = show
}

// SetPadding sets the padding inside the container
func (c *Container) SetPadding(padding Margin) {
	c.padding = padding
}

// SetBorderStyle sets the border style
// Options: "single", "double", "heavy", "rounded"
func (c *Container) SetBorderStyle(style string) {
	c.borderStyle = style
}

// SetUseSurface sets whether the container should use Surface color instead of Background
// This should be set to true for containers inside modals
func (c *Container) SetUseSurface(useSurface bool) {
	c.useSurface = useSurface
}

// AddBorderElement adds an inline element to the border
func (c *Container) AddBorderElement(element InlineElement, position BorderPosition, alignment BorderAlignment) {
	c.borderElements = append(c.borderElements, BorderElement{
		Element:   element,
		Position:  position,
		Alignment: alignment,
	})
}

// AddBorderElementWithOffset adds an inline element with a specific offset
func (c *Container) AddBorderElementWithOffset(element InlineElement, position BorderPosition, alignment BorderAlignment, offset int) {
	c.borderElements = append(c.borderElements, BorderElement{
		Element:   element,
		Position:  position,
		Alignment: alignment,
		Offset:    offset,
	})
}

// ClearBorderElements removes all border elements
func (c *Container) ClearBorderElements() {
	c.borderElements = []BorderElement{}
}

// SetTitle is a convenience method that adds a title as a border element
func (c *Container) SetTitle(title string) {
	c.title = title
	// Remove any existing title elements
	newElements := []BorderElement{}
	for _, elem := range c.borderElements {
		if _, isText := elem.Element.(*TextElement); !isText || elem.Position != BorderTop || elem.Alignment != BorderAlignLeft {
			newElements = append(newElements, elem)
		}
	}
	c.borderElements = newElements

	// Add title as a border element if not empty
	if title != "" {
		c.AddBorderElementWithOffset(NewTextElement(title), BorderTop, BorderAlignLeft, 2)
	}
}

// Draw renders the container to the screen
func (c *Container) Draw(screen *Screen, x, y int, theme *Theme) {
	c.draw(screen, x, y, c.width, c.height, theme)
}

// DrawWithBounds draws the container with specific bounds
func (c *Container) DrawWithBounds(screen *Screen, x, y, width, height int, theme *Theme) {
	c.draw(screen, x, y, width, height, theme)
}

func (c *Container) draw(screen *Screen, x, y, width, height int, theme *Theme) {
	if width <= 0 || height <= 0 {
		return
	}

	// Clear the entire container area with appropriate background
	if c.useSurface {
		ClearSurfaceArea(screen, x, y, width, height, theme)
	} else {
		ClearComponentArea(screen, x, y, width, height, theme)
	}

	contentX, contentY := x, y
	contentWidth, contentHeight := width, height

	// Draw border if enabled
	if c.showBorder {
		borderColor := theme.Palette.Border
		if c.focused {
			borderColor = theme.Palette.Primary
		}

		borderStyle := lipgloss.NewStyle().
			Foreground(borderColor)

		// Draw border based on style
		switch c.borderStyle {
		case "double":
			c.drawDoubleBorder(screen, x, y, width, height, borderStyle, theme)
		case "heavy", "bold":
			if c.focused {
				c.drawHeavyBorder(screen, x, y, width, height, borderStyle, theme)
			} else {
				c.drawSingleBorder(screen, x, y, width, height, borderStyle, theme)
			}
		case "rounded":
			c.drawRoundedBorder(screen, x, y, width, height, borderStyle, theme)
		default:
			c.drawSingleBorder(screen, x, y, width, height, borderStyle, theme)
		}

		// Adjust content area for border
		contentX++
		contentY++
		contentWidth -= 2
		contentHeight -= 2
	}

	// Apply padding
	contentX += c.padding.Left
	contentY += c.padding.Top
	contentWidth -= c.padding.Left + c.padding.Right
	contentHeight -= c.padding.Top + c.padding.Bottom

	// Draw content if it fits
	if contentWidth > 0 && contentHeight > 0 && c.content != nil {
		// Draw content with bounds
		if drawer, ok := c.content.(interface {
			DrawWithBounds(*Screen, int, int, int, int, *Theme)
		}); ok {
			drawer.DrawWithBounds(screen, contentX, contentY, contentWidth, contentHeight, theme)
		} else {
			c.content.Draw(screen, contentX, contentY, theme)
		}
	}
}

// drawHorizontalBorderWithElements draws a horizontal border line with embedded elements
func (c *Container) drawHorizontalBorderWithElements(screen *Screen, x, y, width int, position BorderPosition, borderChar rune, borderStyle lipgloss.Style, theme *Theme) {
	// First, draw the entire border line
	for i := 1; i < width-1; i++ {
		screen.DrawRune(x+i, y, borderChar, borderStyle)
	}

	// Then draw elements on top
	elements := c.getElementsForPosition(position)
	if len(elements) == 0 {
		return
	}

	// Group elements by alignment
	leftElements := []BorderElement{}
	centerElements := []BorderElement{}
	rightElements := []BorderElement{}

	for _, elem := range elements {
		switch elem.Alignment {
		case BorderAlignLeft:
			leftElements = append(leftElements, elem)
		case BorderAlignCenter:
			centerElements = append(centerElements, elem)
		case BorderAlignRight:
			rightElements = append(rightElements, elem)
		}
	}

	// Draw left-aligned elements
	currentX := x + 1
	for _, elem := range leftElements {
		currentX += elem.Offset
		if currentX+elem.Element.Width() < x+width-1 {
			width := elem.Element.Draw(screen, currentX, y, theme, c.focused)
			currentX += width
		}
	}

	// Draw center-aligned elements
	for _, elem := range centerElements {
		elemWidth := elem.Element.Width()
		centerX := x + (width-elemWidth)/2 + elem.Offset
		if centerX > x && centerX+elemWidth < x+width-1 {
			elem.Element.Draw(screen, centerX, y, theme, c.focused)
		}
	}

	// Draw right-aligned elements
	for i := len(rightElements) - 1; i >= 0; i-- {
		elem := rightElements[i]
		elemWidth := elem.Element.Width()
		rightX := x + width - 1 - elemWidth - elem.Offset
		if rightX > x {
			elem.Element.Draw(screen, rightX, y, theme, c.focused)
		}
	}
}

// getElementsForPosition returns all elements for a given border position
func (c *Container) getElementsForPosition(position BorderPosition) []BorderElement {
	result := []BorderElement{}
	for _, elem := range c.borderElements {
		if elem.Position == position {
			result = append(result, elem)
		}
	}
	return result
}

func (c *Container) drawSingleBorder(screen *Screen, x, y, width, height int, style lipgloss.Style, theme *Theme) {
	// Corners
	screen.DrawRune(x, y, '┌', style)
	screen.DrawRune(x+width-1, y, '┐', style)
	screen.DrawRune(x, y+height-1, '└', style)
	screen.DrawRune(x+width-1, y+height-1, '┘', style)

	// Top border
	c.drawHorizontalBorderWithElements(screen, x, y, width, BorderTop, '─', style, theme)

	// Bottom border
	c.drawHorizontalBorderWithElements(screen, x, y+height-1, width, BorderBottom, '─', style, theme)

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '│', style)
		screen.DrawRune(x+width-1, y+i, '│', style)
	}
}

func (c *Container) drawDoubleBorder(screen *Screen, x, y, width, height int, style lipgloss.Style, theme *Theme) {
	// Corners
	screen.DrawRune(x, y, '╔', style)
	screen.DrawRune(x+width-1, y, '╗', style)
	screen.DrawRune(x, y+height-1, '╚', style)
	screen.DrawRune(x+width-1, y+height-1, '╝', style)

	// Top border
	c.drawHorizontalBorderWithElements(screen, x, y, width, BorderTop, '═', style, theme)

	// Bottom border
	c.drawHorizontalBorderWithElements(screen, x, y+height-1, width, BorderBottom, '═', style, theme)

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '║', style)
		screen.DrawRune(x+width-1, y+i, '║', style)
	}
}

func (c *Container) drawHeavyBorder(screen *Screen, x, y, width, height int, style lipgloss.Style, theme *Theme) {
	// Corners
	screen.DrawRune(x, y, '┏', style)
	screen.DrawRune(x+width-1, y, '┓', style)
	screen.DrawRune(x, y+height-1, '┗', style)
	screen.DrawRune(x+width-1, y+height-1, '┛', style)

	// Top border
	c.drawHorizontalBorderWithElements(screen, x, y, width, BorderTop, '━', style, theme)

	// Bottom border
	c.drawHorizontalBorderWithElements(screen, x, y+height-1, width, BorderBottom, '━', style, theme)

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '┃', style)
		screen.DrawRune(x+width-1, y+i, '┃', style)
	}
}

func (c *Container) drawRoundedBorder(screen *Screen, x, y, width, height int, style lipgloss.Style, theme *Theme) {
	// Corners
	screen.DrawRune(x, y, '╭', style)
	screen.DrawRune(x+width-1, y, '╮', style)
	screen.DrawRune(x, y+height-1, '╰', style)
	screen.DrawRune(x+width-1, y+height-1, '╯', style)

	// Top border
	c.drawHorizontalBorderWithElements(screen, x, y, width, BorderTop, '─', style, theme)

	// Bottom border
	c.drawHorizontalBorderWithElements(screen, x, y+height-1, width, BorderBottom, '─', style, theme)

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '│', style)
		screen.DrawRune(x+width-1, y+i, '│', style)
	}
}

// Focus gives keyboard focus to this component
func (c *Container) Focus() {
	c.focused = true
	if c.content != nil {
		if focusable, ok := c.content.(Focusable); ok {
			focusable.Focus()
		}
	}
}

// Blur removes keyboard focus from this component
func (c *Container) Blur() {
	c.focused = false
	if c.content != nil {
		if focusable, ok := c.content.(Focusable); ok {
			focusable.Blur()
		}
	}
}

// IsFocused returns whether this component currently has focus
func (c *Container) IsFocused() bool {
	return c.focused
}

// HandleKey processes keyboard input when focused
func (c *Container) HandleKey(key string) bool {
	// Pass key events to content
	if c.content != nil {
		if handler, ok := c.content.(interface{ HandleKey(string) bool }); ok {
			return handler.HandleKey(key)
		}
	}
	return false
}

// SetSize sets the width and height of the component
func (c *Container) SetSize(width, height int) {
	c.width = width
	c.height = height
}

// GetSize returns the current width and height
func (c *Container) GetSize() (width, height int) {
	return c.width, c.height
}

// DrawWithBorder draws the component with a border and optional title
func (c *Container) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	oldTitle := c.title
	oldShowBorder := c.showBorder

	c.title = title
	c.showBorder = true
	c.Draw(screen, x, y, theme)

	c.title = oldTitle
	c.showBorder = oldShowBorder
}

// Panel is a convenience function to create a container with common settings
func Panel(title string, content Component) *Container {
	container := NewContainer()
	container.SetTitle(title)
	container.SetContent(content)
	container.SetPadding(NewMargin(1))
	return container
}

// Box is a convenience function to create a borderless container
func Box(content Component, padding int) *Container {
	container := NewContainer()
	container.SetContent(content)
	container.SetShowBorder(false)
	container.SetPadding(NewMargin(padding))
	return container
}
