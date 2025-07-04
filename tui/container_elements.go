package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// InlineElement represents an element that can be drawn inline within a container border
type InlineElement interface {
	// Draw renders the element and returns the width it consumed
	Draw(screen *Screen, x, y int, theme *Theme, focused bool) int
	// Width returns the width this element will consume
	Width() int
}

// BorderPosition represents where in the border an element is placed
type BorderPosition string

const (
	BorderTop    BorderPosition = "top"
	BorderBottom BorderPosition = "bottom"
	BorderLeft   BorderPosition = "left"
	BorderRight  BorderPosition = "right"
)

// BorderAlignment represents how elements are aligned within a border
type BorderAlignment string

const (
	BorderAlignLeft   BorderAlignment = "left"
	BorderAlignCenter BorderAlignment = "center"
	BorderAlignRight  BorderAlignment = "right"
)

// BorderElement wraps an InlineElement with positioning information
type BorderElement struct {
	Element   InlineElement
	Position  BorderPosition
	Alignment BorderAlignment
	Offset    int // Optional offset from the alignment edge
}

// TextElement is a simple text element
type TextElement struct {
	text  string
	style lipgloss.Style
}

// NewTextElement creates a new text element
func NewTextElement(text string) *TextElement {
	return &TextElement{
		text:  text,
		style: lipgloss.NewStyle(),
	}
}

// SetStyle sets the style for the text element
func (t *TextElement) SetStyle(style lipgloss.Style) {
	t.style = style
}

// Draw renders the text element
func (t *TextElement) Draw(screen *Screen, x, y int, theme *Theme, focused bool) int {
	// Add spaces around text for better appearance
	text := " " + t.text + " "
	
	// Apply default style with theme colors
	style := t.style
	if focused {
		style = style.Foreground(theme.Palette.Text).Bold(true)
	} else {
		style = style.Foreground(theme.Palette.Text)
	}
	
	screen.DrawString(x, y, text, style)
	return StringWidth(text)
}

// Width returns the width this element will consume
func (t *TextElement) Width() int {
	return StringWidth(t.text) + 2 // +2 for padding spaces
}

// TabsElement represents tabs that can be embedded in a border
type TabsElement struct {
	tabs       []string
	activeTab  int
	style      lipgloss.Style
	activeStyle lipgloss.Style
}

// NewTabsElement creates a new tabs element
func NewTabsElement(tabs []string) *TabsElement {
	return &TabsElement{
		tabs:      tabs,
		activeTab: 0,
		style:     lipgloss.NewStyle(),
		activeStyle: lipgloss.NewStyle(),
	}
}

// SetActiveTab sets the active tab index
func (t *TabsElement) SetActiveTab(index int) {
	if index >= 0 && index < len(t.tabs) {
		t.activeTab = index
	}
}

// Draw renders the tabs element
func (t *TabsElement) Draw(screen *Screen, x, y int, theme *Theme, focused bool) int {
	currentX := x
	
	for i, tab := range t.tabs {
		// Add separator before tab (except first)
		if i > 0 {
			sepStyle := lipgloss.NewStyle().Foreground(theme.Palette.Border)
			screen.DrawRune(currentX, y, '┃', sepStyle)
			currentX++
		}
		
		// Prepare tab text
		tabText := " " + tab + " "
		
		// Determine style
		var style lipgloss.Style
		if i == t.activeTab {
			if focused {
				style = lipgloss.NewStyle().
					Foreground(theme.Palette.Background).
					Background(theme.Palette.Primary).
					Bold(true)
			} else {
				style = lipgloss.NewStyle().
					Foreground(theme.Palette.Text).
					Bold(true)
			}
		} else {
			style = lipgloss.NewStyle().
				Foreground(theme.Palette.TextMuted)
		}
		
		// Draw tab
		screen.DrawString(currentX, y, tabText, style)
		currentX += StringWidth(tabText)
	}
	
	return currentX - x
}

// Width returns the width this element will consume
func (t *TabsElement) Width() int {
	width := 0
	for i, tab := range t.tabs {
		if i > 0 {
			width++ // separator
		}
		width += StringWidth(tab) + 2 // tab text with padding
	}
	return width
}

// IconElement represents an icon that can be embedded in a border
type IconElement struct {
	icon  rune
	style lipgloss.Style
}

// NewIconElement creates a new icon element
func NewIconElement(icon rune) *IconElement {
	return &IconElement{
		icon:  icon,
		style: lipgloss.NewStyle(),
	}
}

// Draw renders the icon element
func (i *IconElement) Draw(screen *Screen, x, y int, theme *Theme, focused bool) int {
	style := i.style.Foreground(theme.Palette.Primary)
	
	// Draw icon with spaces for padding
	screen.DrawRune(x, y, ' ', style)
	screen.DrawRune(x+1, y, i.icon, style)
	screen.DrawRune(x+2, y, ' ', style)
	return 3
}

// Width returns the width this element will consume
func (i *IconElement) Width() int {
	return 3 // icon + padding
}

// SpacerElement represents a flexible spacer
type SpacerElement struct {
	minWidth int
}

// NewSpacerElement creates a new spacer element
func NewSpacerElement(minWidth int) *SpacerElement {
	return &SpacerElement{minWidth: minWidth}
}

// Draw renders the spacer (does nothing)
func (s *SpacerElement) Draw(screen *Screen, x, y int, theme *Theme, focused bool) int {
	// Spacers don't draw anything
	return s.minWidth
}

// Width returns the minimum width for the spacer
func (s *SpacerElement) Width() int {
	return s.minWidth
}

// StatusElement represents a status indicator
type StatusElement struct {
	status string
	style  lipgloss.Style
}

// NewStatusElement creates a new status element
func NewStatusElement(status string) *StatusElement {
	return &StatusElement{
		status: status,
		style:  lipgloss.NewStyle(),
	}
}

// Draw renders the status element
func (s *StatusElement) Draw(screen *Screen, x, y int, theme *Theme, focused bool) int {
	text := " [" + s.status + "] "
	
	// Color based on status
	var style lipgloss.Style
	switch s.status {
	case "OK", "Ready", "Connected":
		style = s.style.Foreground(theme.Palette.Pine)
	case "Error", "Failed", "Disconnected":
		style = s.style.Foreground(theme.Palette.Love)
	case "Warning", "Pending":
		style = s.style.Foreground(theme.Palette.Gold)
	default:
		style = s.style.Foreground(theme.Palette.TextMuted)
	}
	
	screen.DrawString(x, y, text, style)
	return StringWidth(text)
}

// Width returns the width this element will consume
func (s *StatusElement) Width() int {
	return StringWidth(s.status) + 4 // status + "[ ]" + spaces
}

// BadgeElement represents a badge or count indicator
type BadgeElement struct {
	text  string
	style lipgloss.Style
}

// NewBadgeElement creates a new badge element
func NewBadgeElement(text string) *BadgeElement {
	return &BadgeElement{
		text:  text,
		style: lipgloss.NewStyle(),
	}
}

// Draw renders the badge element
func (b *BadgeElement) Draw(screen *Screen, x, y int, theme *Theme, focused bool) int {
	// Format as a badge
	text := " (" + b.text + ") "
	
	style := b.style.Foreground(theme.Palette.Primary).Bold(true)
	
	screen.DrawString(x, y, text, style)
	return StringWidth(text)
}

// Width returns the width this element will consume
func (b *BadgeElement) Width() int {
	return StringWidth(b.text) + 4 // text + "()" + spaces
}