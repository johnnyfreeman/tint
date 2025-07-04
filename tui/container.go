package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Container represents a component that can contain another component
// with optional borders, padding, and title
type Container struct {
	content    Component
	title      string
	showBorder bool
	padding    Margin
	width      int
	height     int
	focused    bool
	borderStyle string // "single", "double", "heavy", "rounded"
}

// NewContainer creates a new container
func NewContainer() *Container {
	return &Container{
		showBorder:  true,
		padding:     NewMargin(1),
		borderStyle: "single",
	}
}

// SetContent sets the content component
func (c *Container) SetContent(component Component) {
	c.content = component
}

// SetTitle sets the container title
func (c *Container) SetTitle(title string) {
	c.title = title
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

	// Clear the container area
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	ClearArea(screen, x, y, width, height, bgStyle)

	contentX, contentY := x, y
	contentWidth, contentHeight := width, height

	// Draw border if enabled
	if c.showBorder {
		borderColor := theme.Palette.Border
		if c.focused {
			borderColor = theme.Palette.Primary
		}
		
		borderStyle := lipgloss.NewStyle().
			Foreground(borderColor).
			Background(theme.Palette.Background)
		
		// Draw border based on style
		switch c.borderStyle {
		case "double":
			c.drawDoubleBorder(screen, x, y, width, height, borderStyle)
		case "heavy", "bold":
			if c.focused {
				c.drawHeavyBorder(screen, x, y, width, height, borderStyle)
			} else {
				c.drawSingleBorder(screen, x, y, width, height, borderStyle)
			}
		case "rounded":
			c.drawRoundedBorder(screen, x, y, width, height, borderStyle)
		default:
			c.drawSingleBorder(screen, x, y, width, height, borderStyle)
		}

		// Draw title if present
		if c.title != "" {
			titleStyle := lipgloss.NewStyle().
				Foreground(theme.Palette.Text).
				Background(theme.Palette.Background).
				Bold(true)
			
			// Calculate title position
			titleText := " " + c.title + " "
			titleX := x + 2
			if titleX+len(titleText) > x+width-2 {
				// Truncate title if too long
				maxLen := width - 4
				if maxLen > 0 && len(titleText) > maxLen {
					titleText = titleText[:maxLen-3] + "..."
				}
			}
			
			screen.DrawString(titleX, y, titleText, titleStyle)
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

func (c *Container) drawSingleBorder(screen *Screen, x, y, width, height int, style lipgloss.Style) {
	// Corners
	screen.DrawRune(x, y, '┌', style)
	screen.DrawRune(x+width-1, y, '┐', style)
	screen.DrawRune(x, y+height-1, '└', style)
	screen.DrawRune(x+width-1, y+height-1, '┘', style)

	// Horizontal lines
	for i := 1; i < width-1; i++ {
		screen.DrawRune(x+i, y, '─', style)
		screen.DrawRune(x+i, y+height-1, '─', style)
	}

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '│', style)
		screen.DrawRune(x+width-1, y+i, '│', style)
	}
}

func (c *Container) drawDoubleBorder(screen *Screen, x, y, width, height int, style lipgloss.Style) {
	// Corners
	screen.DrawRune(x, y, '╔', style)
	screen.DrawRune(x+width-1, y, '╗', style)
	screen.DrawRune(x, y+height-1, '╚', style)
	screen.DrawRune(x+width-1, y+height-1, '╝', style)

	// Horizontal lines
	for i := 1; i < width-1; i++ {
		screen.DrawRune(x+i, y, '═', style)
		screen.DrawRune(x+i, y+height-1, '═', style)
	}

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '║', style)
		screen.DrawRune(x+width-1, y+i, '║', style)
	}
}

func (c *Container) drawHeavyBorder(screen *Screen, x, y, width, height int, style lipgloss.Style) {
	// Corners
	screen.DrawRune(x, y, '┏', style)
	screen.DrawRune(x+width-1, y, '┓', style)
	screen.DrawRune(x, y+height-1, '┗', style)
	screen.DrawRune(x+width-1, y+height-1, '┛', style)

	// Horizontal lines
	for i := 1; i < width-1; i++ {
		screen.DrawRune(x+i, y, '━', style)
		screen.DrawRune(x+i, y+height-1, '━', style)
	}

	// Vertical lines
	for i := 1; i < height-1; i++ {
		screen.DrawRune(x, y+i, '┃', style)
		screen.DrawRune(x+width-1, y+i, '┃', style)
	}
}

func (c *Container) drawRoundedBorder(screen *Screen, x, y, width, height int, style lipgloss.Style) {
	// Corners
	screen.DrawRune(x, y, '╭', style)
	screen.DrawRune(x+width-1, y, '╮', style)
	screen.DrawRune(x, y+height-1, '╰', style)
	screen.DrawRune(x+width-1, y+height-1, '╯', style)

	// Horizontal lines
	for i := 1; i < width-1; i++ {
		screen.DrawRune(x+i, y, '─', style)
		screen.DrawRune(x+i, y+height-1, '─', style)
	}

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