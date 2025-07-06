package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// TokyoNight creates a Tokyo Night theme using colors from the glamour package
func TokyoNight() types.Theme {
	// Tokyo Night color definitions (extracted from glamour/styles)
	bg := lipgloss.Color("#1a1b26")
	bgDark := lipgloss.Color("#16161e")
	bgHighlight := lipgloss.Color("#292e42")
	fg := lipgloss.Color("#a9b1d6")
	fgGutter := lipgloss.Color("#363b54")
	dark3 := lipgloss.Color("#545c7e")
	comment := lipgloss.Color("#565f89")
	blue := lipgloss.Color("#7aa2f7")
	cyan := lipgloss.Color("#7dcfff")
	magenta := lipgloss.Color("#bb9af7")
	purple := lipgloss.Color("#9d7cd8")
	orange := lipgloss.Color("#ff9e64")
	yellow := lipgloss.Color("#e0af68")
	green := lipgloss.Color("#9ece6a")
	red := lipgloss.Color("#f7768e")

	return types.Theme{
		Name: "Tokyo Night",
		Palette: types.Palette{
			Background: bg,
			Surface:    bg,
			Overlay:    bgHighlight,
			Text:       fg,
			TextMuted:  comment,
			TextSubtle: dark3,
			Primary:    blue,
			Secondary:  purple,
			Love:       red,
			Gold:       yellow,
			Rose:       orange,
			Pine:       green,
			Foam:       cyan,
			Iris:       magenta,
			Border:     dark3,
			Shadow:     bgDark,
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       fg,
					Background: bg,
					Border:     dark3,
				},
				Hover: types.StateColors{
					Text:       cyan,
					Background: bg,
					Border:     cyan,
				},
				Selected: types.StateColors{
					Text:       blue,
					Background: bg,
					Border:     blue,
				},
				Disabled: types.StateColors{
					Text:       dark3,
					Background: bg,
					Border:     fgGutter,
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       comment,
					Background: bg,
					Border:     dark3,
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       blue,
						Background: bg,
						Border:     blue,
					},
					Unfocused: types.StateColors{
						Text:       fg,
						Background: bg,
						Border:     dark3,
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       blue,
						Background: bg,
						Border:     blue,
					},
					Unfocused: types.StateColors{
						Text:       fg,
						Background: bg,
						Border:     dark3,
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       blue,
						Background: bg,
						Border:     bg,
					},
					Unfocused: types.StateColors{
						Text:       comment,
						Background: bg,
						Border:     bg,
					},
				},
			},
		},
	}
}