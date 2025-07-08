package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// Monochrome creates a monochrome theme for terminals with limited color support
func Monochrome() types.Theme {
	// Monochrome color definitions using CompleteColor for better compatibility
	white := lipgloss.CompleteColor{TrueColor: "#ffffff", ANSI256: "255", ANSI: "7"}
	lightGray := lipgloss.CompleteColor{TrueColor: "#bcbcbc", ANSI256: "250", ANSI: "7"}
	gray := lipgloss.CompleteColor{TrueColor: "#585858", ANSI256: "240", ANSI: "8"}
	darkGray := lipgloss.CompleteColor{TrueColor: "#262626", ANSI256: "235", ANSI: "0"}
	veryDarkGray := lipgloss.CompleteColor{TrueColor: "#1c1c1c", ANSI256: "234", ANSI: "0"}
	black := lipgloss.CompleteColor{TrueColor: "#121212", ANSI256: "233", ANSI: "0"}

	return types.Theme{
		Name: "Monochrome",
		Palette: types.Palette{
			Background: veryDarkGray,
			Surface:    darkGray,
			Overlay:    gray,
			Text:       white,
			TextMuted:  gray,
			TextSubtle: gray,
			Primary:    white,
			Secondary:  lightGray,
			Love:       white,
			Gold:       white,
			Rose:       white,
			Pine:       white,
			Foam:       white,
			Iris:       white,
			Border:     gray,
			Shadow:     black,
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       white,
					Background: veryDarkGray,
					Border:     gray,
				},
				Hover: types.StateColors{
					Text:       white,
					Background: veryDarkGray,
					Border:     white,
				},
				Selected: types.StateColors{
					Text:       white,
					Background: veryDarkGray,
					Border:     white,
				},
				Disabled: types.StateColors{
					Text:       gray,
					Background: veryDarkGray,
					Border:     darkGray,
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       gray,
					Background: veryDarkGray,
					Border:     gray,
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     white,
					},
					Unfocused: types.StateColors{
						Text:       lightGray,
						Background: veryDarkGray,
						Border:     gray,
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     white,
					},
					Unfocused: types.StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     gray,
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     veryDarkGray,
					},
					Unfocused: types.StateColors{
						Text:       gray,
						Background: veryDarkGray,
						Border:     veryDarkGray,
					},
				},
			},
		},
	}
}