package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Tab represents a single tab with title and content
type Tab struct {
	Title   string
	Content interface{} // Can be string or Component
}

// TabsComponent represents a tabbed container component
type TabsComponent struct {
	tabs        []Tab
	active      int
	width       int
	height      int
	focused     bool
	x, y        int
	renderStyle TabRenderStyle
}

// TabRenderStyle defines how tabs are rendered
type TabRenderStyle int

const (
	TabsOnTop TabRenderStyle = iota
	TabsOnBottom
)

// NewTabs creates a new tabs component
func NewTabs() *TabsComponent {
	return &TabsComponent{
		tabs:        []Tab{},
		active:      0,
		width:       40,
		height:      10,
		focused:     false,
		renderStyle: TabsOnTop,
	}
}

// AddTab adds a new tab
func (t *TabsComponent) AddTab(title string, content interface{}) {
	t.tabs = append(t.tabs, Tab{Title: title, Content: content})
}

// SetTabs replaces all tabs
func (t *TabsComponent) SetTabs(tabs []Tab) {
	t.tabs = tabs
	if t.active >= len(t.tabs) {
		t.active = len(t.tabs) - 1
	}
	if t.active < 0 && len(t.tabs) > 0 {
		t.active = 0
	}
}

// SetActive sets the active tab by index
func (t *TabsComponent) SetActive(index int) {
	if index >= 0 && index < len(t.tabs) {
		t.active = index
	}
}

// GetActive returns the active tab index
func (t *TabsComponent) GetActive() int {
	return t.active
}

// GetActiveTab returns the active tab
func (t *TabsComponent) GetActiveTab() *Tab {
	if t.active >= 0 && t.active < len(t.tabs) {
		return &t.tabs[t.active]
	}
	return nil
}

// NextTab switches to the next tab
func (t *TabsComponent) NextTab() {
	if len(t.tabs) > 0 {
		t.active = (t.active + 1) % len(t.tabs)
	}
}

// PrevTab switches to the previous tab
func (t *TabsComponent) PrevTab() {
	if len(t.tabs) > 0 {
		t.active = (t.active - 1 + len(t.tabs)) % len(t.tabs)
	}
}

// SetRenderStyle sets how tabs are rendered
func (t *TabsComponent) SetRenderStyle(style TabRenderStyle) {
	t.renderStyle = style
}

// Draw renders the tabs component to the screen
func (t *TabsComponent) Draw(screen *Screen, x, y int, theme *Theme) {
	t.drawAtPosition(screen, x, y, t.width, t.height, theme)
}

// drawAtPosition draws the tabs at a specific position with given dimensions
func (t *TabsComponent) drawAtPosition(screen *Screen, x, y, width, height int, theme *Theme) {
	if len(t.tabs) == 0 {
		return
	}

	// Get container styles
	var borderColors StateColors
	if t.focused {
		borderColors = theme.Components.Container.Border.Focused
	} else {
		borderColors = theme.Components.Container.Border.Unfocused
	}
	borderStyle := lipgloss.NewStyle().
		Foreground(borderColors.Border).
		Background(theme.Palette.Background)

	// Content text style with background
	contentStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Background)

	// Fill background first
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}

	// Draw based on render style
	if t.renderStyle == TabsOnTop {
		t.drawTabsOnTop(screen, x, y, width, height, theme, borderStyle, contentStyle)
	} else {
		t.drawTabsOnBottom(screen, x, y, width, height, theme, borderStyle, contentStyle)
	}
}

