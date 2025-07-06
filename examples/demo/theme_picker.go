package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type ThemePicker struct {
	visible      bool
	themes       []string
	selected     int
	hovered      int
	width        int
	height       int
	previewTheme string // Theme being previewed on hover
}

func NewThemePicker() *ThemePicker {
	themes := tui.GetAvailableThemes()
	// Find monochrome index for default
	monochromeIndex := 0
	for i, theme := range themes {
		if theme == "monochrome" {
			monochromeIndex = i
			break
		}
	}
	
	return &ThemePicker{
		visible:  false,
		themes:   themes,
		selected: monochromeIndex,
		hovered:  monochromeIndex,
		width:    45, // Wider to accommodate longer theme names
		height:   len(themes) + 4, // Dynamic height based on theme count
	}
}

func (tp *ThemePicker) DrawWithTheme(screen *tui.Screen, currentTheme *tui.Theme, focused bool) {
	if !tp.visible {
		return
	}

	// Calculate center position
	x := (screen.Width() - tp.width) / 2
	y := (screen.Height() - tp.height) / 2

	// Get container styles
	var borderColors, titleColors tui.StateColors
	if focused {
		borderColors = currentTheme.Components.Container.Border.Focused
		titleColors = currentTheme.Components.Container.Title.Focused
	} else {
		borderColors = currentTheme.Components.Container.Border.Unfocused
		titleColors = currentTheme.Components.Container.Title.Unfocused
	}

	bgStyle := lipgloss.NewStyle().
		Background(currentTheme.Palette.Surface).
		Foreground(currentTheme.Palette.Text)
	borderStyle := lipgloss.NewStyle().
		Foreground(borderColors.Border).
		Background(currentTheme.Palette.Surface)
	titleStyle := lipgloss.NewStyle().
		Foreground(titleColors.Text).
		Background(currentTheme.Palette.Surface).
		Bold(true)

	// Draw block shadow with 1 cell offset (same as modal)
	shadowStyle := lipgloss.NewStyle().
		Foreground(currentTheme.Palette.Shadow).
		Background(currentTheme.Palette.Background)
	shadowOffsetX := 1
	shadowOffsetY := 1
	screen.DrawBlockShadow(x, y, tp.width, tp.height, shadowStyle, shadowOffsetX, shadowOffsetY)

	// Draw background
	for dy := 0; dy < tp.height; dy++ {
		for dx := 0; dx < tp.width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}

	// Draw border with title - use heavy borders when focused
	if focused {
		screen.DrawBrutalistBoxWithTitle(x, y, tp.width, tp.height, "Choose Theme", borderStyle, titleStyle)
	} else {
		screen.DrawBoxWithTitle(x, y, tp.width, tp.height, "Choose Theme", borderStyle, titleStyle)
	}

	// Draw theme options with color swatches
	for i, themeName := range tp.themes {
		theme := tui.GetTheme(themeName)
		lineY := y + 2 + i

		// Clear line
		for dx := x + 1; dx < x+tp.width-1; dx++ {
			screen.DrawRune(dx, lineY, ' ', bgStyle)
		}

		// Determine item state and style
		itemX := x + 2
		prefix := "  "
		var style lipgloss.Style

		if i == tp.selected {
			// Selected theme
			colors := currentTheme.Components.Interactive.Selected
			style = lipgloss.NewStyle().
				Foreground(colors.Text).
				Background(currentTheme.Palette.Surface).
				Bold(true)
			prefix = "◉ " // Filled circle for selected
		} else if i == tp.hovered {
			// Hovered theme
			colors := currentTheme.Components.Interactive.Hover
			style = lipgloss.NewStyle().
				Foreground(colors.Text).
				Background(currentTheme.Palette.Surface)
			prefix = "○ " // Empty circle for hoverable
		} else {
			// Normal theme
			colors := currentTheme.Components.Interactive.Normal
			style = lipgloss.NewStyle().
				Foreground(colors.Text).
				Background(currentTheme.Palette.Surface)
			prefix = "○ " // Empty circle for normal
		}

		screen.DrawString(itemX, lineY, prefix+theme.Name, style)

		// Draw color swatches
		swatchX := x + tp.width - 14
		screen.DrawString(swatchX, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Love).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+2, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Gold).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+4, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Rose).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+6, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Pine).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+8, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Foam).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+10, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Iris).Background(currentTheme.Palette.Surface))
	}
}

func (tp *ThemePicker) Toggle() {
	tp.visible = !tp.visible
	if tp.visible {
		tp.hovered = tp.selected
	}
}

func (tp *ThemePicker) IsVisible() bool {
	return tp.visible
}

func (tp *ThemePicker) MoveUp() {
	if tp.hovered > 0 {
		tp.hovered--
		tp.previewTheme = tp.themes[tp.hovered]
	}
}

func (tp *ThemePicker) MoveDown() {
	if tp.hovered < len(tp.themes)-1 {
		tp.hovered++
		tp.previewTheme = tp.themes[tp.hovered]
	}
}

func (tp *ThemePicker) Select() {
	tp.selected = tp.hovered
	tp.visible = false
}

func (tp *ThemePicker) GetSelectedTheme() string {
	return tp.themes[tp.selected]
}

func (tp *ThemePicker) GetPreviewTheme() string {
	if tp.visible && tp.previewTheme != "" {
		return tp.previewTheme
	}
	return tp.GetSelectedTheme()
}
