package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type model struct {
	screen        *tui.Screen
	currentLayout int
	layouts       []layoutExample
}

type layoutExample struct {
	name string
	draw func(*tui.Screen, *tui.Theme)
}

func main() {
	// Create model
	m := initialModel()

	// Start tea program
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	theme := tui.GetTheme("default")
	m := model{
		screen:        tui.NewScreen(80, 24, theme),
		currentLayout: 0,
	}

	// Define layout examples
	m.layouts = []layoutExample{
		{"Linear Layout - Horizontal", m.drawLinearHorizontal},
		{"Linear Layout - Vertical", m.drawLinearVertical},
		{"Split Layout with Constraints", m.drawSplitLayout},
		{"Stack Layout - Modals", m.drawStackLayout},
		{"Responsive Layout", m.drawResponsiveLayout},
		{"Complex Layout - Fuzzy Finder", m.drawFuzzyFinder},
	}

	return m
}

// Init initializes the model
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Recreate screen with new size
		theme := tui.GetTheme("default")
		m.screen = tui.NewScreen(msg.Width, msg.Height, theme)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "left":
			if m.currentLayout > 0 {
				m.currentLayout--
			}
		case "right":
			if m.currentLayout < len(m.layouts)-1 {
				m.currentLayout++
			}
		}
	}

	return m, nil
}

// View renders the view
func (m model) View() string {
	// Get theme
	theme := tui.GetTheme("default")

	// Clear screen
	m.screen.Clear()

	// Draw current layout
	m.layouts[m.currentLayout].draw(m.screen, &theme)

	// Draw navigation hint
	width := m.screen.Width()
	height := m.screen.Height()
	hint := fmt.Sprintf(" Layout %d/%d: %s | ← → Navigate | Q Quit ",
		m.currentLayout+1, len(m.layouts), m.layouts[m.currentLayout].name)
	hintX := (width - len(hint)) / 2
	hintStyle := lipgloss.NewStyle().Foreground(theme.Palette.Text)
	m.screen.DrawString(hintX, height-1, hint, hintStyle)

	// Render
	return m.screen.Render()
}

func (m *model) drawLinearHorizontal(screen *tui.Screen, theme *tui.Theme) {
	width := screen.Width()
	height := screen.Height()

	// Create main container
	main := tui.NewContainer()
	main.SetTitle("Horizontal Linear Layout")
	main.SetSize(width-4, height-4)

	// Create horizontal layout
	layout := tui.HBox()
	layout.SetSpacing(1)
	layout.SetPadding(tui.NewMargin(1))

	// Add fixed-width sidebar
	sidebar := createPanel("Sidebar", "Fixed 20 cells", theme.Palette.Surface)
	layout.AddFixed(sidebar, 20)

	// Add flexible content area
	content := createPanel("Content", "Flex 1", theme.Palette.Background)
	layout.AddFlex(content, 1)

	// Add percentage-based panel
	info := createPanel("Info Panel", "30% width", theme.Palette.Surface)
	layout.AddPercentage(info, 0.3)

	main.SetContent(layout)
	main.Draw(screen, 2, 1, theme)
}

func (m *model) drawLinearVertical(screen *tui.Screen, theme *tui.Theme) {
	width := screen.Width()
	height := screen.Height()

	// Create main container
	main := tui.NewContainer()
	main.SetTitle("Vertical Linear Layout")
	main.SetSize(width-4, height-4)

	// Create vertical layout
	layout := tui.VBox()
	layout.SetSpacing(1)
	layout.SetPadding(tui.NewMargin(1))

	// Add fixed header
	header := createPanel("Header", "Fixed 3 rows", theme.Palette.Primary)
	layout.AddFixed(header, 3)

	// Add flexible content
	content := createPanel("Content", "Flex 2", theme.Palette.Background)
	layout.AddFlex(content, 2)

	// Add another flex area
	details := createPanel("Details", "Flex 1", theme.Palette.Surface)
	layout.AddFlex(details, 1)

	// Add fixed footer
	footer := createPanel("Footer", "Fixed 3 rows", theme.Palette.Secondary)
	layout.AddFixed(footer, 3)

	main.SetContent(layout)
	main.Draw(screen, 2, 1, theme)
}

