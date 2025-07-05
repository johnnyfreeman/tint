package tui

import (
	"strings"
	"github.com/charmbracelet/lipgloss"
)

// TextArea represents a multi-line text editor
type TextArea struct {
	lines       []string
	cursorRow   int
	cursorCol   int  // Visual column position
	offsetRow   int  // For vertical scrolling
	offsetCol   int  // Visual column offset for horizontal scrolling
	width       int
	height      int
	focused     bool
	placeholder string
}

// NewTextArea creates a new text area
func NewTextArea() *TextArea {
	return &TextArea{
		lines:       []string{""},
		cursorRow:   0,
		cursorCol:   0,
		offsetRow:   0,
		offsetCol:   0,
		width:       40,
		height:      10,
		focused:     false,
		placeholder: "",
	}
}

// SetSize sets the display dimensions
func (t *TextArea) SetSize(width, height int) {
	t.width = width
	t.height = height
	t.adjustOffset()
}

// SetPlaceholder sets the placeholder text
func (t *TextArea) SetPlaceholder(placeholder string) {
	t.placeholder = placeholder
}

// SetValue sets the content and moves cursor to end
func (t *TextArea) SetValue(value string) {
	if value == "" {
		t.lines = []string{""}
		t.cursorRow = 0
		t.cursorCol = 0
	} else {
		t.lines = strings.Split(value, "\n")
		t.cursorRow = len(t.lines) - 1
		t.cursorCol = StringWidth(t.lines[t.cursorRow])
	}
	t.adjustOffset()
}

// Value returns the current content
func (t *TextArea) Value() string {
	return strings.Join(t.lines, "\n")
}

// Focus sets the focus state
func (t *TextArea) Focus() {
	t.focused = true
}

// Blur removes focus
func (t *TextArea) Blur() {
	t.focused = false
}

// IsFocused returns whether the text area is focused
func (t *TextArea) IsFocused() bool {
	return t.focused
}

// HandleInput processes keyboard input
func (t *TextArea) HandleInput(key string) {
	switch key {
	case "up", "ctrl+p":
		if t.cursorRow > 0 {
			t.cursorRow--
			// Adjust column if the new line is shorter
			lineWidth := StringWidth(t.lines[t.cursorRow])
			if t.cursorCol > lineWidth {
				t.cursorCol = lineWidth
			}
			t.adjustOffset()
		}
	case "down", "ctrl+n":
		if t.cursorRow < len(t.lines)-1 {
			t.cursorRow++
			// Adjust column if the new line is shorter
			lineWidth := StringWidth(t.lines[t.cursorRow])
			if t.cursorCol > lineWidth {
				t.cursorCol = lineWidth
			}
			t.adjustOffset()
		}
	case "left", "ctrl+b":
		t.moveCursorLeft()
		t.adjustOffset()
	case "right", "ctrl+f":
		t.moveCursorRight()
		t.adjustOffset()
	case "home", "ctrl+a":
		t.cursorCol = 0
		t.adjustOffset()
	case "end", "ctrl+e":
		t.cursorCol = StringWidth(t.lines[t.cursorRow])
		t.adjustOffset()
	case "enter":
		t.splitLineAtCursor()
		t.adjustOffset()
	case "backspace", "ctrl+h":
		t.deleteBeforeCursor()
		t.adjustOffset()
	case "delete", "ctrl+d":
		t.deleteAtCursor()
	default:
		// Handle regular character input
		if len(key) == 1 && key[0] >= 32 && key[0] < 127 {
			t.insertAtCursor(key)
			t.adjustOffset()
		}
	}
}

// adjustOffset ensures the cursor is visible
func (t *TextArea) adjustOffset() {
	// Vertical scrolling
	if t.cursorRow < t.offsetRow {
		t.offsetRow = t.cursorRow
	}
	if t.cursorRow >= t.offsetRow+t.height {
		t.offsetRow = t.cursorRow - t.height + 1
	}
	
	// Horizontal scrolling
	if t.cursorCol < t.offsetCol {
		t.offsetCol = t.cursorCol
	}
	if t.cursorCol >= t.offsetCol+t.width {
		t.offsetCol = t.cursorCol - t.width + 1
	}
	
	// Don't scroll past the beginning
	if t.offsetRow < 0 {
		t.offsetRow = 0
	}
	if t.offsetCol < 0 {
		t.offsetCol = 0
	}
}

