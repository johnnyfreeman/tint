package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type ThemePicker struct {
	modal        *tui.Modal
	container    *tui.Container
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
	
	// Create modal (Modal → Container → Content pattern)
	width := 45 // Wider to accommodate longer theme names
	height := len(themes) + 4 // Dynamic height based on theme count
	
	modal := tui.NewModal()
	modal.SetSize(width, height)
	modal.SetCentered(true)

	// Create container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Choose Theme")
	container.SetSize(width, height) // Fill the entire modal surface
	container.SetPadding(tui.NewMargin(1))
	container.SetUseSurface(true) // Use surface color for modal
	
	return &ThemePicker{
		modal:        modal,
		container:    container,
		themes:       themes,
		selected:     monochromeIndex,
		hovered:      monochromeIndex,
		width:        width,
		height:       height,
	}
}

func (tp *ThemePicker) DrawWithTheme(screen *tui.Screen, currentTheme *tui.Theme, focused bool) {
	if !tp.modal.IsVisible() {
		return
	}

	// Draw modal surface (provides backdrop and elevation)
	tp.modal.Draw(screen, 0, 0, currentTheme)

	// Get modal position for container placement
	modalWidth, modalHeight := tp.modal.GetSize()
	modalX := (screen.Width() - modalWidth) / 2
	modalY := (screen.Height() - modalHeight) / 2

	// Focus the container when theme picker is focused
	if focused {
		tp.container.Focus()
	} else {
		tp.container.Blur()
	}

	// Draw container filling the entire modal surface
	tp.container.Draw(screen, modalX, modalY, currentTheme)

	// Draw theme options with color swatches inside container
	for i, themeName := range tp.themes {
		theme := tui.GetTheme(themeName)
		lineY := modalY + 2 + i // Position relative to modal

		// Determine item state and style
		itemX := modalX + 2 // Position relative to modal
		prefix := "  "
		var style lipgloss.Style

		if i == tp.selected {
			// Selected theme - highlight full row
			for dx := modalX + 1; dx < modalX+tp.width-1; dx++ {
				screen.SetCell(dx, lineY, tui.Cell{
					Rune:       ' ',
					Background: currentTheme.Palette.Primary,
				})
			}
			colors := currentTheme.Components.Interactive.Selected
			style = lipgloss.NewStyle().
				Foreground(colors.Text).
				Background(currentTheme.Palette.Primary).
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
		swatchX := modalX + tp.width - 14
		screen.DrawString(swatchX, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Love).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+2, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Gold).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+4, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Rose).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+6, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Pine).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+8, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Foam).Background(currentTheme.Palette.Surface))
		screen.DrawString(swatchX+10, lineY, "●", lipgloss.NewStyle().Foreground(theme.Palette.Iris).Background(currentTheme.Palette.Surface))
	}
}

func (tp *ThemePicker) Toggle() {
	tp.modal.Toggle()
	if tp.modal.IsVisible() {
		tp.hovered = tp.selected
		tp.modal.Focus()
		tp.container.Focus()
	} else {
		tp.modal.Blur()
		tp.container.Blur()
	}
}

func (tp *ThemePicker) IsVisible() bool {
	return tp.modal.IsVisible()
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
	tp.modal.Hide()
}

func (tp *ThemePicker) GetSelectedTheme() string {
	return tp.themes[tp.selected]
}

func (tp *ThemePicker) GetPreviewTheme() string {
	if tp.modal.IsVisible() && tp.previewTheme != "" {
		return tp.previewTheme
	}
	return tp.GetSelectedTheme()
}