// drawTabsOnTop draws tabs on the top border
func (t *TabsComponent) drawTabsOnTop(screen *Screen, x, y, width, height int, theme *Theme, borderStyle, contentStyle lipgloss.Style) {
	// Draw top border with embedded tab titles
	if t.focused {
		screen.DrawRune(x, y, '┏', borderStyle)
	} else {
		screen.DrawRune(x, y, '┌', borderStyle)
	}

	currentX := x + 1
	for i, tab := range t.tabs {
		if i > 0 {
			// Draw separator between tabs - bold when focused
			if t.focused {
				screen.DrawRune(currentX, y, '┃', borderStyle)
			} else {
				screen.DrawRune(currentX, y, '─', borderStyle)
			}
			currentX++
		}

		// Determine tab style based on state
		title := " " + tab.Title + " "
		var tabStyle lipgloss.Style

		if i == t.active {
			// Active tab
			var activeColors StateColors
			if t.focused {
				activeColors = theme.Components.Tab.Active.Focused
			} else {
				activeColors = theme.Components.Tab.Active.Unfocused
			}
			tabStyle = lipgloss.NewStyle().
				Foreground(activeColors.Text).
				Background(theme.Palette.Background).
				Bold(true)
		} else {
			// Inactive tab
			tabStyle = lipgloss.NewStyle().
				Foreground(theme.Components.Tab.Inactive.Text).
				Background(theme.Palette.Background)
		}

		screen.DrawString(currentX, y, title, tabStyle)
		currentX += StringWidth(title)
	}

	// Fill rest of top border
	if t.focused {
		for currentX < x+width-1 {
			screen.DrawRune(currentX, y, '━', borderStyle)
			currentX++
		}
		screen.DrawRune(x+width-1, y, '┓', borderStyle)
	} else {
		for currentX < x+width-1 {
			screen.DrawRune(currentX, y, '─', borderStyle)
			currentX++
		}
		screen.DrawRune(x+width-1, y, '┐', borderStyle)
	}

	// Draw sides - heavy when focused
	for i := 1; i < height-1; i++ {
		if t.focused {
			screen.DrawRune(x, y+i, '┃', borderStyle)
			screen.DrawRune(x+width-1, y+i, '┃', borderStyle)
		} else {
			screen.DrawRune(x, y+i, '│', borderStyle)
			screen.DrawRune(x+width-1, y+i, '│', borderStyle)
		}
	}

	// Draw content
	t.drawContent(screen, x, y+1, width, height-2, theme, contentStyle)

	// Draw bottom border - heavy when focused
	if t.focused {
		screen.DrawRune(x, y+height-1, '┗', borderStyle)
		for i := 1; i < width-1; i++ {
			screen.DrawRune(x+i, y+height-1, '━', borderStyle)
		}
		screen.DrawRune(x+width-1, y+height-1, '┛', borderStyle)
	} else {
		screen.DrawRune(x, y+height-1, '└', borderStyle)
		for i := 1; i < width-1; i++ {
			screen.DrawRune(x+i, y+height-1, '─', borderStyle)
		}
		screen.DrawRune(x+width-1, y+height-1, '┘', borderStyle)
	}
}

