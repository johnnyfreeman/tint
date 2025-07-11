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
func (c *Container) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	// Container decides to use available space (containers are typically space-filling)
	containerWidth := availableWidth
	containerHeight := availableHeight
	
	c.draw(screen, x, y, containerWidth, containerHeight, theme)
}

func (c *Container) draw(screen *Screen, x, y, width, height int, theme *Theme) {
	if width <= 0 || height <= 0 {
		return
	}

	ClearSurfaceArea(screen, x, y, width, height, theme)

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
		c.content.Draw(screen, contentX, contentY, contentWidth, contentHeight, theme)
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

// HandleInput processes keyboard input
func (c *Container) HandleInput(key string) {
	// Pass input to content
	if c.content != nil {
		c.content.HandleInput(key)
	}
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
