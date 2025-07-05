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
	modal  *tui.Modal
	showModal bool
}

func initialModel() *model {
	theme := tui.GetTheme("tokyonight")
	
	// Create a modal
	modal := tui.NewModal()
	modal.SetSize(40, 12)
	modal.SetCentered(true)
	
	return &model{
		screen: tui.NewScreen(80, 24, theme),
		width:  80,
		height: 24,
		theme:  theme,
		modal:  modal,
		showModal: false,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Simple Modal Example")
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
		case "space", "enter":
			m.showModal = !m.showModal
			if m.showModal {
				m.modal.Show()
			} else {
				m.modal.Hide()
			}
		case "escape", "esc":
			m.showModal = false
			m.modal.Hide()
		}
	}
	
	return m, nil
}

func (m *model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}
	
	m.screen.Clear()
	
	// Draw background content
	style := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Background)
	
	// Fill screen with pattern
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if (x+y)%4 == 0 {
				m.screen.DrawString(x, y, "·", style)
			}
		}
	}
	
	// Draw instructions
	instructionStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Primary).
		Background(m.theme.Palette.Background).
		Bold(true)
	
	instructions := []string{
		"Simple Modal Example",
		"",
		"Press SPACE or ENTER to toggle modal",
		"Press ESC to close modal",
		"Press Q to quit",
	}
	
	for i, line := range instructions {
		x := (m.width - len(line)) / 2
		m.screen.DrawString(x, 2+i, line, instructionStyle)
	}
	
	// Draw modal if visible
	if m.showModal {
		m.modal.Draw(m.screen, 0, 0, &m.theme)
		
		// Get modal position for container
		modalWidth, modalHeight := m.modal.GetSize()
		modalX := (m.width - modalWidth) / 2
		modalY := (m.height - modalHeight) / 2
		
		// Create a container that fills the modal
		container := tui.NewContainer()
		container.SetTitle("Hello Modal!")
		container.SetSize(modalWidth, modalHeight)
		container.SetPadding(tui.NewMargin(1))
		
		// Add content
		textarea := tui.NewTextArea()
		textarea.SetValue("This is a simple modal example.\n\nThe modal provides:\n• Elevated surface\n• Drop shadow\n\nThe container provides:\n• Border and title\n• Padding")
		textarea.SetSize(modalWidth-4, modalHeight-4)
		container.SetContent(textarea)
		
		// Draw the container
		container.Draw(m.screen, modalX, modalY, &m.theme)
	}
	
	// Show status
	statusStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Background)
	
	status := "Modal is hidden"
	if m.showModal {
		status = "Modal is visible"
	}
	m.screen.DrawString(2, m.height-2, status, statusStyle)
	
	return m.screen.Render()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}