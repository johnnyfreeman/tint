package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

// This example showcases all the new layout components and how to use them

type model struct {
	screen       *tui.Screen
	currentDemo  int
	demos        []layoutDemo
	width        int
	height       int
}

type layoutDemo struct {
	name        string
	description string
	draw        func(*model, *tui.Theme)
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	theme := tui.GetTheme("default")
	m := model{
		screen:      tui.NewScreen(80, 24, theme),
		currentDemo: 0,
		width:       80,
		height:      24,
	}
	
	m.demos = []layoutDemo{
		{
			name:        "Linear Layout (HBox/VBox)",
			description: "Arrange components in rows or columns with flexible sizing",
			draw:        drawLinearDemo,
		},
		{
			name:        "Split Layout",
			description: "Split screen into two panes with constraints",
			draw:        drawSplitDemo,
		},
		{
			name:        "Stack Layout",
			description: "Layer components on top of each other",
			draw:        drawStackDemo,
		},
		{
			name:        "Responsive Layout",
			description: "Different layouts based on screen size",
			draw:        drawResponsiveDemo,
		},
		{
			name:        "Complex App Layout",
			description: "Combining multiple layouts for a real app",
			draw:        drawComplexDemo,
		},
	}
	
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		theme := tui.GetTheme("default")
		m.screen = tui.NewScreen(msg.Width, msg.Height, theme)
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "left", "h":
			if m.currentDemo > 0 {
				m.currentDemo--
			}
		case "right", "l":
			if m.currentDemo < len(m.demos)-1 {
				m.currentDemo++
			}
		case "1", "2", "3", "4", "5":
			idx := int(msg.String()[0] - '1')
			if idx >= 0 && idx < len(m.demos) {
				m.currentDemo = idx
			}
		}
	}
	
	return m, nil
}

func (m model) View() string {
	theme := tui.GetTheme("default")
	m.screen.Clear()
	
	// Draw current demo
	m.demos[m.currentDemo].draw(&m, &theme)
	
	// Draw navigation footer
	footer := fmt.Sprintf(" %d/%d: %s | ← → Navigate | 1-5 Jump | Q Quit ",
		m.currentDemo+1, len(m.demos), m.demos[m.currentDemo].name)
	footerStyle := lipgloss.NewStyle().
		Background(theme.Palette.Surface).
		Foreground(theme.Palette.Text)
	
	footerX := (m.width - len(footer)) / 2
	m.screen.DrawString(footerX, m.height-1, footer, footerStyle)
	
	return m.screen.Render()
}

// Demo 1: Linear Layout
func drawLinearDemo(m *model, theme *tui.Theme) {
	// Create main container
	main := tui.NewContainer()
	main.SetTitle("Linear Layout Demo")
	main.SetSize(m.width-4, m.height-4)
	
	// Create vertical layout
	vbox := tui.VBox()
	vbox.SetSpacing(1)
	vbox.SetPadding(tui.NewMargin(1))
	
	// Add description
	desc := createTextPanel("Description", m.demos[m.currentDemo].description)
	vbox.AddFixed(desc, 4)
	
	// Create horizontal layout for examples
	hbox := tui.HBox()
	hbox.SetSpacing(2)
	
	// Add three panels with different sizing
	panel1 := createTextPanel("Fixed 20", "Width: 20 cells")
	hbox.AddFixed(panel1, 20)
	
	panel2 := createTextPanel("Flex 2", "Takes 2 shares of remaining space")
	hbox.AddFlex(panel2, 2)
	
	panel3 := createTextPanel("Flex 1", "Takes 1 share of remaining space")
	hbox.AddFlex(panel3, 1)
	
	vbox.AddFlex(hbox, 1)
	
	// Add bottom panel
	bottom := createTextPanel("Percentage", "Takes 30% of vertical space")
	vbox.AddPercentage(bottom, 0.3)
	
	main.SetContent(vbox)
	main.Draw(m.screen, 2, 1, theme)
}

// Demo 2: Split Layout
func drawSplitDemo(m *model, theme *tui.Theme) {
	// Create main container
	main := tui.NewContainer()
	main.SetTitle("Split Layout Demo")
	main.SetSize(m.width-4, m.height-4)
	
	// Create vertical split
	vsplit := tui.NewVSplit()
	vsplit.SetConstraint(
		tui.NewConstraintSet(tui.NewPercentage(0.3)).
			WithMin(15).
			WithMax(40),
	)
	
	// Left pane
	left := createTextPanel("Left Pane", "30% width\nMin: 15 cells\nMax: 40 cells")
	vsplit.SetFirst(left)
	
	// Right pane with horizontal split
	hsplit := tui.NewHSplit()
	hsplit.SetFixed(8)
	
	top := createTextPanel("Top Pane", "Remaining width\nFixed 8 rows height")
	hsplit.SetFirst(top)
	
	bottom := createTextPanel("Bottom Pane", "Remaining width\nRemaining height")
	hsplit.SetSecond(bottom)
	
	vsplit.SetSecond(hsplit)
	
	main.SetContent(vsplit)
	main.Draw(m.screen, 2, 1, theme)
}