// Draw renders the text area to the screen
func (t *TextArea) Draw(screen *Screen, x, y int, theme *Theme) {
	// Clear the entire text area with theme background
	ClearComponentArea(screen, x, y, t.width, t.height, theme)
	
	textStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Background)
	
	// Show placeholder if empty
	if len(t.lines) == 1 && t.lines[0] == "" && t.placeholder != "" && !t.focused {
		placeholderStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background).
			Italic(true)
		
		lines := strings.Split(t.placeholder, "\n")
		for i, line := range lines {
			if i >= t.height {
				break
			}
			if StringWidth(line) > t.width {
				line = TruncateWithEllipsis(line, t.width)
			}
			screen.DrawString(x, y+i, line, placeholderStyle)
		}
		return
	}
	
	// Draw visible lines
	for row := 0; row < t.height; row++ {
		lineIndex := t.offsetRow + row
		if lineIndex >= len(t.lines) {
			// Empty rows are already cleared by ClearComponentArea
			continue
		}
		
		line := t.lines[lineIndex]
		
		// Extract visible portion of the line
		visibleLine := t.getVisibleLine(line)
		
		// Draw the line
		screen.DrawString(x, y+row, visibleLine, textStyle)
		
		// Draw cursor if on this line and focused
		if t.focused && lineIndex == t.cursorRow {
			cursorScreenCol := t.getCursorScreenCol()
			if cursorScreenCol >= 0 && cursorScreenCol < t.width {
				cursorStyle := lipgloss.NewStyle().
					Foreground(theme.Palette.Background).
					Background(theme.Palette.Text)
				
				// Get character under cursor
				cursorChar, found := GetCharAtVisualCol(line, t.cursorCol)
				if !found {
					cursorChar = ' '
				}
				screen.DrawRune(x+cursorScreenCol, y+row, cursorChar, cursorStyle)
			}
		}
	}
	
	// Draw line numbers or scroll indicators if needed
	if len(t.lines) > t.height && t.focused {
		scrollStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background)
		
		// Top indicator
		if t.offsetRow > 0 {
			screen.DrawRune(x+t.width-1, y, '↑', scrollStyle)
		}
		// Bottom indicator
		if t.offsetRow+t.height < len(t.lines) {
			screen.DrawRune(x+t.width-1, y+t.height-1, '↓', scrollStyle)
		}
	}
}

// DrawInBox renders the text area inside a box with a title
func (t *TextArea) DrawInBox(screen *Screen, x, y, width, height int, title string, theme *Theme) {
	var borderColors, titleColors StateColors
	if t.focused {
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
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}
	
	// Draw box with title - use heavy borders when focused
	if t.focused {
		screen.DrawBrutalistBoxWithTitle(x, y, width, height, title, borderStyle, titleStyle)
	} else {
		screen.DrawBoxWithTitle(x, y, width, height, title, borderStyle, titleStyle)
	}
	
	// Set text area size based on box dimensions
	t.SetSize(width-4, height-2) // -4 for borders and padding, -2 for top/bottom borders
	
	// Draw the text area inside the box
	t.Draw(screen, x+2, y+1, theme)
}

// HandleKey processes keyboard input when focused
func (t *TextArea) HandleKey(key string) bool {
	if !t.focused {
		return false
	}
	t.HandleInput(key)
	return true
}


// GetSize returns the current width and height
func (t *TextArea) GetSize() (width, height int) {
	return t.width, t.height
}

// DrawWithBorder draws the component with a border and optional title
func (t *TextArea) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	t.DrawInBox(screen, x, y, t.width+4, t.height+2, title, theme)
}