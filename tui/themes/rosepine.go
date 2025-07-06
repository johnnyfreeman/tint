package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// RosePine creates the standard Rosé Pine theme
func RosePine() types.Theme {
	// Rosé Pine color definitions (official colors)
	base := lipgloss.Color("#191724")
	surface := lipgloss.Color("#1f1d2e")
	overlay := lipgloss.Color("#26233a")
	muted := lipgloss.Color("#6e6a86")
	subtle := lipgloss.Color("#908caa")
	text := lipgloss.Color("#e0def4")
	love := lipgloss.Color("#eb6f92")
	gold := lipgloss.Color("#f6c177")
	rose := lipgloss.Color("#ebbcba")
	pine := lipgloss.Color("#31748f")
	foam := lipgloss.Color("#9ccfd8")
	iris := lipgloss.Color("#c4a7e7")
	highlightMed := lipgloss.Color("#403d52")
	shadow := lipgloss.Color("#110f18")

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
	base := lipgloss.Color("#faf4ed")
	surface := lipgloss.Color("#fffaf3")
	overlay := lipgloss.Color("#f2e9e1")
	muted := lipgloss.Color("#9893a5")
	subtle := lipgloss.Color("#797593")
	text := lipgloss.Color("#575279")
	love := lipgloss.Color("#b4637a")
	gold := lipgloss.Color("#ea9d34")
	rose := lipgloss.Color("#d7827e")
	pine := lipgloss.Color("#286983")
	foam := lipgloss.Color("#56949f")
	iris := lipgloss.Color("#907aa9")
	highlightMed := lipgloss.Color("#dfdad9")
	shadow := lipgloss.Color("#f4ede8")

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
	base := lipgloss.Color("#232136")
	surface := lipgloss.Color("#2a273f")
	overlay := lipgloss.Color("#393552")
	muted := lipgloss.Color("#6e6a86")
	subtle := lipgloss.Color("#908caa")
	text := lipgloss.Color("#e0def4")
	love := lipgloss.Color("#eb6f92")
	gold := lipgloss.Color("#f6c177")
	rose := lipgloss.Color("#ea9a97")
	pine := lipgloss.Color("#3e8fb0")
	foam := lipgloss.Color("#9ccfd8")
	iris := lipgloss.Color("#c4a7e7")
	highlightMed := lipgloss.Color("#44415a")
	shadow := lipgloss.Color("#1f1d2e")

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