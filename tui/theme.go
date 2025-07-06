package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/themes"
	"github.com/johnnyfreeman/tint/tui/types"
)

// Re-export types from types package for backward compatibility
type (
	StateColors      = types.StateColors
	InteractiveStyle = types.InteractiveStyle
	FocusableStyle   = types.FocusableStyle
	TabStyle         = types.TabStyle
	ContainerStyle   = types.ContainerStyle
	Components       = types.Components
	Palette          = types.Palette
	Theme            = types.Theme
)

// NotificationStyle defines colors for a notification type
type NotificationStyle struct {
	Icon   string
	Border lipgloss.TerminalColor
	Title  lipgloss.TerminalColor
	Text   lipgloss.TerminalColor
}

// NotificationStyles provides semantic styles for different notification types
var NotificationStyles = struct {
	Success NotificationStyle
	Warning NotificationStyle
	Error   NotificationStyle
	Info    NotificationStyle
}{
	Success: NotificationStyle{
		Icon:   "✓",
		Border: lipgloss.Color("#9ece6a"), // Green
		Title:  lipgloss.Color("#9ece6a"),
		Text:   lipgloss.Color("#c0caf5"),
	},
	Warning: NotificationStyle{
		Icon:   "⚠",
		Border: lipgloss.Color("#e0af68"), // Yellow
		Title:  lipgloss.Color("#e0af68"),
		Text:   lipgloss.Color("#c0caf5"),
	},
	Error: NotificationStyle{
		Icon:   "✗",
		Border: lipgloss.Color("#f7768e"), // Red
		Title:  lipgloss.Color("#f7768e"),
		Text:   lipgloss.Color("#c0caf5"),
	},
	Info: NotificationStyle{
		Icon:   "ℹ",
		Border: lipgloss.Color("#7aa2f7"), // Blue
		Title:  lipgloss.Color("#7aa2f7"),
		Text:   lipgloss.Color("#c0caf5"),
	},
}

// Themes provides access to all available themes for backward compatibility
var Themes = map[string]Theme{
	"tokyonight": themes.GetTheme("tokyonight"),
	"rosepine":   themes.GetTheme("rosepine"),
	"catppuccin": themes.GetTheme("catppuccin"),
	"monochrome": themes.GetTheme("monochrome"),
}

// DefaultTheme is the default theme used when no theme is specified
var DefaultTheme = themes.GetTheme("monochrome")

// GetTheme returns a theme by name using the themes registry
func GetTheme(name string) Theme {
	return themes.GetTheme(name)
}

// GetAvailableThemes returns a list of all available theme names
func GetAvailableThemes() []string {
	return themes.GetAvailableThemes()
}

// ThemeExists checks if a theme with the given name exists
func ThemeExists(name string) bool {
	return themes.ThemeExists(name)
}
