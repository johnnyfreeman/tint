package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

const (
	defaultWidth  = 100
	defaultHeight = 30
)

type model struct {
	screen *tui.Screen
	width  int
	height int
	theme  tui.Theme

	// Different modal types
	activeModal string

	// Modal components (following Modal → Container → Content pattern)
	simpleModal  *simpleModalComponent
	fuzzyModal   *fuzzyModalComponent
	confirmModal *confirmModalComponent
	formModal    *formModalComponent
}

// Simple modal with single container
type simpleModalComponent struct {
	modal     *tui.Modal
	container *tui.Container
	viewer    *tui.Viewer
}

// Fuzzy finder modal with 3 containers
type fuzzyModalComponent struct {
	modal            *tui.Modal
	searchContainer  *tui.Container
	resultsContainer *tui.Container
	previewContainer *tui.Container
	input            *tui.Input
	items            []string
	filtered         []string
	selectedIdx      int
}

// Confirm dialog modal
type confirmModalComponent struct {
	modal       *tui.Modal
	container   *tui.Container
	message     string
	confirmText string
	cancelText  string
	selected    int // 0 = cancel, 1 = confirm
}

// Form modal with multiple inputs
type formModalComponent struct {
	modal     *tui.Modal
	container *tui.Container
	inputs    []*tui.Input
	labels    []string
	activeIdx int
}

