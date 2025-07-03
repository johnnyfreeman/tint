package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Modal represents a modal dialog component
type Modal struct {
	visible  bool
	width    int
	height   int
	title    string
	content  string
	focused  bool
	centered bool
	x, y     int // Position (used when not centered)
}

// NewModal creates a new modal dialog
func NewModal() *Modal {
	return &Modal{
		visible:  false,
		width:    40,
		height:   10,
		title:    "",
		content:  "",
		focused:  false,
		centered: true,
	}
}

// SetTitle sets the modal title
func (m *Modal) SetTitle(title string) {
	m.title = title
}

// SetContent sets the modal content
func (m *Modal) SetContent(content string) {
	m.content = content
}

// Show makes the modal visible
func (m *Modal) Show() {
	m.visible = true
}

// Hide makes the modal invisible
func (m *Modal) Hide() {
	m.visible = false
}

// Toggle switches the modal visibility
func (m *Modal) Toggle() {
	m.visible = !m.visible
}

// IsVisible returns whether the modal is visible
func (m *Modal) IsVisible() bool {
	return m.visible
}

// SetCentered sets whether the modal should be centered on screen
func (m *Modal) SetCentered(centered bool) {
	m.centered = centered
}

// SetPosition sets the modal position (when not centered)
func (m *Modal) SetPosition(x, y int) {
	m.x = x
	m.y = y
	m.centered = false
}

// Draw renders the modal to the screen at the specified position
func (m *Modal) Draw(screen *Screen, x, y int, theme *Theme) {
	if !m.visible {
		return
	}

	// Calculate actual position
	actualX, actualY := x, y
	if m.centered {
		actualX = (screen.Width() - m.width) / 2
		actualY = (screen.Height() - m.height) / 2
	} else if x == 0 && y == 0 {
		actualX, actualY = m.x, m.y
	}

	// Draw block shadow with neo-brutalist style (offset by 1 cell)
	shadowStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Shadow).
		Background(theme.Palette.Background)
	shadowOffsetX := 1
	shadowOffsetY := 1
	
	// Use the new DrawBlockShadow method
	screen.DrawBlockShadow(actualX, actualY, m.width, m.height, shadowStyle, shadowOffsetX, shadowOffsetY)

	// Modal background uses surface color
	bgStyle := lipgloss.NewStyle().
		Background(theme.Palette.Surface).
		Foreground(theme.Palette.Text)

	// Get container styles for border and title
	var borderColors, titleColors StateColors
	if m.focused {
		borderColors = theme.Components.Container.Border.Focused
		titleColors = theme.Components.Container.Title.Focused
	} else {
		borderColors = theme.Components.Container.Border.Unfocused
		titleColors = theme.Components.Container.Title.Unfocused
	}

	borderStyle := lipgloss.NewStyle().
		Foreground(borderColors.Border).
		Background(theme.Palette.Surface)
	titleStyle := lipgloss.NewStyle().
		Foreground(titleColors.Text).
		Background(theme.Palette.Surface)

	// Draw background fill
	for dy := 0; dy < m.height; dy++ {
		for dx := 0; dx < m.width; dx++ {
			screen.DrawRune(actualX+dx, actualY+dy, ' ', bgStyle)
		}
	}

	// Draw border with title - use heavy borders when focused
	if m.focused {
		screen.DrawBrutalistBoxWithTitle(actualX, actualY, m.width, m.height, m.title, borderStyle, titleStyle)
	} else {
		screen.DrawBoxWithTitle(actualX, actualY, m.width, m.height, m.title, borderStyle, titleStyle)
	}

	// Draw content
	contentStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Surface)
	lines := strings.Split(m.content, "\n")
	contentY := actualY + 1
	maxLines := m.height - 2
	
	for i := 0; i < maxLines && i < len(lines); i++ {
		line := lines[i]
		maxWidth := m.width - 4
		if len(line) > maxWidth {
			line = line[:maxWidth]
		}
		// Center align the content
		padding := (maxWidth - len(line)) / 2
		paddedLine := strings.Repeat(" ", padding) + line
		screen.DrawString(actualX+2, contentY+i, paddedLine, contentStyle)
	}
}

// Focus gives keyboard focus to this component
func (m *Modal) Focus() {
	m.focused = true
}

// Blur removes keyboard focus from this component
func (m *Modal) Blur() {
	m.focused = false
}

// IsFocused returns whether this component currently has focus
func (m *Modal) IsFocused() bool {
	return m.focused
}

// HandleKey processes keyboard input when focused
func (m *Modal) HandleKey(key string) bool {
	switch key {
	case "esc", "enter":
		m.Hide()
		return true
	}
	return false
}

// SetSize sets the width and height of the component
func (m *Modal) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// GetSize returns the current width and height
func (m *Modal) GetSize() (width, height int) {
	return m.width, m.height
}

// DrawWithBorder draws the component with a border and optional title
func (m *Modal) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	if title != "" {
		m.title = title
	}
	m.Draw(screen, x, y, theme)
}