package main

import (
	"github.com/johnnyfreeman/tint/tui"
)

type Sidebar struct {
	container *tui.Container
	viewer    *tui.Viewer
	visible   bool
	width     int
	items     []string
	selected  int
	hovered   int // For keyboard navigation (not yet selected)
}

func NewSidebar() *Sidebar {
	// Create container for sidebar
	container := tui.NewContainer()
	container.SetTitle("Sidebar")
	container.SetSize(20, 0) // Height will be set dynamically
	container.SetPadding(tui.NewMargin(1))
	
	// Create viewer for sidebar content
	viewer := tui.NewViewer()
	container.SetContent(viewer)
	
	return &Sidebar{
		container: container,
		viewer:    viewer,
		visible:   true,
		width:     20,
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

	// Set container size and focus state
	s.container.SetSize(s.width, height)
	if focused {
		s.container.Focus()
	} else {
		s.container.Blur()
	}

	// Build sidebar content with styled items
	content := ""
	for i, item := range s.items {
		var prefix string
		if i == s.selected {
			prefix = "▶ " // Solid right-pointing triangle
		} else {
			prefix = "  "
		}
		
		if i > 0 {
			content += "\n"
		}
		content += prefix + item
	}
	
	// Update viewer content
	s.viewer.SetContent(content)
	
	// Draw the container (which handles borders, focus, etc.)
	s.container.Draw(screen, x, y, s.width, height, &theme)
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
