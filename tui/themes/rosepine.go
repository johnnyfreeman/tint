package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// RosePine creates the standard Rosé Pine theme
func RosePine() types.Theme {
	// Rosé Pine color definitions (official colors)
	base := lipgloss.CompleteColor{TrueColor: "#191724", ANSI256: "235", ANSI: "0"}
	surface := lipgloss.CompleteColor{TrueColor: "#1f1d2e", ANSI256: "236", ANSI: "0"}
	overlay := lipgloss.CompleteColor{TrueColor: "#26233a", ANSI256: "237", ANSI: "8"}
	muted := lipgloss.CompleteColor{TrueColor: "#6e6a86", ANSI256: "243", ANSI: "8"}
	subtle := lipgloss.CompleteColor{TrueColor: "#908caa", ANSI256: "245", ANSI: "7"}
	text := lipgloss.CompleteColor{TrueColor: "#e0def4", ANSI256: "189", ANSI: "7"}
	love := lipgloss.CompleteColor{TrueColor: "#eb6f92", ANSI256: "210", ANSI: "1"}
	gold := lipgloss.CompleteColor{TrueColor: "#f6c177", ANSI256: "179", ANSI: "3"}
	rose := lipgloss.CompleteColor{TrueColor: "#ebbcba", ANSI256: "224", ANSI: "7"}
	pine := lipgloss.CompleteColor{TrueColor: "#31748f", ANSI256: "67", ANSI: "4"}
	foam := lipgloss.CompleteColor{TrueColor: "#9ccfd8", ANSI256: "152", ANSI: "6"}
	iris := lipgloss.CompleteColor{TrueColor: "#c4a7e7", ANSI256: "183", ANSI: "5"}
	highlightMed := lipgloss.CompleteColor{TrueColor: "#403d52", ANSI256: "238", ANSI: "8"}
	shadow := lipgloss.CompleteColor{TrueColor: "#110f18", ANSI256: "234", ANSI: "0"}

	return types.Theme{
		Name: "Rosé Pine",
		Palette: types.Palette{
			Background: base,
			Surface:    surface,
			Overlay:    overlay,
			Text:       text,
			TextMuted:  muted,
			TextSubtle: subtle,
			Primary:    gold, // Gold as primary for focused elements
			Secondary:  rose, // Rose for selections
			Love:       love,
			Gold:       gold,
			Rose:       rose,
			Pine:       pine,
			Foam:       foam,
			Iris:       iris,
			Border:     subtle,
			Shadow:     shadow,
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       text,
					Background: base,
					Border:     subtle,
				},
				Hover: types.StateColors{
					Text:       rose,
					Background: base,
					Border:     rose,
				},
				Selected: types.StateColors{
					Text:       rose, // Rose for selected items
					Background: base,
					Border:     rose,
				},
				Disabled: types.StateColors{
					Text:       muted,
					Background: base,
					Border:     highlightMed,
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       muted,
					Background: base,
					Border:     subtle,
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold, // Gold for focused active tab
						Background: base,
						Border:     gold,
					},
					Unfocused: types.StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold, // Gold for focused borders
						Background: base,
						Border:     gold,
					},
					Unfocused: types.StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold, // Gold for focused titles
						Background: base,
						Border:     base,
					},
					Unfocused: types.StateColors{
						Text:       muted,
						Background: base,
						Border:     base,
					},
				},
			},
		},
	}
}

