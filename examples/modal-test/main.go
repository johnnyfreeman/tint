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

	// Modal states
	currentModal int
	modals       []modalConfig
}

type modalConfig struct {
	modal       *tui.Modal
	description string
	drawFunc    func(screen *tui.Screen, x, y, width, height int, theme *tui.Theme) // Optional custom draw function
}

func initialModel() *model {
	theme := tui.GetTheme("tokyonight")

	// Create different modal configurations to test
	modals := []modalConfig{
		{
			modal:       createSimpleModal(),
			description: "Modal with container and text",
			drawFunc:    drawSimpleModal,
		},
		{
			modal:       createLargeModal(),
			description: "Large modal with container",
			drawFunc:    drawLargeModal,
		},
		{
			modal:       createSmallModal(),
			description: "Small modal with container",
			drawFunc:    drawSmallModal,
		},
		{
			modal:       createModalWithNewlines(),
			description: "Modal with formatted content",
			drawFunc:    drawFormattedModal,
		},
		{
			modal:       createEmptyModal(),
			description: "Empty modal (no content)",
			drawFunc:    nil, // Just the modal surface, nothing inside
		},
		{
			modal:       createModalForContainers(),
			description: "Modal with 3 containers inside",
			drawFunc:    drawMultipleContainers,
		},
	}

	// Show the first modal
	modals[0].modal.Show()

	return &model{
		screen:       tui.NewScreen(100, 30, theme),
		width:        100,
		height:       30,
		theme:        theme,
		currentModal: 0,
		modals:       modals,
	}
}

func createSimpleModal() *tui.Modal {
	modal := tui.NewModal()
	modal.SetSize(40, 8)
	modal.SetCentered(true)
	return modal
}

func createLargeModal() *tui.Modal {
	modal := tui.NewModal()
	modal.SetSize(60, 16)
	modal.SetCentered(true)
	return modal
}

func createSmallModal() *tui.Modal {
	modal := tui.NewModal()
	modal.SetSize(20, 6)
	modal.SetCentered(true)
	return modal
}

func createModalWithNewlines() *tui.Modal {
	modal := tui.NewModal()
	modal.SetSize(45, 12)
	modal.SetCentered(true)
	return modal
}

func createEmptyModal() *tui.Modal {
	modal := tui.NewModal()
	modal.SetSize(50, 15)
	modal.SetCentered(true)
	return modal
}

func createModalForContainers() *tui.Modal {
	modal := tui.NewModal()
	modal.SetSize(70, 20)
	modal.SetCentered(true)
	return modal
}

// Draw functions for modal content
func drawSimpleModal(screen *tui.Screen, modalX, modalY, width, height int, theme *tui.Theme) {
	// Create a container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Simple Modal")
	container.SetSize(40, 8) // Same size as modal
	container.SetPadding(tui.NewMargin(1))

	// Add text content
	textarea := tui.NewTextArea()
	textarea.SetValue("This is a simple modal with basic text content.")
	textarea.SetSize(36, 4) // Adjusted for full container
	container.SetContent(textarea)

	// Draw container at modal position (no extra offset)
	container.Draw(screen, modalX, modalY, width, height, theme)
}

func drawLargeModal(screen *tui.Screen, modalX, modalY, width, height int, theme *tui.Theme) {
	// Create a container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Large Modal")
	container.SetSize(60, 16) // Same size as modal
	container.SetPadding(tui.NewMargin(1))

	// Add text content
	textarea := tui.NewTextArea()
	textarea.SetValue("This is a large modal with more content.\n\nIt demonstrates how modals handle\nmultiple lines of text and how the\nSurface color provides visual elevation\nfrom the background.\n\nThe modal should clip content if it\nexceeds the available space.")
	textarea.SetSize(56, 12) // Adjusted for full container
	container.SetContent(textarea)

	// Draw container at modal position (no extra offset)
	container.Draw(screen, modalX, modalY, width, height, theme)
}

func drawSmallModal(screen *tui.Screen, modalX, modalY, width, height int, theme *tui.Theme) {
	// Create a container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Tiny")
	container.SetSize(20, 6) // Same size as modal
	container.SetPadding(tui.NewMargin(1))

	// Add text content
	textarea := tui.NewTextArea()
	textarea.SetValue("Small!")
	textarea.SetSize(16, 2) // Adjusted for full container
	container.SetContent(textarea)

	// Draw container at modal position (no extra offset)
	container.Draw(screen, modalX, modalY, width, height, theme)
}

func drawFormattedModal(screen *tui.Screen, modalX, modalY, width, height int, theme *tui.Theme) {
	// Create a container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Formatted Text")
	container.SetSize(45, 12) // Same size as modal
	container.SetPadding(tui.NewMargin(1))

	// Add formatted text content
	textarea := tui.NewTextArea()
	textarea.SetValue("Line 1: First line of text\n\nLine 3: After blank line\n\nLine 5: Another paragraph\n      - Indented item\n      - Another item")
	textarea.SetSize(41, 8) // Adjusted for full container
	container.SetContent(textarea)

	// Draw container at modal position (no extra offset)
	container.Draw(screen, modalX, modalY, width, height, theme)
}

