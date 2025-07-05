package main

import (
	"fmt"
	"strings"

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
	showInfoModal    bool
	showConfirmModal bool
	showFormModal    bool

	// Form data
	formName  string
	formEmail string

	// UI components
	infoModal    *tui.Modal
	confirmModal *tui.Modal
	formModal    *tui.Modal
}

func initialModel() *model {
	theme := tui.GetTheme("tokyonight")

	// Create modals
	infoModal := tui.NewModal()
	infoModal.SetSize(50, 15)
	infoModal.SetCentered(true)

	confirmModal := tui.NewModal()
	confirmModal.SetSize(40, 10)
	confirmModal.SetCentered(true)

	formModal := tui.NewModal()
	formModal.SetSize(60, 20)
	formModal.SetCentered(true)

	return &model{
		screen:       tui.NewScreen(80, 24, theme),
		width:        80,
		height:       24,
		theme:        theme,
		infoModal:    infoModal,
		confirmModal: confirmModal,
		formModal:    formModal,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Modal/Container Pattern Demo")
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.screen = tui.NewScreen(m.width, m.height, m.theme)

	case tea.KeyMsg:
		// Handle modal-specific input first
		if m.showFormModal {
			switch msg.String() {
			case "escape", "esc":
				m.showFormModal = false
				return m, nil
			case "enter":
				// Submit form
				m.showFormModal = false
				return m, nil
			default:
				// Handle form input
				if strings.Contains("abcdefghijklmnopqrstuvwxyz @.", msg.String()) {
					if len(m.formEmail) < 30 {
						m.formEmail += msg.String()
					}
				}
				if msg.String() == "backspace" && len(m.formEmail) > 0 {
					m.formEmail = m.formEmail[:len(m.formEmail)-1]
				}
			}
			return m, nil
		}

		if m.showConfirmModal {
			switch msg.String() {
			case "y", "Y":
				// Confirmed
				m.showConfirmModal = false
				return m, nil
			case "n", "N", "escape", "esc":
				// Cancelled
				m.showConfirmModal = false
				return m, nil
			}
			return m, nil
		}

		if m.showInfoModal {
			if msg.String() == "escape" || msg.String() == "esc" || msg.String() == "enter" {
				m.showInfoModal = false
			}
			return m, nil
		}

		// Main menu input
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.showInfoModal = true
		case "2":
			m.showConfirmModal = true
		case "3":
			m.showFormModal = true
			m.formEmail = "" // Reset form
		}
	}

	return m, nil
}

func (m *model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	m.screen.Clear()

	// Draw main content
	m.drawMainContent()

	// Draw modals on top
	if m.showInfoModal {
		m.drawInfoModal()
	}
	if m.showConfirmModal {
		m.drawConfirmModal()
	}
	if m.showFormModal {
		m.drawFormModal()
	}

	return m.screen.Render()
}

func (m *model) drawMainContent() {
	// Title
	titleStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Primary).
		Bold(true)
	m.screen.DrawString(2, 2, "Modal/Container Pattern Demo", titleStyle)

	// Instructions
	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text)
	instructions := []string{
		"This demo shows the proper pattern for modals:",
		"",
		"• Modal provides elevated surface with shadow",
		"• Container provides structure (border, title, padding)",
		"• Content goes inside the container",
		"",
		"Press number keys to show different modals:",
		"",
		"1. Information Modal - Simple content display",
		"2. Confirmation Modal - Yes/No dialog",
		"3. Form Modal - Input fields with labels",
		"",
		"Press 'q' to quit",
	}

	for i, line := range instructions {
		m.screen.DrawString(4, 4+i, line, textStyle)
	}

	// Show current state
	stateStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted)
	state := "No modal active"
	if m.showInfoModal {
		state = "Info modal is open (ESC to close)"
	} else if m.showConfirmModal {
		state = "Confirm modal is open (Y/N to respond)"
	} else if m.showFormModal {
		state = "Form modal is open (Type to edit, ENTER to submit)"
	}
	m.screen.DrawString(4, m.height-3, state, stateStyle)
}