func (m *model) drawSplitLayout(screen *tui.Screen, theme *tui.Theme) {
	width := screen.Width()
	height := screen.Height()

	// Create main container
	main := tui.NewContainer()
	main.SetTitle("Split Layout with Constraints")
	main.SetSize(width-4, height-4)

	// Create vertical split
	vsplit := tui.NewVSplit()
	vsplit.SetConstraint(
		tui.NewConstraintSet(tui.NewPercentage(0.3)).
			WithMin(20).
			WithMax(40),
	)

	// Left side - file tree
	fileTree := createPanel("File Tree", "30% (min 20, max 40)", theme.Palette.Surface)
	vsplit.SetFirst(fileTree)

	// Right side - horizontal split
	hsplit := tui.NewHSplit()
	hsplit.SetFixed(10) // Fixed height terminal

	// Editor
	editor := createPanel("Editor", "Remaining space", theme.Palette.Background)
	hsplit.SetFirst(editor)

	// Terminal
	terminal := createPanel("Terminal", "Fixed 10 rows", theme.Palette.Surface)
	hsplit.SetSecond(terminal)

	vsplit.SetSecond(hsplit)

	main.SetContent(vsplit)
	main.Draw(screen, 2, 1, theme)
}

func (m *model) drawStackLayout(screen *tui.Screen, theme *tui.Theme) {
	width := screen.Width()
	height := screen.Height()

	// Create stack
	stack := tui.NewStack()
	stack.SetSize(width, height)

	// Background layer
	bg := createColorPanel("Background", theme.Palette.Background)
	stack.AddFull(bg)

	// Main content
	content := tui.NewContainer()
	content.SetTitle("Main Content")
	content.SetPadding(tui.NewMargin(2))
	contentViewer := tui.NewViewer()
	contentViewer.SetContent("This is the main content area.\nThe modal appears on top.")
	content.SetContent(contentViewer)
	stack.AddAnchored(content, tui.AlignCenter, tui.AlignCenter,
		tui.NewConstraintSet(tui.NewPercentage(0.8)),
		tui.NewConstraintSet(tui.NewPercentage(0.6)))

	// Modal overlay
	modal := tui.NewModal()
	modal.SetSize(60, 20)

	modalContainer := tui.NewContainer()
	modalContainer.SetTitle("Modal Dialog")
	modalContainer.SetSize(60, 20)
	modalContainer.SetPadding(tui.NewMargin(2))

	modalViewer := tui.NewViewer()
	modalViewer.SetContent("This is a modal dialog.\nIt appears on top of the content.")
	modalContainer.SetContent(modalViewer)

	stack.AddCentered(modal,
		tui.NewConstraintSet(tui.NewLength(60)),
		tui.NewConstraintSet(tui.NewLength(20)))

	stack.Draw(screen, 0, 0, theme)
}

func (m *model) drawResponsiveLayout(screen *tui.Screen, theme *tui.Theme) {
	width := screen.Width()
	height := screen.Height()

	// Create responsive layout
	responsive := tui.NewResponsiveLayout()
	responsive.SetSize(width, height)

	// Mobile layout (< 80 columns)
	mobileLayout := tui.VBox()
	mobileLayout.SetPadding(tui.NewMargin(1))

	mobileHeader := createPanel("Mobile Header", "Full width", theme.Palette.Primary)
	mobileLayout.AddFixed(mobileHeader, 3)

	mobileContent := createPanel("Mobile Content", "Stacked layout", theme.Palette.Background)
	mobileLayout.AddFlex(mobileContent, 1)

	responsive.AddMobile(mobileLayout)

	// Desktop layout (>= 80 columns)
	desktopLayout := tui.HBox()
	desktopLayout.SetPadding(tui.NewMargin(2))
	desktopLayout.SetSpacing(2)

	sidebar := createPanel("Sidebar", "Desktop only", theme.Palette.Surface)
	desktopLayout.AddFixed(sidebar, 30)

	content := createPanel("Main Content", fmt.Sprintf("Width: %d", width), theme.Palette.Background)
	desktopLayout.AddFlex(content, 1)

	responsive.AddDesktop(desktopLayout)

	// Draw with info
	responsive.Draw(screen, 0, 0, theme)

	// Show current mode
	mode := "Desktop"
	if width < 80 {
		mode = "Mobile"
	}
	info := fmt.Sprintf(" Responsive Mode: %s (width: %d) ", mode, width)
	infoStyle := lipgloss.NewStyle().Foreground(theme.Palette.Text)
	screen.DrawString(2, height-3, info, infoStyle)
}