func initialModel() *model {
	// Use monochrome theme for consistent styling
	theme := tui.GetTheme("monochrome")

	// Create simple modal (Modal → Container → Content)
	simpleModal := tui.NewModal()
	simpleModal.SetSize(50, 15)
	simpleModal.SetCentered(true)

	simpleContainer := tui.NewContainer()
	simpleContainer.SetTitle("About")
	simpleContainer.SetSize(50, 15) // Fill the entire modal surface
	simpleContainer.SetPadding(tui.NewMargin(2))

	simpleViewer := tui.NewViewer()
	simpleViewer.SetContent(`Modal Examples Demo

This example demonstrates various modal patterns:

1. Simple Modal - Single container with content
2. Fuzzy Finder - Three containers layout
3. Confirm Dialog - Action confirmation
4. Form Modal - Multiple input fields

Press the number keys (1-4) to show each modal.
Press Escape to close any modal.`)
	simpleViewer.SetWrapText(true)
	simpleContainer.SetContent(simpleViewer)

	simpleComp := &simpleModalComponent{
		modal:     simpleModal,
		container: simpleContainer,
		viewer:    simpleViewer,
	}

	// Create fuzzy finder modal (Modal → Multiple Containers → Content)
	fuzzyModal := tui.NewModal()
	fuzzyModal.SetSize(80, 25)
	fuzzyModal.SetCentered(true)

	searchContainer := tui.NewContainer()
	searchContainer.SetTitle("Search")
	searchContainer.SetSize(34, 3)
	searchContainer.SetPadding(tui.NewMargin(1))

	searchInput := tui.NewInput()
	searchInput.SetPlaceholder("Type to filter...")
	searchInput.SetWidth(30) // Account for container padding
	searchContainer.SetContent(searchInput)

	resultsContainer := tui.NewContainer()
	resultsContainer.SetTitle("Results")
	resultsContainer.SetSize(34, 22)
	resultsContainer.SetPadding(tui.NewMargin(1))

	previewContainer := tui.NewContainer()
	previewContainer.SetTitle("Preview")
	previewContainer.SetSize(45, 25)
	previewContainer.SetPadding(tui.NewMargin(1))

	fuzzyComp := &fuzzyModalComponent{
		modal:            fuzzyModal,
		searchContainer:  searchContainer,
		resultsContainer: resultsContainer,
		previewContainer: previewContainer,
		input:            searchInput,
		items: []string{
			"main.go", "config.go", "utils.go", "server.go", "client.go",
			"auth.go", "database.go", "middleware.go", "routes.go", "models.go",
			"README.md", "LICENSE", "Makefile", "docker-compose.yml",
		},
		filtered: []string{},
	}

	// Create confirm modal (Modal → Container → Content)
	confirmModal := tui.NewModal()
	confirmModal.SetSize(40, 10)
	confirmModal.SetCentered(true)

	confirmContainer := tui.NewContainer()
	confirmContainer.SetTitle("Confirm Action")
	confirmContainer.SetSize(40, 10) // Fill the entire modal surface
	confirmContainer.SetPadding(tui.NewMargin(2))

	confirmComp := &confirmModalComponent{
		modal:       confirmModal,
		container:   confirmContainer,
		message:     "Are you sure you want to proceed?",
		confirmText: "Yes",
		cancelText:  "No",
		selected:    0,
	}

	// Create form modal (Modal → Container → Content)
	formModal := tui.NewModal()
	formModal.SetSize(50, 20)
	formModal.SetCentered(true)

	formContainer := tui.NewContainer()
	formContainer.SetTitle("User Form")
	formContainer.SetSize(50, 20) // Fill the entire modal surface
	formContainer.SetPadding(tui.NewMargin(2))

	nameInput := tui.NewInput()
	nameInput.SetPlaceholder("Enter your name...")
	nameInput.SetWidth(42)

	emailInput := tui.NewInput()
	emailInput.SetPlaceholder("Enter your email...")
	emailInput.SetWidth(42)

	messageInput := tui.NewInput()
	messageInput.SetPlaceholder("Enter a message...")
	messageInput.SetWidth(42)

	formComp := &formModalComponent{
		modal:     formModal,
		container: formContainer,
		inputs:    []*tui.Input{nameInput, emailInput, messageInput},
		labels:    []string{"Name:", "Email:", "Message:"},
		activeIdx: 0,
	}

	return &model{
		screen:       tui.NewScreen(defaultWidth, defaultHeight, theme),
		width:        defaultWidth,
		height:       defaultHeight,
		theme:        theme,
		activeModal:  "",
		simpleModal:  simpleComp,
		fuzzyModal:   fuzzyComp,
		confirmModal: confirmComp,
		formModal:    formComp,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Tint Modal Examples")
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.screen = tui.NewScreen(m.width, m.height, m.theme)

	case tea.KeyMsg:
		// Handle escape to close modals
		if msg.String() == "escape" || msg.String() == "esc" {
			m.activeModal = ""
			m.simpleModal.modal.Hide()
			m.simpleModal.container.Blur()
			m.fuzzyModal.modal.Hide()
			m.fuzzyModal.searchContainer.Blur()
			m.fuzzyModal.resultsContainer.Blur()
			m.fuzzyModal.previewContainer.Blur()
			m.confirmModal.modal.Hide()
			m.confirmModal.container.Blur()
			m.formModal.modal.Hide()
			m.formModal.container.Blur()
			return m, nil
		}

		// Route input to active modal
		switch m.activeModal {
		case "fuzzy":
			m.handleFuzzyInput(msg)
			return m, nil
		case "confirm":
			m.handleConfirmInput(msg)
			return m, nil
		case "form":
			m.handleFormInput(msg)
			return m, nil
		}

		// Global shortcuts to show modals
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "1":
			m.showModal("simple")
		case "2":
			m.showModal("fuzzy")
		case "3":
			m.showModal("confirm")
		case "4":
			m.showModal("form")
		}
	}

	return m, nil
}