// Demo 3: Stack Layout
func drawStackDemo(m *model, theme *tui.Theme) {
	// Create stack
	stack := tui.NewStack()
	stack.SetSize(m.width, m.height-1)
	
	// Background layer
	bg := createColorPanel("Background Layer", theme.Palette.Background)
	stack.AddFull(bg)
	
	// Middle layer - large panel
	middle := tui.NewContainer()
	middle.SetTitle("Middle Layer")
	middle.SetPadding(tui.NewMargin(2))
	middleText := tui.NewViewer()
	middleText.SetContent("This panel is 70% width and 60% height\nCentered on screen")
	middle.SetContent(middleText)
	
	stack.AddCentered(middle,
		tui.NewConstraintSet(tui.NewPercentage(0.7)),
		tui.NewConstraintSet(tui.NewPercentage(0.6)))
	
	// Top layer - modal
	modal := tui.NewModal()
	modal.SetSize(50, 12)
	modal.Show()
	
	modalContainer := tui.NewContainer()
	modalContainer.SetTitle("Top Layer (Modal)")
	modalContainer.SetSize(50, 12)
	modalContainer.SetPadding(tui.NewMargin(1))
	
	modalText := tui.NewViewer()
	modalText.SetContent("This is a modal on top\nIt has a shadow effect\nPress ESC to close (in real app)")
	modalContainer.SetContent(modalText)
	
	// Create a component that draws both modal and container
	modalWrapper := &modalWithContainer{
		modal:     modal,
		container: modalContainer,
	}
	
	stack.AddCentered(modalWrapper,
		tui.NewConstraintSet(tui.NewLength(50)),
		tui.NewConstraintSet(tui.NewLength(12)))
	
	stack.Draw(m.screen, 0, 0, theme)
}

// Demo 4: Responsive Layout
func drawResponsiveDemo(m *model, theme *tui.Theme) {
	responsive := tui.NewResponsiveLayout()
	responsive.SetSize(m.width, m.height-1)
	
	// Mobile layout (< 60 columns)
	mobileLayout := tui.VBox()
	mobileLayout.SetPadding(tui.NewMargin(2))
	mobileLayout.SetSpacing(1)
	
	mobileHeader := createTextPanel("Mobile Layout", "Screen width < 60")
	mobileLayout.AddFixed(mobileHeader, 5)
	
	mobileContent := createTextPanel("Stacked Content", "Everything is stacked vertically\nfor narrow screens")
	mobileLayout.AddFlex(mobileContent, 1)
	
	responsive.AddMaxSize(mobileLayout, 59, 9999)
	
	// Tablet layout (60-100 columns)
	tabletLayout := tui.HBox()
	tabletLayout.SetPadding(tui.NewMargin(2))
	tabletLayout.SetSpacing(2)
	
	tabletSidebar := createTextPanel("Sidebar", "Tablet\nLayout")
	tabletLayout.AddFixed(tabletSidebar, 20)
	
	tabletContent := createTextPanel("Tablet Content", "60 ≤ width < 100\nSidebar + Content")
	tabletLayout.AddFlex(tabletContent, 1)
	
	responsive.AddWidthRange(tabletLayout, 60, 99)
	
	// Desktop layout (≥ 100 columns)
	desktopLayout := tui.HBox()
	desktopLayout.SetPadding(tui.NewMargin(2))
	desktopLayout.SetSpacing(2)
	
	desktopLeft := createTextPanel("Left", "Desktop")
	desktopLayout.AddFixed(desktopLeft, 25)
	
	desktopCenter := createTextPanel("Center", "Width ≥ 100\nThree column layout")
	desktopLayout.AddFlex(desktopCenter, 1)
	
	desktopRight := createTextPanel("Right", "Info\nPanel")
	desktopLayout.AddFixed(desktopRight, 25)
	
	responsive.AddMinSize(desktopLayout, 100, 0)
	
	// Add current size indicator
	sizeInfo := fmt.Sprintf("Current size: %d×%d", m.width, m.height)
	style := lipgloss.NewStyle().Foreground(theme.Palette.Primary)
	m.screen.DrawString(2, 0, sizeInfo, style)
	
	responsive.Draw(m.screen, 0, 1, theme)
}

