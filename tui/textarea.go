package tui

import (
	"strings"
	"github.com/charmbracelet/lipgloss"
)

// TextArea represents a multi-line text editor
type TextArea struct {
	lines       []string
	cursorRow   int
	cursorCol   int
	offsetRow   int // For vertical scrolling
	offsetCol   int // For horizontal scrolling
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
		t.cursorCol = len(t.lines[t.cursorRow])
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
			if t.cursorCol > len(t.lines[t.cursorRow]) {
				t.cursorCol = len(t.lines[t.cursorRow])
			}
			t.adjustOffset()
		}
	case "down", "ctrl+n":
		if t.cursorRow < len(t.lines)-1 {
			t.cursorRow++
			// Adjust column if the new line is shorter
			if t.cursorCol > len(t.lines[t.cursorRow]) {
				t.cursorCol = len(t.lines[t.cursorRow])
			}
			t.adjustOffset()
		}
	case "left", "ctrl+b":
		if t.cursorCol > 0 {
			t.cursorCol--
		} else if t.cursorRow > 0 {
			// Move to end of previous line
			t.cursorRow--
			t.cursorCol = len(t.lines[t.cursorRow])
		}
		t.adjustOffset()
	case "right", "ctrl+f":
		if t.cursorCol < len(t.lines[t.cursorRow]) {
			t.cursorCol++
		} else if t.cursorRow < len(t.lines)-1 {
			// Move to beginning of next line
			t.cursorRow++
			t.cursorCol = 0
		}
		t.adjustOffset()
	case "home", "ctrl+a":
		t.cursorCol = 0
		t.adjustOffset()
	case "end", "ctrl+e":
		t.cursorCol = len(t.lines[t.cursorRow])
		t.adjustOffset()
	case "enter":
		// Split the current line at cursor position
		currentLine := t.lines[t.cursorRow]
		before := currentLine[:t.cursorCol]
		after := currentLine[t.cursorCol:]
		
		// Update current line and insert new line
		t.lines[t.cursorRow] = before
		newLines := append(t.lines[:t.cursorRow+1], append([]string{after}, t.lines[t.cursorRow+1:]...)...)
		t.lines = newLines
		
		// Move cursor to beginning of new line
		t.cursorRow++
		t.cursorCol = 0
		t.adjustOffset()
	case "backspace", "ctrl+h":
		if t.cursorCol > 0 {
			// Delete character before cursor
			line := t.lines[t.cursorRow]
			t.lines[t.cursorRow] = line[:t.cursorCol-1] + line[t.cursorCol:]
			t.cursorCol--
		} else if t.cursorRow > 0 {
			// Join with previous line
			prevLine := t.lines[t.cursorRow-1]
			currentLine := t.lines[t.cursorRow]
			t.cursorCol = len(prevLine)
			t.lines[t.cursorRow-1] = prevLine + currentLine
			// Remove current line
			t.lines = append(t.lines[:t.cursorRow], t.lines[t.cursorRow+1:]...)
			t.cursorRow--
		}
		t.adjustOffset()
	case "delete", "ctrl+d":
		line := t.lines[t.cursorRow]
		if t.cursorCol < len(line) {
			// Delete character at cursor
			t.lines[t.cursorRow] = line[:t.cursorCol] + line[t.cursorCol+1:]
		} else if t.cursorRow < len(t.lines)-1 {
			// Join with next line
			t.lines[t.cursorRow] = line + t.lines[t.cursorRow+1]
			// Remove next line
			t.lines = append(t.lines[:t.cursorRow+1], t.lines[t.cursorRow+2:]...)
		}
	default:
		// Handle regular character input
		if len(key) == 1 && key[0] >= 32 && key[0] < 127 {
			line := t.lines[t.cursorRow]
			t.lines[t.cursorRow] = line[:t.cursorCol] + key + line[t.cursorCol:]
			t.cursorCol++
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
			if len(line) > t.width {
				line = line[:t.width]
			}
			screen.DrawString(x, y+i, line, placeholderStyle)
		}
		return
	}
	
	// Draw visible lines
	for row := 0; row < t.height; row++ {
		lineIndex := t.offsetRow + row
		if lineIndex >= len(t.lines) {
			// Fill empty rows
			emptyStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
			screen.DrawString(x, y+row, strings.Repeat(" ", t.width), emptyStyle)
			continue
		}
		
		line := t.lines[lineIndex]
		
		// Extract visible portion of the line
		visibleLine := ""
		if t.offsetCol < len(line) {
			endCol := t.offsetCol + t.width
			if endCol > len(line) {
				endCol = len(line)
			}
			visibleLine = line[t.offsetCol:endCol]
		}
		
		// Draw the line
		screen.DrawString(x, y+row, visibleLine, textStyle)
		
		// Fill the rest of the row
		remainingWidth := t.width - len(visibleLine)
		if remainingWidth > 0 {
			emptyStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
			screen.DrawString(x+len(visibleLine), y+row, strings.Repeat(" ", remainingWidth), emptyStyle)
		}
		
		// Draw cursor if on this line and focused
		if t.focused && lineIndex == t.cursorRow {
			cursorScreenCol := t.cursorCol - t.offsetCol
			if cursorScreenCol >= 0 && cursorScreenCol < t.width {
				cursorStyle := lipgloss.NewStyle().
					Foreground(theme.Palette.Background).
					Background(theme.Palette.Text)
				
				// Get character under cursor
				var cursorChar rune = ' '
				if t.cursorCol < len(line) {
					cursorChar = rune(line[t.cursorCol])
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