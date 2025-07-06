package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// Monochrome creates a monochrome theme for terminals with limited color support
func Monochrome() types.Theme {
	// Monochrome color definitions
	white := lipgloss.Color("255")
	lightGray := lipgloss.Color("250")
	gray := lipgloss.Color("240")
	darkGray := lipgloss.Color("235")
	veryDarkGray := lipgloss.Color("234")
	black := lipgloss.Color("233")

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