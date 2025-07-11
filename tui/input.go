package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Input represents a single-line text input field
type Input struct {
	value       string
	cursor      int // Visual column position
	offset      int // Visual column offset for horizontal scrolling
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
	i.cursor = StringWidth(value)
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
		i.moveCursorLeft()
		i.adjustOffset()
	case "right", "ctrl+f":
		i.moveCursorRight()
		i.adjustOffset()
	case "home", "ctrl+a":
		i.cursor = 0
		i.adjustOffset()
	case "end", "ctrl+e":
		i.cursor = StringWidth(i.value)
		i.adjustOffset()
	case "backspace", "ctrl+h":
		i.deleteBeforeCursorInput()
		i.adjustOffset()
	case "delete", "ctrl+d":
		i.deleteAtCursorInput()
	case "ctrl+k": // Kill to end of line
		i.killToEndOfLine()
	case "ctrl+u": // Kill to beginning of line
		i.killToBeginningOfLine()
		i.adjustOffset()
	default:
		// Handle regular character input
		if len(key) == 1 && key[0] >= 32 && key[0] < 127 {
			i.insertAtCursorInput(key)
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
func (i *Input) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	// Input decides to use available width but always height=1
	inputWidth := availableWidth
	inputHeight := 1
	
	// Update internal width from decision
	i.width = inputWidth
	i.adjustOffset()
	
	// Clear the entire input area with theme background
	ClearComponentArea(screen, x, y, inputWidth, inputHeight, theme)

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
		if StringWidth(displayText) > inputWidth {
			displayText = TruncateWithEllipsis(displayText, inputWidth)
		}
		screen.DrawString(x, y, displayText, placeholderStyle)
	} else {
		// Show value
		displayText = i.getVisibleValuePortion()
		screen.DrawString(x, y, displayText, style)
	}

	// Draw cursor if focused
	if i.focused && i.cursor >= i.offset && i.cursor <= i.offset+inputWidth {
		cursorX := x + i.cursor - i.offset
		cursorStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.Background).
			Background(theme.Palette.Text)

		// Get the character under the cursor or use space if at end
		cursorChar, found := GetCharAtVisualCol(i.value, i.cursor)
		if !found {
			cursorChar = ' '
		}
		screen.DrawRune(cursorX, y, cursorChar, cursorStyle)
	}

	// The rest of the input width is already cleared by ClearComponentArea
}


