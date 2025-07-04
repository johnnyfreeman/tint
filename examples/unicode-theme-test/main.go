package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/example/tint/tui"
)

func main() {
	// Create a small screen to test theme picker rendering
	theme := tui.GetTheme("monochrome")
	screen := tui.NewScreen(50, 15, theme)
	
	// Simulate theme picker states
	themes := []string{"tokyonight", "rosepine", "catppuccin", "monochrome"}
	selected := 1 // rosepine selected
	
	// Draw a mock theme picker
	x, y := 5, 2
	width, height := 40, 10
	
	// Draw border
	borderStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Primary)
	screen.DrawBoxWithTitle(x, y, width, height, "🎨 Choose Theme", borderStyle, borderStyle)
	
	// Draw theme options with unicode selectors
	for i, themeName := range themes {
		themeObj := tui.GetTheme(themeName)
		lineY := y + 2 + i
		itemX := x + 2
		
		var prefix string
		var style lipgloss.Style
		
		if i == selected {
			// Selected with filled circle
			prefix = "◉ "
			style = lipgloss.NewStyle().
				Foreground(theme.Palette.Primary).
				Bold(true)
		} else {
			// Unselected with empty circle
			prefix = "○ "
			style = lipgloss.NewStyle().
				Foreground(theme.Palette.Text)
		}
		
		// Draw selector and theme name
		screen.DrawString(itemX, lineY, prefix+themeObj.Name, style)
		
		// Draw color swatches
		swatchX := itemX + 20
		colors := []lipgloss.TerminalColor{
			themeObj.Palette.Primary,
			themeObj.Palette.Secondary,
			themeObj.Palette.Pine,
			themeObj.Palette.Gold,
		}
		for j, color := range colors {
			swatchStyle := lipgloss.NewStyle().
				Foreground(color).
				Background(theme.Palette.Background)
			screen.DrawString(swatchX+j*2, lineY, "█", swatchStyle)
		}
	}
	
	// Render
	fmt.Print("\033[2J\033[H") // Clear screen
	fmt.Println("Unicode Theme Picker Test:\n")
	fmt.Print(screen.Render())
	fmt.Println("\n\nThe theme picker now uses:")
	fmt.Println("  ◉ = Selected theme")
	fmt.Println("  ○ = Available themes")
	fmt.Println("  ▶ = Selected sidebar item")
}