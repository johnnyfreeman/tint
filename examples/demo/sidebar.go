package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type Sidebar struct {
	visible  bool
	width    int
	items    []string
	selected int
	hovered  int // For keyboard navigation (not yet selected)
}

func NewSidebar() *Sidebar {
	return &Sidebar{
		visible: true,
		width:   20,
		items: []string{
			"Dashboard",
			"Files",
			"Settings",
			"About",
		},
		selected: 0,
		hovered:  0,
	}
}

func (s *Sidebar) DrawWithTheme(screen *tui.Screen, x, y, height int, theme tui.Theme, focused bool) {
	if !s.visible {
		return
	}

	// Get appropriate container styles based on focus
	var containerStyle tui.FocusableStyle
	if focused {
		containerStyle = tui.FocusableStyle{
			Focused:   theme.Components.Container.Border.Focused,
			Unfocused: theme.Components.Container.Border.Focused,
		}
	} else {
		containerStyle = tui.FocusableStyle{
			Focused:   theme.Components.Container.Border.Unfocused,
			Unfocused: theme.Components.Container.Border.Unfocused,
		}
	}

	borderStyle := lipgloss.NewStyle().
		Foreground(containerStyle.Unfocused.Border).
		Background(theme.Palette.Background)
	
	// Title style based on focus
	var titleColors tui.StateColors
	if focused {
		titleColors = theme.Components.Container.Title.Focused
	} else {
		titleColors = theme.Components.Container.Title.Unfocused
	}
	titleStyle := lipgloss.NewStyle().
		Foreground(titleColors.Text).
		Background(theme.Palette.Background)

	// Fill background first
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < s.width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}
	
	// Draw box with title
	screen.DrawBoxWithTitle(x, y, s.width, height, "Sidebar", borderStyle, titleStyle)

	// Normal item style from theme
	normalItemStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Background)

	// Draw items
	itemY := y + 2 // Start after border and spacing
	for i, item := range s.items {
		if itemY >= y+height-1 {
			break // Don't draw past bottom border
		}

		// Clear the line first (inside the box) with background color
		clearStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.Text).
			Background(theme.Palette.Background)
		for dx := x + 1; dx < x + s.width - 1; dx++ {
			screen.DrawRune(dx, itemY, ' ', clearStyle)
		}

		// Determine item state and get appropriate style
		var itemStyle lipgloss.Style
		var prefix string
		
		if i == s.selected {
			// Selected item
			if focused {
				// Sidebar is focused, use interactive selected style
				colors := theme.Components.Interactive.Selected
				itemStyle = lipgloss.NewStyle().
					Foreground(colors.Text).
					Background(theme.Palette.Background).
					Bold(true)
			} else {
				// Sidebar not focused, use dimmer version
				itemStyle = lipgloss.NewStyle().
					Foreground(theme.Palette.TextMuted).
					Background(theme.Palette.Background).
					Bold(true)
			}
			prefix = "▸ "
		} else if focused && i == s.hovered {
			// Hovered item (only when sidebar is focused)
			colors := theme.Components.Interactive.Hover
			itemStyle = lipgloss.NewStyle().
				Foreground(colors.Text).
				Background(theme.Palette.Background)
			prefix = "  "
		} else {
			// Normal item
			itemStyle = normalItemStyle
			prefix = "  "
		}

		// Draw item
		itemX := x + 2
		screen.DrawString(itemX, itemY, prefix, itemStyle)
		screen.DrawString(itemX+2, itemY, item, itemStyle)
		
		itemY++
	}
}

func (s *Sidebar) Toggle() {
	s.visible = !s.visible
}

func (s *Sidebar) IsVisible() bool {
	return s.visible
}

func (s *Sidebar) Width() int {
	if s.visible {
		return s.width
	}
	return 0
}

func (s *Sidebar) MoveUp() {
	if s.selected > 0 {
		s.selected--
		s.hovered = s.selected
	}
}

func (s *Sidebar) MoveDown() {
	if s.selected < len(s.items)-1 {
		s.selected++
		s.hovered = s.selected
	}
}

func (s *Sidebar) HoverUp() {
	if s.hovered > 0 {
		s.hovered--
	}
}

func (s *Sidebar) HoverDown() {
	if s.hovered < len(s.items)-1 {
		s.hovered++
	}
}

func (s *Sidebar) Select() {
	s.selected = s.hovered
}