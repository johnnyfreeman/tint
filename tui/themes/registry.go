package themes

import "github.com/johnnyfreeman/tint/tui/types"

// Registry holds all available themes
var Registry = map[string]func() types.Theme{
	// Tokyo Night
	"tokyonight": TokyoNight,

	// Rosé Pine variants
	"rosepine":      RosePine,
	"rosepine-dawn": RosePineDawn,
	"rosepine-moon": RosePineMoon,

	// Catppuccin variants
	"catppuccin-latte":     CatppuccinLatte,
	"catppuccin-frappe":    CatppuccinFrappe,
	"catppuccin-macchiato": CatppuccinMacchiato,
	"catppuccin-mocha":     CatppuccinMocha,

	// Custom themes
	"monochrome": Monochrome,
}

// GetTheme returns a theme by name, or the default theme if not found
func GetTheme(name string) types.Theme {
	if themeFunc, ok := Registry[name]; ok {
		return themeFunc()
	}
	// Default to monochrome if theme not found
	return Monochrome()
}

// GetAvailableThemes returns a list of all available theme names
func GetAvailableThemes() []string {
	themes := make([]string, 0, len(Registry))
	for name := range Registry {
		themes = append(themes, name)
	}
	return themes
}

// ThemeExists checks if a theme with the given name exists
func ThemeExists(name string) bool {
	_, exists := Registry[name]
	return exists
}