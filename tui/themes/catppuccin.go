package themes

import (
	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui/types"
)

// CatppuccinMocha creates a Catppuccin Mocha theme using the official colors
func CatppuccinMocha() types.Theme {
	palette := catppuccin.Mocha

	return types.Theme{
		Name: "Catppuccin Mocha",
		Palette: types.Palette{
			Background: lipgloss.Color(palette.Base().Hex),
			Surface:    lipgloss.Color(palette.Surface0().Hex),
			Overlay:    lipgloss.Color(palette.Surface1().Hex),
			Text:       lipgloss.Color(palette.Text().Hex),
			TextMuted:  lipgloss.Color(palette.Overlay1().Hex),
			TextSubtle: lipgloss.Color(palette.Overlay0().Hex),
			Primary:    lipgloss.Color(palette.Blue().Hex),
			Secondary:  lipgloss.Color(palette.Mauve().Hex),
			Love:       lipgloss.Color(palette.Red().Hex),
			Gold:       lipgloss.Color(palette.Yellow().Hex),
			Rose:       lipgloss.Color(palette.Rosewater().Hex),
			Pine:       lipgloss.Color(palette.Green().Hex),
			Foam:       lipgloss.Color(palette.Sapphire().Hex),
			Iris:       lipgloss.Color(palette.Mauve().Hex),
			Border:     lipgloss.Color(palette.Overlay0().Hex),
			Shadow:     lipgloss.Color(palette.Crust().Hex),
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.Color(palette.Text().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Hover: types.StateColors{
					Text:       lipgloss.Color(palette.Lavender().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Lavender().Hex),
				},
				Selected: types.StateColors{
					Text:       lipgloss.Color(palette.Blue().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Blue().Hex),
				},
				Disabled: types.StateColors{
					Text:       lipgloss.Color(palette.Surface2().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Surface1().Hex),
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.Color(palette.Overlay1().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Overlay1().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
				},
			},
		},
	}
}

// CatppuccinLatte creates a Catppuccin Latte theme (light variant)
func CatppuccinLatte() types.Theme {
	palette := catppuccin.Latte

	return types.Theme{
		Name: "Catppuccin Latte",
		Palette: types.Palette{
			Background: lipgloss.Color(palette.Base().Hex),
			Surface:    lipgloss.Color(palette.Surface0().Hex),
			Overlay:    lipgloss.Color(palette.Surface1().Hex),
			Text:       lipgloss.Color(palette.Text().Hex),
			TextMuted:  lipgloss.Color(palette.Overlay1().Hex),
			TextSubtle: lipgloss.Color(palette.Overlay0().Hex),
			Primary:    lipgloss.Color(palette.Blue().Hex),
			Secondary:  lipgloss.Color(palette.Mauve().Hex),
			Love:       lipgloss.Color(palette.Red().Hex),
			Gold:       lipgloss.Color(palette.Yellow().Hex),
			Rose:       lipgloss.Color(palette.Rosewater().Hex),
			Pine:       lipgloss.Color(palette.Green().Hex),
			Foam:       lipgloss.Color(palette.Sapphire().Hex),
			Iris:       lipgloss.Color(palette.Mauve().Hex),
			Border:     lipgloss.Color(palette.Overlay0().Hex),
			Shadow:     lipgloss.Color(palette.Crust().Hex),
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.Color(palette.Text().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Hover: types.StateColors{
					Text:       lipgloss.Color(palette.Lavender().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Lavender().Hex),
				},
				Selected: types.StateColors{
					Text:       lipgloss.Color(palette.Blue().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Blue().Hex),
				},
				Disabled: types.StateColors{
					Text:       lipgloss.Color(palette.Surface2().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Surface1().Hex),
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.Color(palette.Overlay1().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Overlay1().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
				},
			},
		},
	}
}

// CatppuccinFrappe creates a Catppuccin Frappe theme
func CatppuccinFrappe() types.Theme {
	palette := catppuccin.Frappe

	return types.Theme{
		Name: "Catppuccin Frappe",
		Palette: types.Palette{
			Background: lipgloss.Color(palette.Base().Hex),
			Surface:    lipgloss.Color(palette.Surface0().Hex),
			Overlay:    lipgloss.Color(palette.Surface1().Hex),
			Text:       lipgloss.Color(palette.Text().Hex),
			TextMuted:  lipgloss.Color(palette.Overlay1().Hex),
			TextSubtle: lipgloss.Color(palette.Overlay0().Hex),
			Primary:    lipgloss.Color(palette.Blue().Hex),
			Secondary:  lipgloss.Color(palette.Mauve().Hex),
			Love:       lipgloss.Color(palette.Red().Hex),
			Gold:       lipgloss.Color(palette.Yellow().Hex),
			Rose:       lipgloss.Color(palette.Rosewater().Hex),
			Pine:       lipgloss.Color(palette.Green().Hex),
			Foam:       lipgloss.Color(palette.Sapphire().Hex),
			Iris:       lipgloss.Color(palette.Mauve().Hex),
			Border:     lipgloss.Color(palette.Overlay0().Hex),
			Shadow:     lipgloss.Color(palette.Crust().Hex),
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.Color(palette.Text().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Hover: types.StateColors{
					Text:       lipgloss.Color(palette.Lavender().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Lavender().Hex),
				},
				Selected: types.StateColors{
					Text:       lipgloss.Color(palette.Blue().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Blue().Hex),
				},
				Disabled: types.StateColors{
					Text:       lipgloss.Color(palette.Surface2().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Surface1().Hex),
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.Color(palette.Overlay1().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Overlay1().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
				},
			},
		},
	}
}

// CatppuccinMacchiato creates a Catppuccin Macchiato theme
func CatppuccinMacchiato() types.Theme {
	palette := catppuccin.Macchiato

	return types.Theme{
		Name: "Catppuccin Macchiato",
		Palette: types.Palette{
			Background: lipgloss.Color(palette.Base().Hex),
			Surface:    lipgloss.Color(palette.Surface0().Hex),
			Overlay:    lipgloss.Color(palette.Surface1().Hex),
			Text:       lipgloss.Color(palette.Text().Hex),
			TextMuted:  lipgloss.Color(palette.Overlay1().Hex),
			TextSubtle: lipgloss.Color(palette.Overlay0().Hex),
			Primary:    lipgloss.Color(palette.Blue().Hex),
			Secondary:  lipgloss.Color(palette.Mauve().Hex),
			Love:       lipgloss.Color(palette.Red().Hex),
			Gold:       lipgloss.Color(palette.Yellow().Hex),
			Rose:       lipgloss.Color(palette.Rosewater().Hex),
			Pine:       lipgloss.Color(palette.Green().Hex),
			Foam:       lipgloss.Color(palette.Sapphire().Hex),
			Iris:       lipgloss.Color(palette.Mauve().Hex),
			Border:     lipgloss.Color(palette.Overlay0().Hex),
			Shadow:     lipgloss.Color(palette.Crust().Hex),
		},
		Components: types.Components{
			Interactive: types.InteractiveStyle{
				Normal: types.StateColors{
					Text:       lipgloss.Color(palette.Text().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Hover: types.StateColors{
					Text:       lipgloss.Color(palette.Lavender().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Lavender().Hex),
				},
				Selected: types.StateColors{
					Text:       lipgloss.Color(palette.Blue().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Blue().Hex),
				},
				Disabled: types.StateColors{
					Text:       lipgloss.Color(palette.Surface2().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Surface1().Hex),
				},
			},
			Tab: types.TabStyle{
				Inactive: types.StateColors{
					Text:       lipgloss.Color(palette.Overlay1().Hex),
					Background: lipgloss.Color(palette.Base().Hex),
					Border:     lipgloss.Color(palette.Overlay0().Hex),
				},
				Active: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
			},
			Container: types.ContainerStyle{
				Border: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Blue().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Text().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Overlay0().Hex),
					},
				},
				Title: types.FocusableStyle{
					Focused: types.StateColors{
						Text:       lipgloss.Color(palette.Blue().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
					Unfocused: types.StateColors{
						Text:       lipgloss.Color(palette.Overlay1().Hex),
						Background: lipgloss.Color(palette.Base().Hex),
						Border:     lipgloss.Color(palette.Base().Hex),
					},
				},
			},
		},
	}
}