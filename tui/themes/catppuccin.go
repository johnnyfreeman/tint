package themes

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// CatppuccinMocha creates a Catppuccin Mocha theme using static hex colors
func CatppuccinMocha() types.Theme {
	return types.Theme{
		Name: "Catppuccin Mocha",
		Palette: types.Palette{
			Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
			Surface:    lipgloss.CompleteColor{TrueColor: "#313244", ANSI256: "236", ANSI: "0"},
			Overlay:    lipgloss.CompleteColor{TrueColor: "#45475a", ANSI256: "237", ANSI: "0"},
			Text:       lipgloss.CompleteColor{TrueColor: "#cdd6f4", ANSI256: "189", ANSI: "7"},
			TextMuted:  lipgloss.CompleteColor{TrueColor: "#6c7086", ANSI256: "244", ANSI: "7"},
			TextSubtle: lipgloss.CompleteColor{TrueColor: "#585b70", ANSI256: "243", ANSI: "8"},
			Primary:    lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
			Secondary:  lipgloss.CompleteColor{TrueColor: "#cba6f7", ANSI256: "183", ANSI: "5"},
			Love:       lipgloss.CompleteColor{TrueColor: "#f38ba8", ANSI256: "210", ANSI: "1"},
			Gold:       lipgloss.CompleteColor{TrueColor: "#f9e2af", ANSI256: "179", ANSI: "3"},
			Rose:       lipgloss.CompleteColor{TrueColor: "#f5e0dc", ANSI256: "224", ANSI: "17"},
			Pine:       lipgloss.CompleteColor{TrueColor: "#a6e3a1", ANSI256: "149", ANSI: "2"},
			Foam:       lipgloss.CompleteColor{TrueColor: "#74c7ec", ANSI256: "74", ANSI: "6"},
			Iris:       lipgloss.CompleteColor{TrueColor: "#cba6f7", ANSI256: "183", ANSI: "5"},
			Border:     lipgloss.CompleteColor{TrueColor: "#585b70", ANSI256: "243", ANSI: "0"},
			Shadow:     lipgloss.CompleteColor{TrueColor: "#11111b", ANSI256: "234", ANSI: "0"},
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#cdd6f4", ANSI256: "189", ANSI: "7"},
					Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#585b70", ANSI256: "243", ANSI: "8"},
				},
				Hover: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#b4befe", ANSI256: "147", ANSI: "6"},
					Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#b4befe", ANSI256: "147", ANSI: "6"},
				},
				Selected: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
					Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
				},
				Disabled: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#7f849c", ANSI256: "245", ANSI: "8"},
					Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#45475a", ANSI256: "237", ANSI: "8"},
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#6c7086", ANSI256: "244", ANSI: "8"},
					Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#585b70", ANSI256: "243", ANSI: "8"},
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#cdd6f4", ANSI256: "189", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#585b70", ANSI256: "243", ANSI: "8"},
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#cdd6f4", ANSI256: "189", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#585b70", ANSI256: "243", ANSI: "8"},
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#89b4fa", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#6c7086", ANSI256: "244", ANSI: "8"},
						Background: lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#1e1e2e", ANSI256: "235", ANSI: "0"},
					},
				},
			},
		},
	}
}

