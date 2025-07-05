package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type model struct {
	screen *tui.Screen
	width  int
	height int
	theme  tui.Theme
}

func initialModel() *model {
	theme := tui.GetTheme("tokyonight")
	return &model{
		screen: tui.NewScreen(80, 24, theme),
		width:  80,
		height: 24,
		theme:  theme,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Tint Screen Example")
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.screen = tui.NewScreen(m.width, m.height, m.theme)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *model) View() string {
	// Clear screen - this fills ALL cells with theme background
	m.screen.Clear()

	// Just draw some text to show the screen is working
	style := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Background)

	m.screen.DrawString(2, 2, "Screen with Container and Modal Example", style)
	m.screen.DrawString(2, 4, "Press 'q' to quit", style)
	m.screen.DrawString(2, 6, fmt.Sprintf("Screen size: %dx%d", m.width, m.height), style)

	// Create and draw a container
	container := tui.NewContainer()
	container.SetTitle("Test Container")
	container.SetSize(40, 10)
	container.SetPadding(tui.NewMargin(2))

	// Draw the container
	container.Draw(m.screen, 10, 10, &m.theme)

	// Create and draw a simple modal with a container inside
	modal1 := tui.NewModal()
	modal1.SetSize(40, 10)
	modal1.Show()

	// Draw the modal
	modal1X := 55
	modal1Y := 10
	modal1.Draw(m.screen, modal1X, modal1Y, &m.theme)

	// Create a container that fills the modal
	container1 := tui.NewContainer()
	container1.SetTitle("Simple Modal")
	container1.SetSize(40, 10) // Same size as modal
	container1.SetPadding(tui.NewMargin(1))

	// Add content to container
	textarea := tui.NewTextArea()
	textarea.SetValue("This is a simple modal\nwith some text content")
	textarea.SetSize(36, 6) // Adjusted for full container
	container1.SetContent(textarea)

	// Draw the container at modal position
	container1.Draw(m.screen, modal1X, modal1Y, &m.theme)

	// Create a modal with a container inside
	modal2 := tui.NewModal()
	modal2.SetSize(45, 12)
	modal2.Show()

	// We need to draw the modal first, then draw a container inside it
	// Calculate position for the modal
	modal2X := 25
	modal2Y := 22
	modal2.Draw(m.screen, modal2X, modal2Y, &m.theme)

	// Create a container that fills the modal
	innerContainer := tui.NewContainer()
	innerContainer.SetTitle("Inner Container")
	innerContainer.SetSize(45, 12) // Same size as modal
	innerContainer.SetPadding(tui.NewMargin(1))

	// Draw the container at modal position
	innerContainer.Draw(m.screen, modal2X, modal2Y, &m.theme)

	return m.screen.Render()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
