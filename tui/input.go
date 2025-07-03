package tui

import (
	"strings"
	"github.com/charmbracelet/lipgloss"
)

// Input represents a single-line text input field
type Input struct {
	value       string
	cursor      int
	offset      int // For horizontal scrolling when text is longer than width
	width       int
	placeholder string
	focused     bool
}

// NewInput creates a new input field
func NewInput() *Input {
	return &Input{
		value:       "",
		cursor:      0,
		offset:      0,
		width:       20,
		placeholder: "",
		focused:     false,
	}
}

// SetWidth sets the display width of the input
func (i *Input) SetWidth(width int) {
	i.width = width
	i.adjustOffset()
}

// SetPlaceholder sets the placeholder text
func (i *Input) SetPlaceholder(placeholder string) {
	i.placeholder = placeholder
}

// SetValue sets the input value and moves cursor to end
func (i *Input) SetValue(value string) {
	i.value = value
	i.cursor = len(value)
	i.adjustOffset()
}

// SetValueWithCursorAtStart sets the input value and moves cursor to start
func (i *Input) SetValueWithCursorAtStart(value string) {
	i.value = value
	i.cursor = 0
	i.offset = 0
}

// Value returns the current input value
func (i *Input) Value() string {
	return i.value
}

// Focus sets the focus state
func (i *Input) Focus() {
	i.focused = true
}

// Blur removes focus
func (i *Input) Blur() {
	i.focused = false
}

// IsFocused returns whether the input is focused
func (i *Input) IsFocused() bool {
	return i.focused
}

// HandleInput processes keyboard input
func (i *Input) HandleInput(key string) {
	switch key {
	case "left", "ctrl+b":
		if i.cursor > 0 {
			i.cursor--
			i.adjustOffset()
		}
	case "right", "ctrl+f":
		if i.cursor < len(i.value) {
			i.cursor++
			i.adjustOffset()
		}
	case "home", "ctrl+a":
		i.cursor = 0
		i.adjustOffset()
	case "end", "ctrl+e":
		i.cursor = len(i.value)
		i.adjustOffset()
	case "backspace", "ctrl+h":
		if i.cursor > 0 {
			i.value = i.value[:i.cursor-1] + i.value[i.cursor:]
			i.cursor--
			i.adjustOffset()
		}
	case "delete", "ctrl+d":
		if i.cursor < len(i.value) {
			i.value = i.value[:i.cursor] + i.value[i.cursor+1:]
		}
	case "ctrl+k": // Kill to end of line
		i.value = i.value[:i.cursor]
	case "ctrl+u": // Kill to beginning of line
		i.value = i.value[i.cursor:]
		i.cursor = 0
		i.adjustOffset()
	default:
		// Handle regular character input
		if len(key) == 1 && key[0] >= 32 && key[0] < 127 {
			i.value = i.value[:i.cursor] + key + i.value[i.cursor:]
			i.cursor++
			i.adjustOffset()
		}
	}
}

// adjustOffset ensures the cursor is visible
func (i *Input) adjustOffset() {
	// If cursor is before the visible area, scroll left
	if i.cursor < i.offset {
		i.offset = i.cursor
	}
	// If cursor is after the visible area, scroll right
	// Leave room for the cursor
	if i.cursor > i.offset+i.width-1 {
		i.offset = i.cursor - i.width + 1
	}
	// Don't scroll past the beginning
	if i.offset < 0 {
		i.offset = 0
	}
}

// Draw renders the input field to the screen
func (i *Input) Draw(screen *Screen, x, y int, theme *Theme) {
	style := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Background)
	
	if i.focused {
		style = style.Underline(true)
	}
	
	// Draw the visible portion of the text or placeholder
	var displayText string
	if i.value == "" && i.placeholder != "" {
		// Show placeholder
		placeholderStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background).
			Italic(true)
		
		displayText = i.placeholder
		if len(displayText) > i.width {
			displayText = displayText[:i.width]
		}
		screen.DrawString(x, y, displayText, placeholderStyle)
	} else {
		// Show value
		visibleEnd := i.offset + i.width
		if visibleEnd > len(i.value) {
			visibleEnd = len(i.value)
		}
		displayText = i.value[i.offset:visibleEnd]
		screen.DrawString(x, y, displayText, style)
	}
	
	// Draw cursor if focused
	if i.focused && i.cursor >= i.offset && i.cursor <= i.offset+i.width {
		cursorX := x + i.cursor - i.offset
		cursorStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.Background).
			Background(theme.Palette.Text)
		
		// Get the character under the cursor or use space if at end
		var cursorChar rune = ' '
		if i.cursor < len(i.value) {
			cursorChar = rune(i.value[i.cursor])
		}
		screen.DrawRune(cursorX, y, cursorChar, cursorStyle)
	}
	
	// Fill the rest of the input width with spaces
	remainingWidth := i.width - len(displayText)
	if remainingWidth > 0 {
		emptyStyle := lipgloss.NewStyle().
			Background(theme.Palette.Background)
		screen.DrawString(x+len(displayText), y, strings.Repeat(" ", remainingWidth), emptyStyle)
	}
}

// DrawInBox renders the input field inside a box with a title
func (i *Input) DrawInBox(screen *Screen, x, y int, title string, theme *Theme) {
	// Determine box width based on input width + padding
	boxWidth := i.width + 4 // 2 chars padding on each side
	boxHeight := 3 // Top border, content, bottom border
	
	var borderColors, titleColors StateColors
	if i.focused {
		borderColors = theme.Components.Container.Border.Focused
		titleColors = theme.Components.Container.Title.Focused
	} else {
		borderColors = theme.Components.Container.Border.Unfocused
		titleColors = theme.Components.Container.Title.Unfocused
	}
	
	borderStyle := lipgloss.NewStyle().
		Foreground(borderColors.Border).
		Background(theme.Palette.Background)
	titleStyle := lipgloss.NewStyle().
		Foreground(titleColors.Text).
		Background(theme.Palette.Background)
	
	// Fill background
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	for dy := 0; dy < boxHeight; dy++ {
		for dx := 0; dx < boxWidth; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}
	
	// Draw box with title
	screen.DrawBoxWithTitle(x, y, boxWidth, boxHeight, title, borderStyle, titleStyle)
	
	// Draw the input inside the box
	i.Draw(screen, x+2, y+1, theme)
}

// GetSize returns the current width and height
func (i *Input) GetSize() (width, height int) {
	return i.width, 1
}

// SetSize sets the width and height of the component
func (i *Input) SetSize(width, height int) {
	i.width = width
	// Input is always 1 line high
}

// DrawWithBorder draws the component with a border and optional title
func (i *Input) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	i.DrawInBox(screen, x, y, title, theme)
}

// HandleKey processes keyboard input when focused (Component interface)
func (i *Input) HandleKey(key string) bool {
	if !i.focused {
		return false
	}
	i.HandleInput(key)
	return true
}