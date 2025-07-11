package tui

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// Viewer represents a read-only scrollable text viewer
type Viewer struct {
	content      string
	lines        []string
	scrollOffset int
	width        int
	height       int
	focused      bool
	wrapText     bool
}

// NewViewer creates a new viewer
func NewViewer() *Viewer {
	return &Viewer{
		content:      "",
		lines:        []string{},
		scrollOffset: 0,
		width:        40,
		height:       10,
		focused:      false,
		wrapText:     true,
	}
}

// SetSize sets the display dimensions
func (v *Viewer) SetSize(width, height int) {
	v.width = width
	v.height = height
	v.processContent()
}

// SetContent sets the content to display
func (v *Viewer) SetContent(content string) {
	v.content = content
	v.scrollOffset = 0
	v.processContent()
}

// SetWrapText enables or disables text wrapping
func (v *Viewer) SetWrapText(wrap bool) {
	v.wrapText = wrap
	v.processContent()
}

// Focus sets the focus state
func (v *Viewer) Focus() {
	v.focused = true
}

// Blur removes focus
func (v *Viewer) Blur() {
	v.focused = false
}

// IsFocused returns whether the viewer is focused
func (v *Viewer) IsFocused() bool {
	return v.focused
}

// processContent splits content into displayable lines
func (v *Viewer) processContent() {
	if v.content == "" {
		v.lines = []string{}
		return
	}

	rawLines := strings.Split(v.content, "\n")
	v.lines = []string{}

	if !v.wrapText {
		v.lines = rawLines
		return
	}

	// Wrap lines that are too long
	for _, line := range rawLines {
		if StringWidth(line) <= v.width {
			v.lines = append(v.lines, line)
		} else {
			// Use the unicode-aware Wrap function
			wrappedLines := Wrap(line, v.width)
			v.lines = append(v.lines, wrappedLines...)
		}
	}
}

// HandleInput processes keyboard input
func (v *Viewer) HandleInput(key string) {
	switch key {
	case "up", "k":
		if v.scrollOffset > 0 {
			v.scrollOffset--
		}
	case "down", "j":
		maxScroll := len(v.lines) - v.height
		if maxScroll < 0 {
			maxScroll = 0
		}
		if v.scrollOffset < maxScroll {
			v.scrollOffset++
		}
	case "pgup":
		v.scrollOffset -= v.height - 1
		if v.scrollOffset < 0 {
			v.scrollOffset = 0
		}
	case "pgdown":
		v.scrollOffset += v.height - 1
		maxScroll := len(v.lines) - v.height
		if maxScroll < 0 {
			maxScroll = 0
		}
		if v.scrollOffset > maxScroll {
			v.scrollOffset = maxScroll
		}
	case "home", "ctrl+home":
		v.scrollOffset = 0
	case "end", "ctrl+end":
		v.scrollOffset = len(v.lines) - v.height
		if v.scrollOffset < 0 {
			v.scrollOffset = 0
		}
	}
}

// Draw renders the viewer to the screen
func (v *Viewer) Draw(screen *Screen, x, y, availableWidth, availableHeight int, theme *Theme) {
	// Viewer decides to use available space for content display
	viewerWidth := availableWidth
	viewerHeight := availableHeight
	
	// Update internal dimensions
	v.width = viewerWidth
	v.height = viewerHeight
	v.processContent()
	
	// Clear the entire viewer area with theme background
	ClearComponentArea(screen, x, y, viewerWidth, viewerHeight, theme)

	textStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Background)

	// Draw visible lines
	for row := 0; row < viewerHeight; row++ {
		lineIndex := v.scrollOffset + row

		if lineIndex >= len(v.lines) {
			// Empty rows are already cleared by ClearComponentArea
			continue
		}

		line := v.lines[lineIndex]

		// Truncate if line is still too long (shouldn't happen with wrapping)
		displayLine := line
		if StringWidth(displayLine) > viewerWidth {
			displayLine = TruncateWithEllipsis(displayLine, viewerWidth)
		}

		// Draw the line
		screen.DrawString(x, y+row, displayLine, textStyle)
	}

	// Draw scroll indicators
	if v.focused && len(v.lines) > viewerHeight {
		scrollStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background)

		// Scroll bar track
		scrollBarX := x + viewerWidth - 1
		for i := 0; i < viewerHeight; i++ {
			screen.DrawRune(scrollBarX, y+i, '│', scrollStyle)
		}

		// Scroll bar thumb
		// Calculate thumb height as a proportion of visible content
		thumbHeight := int(float64(viewerHeight) * float64(viewerHeight) / float64(len(v.lines)))
		if thumbHeight < 1 {
			thumbHeight = 1
		}
		if thumbHeight > viewerHeight {
			thumbHeight = viewerHeight
		}

		// Calculate thumb position more accurately
		var thumbPos int
		maxScroll := len(v.lines) - viewerHeight
		if maxScroll <= 0 {
			thumbPos = 0
		} else {
			// Use floating point for accurate position calculation
			scrollRatio := float64(v.scrollOffset) / float64(maxScroll)
			maxThumbPos := viewerHeight - thumbHeight
			thumbPos = int(scrollRatio * float64(maxThumbPos))
		}

		thumbStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.Primary).
			Background(theme.Palette.Background)

		for i := 0; i < thumbHeight; i++ {
			if thumbPos+i < v.height {
				screen.DrawRune(scrollBarX, y+thumbPos+i, '█', thumbStyle)
			}
		}
	}
}

// DrawInBox renders the viewer inside a box
func (v *Viewer) DrawInBox(screen *Screen, x, y, width, height int, theme *Theme) {
	// Fill background
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}

	// No box border for viewer - it will be drawn by parent
	// Set viewer size to use full area minus padding
	v.SetSize(width-4, height-2) // Leave room for padding

	// Draw the viewer
	v.Draw(screen, x+2, y+1, width-4, height-2, theme)
}

// GetLineCount returns the total number of lines
func (v *Viewer) GetLineCount() int {
	return len(v.lines)
}

// GetVisibleLines returns the number of lines that can be displayed
func (v *Viewer) GetVisibleLines() int {
	return v.height
}

// IsScrollable returns whether the content is scrollable
func (v *Viewer) IsScrollable() bool {
	return len(v.lines) > v.height
}

// HandleKey processes keyboard input when focused
func (v *Viewer) HandleKey(key string) bool {
	if !v.focused {
		return false
	}
	v.HandleInput(key)
	return true
}

// GetSize returns the current width and height
func (v *Viewer) GetSize() (width, height int) {
	return v.width, v.height
}

// DrawWithBorder draws the component with a border and optional title
func (v *Viewer) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	v.DrawInBox(screen, x, y, v.width+4, v.height+2, theme)
}
