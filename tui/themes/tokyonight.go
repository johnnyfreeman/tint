package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// TokyoNight creates a Tokyo Night theme using colors from the glamour package
func TokyoNight() types.Theme {
	// Tokyo Night color definitions (extracted from glamour/styles)
	bg := lipgloss.CompleteColor{TrueColor: "#1a1b26", ANSI256: "235", ANSI: "0"}
	bgDark := lipgloss.CompleteColor{TrueColor: "#16161e", ANSI256: "234", ANSI: "0"}
	bgHighlight := lipgloss.CompleteColor{TrueColor: "#292e42", ANSI256: "237", ANSI: "8"}
	fg := lipgloss.CompleteColor{TrueColor: "#a9b1d6", ANSI256: "146", ANSI: "7"}
	fgGutter := lipgloss.CompleteColor{TrueColor: "#363b54", ANSI256: "238", ANSI: "8"}
	dark3 := lipgloss.CompleteColor{TrueColor: "#545c7e", ANSI256: "243", ANSI: "8"}
	comment := lipgloss.CompleteColor{TrueColor: "#565f89", ANSI256: "244", ANSI: "8"}
	blue := lipgloss.CompleteColor{TrueColor: "#7aa2f7", ANSI256: "75", ANSI: "4"}
	cyan := lipgloss.CompleteColor{TrueColor: "#7dcfff", ANSI256: "123", ANSI: "6"}
	magenta := lipgloss.CompleteColor{TrueColor: "#bb9af7", ANSI256: "177", ANSI: "5"}
	purple := lipgloss.CompleteColor{TrueColor: "#9d7cd8", ANSI256: "141", ANSI: "5"}
	orange := lipgloss.CompleteColor{TrueColor: "#ff9e64", ANSI256: "215", ANSI: "3"}
	yellow := lipgloss.CompleteColor{TrueColor: "#e0af68", ANSI256: "179", ANSI: "3"}
	green := lipgloss.CompleteColor{TrueColor: "#9ece6a", ANSI256: "149", ANSI: "2"}
	red := lipgloss.CompleteColor{TrueColor: "#f7768e", ANSI256: "210", ANSI: "1"}

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