func (m *model) showModal(modalType string) {
	// Hide all modals first
	m.simpleModal.modal.Hide()
	m.simpleModal.container.Blur()
	m.fuzzyModal.modal.Hide()
	m.fuzzyModal.searchContainer.Blur()
	m.fuzzyModal.resultsContainer.Blur()
	m.fuzzyModal.previewContainer.Blur()
	m.confirmModal.modal.Hide()
	m.confirmModal.container.Blur()
	m.formModal.modal.Hide()
	m.formModal.container.Blur()

	// Show selected modal
	m.activeModal = modalType
	switch modalType {
	case "simple":
		m.simpleModal.modal.Show()
		m.simpleModal.modal.Focus()
		m.simpleModal.container.Focus()
	case "fuzzy":
		m.fuzzyModal.modal.Show()
		m.fuzzyModal.modal.Focus()
		m.fuzzyModal.searchContainer.Focus()
		m.fuzzyModal.resultsContainer.Focus()
		m.fuzzyModal.previewContainer.Focus()
		m.fuzzyModal.input.SetValue("")
		m.fuzzyModal.input.Focus()
		m.updateFuzzyResults("")
	case "confirm":
		m.confirmModal.modal.Show()
		m.confirmModal.modal.Focus()
		m.confirmModal.container.Focus()
		m.confirmModal.selected = 0
	case "form":
		m.formModal.modal.Show()
		m.formModal.modal.Focus()
		m.formModal.container.Focus()
		m.formModal.activeIdx = 0
		m.formModal.inputs[0].Focus()
	}
}

func (m *model) handleFuzzyInput(msg tea.KeyMsg) {
	switch msg.String() {
	case "up", "k":
		if m.fuzzyModal.selectedIdx > 0 {
			m.fuzzyModal.selectedIdx--
		}
	case "down", "j":
		if m.fuzzyModal.selectedIdx < len(m.fuzzyModal.filtered)-1 {
			m.fuzzyModal.selectedIdx++
		}
	case "enter":
		// In a real app, this would select the item
		m.activeModal = ""
		m.fuzzyModal.modal.Hide()
		m.fuzzyModal.searchContainer.Blur()
		m.fuzzyModal.resultsContainer.Blur()
		m.fuzzyModal.previewContainer.Blur()
	default:
		m.fuzzyModal.input.HandleInput(msg.String())
		m.updateFuzzyResults(m.fuzzyModal.input.Value())
	}
}

func (m *model) updateFuzzyResults(query string) {
	m.fuzzyModal.filtered = []string{}

	if query == "" {
		m.fuzzyModal.filtered = m.fuzzyModal.items
	} else {
		query = strings.ToLower(query)
		for _, item := range m.fuzzyModal.items {
			if strings.Contains(strings.ToLower(item), query) {
				m.fuzzyModal.filtered = append(m.fuzzyModal.filtered, item)
			}
		}
	}

	m.fuzzyModal.selectedIdx = 0
}

func (m *model) handleConfirmInput(msg tea.KeyMsg) {
	switch msg.String() {
	case "left", "h":
		m.confirmModal.selected = 0
	case "right", "l":
		m.confirmModal.selected = 1
	case "tab":
		m.confirmModal.selected = (m.confirmModal.selected + 1) % 2
	case "enter":
		// In a real app, this would trigger the action
		m.activeModal = ""
		m.confirmModal.modal.Hide()
		m.confirmModal.container.Blur()
	}
}

func (m *model) handleFormInput(msg tea.KeyMsg) {
	switch msg.String() {
	case "tab", "down":
		m.formModal.inputs[m.formModal.activeIdx].Blur()
		m.formModal.activeIdx = (m.formModal.activeIdx + 1) % len(m.formModal.inputs)
		m.formModal.inputs[m.formModal.activeIdx].Focus()
	case "shift+tab", "up":
		m.formModal.inputs[m.formModal.activeIdx].Blur()
		m.formModal.activeIdx--
		if m.formModal.activeIdx < 0 {
			m.formModal.activeIdx = len(m.formModal.inputs) - 1
		}
		m.formModal.inputs[m.formModal.activeIdx].Focus()
	case "enter":
		if m.formModal.activeIdx == len(m.formModal.inputs)-1 {
			// Submit form
			m.activeModal = ""
			m.formModal.modal.Hide()
			m.formModal.container.Blur()
		} else {
			// Move to next field
			m.formModal.inputs[m.formModal.activeIdx].Blur()
			m.formModal.activeIdx++
			m.formModal.inputs[m.formModal.activeIdx].Focus()
		}
	default:
		m.formModal.inputs[m.formModal.activeIdx].HandleInput(msg.String())
	}
}