// RosePineDawn creates the light Rosé Pine Dawn theme
func RosePineDawn() types.Theme {
	// Rosé Pine Dawn color definitions (official colors)
	base := lipgloss.CompleteColor{TrueColor: "#faf4ed", ANSI256: "255", ANSI: "15"}
	surface := lipgloss.CompleteColor{TrueColor: "#fffaf3", ANSI256: "255", ANSI: "15"}
	overlay := lipgloss.CompleteColor{TrueColor: "#f2e9e1", ANSI256: "254", ANSI: "15"}
	muted := lipgloss.CompleteColor{TrueColor: "#9893a5", ANSI256: "246", ANSI: "8"}
	subtle := lipgloss.CompleteColor{TrueColor: "#797593", ANSI256: "243", ANSI: "8"}
	text := lipgloss.CompleteColor{TrueColor: "#575279", ANSI256: "240", ANSI: "0"}
	love := lipgloss.CompleteColor{TrueColor: "#b4637a", ANSI256: "168", ANSI: "1"}
	gold := lipgloss.CompleteColor{TrueColor: "#ea9d34", ANSI256: "172", ANSI: "3"}
	rose := lipgloss.CompleteColor{TrueColor: "#d7827e", ANSI256: "174", ANSI: "7"}
	pine := lipgloss.CompleteColor{TrueColor: "#286983", ANSI256: "24", ANSI: "4"}
	foam := lipgloss.CompleteColor{TrueColor: "#56949f", ANSI256: "66", ANSI: "6"}
	iris := lipgloss.CompleteColor{TrueColor: "#907aa9", ANSI256: "103", ANSI: "5"}
	highlightMed := lipgloss.CompleteColor{TrueColor: "#dfdad9", ANSI256: "253", ANSI: "7"}
	shadow := lipgloss.CompleteColor{TrueColor: "#f4ede8", ANSI256: "254", ANSI: "15"}

	return types.Theme{
		Name: "Rosé Pine Dawn",
		Palette: types.Palette{
			Background: base,
			Surface:    surface,
			Overlay:    overlay,
			Text:       text,
			TextMuted:  muted,
			TextSubtle: subtle,
			Primary:    gold,
			Secondary:  rose,
			Love:       love,
			Gold:       gold,
			Rose:       rose,
			Pine:       pine,
			Foam:       foam,
			Iris:       iris,
			Border:     subtle,
			Shadow:     shadow,
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       text,
					Background: base,
					Border:     subtle,
				},
				Hover: types.StateColors{
					Text:       rose,
					Background: base,
					Border:     rose,
				},
				Selected: types.StateColors{
					Text:       rose,
					Background: base,
					Border:     rose,
				},
				Disabled: types.StateColors{
					Text:       muted,
					Background: base,
					Border:     highlightMed,
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       muted,
					Background: base,
					Border:     subtle,
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold,
						Background: base,
						Border:     gold,
					},
					Unfocused: types.StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold,
						Background: base,
						Border:     gold,
					},
					Unfocused: types.StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold,
						Background: base,
						Border:     base,
					},
					Unfocused: types.StateColors{
						Text:       muted,
						Background: base,
						Border:     base,
					},
				},
			},
		},
	}
}

// RosePineMoon creates the Rosé Pine Moon theme
func RosePineMoon() types.Theme {
	// Rosé Pine Moon color definitions (official colors)
	base := lipgloss.CompleteColor{TrueColor: "#232136", ANSI256: "235", ANSI: "0"}
	surface := lipgloss.CompleteColor{TrueColor: "#2a273f", ANSI256: "236", ANSI: "0"}
	overlay := lipgloss.CompleteColor{TrueColor: "#393552", ANSI256: "237", ANSI: "8"}
	muted := lipgloss.CompleteColor{TrueColor: "#6e6a86", ANSI256: "243", ANSI: "8"}
	subtle := lipgloss.CompleteColor{TrueColor: "#908caa", ANSI256: "245", ANSI: "7"}
	text := lipgloss.CompleteColor{TrueColor: "#e0def4", ANSI256: "189", ANSI: "7"}
	love := lipgloss.CompleteColor{TrueColor: "#eb6f92", ANSI256: "210", ANSI: "1"}
	gold := lipgloss.CompleteColor{TrueColor: "#f6c177", ANSI256: "179", ANSI: "3"}
	rose := lipgloss.CompleteColor{TrueColor: "#ea9a97", ANSI256: "217", ANSI: "7"}
	pine := lipgloss.CompleteColor{TrueColor: "#3e8fb0", ANSI256: "67", ANSI: "4"}
	foam := lipgloss.CompleteColor{TrueColor: "#9ccfd8", ANSI256: "152", ANSI: "6"}
	iris := lipgloss.CompleteColor{TrueColor: "#c4a7e7", ANSI256: "183", ANSI: "5"}
	highlightMed := lipgloss.CompleteColor{TrueColor: "#44415a", ANSI256: "238", ANSI: "8"}
	shadow := lipgloss.CompleteColor{TrueColor: "#1f1d2e", ANSI256: "234", ANSI: "0"}

	return types.Theme{
		Name: "Rosé Pine Moon",
		Palette: types.Palette{
			Background: base,
			Surface:    surface,
			Overlay:    overlay,
			Text:       text,
			TextMuted:  muted,
			TextSubtle: subtle,
			Primary:    gold,
			Secondary:  rose,
			Love:       love,
			Gold:       gold,
			Rose:       rose,
			Pine:       pine,
			Foam:       foam,
			Iris:       iris,
			Border:     subtle,
			Shadow:     shadow,
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       text,
					Background: base,
					Border:     subtle,
				},
				Hover: types.StateColors{
					Text:       rose,
					Background: base,
					Border:     rose,
				},
				Selected: types.StateColors{
					Text:       rose,
					Background: base,
					Border:     rose,
				},
				Disabled: types.StateColors{
					Text:       muted,
					Background: base,
					Border:     highlightMed,
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       muted,
					Background: base,
					Border:     subtle,
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold,
						Background: base,
						Border:     gold,
					},
					Unfocused: types.StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold,
						Background: base,
						Border:     gold,
					},
					Unfocused: types.StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       gold,
						Background: base,
						Border:     base,
					},
					Unfocused: types.StateColors{
						Text:       muted,
						Background: base,
						Border:     base,
					},
				},
			},
		},
	}
}