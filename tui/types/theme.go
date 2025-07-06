package types

import "github.com/charmbracelet/lipgloss"

// StateColors represents colors for a specific UI state
type StateColors struct {
	Text       lipgloss.TerminalColor
	Background lipgloss.TerminalColor
	Border     lipgloss.TerminalColor
}

// InteractiveStyle defines styles for interactive elements
type InteractiveStyle struct {
	Normal   StateColors // Default state
	Hover    StateColors // Mouse hover or keyboard navigation
	Selected StateColors // Selected/active state
	Disabled StateColors // Disabled state
}

// FocusableStyle defines styles for elements that can gain/lose focus
type FocusableStyle struct {
	Focused   StateColors
	Unfocused StateColors
}

// TabStyle defines styles specifically for tabs
type TabStyle struct {
	Inactive StateColors    // Inactive tabs
	Active   FocusableStyle // Active tab (can be focused/unfocused)
}

// ContainerStyle defines styles for container elements
type ContainerStyle struct {
	Border FocusableStyle // Border changes on focus
	Title  FocusableStyle // Title changes on focus
}

// Components groups all component-specific styles
type Components struct {
	Interactive InteractiveStyle
	Tab         TabStyle
	Container   ContainerStyle
}

// Palette defines the core color palette
type Palette struct {
	// Base colors
	Background lipgloss.TerminalColor // Terminal default
	Surface    lipgloss.TerminalColor // Elevated surfaces (modals, cards)
	Overlay    lipgloss.TerminalColor // Tooltips, dropdowns

	// Text hierarchy
	Text       lipgloss.TerminalColor // Primary text
	TextMuted  lipgloss.TerminalColor // Secondary text
	TextSubtle lipgloss.TerminalColor // Disabled/tertiary text

	// Brand colors
	Primary   lipgloss.TerminalColor // Main brand color
	Secondary lipgloss.TerminalColor // Secondary brand color

	// Semantic colors (theme's interpretation)
	Love lipgloss.TerminalColor // Red/pink tones
	Gold lipgloss.TerminalColor // Yellow/gold tones
	Rose lipgloss.TerminalColor // Rose/coral tones
	Pine lipgloss.TerminalColor // Green tones
	Foam lipgloss.TerminalColor // Cyan/teal tones
	Iris lipgloss.TerminalColor // Purple/violet tones

	// UI colors
	Border lipgloss.TerminalColor // Default border
	Shadow lipgloss.TerminalColor // Shadow/depth effect
}

// Theme represents the theme structure
type Theme struct {
	Name       string
	Palette    Palette
	Components Components
}