func (m *model) drawFuzzyFinder(screen *tui.Screen, theme *tui.Theme) {
	width := screen.Width()
	height := screen.Height()

	// Create modal
	modal := tui.NewModal()
	modal.SetSize(90, 30)

	// Main vertical layout
	mainLayout := tui.VBox()
	mainLayout.SetSize(90, 30)
	mainLayout.SetSpacing(1)
	mainLayout.SetPadding(tui.NewMargin(1))

	// Search container
	searchContainer := tui.NewContainer()
	searchContainer.SetTitle("Search")
	searchContainer.SetPadding(tui.NewMargin(1))
	searchViewer := tui.NewViewer()
	searchViewer.SetContent("Type to search files...")
	searchContainer.SetContent(searchViewer)
	mainLayout.AddFixed(searchContainer, 5)

	// Results/Preview split
	split := tui.NewVSplit()
	split.SetFixed(30) // Fixed width results

	// Results container
	resultsContainer := tui.NewContainer()
	resultsContainer.SetTitle("Results")
	resultsContainer.SetPadding(tui.NewMargin(1))
	resultsList := createFileList()
	resultsContainer.SetContent(resultsList)
	split.SetFirst(resultsContainer)

	// Preview container
	previewContainer := tui.NewContainer()
	previewContainer.SetTitle("Preview")
	previewContainer.SetPadding(tui.NewMargin(1))
	previewContent := createPreviewContent()
	previewContainer.SetContent(previewContent)
	split.SetSecond(previewContainer)

	mainLayout.AddFlex(split, 1)

	// Modal is just an elevated surface, draw mainLayout on top of it

	// Center the modal
	modalX := (width - 90) / 2
	modalY := (height - 30) / 2
	modal.Draw(screen, modalX, modalY, theme)

	// Draw the layout on top of the modal
	mainLayout.Draw(screen, modalX, modalY, theme)
}

// Helper functions

func createPanel(title, info string, bgColor lipgloss.TerminalColor) *tui.Container {
	container := tui.NewContainer()
	container.SetTitle(title)
	container.SetPadding(tui.NewMargin(1))

	viewer := tui.NewViewer()
	viewer.SetContent(info)
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
	style := lipgloss.NewStyle().Background(p.color)
	width := screen.Width()
	height := screen.Height()
	for dy := 0; dy < height-y; dy++ {
		for dx := 0; dx < width-x; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', style)
		}
	}
}

func (p *colorPanel) Focus()                    {}
func (p *colorPanel) Blur()                     {}
func (p *colorPanel) IsFocused() bool           { return false }
func (p *colorPanel) HandleKey(key string) bool { return false }

func createFileList() tui.Component {
	files := []string{
		"main.go",
		"config.go",
		"utils.go",
		"server.go",
		"client.go",
		"auth.go",
		"database.go",
		"middleware.go",
		"routes.go",
		"models.go",
	}

	text := ""
	for _, file := range files {
		text += file + "\n"
	}

	viewer := tui.NewViewer()
	viewer.SetContent(text)
	return viewer
}

func createPreviewContent() tui.Component {
	preview := `package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("Hello from main.go")
    
    // Initialize application
    app := NewApp()
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}`

	viewer := tui.NewViewer()
	viewer.SetContent(preview)
	return viewer
}