func (m *model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	// Clear screen with theme background
	m.screen.Clear()

	// Draw main content
	m.drawMainContent()

	// Draw active modal (Modal → Container → Content pattern)
	switch m.activeModal {
	case "simple":
		m.drawSimpleModal()
	case "fuzzy":
		m.drawFuzzyModal()
	case "confirm":
		m.drawConfirmModal()
	case "form":
		m.drawFormModal()
	}

	return m.screen.Render()
}

func (m *model) drawMainContent() {
	// Title
	titleStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Primary).
		Bold(true)
	m.screen.DrawString(2, 1, "Modal Examples", titleStyle)

	// Instructions
	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text)
	instructions := []string{
		"",
		"Press number keys to show different modal types:",
		"",
		"1 - Simple Modal (single container)",
		"2 - Fuzzy Finder (three containers)",
		"3 - Confirm Dialog (action buttons)",
		"4 - Form Modal (multiple inputs)",
		"",
		"Press 'q' to quit",
		"Press 'Escape' to close any modal",
	}

	for i, line := range instructions {
		m.screen.DrawString(4, 3+i, line, textStyle)
	}

	// Status
	statusStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted)
	status := "No modal active"
	if m.activeModal != "" {
		status = fmt.Sprintf("Active modal: %s", m.activeModal)
	}
	m.screen.DrawString(4, m.height-2, status, statusStyle)
}

func (m *model) drawSimpleModal() {
	if !m.simpleModal.modal.IsVisible() {
		return
	}

	// Draw modal surface (provides backdrop and elevation)
	m.simpleModal.modal.Draw(m.screen, 0, 0, m.width, m.height, &m.theme)

	// Get modal position for container placement
	modalWidth, modalHeight := m.simpleModal.modal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Draw container filling the entire modal surface
	m.simpleModal.container.Draw(m.screen, modalX, modalY, modalWidth, modalHeight, &m.theme)
}

func (m *model) drawFuzzyModal() {
	if !m.fuzzyModal.modal.IsVisible() {
		return
	}

	// Draw modal surface (provides backdrop and elevation)
	m.fuzzyModal.modal.Draw(m.screen, 0, 0, m.width, m.height, &m.theme)

	// Get modal position for container placement
	modalWidth, modalHeight := m.fuzzyModal.modal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Draw containers filling modal surface area with 1-column gap
	// Left column: search and results
	m.fuzzyModal.searchContainer.Draw(m.screen, modalX, modalY, 34, 3, &m.theme)
	m.fuzzyModal.resultsContainer.Draw(m.screen, modalX, modalY+3, 34, 22, &m.theme)

	// Right column: preview (with 1-column gap)
	m.fuzzyModal.previewContainer.Draw(m.screen, modalX+35, modalY, 45, 25, &m.theme)

	// Draw results content manually inside results container
	m.drawFuzzyResults(modalX+2, modalY+5, 30, 18)

	// Draw preview content
	m.drawFuzzyPreview(modalX+37, modalY+2, 41, 21)
}

func (m *model) drawFuzzyResults(x, y, width, height int) {
	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)
	selectedStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Background).
		Background(m.theme.Palette.Primary)

	for i, item := range m.fuzzyModal.filtered {
		if i >= height {
			break
		}

		style := textStyle
		if i == m.fuzzyModal.selectedIdx {
			// Highlight full row
			for dx := 0; dx < width; dx++ {
				m.screen.SetCell(x+dx, y+i, tui.Cell{
					Rune:       ' ',
					Background: m.theme.Palette.Primary,
				})
			}
			style = selectedStyle
		}

		displayName := item
		if len(displayName) > width {
			displayName = displayName[:width-3] + "..."
		}

		m.screen.DrawString(x, y+i, displayName, style)
	}
}