// CatppuccinLatte creates a Catppuccin Latte theme using static hex colors
func CatppuccinLatte() types.Theme {
	return types.Theme{
		Name: "Catppuccin Latte",
		Palette: types.Palette{
			Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
			Surface:    lipgloss.CompleteColor{TrueColor: "#ccd0da", ANSI256: "251", ANSI: "15"},
			Overlay:    lipgloss.CompleteColor{TrueColor: "#bcc0cc", ANSI256: "250", ANSI: "15"},
			Text:       lipgloss.CompleteColor{TrueColor: "#4c4f69", ANSI256: "240", ANSI: "0"},
			TextMuted:  lipgloss.CompleteColor{TrueColor: "#8c8fa1", ANSI256: "245", ANSI: "0"},
			TextSubtle: lipgloss.CompleteColor{TrueColor: "#9ca0b0", ANSI256: "247", ANSI: "8"},
			Primary:    lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
			Secondary:  lipgloss.CompleteColor{TrueColor: "#8839ef", ANSI256: "91", ANSI: "5"},
			Love:       lipgloss.CompleteColor{TrueColor: "#d20f39", ANSI256: "160", ANSI: "1"},
			Gold:       lipgloss.CompleteColor{TrueColor: "#df8e1d", ANSI256: "172", ANSI: "3"},
			Rose:       lipgloss.CompleteColor{TrueColor: "#dc8a78", ANSI256: "174", ANSI: "17"},
			Pine:       lipgloss.CompleteColor{TrueColor: "#40a02b", ANSI256: "70", ANSI: "2"},
			Foam:       lipgloss.CompleteColor{TrueColor: "#179299", ANSI256: "30", ANSI: "6"},
			Iris:       lipgloss.CompleteColor{TrueColor: "#8839ef", ANSI256: "91", ANSI: "5"},
			Border:     lipgloss.CompleteColor{TrueColor: "#9ca0b0", ANSI256: "247", ANSI: "0"},
			Shadow:     lipgloss.CompleteColor{TrueColor: "#dce0e8", ANSI256: "253", ANSI: "15"},
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#4c4f69", ANSI256: "240", ANSI: "0"},
					Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					Border:     lipgloss.CompleteColor{TrueColor: "#9ca0b0", ANSI256: "247", ANSI: "8"},
				},
				Hover: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#7287fd", ANSI256: "105", ANSI: "6"},
					Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					Border:     lipgloss.CompleteColor{TrueColor: "#7287fd", ANSI256: "105", ANSI: "6"},
				},
				Selected: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
					Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					Border:     lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
				},
				Disabled: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#acb0be", ANSI256: "248", ANSI: "8"},
					Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					Border:     lipgloss.CompleteColor{TrueColor: "#bcc0cc", ANSI256: "250", ANSI: "15"},
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#8c8fa1", ANSI256: "245", ANSI: "0"},
					Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					Border:     lipgloss.CompleteColor{TrueColor: "#9ca0b0", ANSI256: "247", ANSI: "8"},
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
						Border:     lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#4c4f69", ANSI256: "240", ANSI: "0"},
						Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
						Border:     lipgloss.CompleteColor{TrueColor: "#9ca0b0", ANSI256: "247", ANSI: "8"},
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
						Border:     lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#4c4f69", ANSI256: "240", ANSI: "0"},
						Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
						Border:     lipgloss.CompleteColor{TrueColor: "#9ca0b0", ANSI256: "247", ANSI: "8"},
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#1e66f5", ANSI256: "27", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
						Border:     lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8c8fa1", ANSI256: "245", ANSI: "0"},
						Background: lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
						Border:     lipgloss.CompleteColor{TrueColor: "#eff1f5", ANSI256: "254", ANSI: "15"},
					},
				},
			},
		},
	}
}

// CatppuccinFrappe creates a Catppuccin Frappe theme using static hex colors
func CatppuccinFrappe() types.Theme {
	return types.Theme{
		Name: "Catppuccin Frappe",
		Palette: types.Palette{
			Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
			Surface:    lipgloss.CompleteColor{TrueColor: "#414559", ANSI256: "238", ANSI: "0"},
			Overlay:    lipgloss.CompleteColor{TrueColor: "#51576d", ANSI256: "239", ANSI: "0"},
			Text:       lipgloss.CompleteColor{TrueColor: "#c6d0f5", ANSI256: "189", ANSI: "7"},
			TextMuted:  lipgloss.CompleteColor{TrueColor: "#838ba7", ANSI256: "245", ANSI: "7"},
			TextSubtle: lipgloss.CompleteColor{TrueColor: "#737994", ANSI256: "243", ANSI: "8"},
			Primary:    lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
			Secondary:  lipgloss.CompleteColor{TrueColor: "#ca9ee6", ANSI256: "183", ANSI: "5"},
			Love:       lipgloss.CompleteColor{TrueColor: "#e78284", ANSI256: "210", ANSI: "1"},
			Gold:       lipgloss.CompleteColor{TrueColor: "#e5c890", ANSI256: "179", ANSI: "3"},
			Rose:       lipgloss.CompleteColor{TrueColor: "#f2d5cf", ANSI256: "224", ANSI: "17"},
			Pine:       lipgloss.CompleteColor{TrueColor: "#a6d189", ANSI256: "149", ANSI: "2"},
			Foam:       lipgloss.CompleteColor{TrueColor: "#81c8be", ANSI256: "115", ANSI: "6"},
			Iris:       lipgloss.CompleteColor{TrueColor: "#ca9ee6", ANSI256: "183", ANSI: "5"},
			Border:     lipgloss.CompleteColor{TrueColor: "#737994", ANSI256: "243", ANSI: "0"},
			Shadow:     lipgloss.CompleteColor{TrueColor: "#232634", ANSI256: "235", ANSI: "0"},
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#c6d0f5", ANSI256: "189", ANSI: "7"},
					Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#737994", ANSI256: "243", ANSI: "8"},
				},
				Hover: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#babbf1", ANSI256: "147", ANSI: "6"},
					Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#babbf1", ANSI256: "147", ANSI: "6"},
				},
				Selected: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
					Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
				},
				Disabled: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#626880", ANSI256: "242", ANSI: "8"},
					Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#51576d", ANSI256: "239", ANSI: "0"},
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#838ba7", ANSI256: "245", ANSI: "7"},
					Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#737994", ANSI256: "243", ANSI: "8"},
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#c6d0f5", ANSI256: "189", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#737994", ANSI256: "243", ANSI: "8"},
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#c6d0f5", ANSI256: "189", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#737994", ANSI256: "243", ANSI: "8"},
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8caaee", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#838ba7", ANSI256: "245", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#303446", ANSI256: "237", ANSI: "0"},
					},
				},
			},
		},
	}
}

