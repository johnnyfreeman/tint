package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// StatusBarSegment represents a single segment in the status bar
type StatusBarSegment struct {
	Text      string
	Alignment string // "left", "center", "right"
	Style     lipgloss.Style
}

// StatusBar represents a status bar component typically shown at the bottom
type StatusBar struct {
	segments []StatusBarSegment
	height   int
	style    lipgloss.Style
}

// NewStatusBar creates a new status bar
func NewStatusBar() *StatusBar {
	return &StatusBar{
		segments: []StatusBarSegment{},
		height:   1,
	}
}

// AddSegment adds a segment to the status bar
func (s *StatusBar) AddSegment(text, alignment string) {
	s.segments = append(s.segments, StatusBarSegment{
		Text:      text,
		Alignment: alignment,
	})
}

// SetSegments replaces all segments
func (s *StatusBar) SetSegments(segments []StatusBarSegment) {
	s.segments = segments
}

// Clear removes all segments
func (s *StatusBar) Clear() {
	s.segments = []StatusBarSegment{}
}

// Draw renders the status bar to the screen
func (s *StatusBar) Draw(screen *Screen, x, y int, theme *Theme) {
	width := screen.Width()
	
	// Default style if not set
	bgStyle := lipgloss.NewStyle().
		Background(theme.Palette.Surface).
		Foreground(theme.Palette.TextMuted)
	
	// Clear the status bar area
	for i := 0; i < width; i++ {
		screen.SetCell(x+i, y, Cell{
			Rune:       ' ',
			Background: theme.Palette.Surface,
		})
	}
	
	// Group segments by alignment
	var leftSegments, centerSegments, rightSegments []StatusBarSegment
	for _, segment := range s.segments {
		switch segment.Alignment {
		case "center":
			centerSegments = append(centerSegments, segment)
		case "right":
			rightSegments = append(rightSegments, segment)
		default: // "left" or unspecified
			leftSegments = append(leftSegments, segment)
		}
	}
	
	// Draw left-aligned segments
	currentX := x
	for i, segment := range leftSegments {
		if i > 0 {
			screen.DrawString(currentX, y, " | ", bgStyle)
			currentX += 3
		}
		style := bgStyle
		if segment.Style.GetForeground() != nil || segment.Style.GetBackground() != nil {
			style = segment.Style
		}
		screen.DrawString(currentX, y, segment.Text, style)
		currentX += len(segment.Text)
	}
	
	// Draw center-aligned segments
	if len(centerSegments) > 0 {
		centerText := ""
		for i, segment := range centerSegments {
			if i > 0 {
				centerText += " | "
			}
			centerText += segment.Text
		}
		centerX := x + (width-len(centerText))/2
		screen.DrawString(centerX, y, centerText, bgStyle)
	}
	
	// Draw right-aligned segments
	if len(rightSegments) > 0 {
		rightText := ""
		for i, segment := range rightSegments {
			if i > 0 {
				rightText += " | "
			}
			rightText += segment.Text
		}
		rightX := x + width - len(rightText) - 1
		screen.DrawString(rightX, y, rightText, bgStyle)
	}
}

// SetHeight sets the height of the status bar (usually 1)
func (s *StatusBar) SetHeight(height int) {
	s.height = height
}

// GetHeight returns the height of the status bar
func (s *StatusBar) GetHeight() int {
	return s.height
}

// Helper functions for common status bar patterns

// NewHelpStatusBar creates a status bar with common help text
func NewHelpStatusBar(shortcuts map[string]string) *StatusBar {
	sb := NewStatusBar()
	
	// Build help text from shortcuts map
	var parts []string
	for key, action := range shortcuts {
		parts = append(parts, key+":"+action)
	}
	helpText := strings.Join(parts, " ")
	
	sb.AddSegment(helpText, "right")
	return sb
}

// NewFileStatusBar creates a status bar for file editing
func NewFileStatusBar(filename, mode string, line, col int) *StatusBar {
	sb := NewStatusBar()
	
	// Left: mode and filename
	sb.AddSegment(mode+" | "+filename, "left")
	
	// Center: cursor position
	cursorText := fmt.Sprintf("Ln %d, Col %d", line, col)
	sb.AddSegment(cursorText, "center")
	
	// Right: encoding and settings
	sb.AddSegment("UTF-8 | Spaces: 4", "right")
	
	return sb
}