func (m *model) drawFuzzyPreview(x, y, width, height int) {
	previewStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Surface)

	if m.fuzzyModal.selectedIdx < len(m.fuzzyModal.filtered) {
		selected := m.fuzzyModal.filtered[m.fuzzyModal.selectedIdx]
		preview := fmt.Sprintf("Preview of: %s\n\nThis would show the file contents\nor other relevant information.", selected)

		lines := strings.Split(preview, "\n")
		for i, line := range lines {
			if i >= height {
				break
			}
			if len(line) > width {
				line = line[:width]
			}
			m.screen.DrawString(x, y+i, line, previewStyle)
		}
	}
}

func (m *model) drawConfirmModal() {
	if !m.confirmModal.modal.IsVisible() {
		return
	}

	// Draw modal surface (provides backdrop and elevation)
	m.confirmModal.modal.Draw(m.screen, 0, 0, m.width, m.height, &m.theme)

	// Get modal position for container placement
	modalWidth, modalHeight := m.confirmModal.modal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Draw container filling the entire modal surface
	m.confirmModal.container.Draw(m.screen, modalX, modalY, modalWidth, modalHeight, &m.theme)

	// Draw message inside container
	messageStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)
	messageX := modalX + 2
	messageY := modalY + 2
	m.screen.DrawString(messageX, messageY, m.confirmModal.message, messageStyle)

	// Draw buttons
	buttonY := modalY + modalHeight - 3
	buttonSpacing := 8

	// Cancel button
	cancelStyle := messageStyle
	if m.confirmModal.selected == 0 {
		cancelStyle = lipgloss.NewStyle().
			Foreground(m.theme.Palette.Background).
			Background(m.theme.Palette.Primary).
			Padding(0, 1)
	}
	cancelX := modalX + modalWidth/2 - buttonSpacing - len(m.confirmModal.cancelText)/2
	m.screen.DrawString(cancelX, buttonY, m.confirmModal.cancelText, cancelStyle)

	// Confirm button
	confirmStyle := messageStyle
	if m.confirmModal.selected == 1 {
		confirmStyle = lipgloss.NewStyle().
			Foreground(m.theme.Palette.Background).
			Background(m.theme.Palette.Pine).
			Padding(0, 1)
	}
	confirmX := modalX + modalWidth/2 + buttonSpacing - len(m.confirmModal.confirmText)/2
	m.screen.DrawString(confirmX, buttonY, m.confirmModal.confirmText, confirmStyle)
}

func (m *model) drawFormModal() {
	if !m.formModal.modal.IsVisible() {
		return
	}

	// Draw modal surface (provides backdrop and elevation)
	m.formModal.modal.Draw(m.screen, 0, 0, m.width, m.height, &m.theme)

	// Get modal position for container placement
	modalWidth, modalHeight := m.formModal.modal.GetSize()
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Draw container filling the entire modal surface
	m.formModal.container.Draw(m.screen, modalX, modalY, modalWidth, modalHeight, &m.theme)

	// Draw form fields inside container
	labelStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)

	fieldY := modalY + 3
	fieldSpacing := 4

	for i, label := range m.formModal.labels {
		// Draw label
		m.screen.DrawString(modalX+2, fieldY, label, labelStyle)

		// Draw input
		m.formModal.inputs[i].Draw(m.screen, modalX+2, fieldY+1, 42, 1, &m.theme)

		fieldY += fieldSpacing
	}

	// Draw help text
	helpStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Surface)
	helpText := "Tab: Next field | Enter: Submit | Escape: Cancel"
	helpX := modalX + (modalWidth-len(helpText))/2
	m.screen.DrawString(helpX, modalY+modalHeight-3, helpText, helpStyle)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}