// Demo 5: Complex App Layout
func drawComplexDemo(m *model, theme *tui.Theme) {
	// Main horizontal split for sidebar
	mainSplit := tui.NewVSplit()
	mainSplit.SetFixed(25)
	mainSplit.SetSize(m.width, m.height-1)
	
	// Sidebar
	sidebar := createSidebar()
	mainSplit.SetFirst(sidebar)
	
	// Main content area with vertical layout
	contentLayout := tui.VBox()
	contentLayout.SetSpacing(0)
	
	// Header
	header := createHeader()
	contentLayout.AddFixed(header, 3)
	
	// Content area with tabs
	tabContent := createTabbedContent()
	contentLayout.AddFlex(tabContent, 1)
	
	// Status bar
	statusBar := createStatusBar()
	contentLayout.AddFixed(statusBar, 1)
	
	mainSplit.SetSecond(contentLayout)
	
	// Draw everything
	mainSplit.Draw(m.screen, 0, 0, theme)
	
	// Add floating notification using stack
	if m.width > 80 {
		drawNotification(m, theme)
	}
}

// Helper functions

func createTextPanel(title, content string) *tui.Container {
	container := tui.NewContainer()
	container.SetTitle(title)
	container.SetPadding(tui.NewMargin(1))
	
	viewer := tui.NewViewer()
	viewer.SetContent(content)
	container.SetContent(viewer)
	
	return container
}

func createColorPanel(text string, color lipgloss.TerminalColor) tui.Component {
	return &colorPanel{text: text, color: color}
}

type colorPanel struct {
	text  string
	color lipgloss.TerminalColor
}

func (p *colorPanel) Draw(screen *tui.Screen, x, y int, theme *tui.Theme) {
	// Fill with color
}

func (p *colorPanel) Focus() {}
func (p *colorPanel) Blur() {}
func (p *colorPanel) IsFocused() bool { return false }
func (p *colorPanel) HandleKey(key string) bool { return false }

// Modal with container wrapper
type modalWithContainer struct {
	modal     *tui.Modal
	container *tui.Container
}

func (m *modalWithContainer) Draw(screen *tui.Screen, x, y int, theme *tui.Theme) {
	m.modal.Draw(screen, x, y, theme)
	m.container.Draw(screen, x, y, theme)
}

func (m *modalWithContainer) Focus() { m.container.Focus() }
func (m *modalWithContainer) Blur() { m.container.Blur() }
func (m *modalWithContainer) IsFocused() bool { return m.container.IsFocused() }
func (m *modalWithContainer) HandleKey(key string) bool { return m.container.HandleKey(key) }

// Complex demo helpers

func createSidebar() *tui.Container {
	sidebar := tui.NewContainer()
	sidebar.SetTitle("Navigation")
	sidebar.SetPadding(tui.NewMargin(1))
	
	// Create menu using VBox
	menu := tui.VBox()
	menu.SetSpacing(1)
	
	items := []string{"Dashboard", "Files", "Search", "Settings", "Help"}
	for i, item := range items {
		menuItem := tui.NewViewer()
		if i == 0 {
			menuItem.SetContent("> " + item) // Selected
		} else {
			menuItem.SetContent("  " + item)
		}
		menu.AddFixed(menuItem, 1)
	}
	
	sidebar.SetContent(menu)
	return sidebar
}

func createHeader() *tui.Container {
	header := tui.NewContainer()
	header.SetTitle("Tint Layout System Demo")
	header.SetPadding(tui.NewMargin(1))
	
	viewer := tui.NewViewer()
	viewer.SetContent("Welcome to the comprehensive layout demo!")
	header.SetContent(viewer)
	
	return header
}

func createTabbedContent() *tui.Container {
	container := tui.NewContainer()
	container.SetPadding(tui.NewMargin(1))
	
	// Add tabs to border
	tabs := tui.NewTabsElement([]string{"Code", "Preview", "Console"})
	tabs.SetActiveTab(0)
	container.AddBorderElement(tabs, tui.BorderTop, tui.BorderAlignLeft)
	
	// Content
	content := tui.NewViewer()
	content.SetContent(`func main() {
    // This is the code tab
    fmt.Println("Hello, Tint!")
    
    // The new layout system makes it easy to create
    // complex, responsive UIs with minimal code
}`)
	
	container.SetContent(content)
	return container
}

func createStatusBar() *tui.Container {
	status := tui.NewContainer()
	status.SetShowBorder(false)
	status.SetPadding(tui.NewMargin(0))
	
	text := tui.NewViewer()
	text.SetContent(" Ready | Line 1, Col 1 | Go | UTF-8 | LF ")
	status.SetContent(text)
	
	return status
}

func drawNotification(m *model, theme *tui.Theme) {
	// Create a small notification in top-right
	notif := tui.NewContainer()
	notif.SetTitle("Notification")
	notif.SetPadding(tui.NewMargin(1))
	notif.SetBorderStyle("rounded")
	
	text := tui.NewViewer()
	text.SetContent("Layout system ready!")
	notif.SetContent(text)
	
	notif.SetSize(25, 5)
	notif.Draw(m.screen, m.width-27, 2, theme)
}