// CatppuccinMacchiato creates a Catppuccin Macchiato theme using static hex colors
func CatppuccinMacchiato() types.Theme {
	return types.Theme{
		Name: "Catppuccin Macchiato",
		Palette: types.Palette{
			Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
			Surface:    lipgloss.CompleteColor{TrueColor: "#363a4f", ANSI256: "237", ANSI: "0"},
			Overlay:    lipgloss.CompleteColor{TrueColor: "#494d64", ANSI256: "238", ANSI: "0"},
			Text:       lipgloss.CompleteColor{TrueColor: "#cad3f5", ANSI256: "189", ANSI: "7"},
			TextMuted:  lipgloss.CompleteColor{TrueColor: "#8087a2", ANSI256: "245", ANSI: "7"},
			TextSubtle: lipgloss.CompleteColor{TrueColor: "#6e738d", ANSI256: "243", ANSI: "8"},
			Primary:    lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
			Secondary:  lipgloss.CompleteColor{TrueColor: "#c6a0f6", ANSI256: "183", ANSI: "5"},
			Love:       lipgloss.CompleteColor{TrueColor: "#ed8796", ANSI256: "210", ANSI: "1"},
			Gold:       lipgloss.CompleteColor{TrueColor: "#eed49f", ANSI256: "179", ANSI: "3"},
			Rose:       lipgloss.CompleteColor{TrueColor: "#f4dbd6", ANSI256: "224", ANSI: "17"},
			Pine:       lipgloss.CompleteColor{TrueColor: "#a6da95", ANSI256: "149", ANSI: "2"},
			Foam:       lipgloss.CompleteColor{TrueColor: "#8bd5ca", ANSI256: "115", ANSI: "6"},
			Iris:       lipgloss.CompleteColor{TrueColor: "#c6a0f6", ANSI256: "183", ANSI: "5"},
			Border:     lipgloss.CompleteColor{TrueColor: "#6e738d", ANSI256: "243", ANSI: "0"},
			Shadow:     lipgloss.CompleteColor{TrueColor: "#181926", ANSI256: "235", ANSI: "0"},
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#cad3f5", ANSI256: "189", ANSI: "7"},
					Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#6e738d", ANSI256: "243", ANSI: "8"},
				},
				Hover: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#b7bdf8", ANSI256: "147", ANSI: "6"},
					Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#b7bdf8", ANSI256: "147", ANSI: "6"},
				},
				Selected: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
					Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
				},
				Disabled: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#5b6078", ANSI256: "242", ANSI: "8"},
					Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#494d64", ANSI256: "238", ANSI: "0"},
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.CompleteColor{TrueColor: "#8087a2", ANSI256: "245", ANSI: "7"},
					Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					Border:     lipgloss.CompleteColor{TrueColor: "#6e738d", ANSI256: "243", ANSI: "8"},
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#cad3f5", ANSI256: "189", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#6e738d", ANSI256: "243", ANSI: "8"},
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#cad3f5", ANSI256: "189", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#6e738d", ANSI256: "243", ANSI: "8"},
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8aadf4", ANSI256: "117", ANSI: "4"},
						Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.CompleteColor{TrueColor: "#8087a2", ANSI256: "245", ANSI: "7"},
						Background: lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
						Border:     lipgloss.CompleteColor{TrueColor: "#24273a", ANSI256: "236", ANSI: "0"},
					},
				},
			},
		},
	}
}