// drawTabsOnBottom draws tabs on the bottom border
func (t *TabsComponent) drawTabsOnBottom(screen *Screen, x, y, width, height int, theme *Theme, borderStyle, contentStyle lipgloss.Style) {
	// Draw top border - heavy when focused
	if t.focused {
		screen.DrawRune(x, y, '┏', borderStyle)
		for i := 1; i < width-1; i++ {
			screen.DrawRune(x+i, y, '━', borderStyle)
		}
		screen.DrawRune(x+width-1, y, '┓', borderStyle)
	} else {
		screen.DrawRune(x, y, '┌', borderStyle)
		for i := 1; i < width-1; i++ {
			screen.DrawRune(x+i, y, '─', borderStyle)
		}
		screen.DrawRune(x+width-1, y, '┐', borderStyle)
	}

	// Draw sides - heavy when focused
	for i := 1; i < height-1; i++ {
		if t.focused {
			screen.DrawRune(x, y+i, '┃', borderStyle)
			screen.DrawRune(x+width-1, y+i, '┃', borderStyle)
		} else {
			screen.DrawRune(x, y+i, '│', borderStyle)
			screen.DrawRune(x+width-1, y+i, '│', borderStyle)
		}
	}

	// Draw content
	t.drawContent(screen, x, y+1, width, height-2, theme, contentStyle)

	// Draw bottom border with embedded tab titles
	if t.focused {
		screen.DrawRune(x, y+height-1, '┗', borderStyle)
	} else {
		screen.DrawRune(x, y+height-1, '└', borderStyle)
	}

	currentX := x + 1
	for i, tab := range t.tabs {
		if i > 0 {
			// Draw separator between tabs - bold when focused
			if t.focused {
				screen.DrawRune(currentX, y+height-1, '┃', borderStyle)
			} else {
				screen.DrawRune(currentX, y+height-1, '─', borderStyle)
			}
			currentX++
		}

		// Determine tab style based on state
		title := " " + tab.Title + " "
		var tabStyle lipgloss.Style

		if i == t.active {
			// Active tab
			var activeColors StateColors
			if t.focused {
				activeColors = theme.Components.Tab.Active.Focused
			} else {
				activeColors = theme.Components.Tab.Active.Unfocused
			}
			tabStyle = lipgloss.NewStyle().
				Foreground(activeColors.Text).
				Background(theme.Palette.Background).
				Bold(true)
		} else {
			// Inactive tab
			tabStyle = lipgloss.NewStyle().
				Foreground(theme.Components.Tab.Inactive.Text).
				Background(theme.Palette.Background)
		}

		screen.DrawString(currentX, y+height-1, title, tabStyle)
		currentX += StringWidth(title)
	}

	// Fill rest of bottom border - heavy when focused
	if t.focused {
		for currentX < x+width-1 {
			screen.DrawRune(currentX, y+height-1, '━', borderStyle)
			currentX++
		}
		screen.DrawRune(x+width-1, y+height-1, '┛', borderStyle)
	} else {
		for currentX < x+width-1 {
			screen.DrawRune(currentX, y+height-1, '─', borderStyle)
			currentX++
		}
		screen.DrawRune(x+width-1, y+height-1, '┘', borderStyle)
	}
}

// drawContent draws the content of the active tab
func (t *TabsComponent) drawContent(screen *Screen, x, y, width, height int, theme *Theme, contentStyle lipgloss.Style) {
	activeTab := t.GetActiveTab()
	if activeTab == nil {
		return
	}

	// Handle different content types
	switch content := activeTab.Content.(type) {
	case string:
		// Draw string content
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			if i >= height {
				break
			}
			// Clear the line first
			for dx := x + 1; dx < x + width - 1; dx++ {
				screen.DrawRune(dx, y+i, ' ', contentStyle)
			}
			// Draw content
			if StringWidth(line) > width-4 {
				line = Truncate(line, width-4)
			}
			screen.DrawString(x+2, y+i, line, contentStyle)
		}
	case Component:
		// Draw component content
		content.Draw(screen, x+1, y, theme)
	}
}

// Focus gives keyboard focus to this component
func (t *TabsComponent) Focus() {
	t.focused = true
}

// Blur removes keyboard focus from this component
func (t *TabsComponent) Blur() {
	t.focused = false
}

// IsFocused returns whether this component currently has focus
func (t *TabsComponent) IsFocused() bool {
	return t.focused
}

// HandleKey processes keyboard input when focused
func (t *TabsComponent) HandleKey(key string) bool {
	switch key {
	case "left", "h":
		t.PrevTab()
		return true
	case "right", "l":
		t.NextTab()
		return true
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		// Quick tab switching
		index := int(key[0] - '1')
		if index < len(t.tabs) {
			t.SetActive(index)
			return true
		}
	}
	
	// Pass key to active tab content if it's a component
	if activeTab := t.GetActiveTab(); activeTab != nil {
		if component, ok := activeTab.Content.(Component); ok {
			return component.HandleKey(key)
		}
	}
	
	return false
}

// SetSize sets the width and height of the component
func (t *TabsComponent) SetSize(width, height int) {
	t.width = width
	t.height = height
}

// GetSize returns the current width and height
func (t *TabsComponent) GetSize() (width, height int) {
	return t.width, t.height
}

// DrawWithBorder draws the component with a border and optional title
func (t *TabsComponent) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	// Tabs already have borders, so just call Draw
	t.Draw(screen, x, y, theme)
}