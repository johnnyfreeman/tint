package tui

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

// Themes contains all available themes
var Themes = map[string]Theme{
	"tokyonight": createTokyoNightTheme(),
	"rosepine":   createRosePineTheme(),
	"catppuccin": createCatppuccinTheme(),
	"monochrome": createMonochromeTheme(),
}

func createTokyoNightTheme() Theme {
	// Tokyo Night color definitions
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

	return Theme{
		Name: "Tokyo Night",
		Palette: Palette{
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
		Components: Components{
			Interactive: InteractiveStyle{
				Normal: StateColors{
					Text:       fg,
					Background: bg,
					Border:     dark3,
				},
				Hover: StateColors{
					Text:       cyan,
					Background: bg,
					Border:     cyan,
				},
				Selected: StateColors{
					Text:       blue,
					Background: bg,
					Border:     blue,
				},
				Disabled: StateColors{
					Text:       dark3,
					Background: bg,
					Border:     fgGutter,
				},
			},
			Tab: TabStyle{
				Inactive: StateColors{
					Text:       comment,
					Background: bg,
					Border:     dark3,
				},
				Active: FocusableStyle{
					Focused: StateColors{
						Text:       blue,
						Background: bg,
						Border:     blue,
					},
					Unfocused: StateColors{
						Text:       fg,
						Background: bg,
						Border:     dark3,
					},
				},
			},
			Container: ContainerStyle{
				Border: FocusableStyle{
					Focused: StateColors{
						Text:       blue,
						Background: bg,
						Border:     blue,
					},
					Unfocused: StateColors{
						Text:       fg,
						Background: bg,
						Border:     dark3,
					},
				},
				Title: FocusableStyle{
					Focused: StateColors{
						Text:       blue,
						Background: bg,
						Border:     bg,
					},
					Unfocused: StateColors{
						Text:       comment,
						Background: bg,
						Border:     bg,
					},
				},
			},
		},
	}
}

func createRosePineTheme() Theme {
	// Rosé Pine color definitions
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

	return Theme{
		Name: "Rosé Pine",
		Palette: Palette{
			Background: base,
			Surface:    surface,
			Overlay:    overlay,
			Text:       text,
			TextMuted:  muted,
			TextSubtle: subtle,
			Primary:    gold,     // Gold as primary for focused elements
			Secondary:  rose,     // Rose for selections
			Love:       love,
			Gold:       gold,
			Rose:       rose,
			Pine:       pine,
			Foam:       foam,
			Iris:       iris,
			Border:     subtle,
			Shadow:     shadow,
		},
		Components: Components{
			Interactive: InteractiveStyle{
				Normal: StateColors{
					Text:       text,
					Background: base,
					Border:     subtle,
				},
				Hover: StateColors{
					Text:       rose,
					Background: base,
					Border:     rose,
				},
				Selected: StateColors{
					Text:       rose,      // Rose for selected items
					Background: base,
					Border:     rose,
				},
				Disabled: StateColors{
					Text:       muted,
					Background: base,
					Border:     highlightMed,
				},
			},
			Tab: TabStyle{
				Inactive: StateColors{
					Text:       muted,
					Background: base,
					Border:     subtle,
				},
				Active: FocusableStyle{
					Focused: StateColors{
						Text:       gold,     // Gold for focused active tab
						Background: base,
						Border:     gold,
					},
					Unfocused: StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
			},
			Container: ContainerStyle{
				Border: FocusableStyle{
					Focused: StateColors{
						Text:       gold,     // Gold for focused borders
						Background: base,
						Border:     gold,
					},
					Unfocused: StateColors{
						Text:       text,
						Background: base,
						Border:     subtle,
					},
				},
				Title: FocusableStyle{
					Focused: StateColors{
						Text:       gold,     // Gold for focused titles
						Background: base,
						Border:     base,
					},
					Unfocused: StateColors{
						Text:       muted,
						Background: base,
						Border:     base,
					},
				},
			},
		},
	}
}

func createCatppuccinTheme() Theme {
	// Catppuccin Mocha color definitions
	rosewater := lipgloss.Color("#f5e0dc")
	mauve := lipgloss.Color("#cba6f7")
	red := lipgloss.Color("#f38ba8")
	yellow := lipgloss.Color("#f9e2af")
	green := lipgloss.Color("#a6e3a1")
	sapphire := lipgloss.Color("#74c7ec")
	blue := lipgloss.Color("#89b4fa")
	lavender := lipgloss.Color("#b4befe")
	text := lipgloss.Color("#cdd6f4")
	overlay1 := lipgloss.Color("#7f849c")
	overlay0 := lipgloss.Color("#6c7086")
	surface2 := lipgloss.Color("#585b70")
	surface1 := lipgloss.Color("#45475a")
	surface0 := lipgloss.Color("#313244")
	base := lipgloss.Color("#1e1e2e")
	crust := lipgloss.Color("#11111b")

	return Theme{
		Name: "Catppuccin",
		Palette: Palette{
			Background: base,
			Surface:    surface0,
			Overlay:    surface1,
			Text:       text,
			TextMuted:  overlay1,
			TextSubtle: overlay0,
			Primary:    blue,
			Secondary:  mauve,
			Love:       red,
			Gold:       yellow,
			Rose:       rosewater,
			Pine:       green,
			Foam:       sapphire,
			Iris:       mauve,
			Border:     overlay0,
			Shadow:     crust,
		},
		Components: Components{
			Interactive: InteractiveStyle{
				Normal: StateColors{
					Text:       text,
					Background: base,
					Border:     overlay0,
				},
				Hover: StateColors{
					Text:       lavender,
					Background: base,
					Border:     lavender,
				},
				Selected: StateColors{
					Text:       blue,
					Background: base,
					Border:     blue,
				},
				Disabled: StateColors{
					Text:       surface2,
					Background: base,
					Border:     surface1,
				},
			},
			Tab: TabStyle{
				Inactive: StateColors{
					Text:       overlay1,
					Background: base,
					Border:     overlay0,
				},
				Active: FocusableStyle{
					Focused: StateColors{
						Text:       blue,
						Background: base,
						Border:     blue,
					},
					Unfocused: StateColors{
						Text:       text,
						Background: base,
						Border:     overlay0,
					},
				},
			},
			Container: ContainerStyle{
				Border: FocusableStyle{
					Focused: StateColors{
						Text:       blue,
						Background: base,
						Border:     blue,
					},
					Unfocused: StateColors{
						Text:       text,
						Background: base,
						Border:     overlay0,
					},
				},
				Title: FocusableStyle{
					Focused: StateColors{
						Text:       blue,
						Background: base,
						Border:     base,
					},
					Unfocused: StateColors{
						Text:       overlay1,
						Background: base,
						Border:     base,
					},
				},
			},
		},
	}
}

func createMonochromeTheme() Theme {
	// Monochrome color definitions
	white := lipgloss.Color("255")
	lightGray := lipgloss.Color("250")
	gray := lipgloss.Color("240")
	darkGray := lipgloss.Color("235")
	veryDarkGray := lipgloss.Color("234")
	black := lipgloss.Color("233")

	return Theme{
		Name: "Monochrome",
		Palette: Palette{
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
		Components: Components{
			Interactive: InteractiveStyle{
				Normal: StateColors{
					Text:       white,
					Background: veryDarkGray,
					Border:     gray,
				},
				Hover: StateColors{
					Text:       white,
					Background: veryDarkGray,
					Border:     white,
				},
				Selected: StateColors{
					Text:       white,
					Background: veryDarkGray,
					Border:     white,
				},
				Disabled: StateColors{
					Text:       gray,
					Background: veryDarkGray,
					Border:     darkGray,
				},
			},
			Tab: TabStyle{
				Inactive: StateColors{
					Text:       gray,
					Background: veryDarkGray,
					Border:     gray,
				},
				Active: FocusableStyle{
					Focused: StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     white,
					},
					Unfocused: StateColors{
						Text:       lightGray,
						Background: veryDarkGray,
						Border:     gray,
					},
				},
			},
			Container: ContainerStyle{
				Border: FocusableStyle{
					Focused: StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     white,
					},
					Unfocused: StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     gray,
					},
				},
				Title: FocusableStyle{
					Focused: StateColors{
						Text:       white,
						Background: veryDarkGray,
						Border:     veryDarkGray,
					},
					Unfocused: StateColors{
						Text:       gray,
						Background: veryDarkGray,
						Border:     veryDarkGray,
					},
				},
			},
		},
	}
}

// GetTheme returns a theme by name
func GetTheme(name string) Theme {
	if theme, ok := Themes[name]; ok {
		return theme
	}
	return Themes["monochrome"]
}