func (m *model) drawInfoModal() {
	m.infoModal.Draw(m.screen, 0, 0, &m.theme)

	// Get modal position
	modalWidth, modalHeight := m.infoModal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Create container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Information")
	container.SetSize(modalWidth, modalHeight)
	container.SetPadding(tui.NewMargin(2))

	// Create content
	viewer := tui.NewViewer()
	viewer.SetContent(`Welcome to the Modal Demo!

This modal demonstrates the proper pattern:

1. The Modal component provides:
   • Elevated surface (different background)
   • Drop shadow for depth
   • Centering logic

2. The Container component provides:
   • Border and title
   • Padding for content
   • Focus management

Press ESC or ENTER to close this modal.`)
	viewer.SetWrapText(true)
	viewer.SetSize(modalWidth-6, modalHeight-6)

	container.SetContent(viewer)
	container.Draw(m.screen, modalX, modalY, &m.theme)
}

func (m *model) drawConfirmModal() {
	m.confirmModal.Draw(m.screen, 0, 0, &m.theme)

	// Get modal position
	modalWidth, modalHeight := m.confirmModal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Create container
	container := tui.NewContainer()
	container.SetTitle("Confirm Action")
	container.SetSize(modalWidth, modalHeight)
	container.SetPadding(tui.NewMargin(2))
	container.Draw(m.screen, modalX, modalY, &m.theme)

	// Draw question
	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)

	question := "Are you sure you want to proceed?"
	questionX := modalX + (modalWidth-len(question))/2
	m.screen.DrawString(questionX, modalY+3, question, textStyle)

	// Draw buttons
	buttonY := modalY + modalHeight - 4

	yesStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Background).
		Background(m.theme.Palette.Pine).
		Bold(true).
		Padding(0, 2)

	noStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Background).
		Background(m.theme.Palette.Love).
		Bold(true).
		Padding(0, 2)

	yesBtn := " Y - Yes "
	noBtn := " N - No "

	btnSpacing := 4
	totalBtnWidth := len(yesBtn) + len(noBtn) + btnSpacing
	btnStartX := modalX + (modalWidth-totalBtnWidth)/2

	m.screen.DrawString(btnStartX, buttonY, yesBtn, yesStyle)
	m.screen.DrawString(btnStartX+len(yesBtn)+btnSpacing, buttonY, noBtn, noStyle)
}

func (m *model) drawFormModal() {
	m.formModal.Draw(m.screen, 0, 0, &m.theme)

	// Get modal position
	modalWidth, modalHeight := m.formModal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Create container
	container := tui.NewContainer()
	container.SetTitle("User Registration")
	container.SetSize(modalWidth, modalHeight)
	container.SetPadding(tui.NewMargin(2))
	container.Draw(m.screen, modalX, modalY, &m.theme)

	// Draw form fields
	labelStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)

	fieldY := modalY + 4

	// Name field
	m.screen.DrawString(modalX+4, fieldY, "Name:", labelStyle)
	fieldY += 1
	nameField := tui.NewContainer()
	nameField.SetSize(modalWidth-8, 3)
	nameField.SetBorderStyle("single")
	nameInput := tui.NewInput()
	nameInput.SetValue("John Doe")
	nameInput.SetWidth(modalWidth - 12)
	nameField.SetContent(nameInput)
	nameField.Draw(m.screen, modalX+4, fieldY, &m.theme)
	fieldY += 4

	// Email field (active)
	m.screen.DrawString(modalX+4, fieldY, "Email:", labelStyle)
	fieldY += 1
	emailField := tui.NewContainer()
	emailField.SetSize(modalWidth-8, 3)
	emailField.SetBorderStyle("heavy") // Show it's focused
	emailInput := tui.NewInput()
	emailInput.SetValue(m.formEmail)
	emailInput.SetPlaceholder("user@example.com")
	emailInput.SetWidth(modalWidth - 12)
	emailInput.Focus()
	emailField.SetContent(emailInput)
	emailField.Draw(m.screen, modalX+4, fieldY, &m.theme)
	fieldY += 4

	// Instructions
	helpStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Surface)
	m.screen.DrawString(modalX+4, modalY+modalHeight-4, "Type to edit email, ENTER to submit, ESC to cancel", helpStyle)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