// Custom draw function for the modal with multiple containers
func drawMultipleContainers(screen *tui.Screen, modalX, modalY, width, height int, theme *tui.Theme) {
	// Create first container
	container1 := tui.NewContainer()
	container1.SetTitle("Container 1")
	container1.SetSize(30, 8)
	container1.SetPadding(tui.NewMargin(1))

	// Add some content to container1
	input := tui.NewInput()
	input.SetPlaceholder("Type here...")
	input.SetWidth(25)
	container1.SetContent(input)

	// Create second container
	container2 := tui.NewContainer()
	container2.SetTitle("Container 2")
	container2.SetSize(30, 8)
	container2.SetPadding(tui.NewMargin(1))

	// Add content to container2
	textarea := tui.NewTextArea()
	textarea.SetValue("Some text\nin the\ntext area")
	textarea.SetSize(25, 5)
	container2.SetContent(textarea)

	// Create third container (empty)
	container3 := tui.NewContainer()
	container3.SetTitle("Empty Container")
	container3.SetSize(25, 12)
	container3.SetPadding(tui.NewMargin(2))

	// Draw containers inside the modal with small padding
	container1.Draw(screen, modalX+1, modalY+1, 30, 8, theme)
	container2.Draw(screen, modalX+34, modalY+1, 30, 8, theme)
	container3.Draw(screen, modalX+1, modalY+10, 25, 12, theme)
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Modal Test")
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

		case "escape", "esc":
			// Hide current modal
			m.modals[m.currentModal].modal.Hide()

		case "space", "enter":
			// Show current modal
			m.modals[m.currentModal].modal.Show()

		case "tab", "right", "l":
			// Next modal
			m.modals[m.currentModal].modal.Hide()
			m.currentModal = (m.currentModal + 1) % len(m.modals)
			m.modals[m.currentModal].modal.Show()

		case "shift+tab", "left", "h":
			// Previous modal
			m.modals[m.currentModal].modal.Hide()
			m.currentModal--
			if m.currentModal < 0 {
				m.currentModal = len(m.modals) - 1
			}
			m.modals[m.currentModal].modal.Show()

		case "c":
			// Toggle centered - since we can't read the current state,
			// we'll track it separately if needed
			// For now, just set to not centered
			modal := m.modals[m.currentModal].modal
			modal.SetCentered(false)

		case "1", "2", "3", "4", "5", "6":
			// Direct selection
			idx := int(msg.String()[0] - '1')
			if idx >= 0 && idx < len(m.modals) {
				m.modals[m.currentModal].modal.Hide()
				m.currentModal = idx
				m.modals[m.currentModal].modal.Show()
			}
		}
	}

	return m, nil
}

func (m *model) View() string {
	// Clear screen with theme background
	m.screen.Clear()

	// Draw instructions
	style := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Background)

	m.screen.DrawString(2, 2, "Modal Test - Interactive Demo", style)
	m.screen.DrawString(2, 4, "Controls:", style)
	m.screen.DrawString(4, 5, "Tab/Arrow keys: Switch between modals", style)
	m.screen.DrawString(4, 6, "Space/Enter: Show modal", style)
	m.screen.DrawString(4, 7, "Escape: Hide modal", style)
	m.screen.DrawString(4, 8, "C: Toggle centered/positioned", style)
	m.screen.DrawString(4, 9, "1-6: Select modal directly", style)
	m.screen.DrawString(4, 10, "Q: Quit", style)

	// Show current modal info
	infoStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Primary).
		Background(m.theme.Palette.Background)

	m.screen.DrawString(2, 12, fmt.Sprintf("Current Modal: [%d] %s", m.currentModal+1, m.modals[m.currentModal].description), infoStyle)

	// Draw modal selection indicators
	for i, modalCfg := range m.modals {
		x := 2 + i*25
		y := 14

		if i == m.currentModal {
			// Highlight current
			highlightStyle := lipgloss.NewStyle().
				Foreground(m.theme.Palette.Background).
				Background(m.theme.Palette.Primary)
			m.screen.DrawString(x, y, fmt.Sprintf(" %d: %s ", i+1, modalCfg.description), highlightStyle)
		} else {
			m.screen.DrawString(x, y, fmt.Sprintf(" %d: %s ", i+1, modalCfg.description), style)
		}
	}

	// Draw the current modal
	currentModalConfig := m.modals[m.currentModal]
	currentModal := currentModalConfig.modal
	if currentModal.IsVisible() {
		// First draw the modal itself
		currentModal.Draw(m.screen, 0, 0, m.width, m.height, &m.theme)

		// Then draw any custom content if there's a draw function
		if currentModalConfig.drawFunc != nil {
			// Calculate modal position for centered modals
			modalWidth, modalHeight := currentModal.GetSize()
			modalX := (m.width - modalWidth) / 2
			modalY := (m.height - modalHeight) / 2
			currentModalConfig.drawFunc(m.screen, modalX, modalY, modalWidth, modalHeight, &m.theme)
		}
	}

	// Show modal state
	stateStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Background)

	visible := "Hidden"
	if currentModal.IsVisible() {
		visible = "Visible"
	}

	m.screen.DrawString(2, m.height-3, fmt.Sprintf("State: %s", visible), stateStyle)

	return m.screen.Render()
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
