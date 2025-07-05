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
	activeModal  string
	simpleModal  *simpleModalComponent
	fuzzyModal   *fuzzyModalComponent
	confirmModal *confirmModalComponent
	formModal    *formModalComponent
}

// Simple modal with single container
type simpleModalComponent struct {
	visible   bool
	container *tui.Container
	viewer    *tui.Viewer
}

// Fuzzy finder modal with 3 containers
type fuzzyModalComponent struct {
	visible          bool
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
	visible     bool
	container   *tui.Container
	message     string
	confirmText string
	cancelText  string
	selected    int // 0 = cancel, 1 = confirm
}

// Form modal with multiple inputs
type formModalComponent struct {
	visible   bool
	container *tui.Container
	inputs    []*tui.Input
	labels    []string
	activeIdx int
}

func initialModel() *model {
	// Use monochrome theme for consistent styling
	theme := tui.GetTheme("monochrome")

	// Create simple modal
	simpleContainer := tui.NewContainer()
	simpleContainer.SetTitle("About")
	simpleContainer.SetSize(50, 15)
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

	simpleModal := &simpleModalComponent{
		container: simpleContainer,
		viewer:    simpleViewer,
	}

	// Create fuzzy finder modal
	searchContainer := tui.NewContainer()
	searchContainer.SetTitle("Search")
	searchContainer.SetPadding(tui.NewMargin(1))

	searchInput := tui.NewInput()
	searchInput.SetPlaceholder("Type to filter...")
	searchContainer.SetContent(searchInput)

	resultsContainer := tui.NewContainer()
	resultsContainer.SetTitle("Results")
	resultsContainer.SetPadding(tui.NewMargin(1))

	previewContainer := tui.NewContainer()
	previewContainer.SetTitle("Preview")
	previewContainer.SetPadding(tui.NewMargin(1))

	fuzzyModal := &fuzzyModalComponent{
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

	// Create confirm modal
	confirmContainer := tui.NewContainer()
	confirmContainer.SetTitle("Confirm Action")
	confirmContainer.SetSize(40, 10)
	confirmContainer.SetPadding(tui.NewMargin(2))

	confirmModal := &confirmModalComponent{
		container:   confirmContainer,
		message:     "Are you sure you want to proceed?",
		confirmText: "Yes",
		cancelText:  "No",
		selected:    0,
	}

	// Create form modal
	formContainer := tui.NewContainer()
	formContainer.SetTitle("User Form")
	formContainer.SetSize(50, 20)
	formContainer.SetPadding(tui.NewMargin(2))

	nameInput := tui.NewInput()
	nameInput.SetPlaceholder("Enter your name...")

	emailInput := tui.NewInput()
	emailInput.SetPlaceholder("Enter your email...")

	messageInput := tui.NewInput()
	messageInput.SetPlaceholder("Enter a message...")

	formModal := &formModalComponent{
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
		simpleModal:  simpleModal,
		fuzzyModal:   fuzzyModal,
		confirmModal: confirmModal,
		formModal:    formModal,
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
			m.simpleModal.visible = false
			m.fuzzyModal.visible = false
			m.confirmModal.visible = false
			m.formModal.visible = false
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
	m.simpleModal.visible = false
	m.fuzzyModal.visible = false
	m.confirmModal.visible = false
	m.formModal.visible = false

	// Show selected modal
	m.activeModal = modalType
	switch modalType {
	case "simple":
		m.simpleModal.visible = true
	case "fuzzy":
		m.fuzzyModal.visible = true
		m.fuzzyModal.input.SetValue("")
		m.fuzzyModal.input.Focus()
		m.updateFuzzyResults("")
	case "confirm":
		m.confirmModal.visible = true
		m.confirmModal.selected = 0
	case "form":
		m.formModal.visible = true
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
		m.fuzzyModal.visible = false
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
		m.confirmModal.visible = false
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
			m.formModal.visible = false
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

	// Draw active modal
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

func (m *model) drawModalBackdrop(x, y, width, height int) {
	// Clear the entire modal area first with surface color (for modal background)
	surfaceStyle := lipgloss.NewStyle().Background(m.theme.Palette.Surface)
	tui.ClearArea(m.screen, x, y, width, height, surfaceStyle)

	// Draw shadow with 1 pixel offset to bottom-right
	// This creates a subtle depth effect
	for dy := 0; dy < height; dy++ {
		m.screen.SetCell(x+width, y+dy+1, tui.Cell{
			Rune:       ' ',
			Background: m.theme.Palette.Shadow,
		})
	}
	for dx := 0; dx < width+1; dx++ {
		m.screen.SetCell(x+dx, y+height, tui.Cell{
			Rune:       ' ',
			Background: m.theme.Palette.Shadow,
		})
	}
}

func (m *model) drawSimpleModal() {
	if !m.simpleModal.visible {
		return
	}

	width, height := m.simpleModal.container.GetSize()
	x := (m.width - width) / 2
	y := (m.height - height) / 2

	// Draw backdrop
	m.drawModalBackdrop(x, y, width, height)

	// Draw container
	m.simpleModal.container.Draw(m.screen, x, y, &m.theme)
}

func (m *model) drawFuzzyModal() {
	if !m.fuzzyModal.visible {
		return
	}

	// Modal dimensions
	modalWidth := 80
	modalHeight := 25
	modalX := (m.width - modalWidth) / 2
	modalY := (m.height - modalHeight) / 2

	// Draw backdrop
	m.drawModalBackdrop(modalX, modalY, modalWidth, modalHeight)

	// Layout calculations
	searchHeight := 3
	leftColumnWidth := 30
	rightColumnWidth := modalWidth - leftColumnWidth - 3 // Account for margins
	resultsHeight := modalHeight - searchHeight - 2      // No extra spacing

	// Update container sizes
	m.fuzzyModal.searchContainer.SetSize(leftColumnWidth, searchHeight)
	m.fuzzyModal.resultsContainer.SetSize(leftColumnWidth, resultsHeight)
	m.fuzzyModal.previewContainer.SetSize(rightColumnWidth, modalHeight-2)

	// Update input width
	m.fuzzyModal.input.SetWidth(leftColumnWidth - 4) // Account for container padding

	// Draw containers
	// Left column: search and results
	m.fuzzyModal.searchContainer.Draw(m.screen, modalX+1, modalY+1, &m.theme)
	m.fuzzyModal.resultsContainer.Draw(m.screen, modalX+1, modalY+searchHeight+1, &m.theme) // No gap

	// Right column: preview (full height)
	m.fuzzyModal.previewContainer.Draw(m.screen, modalX+leftColumnWidth+2, modalY+1, &m.theme)

	// Draw results content manually
	m.drawFuzzyResults(modalX+3, modalY+searchHeight+3, leftColumnWidth-4, resultsHeight-4)

	// Draw preview content
	m.drawFuzzyPreview(modalX+leftColumnWidth+4, modalY+3, rightColumnWidth-4, modalHeight-6)
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
	if !m.confirmModal.visible {
		return
	}

	width, height := m.confirmModal.container.GetSize()
	x := (m.width - width) / 2
	y := (m.height - height) / 2

	// Draw backdrop
	m.drawModalBackdrop(x, y, width, height)

	// Draw container
	m.confirmModal.container.Draw(m.screen, x, y, &m.theme)

	// Draw message
	messageStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)
	messageX := x + 2
	messageY := y + 3
	m.screen.DrawString(messageX, messageY, m.confirmModal.message, messageStyle)

	// Draw buttons
	buttonY := y + height - 4
	buttonSpacing := 10

	// Cancel button
	cancelStyle := messageStyle
	if m.confirmModal.selected == 0 {
		cancelStyle = lipgloss.NewStyle().
			Foreground(m.theme.Palette.Background).
			Background(m.theme.Palette.Primary).
			Padding(0, 2)
	}
	cancelX := x + width/2 - buttonSpacing - len(m.confirmModal.cancelText)/2
	m.screen.DrawString(cancelX, buttonY, m.confirmModal.cancelText, cancelStyle)

	// Confirm button
	confirmStyle := messageStyle
	if m.confirmModal.selected == 1 {
		confirmStyle = lipgloss.NewStyle().
			Foreground(m.theme.Palette.Background).
			Background(m.theme.Palette.Pine).
			Padding(0, 2)
	}
	confirmX := x + width/2 + buttonSpacing - len(m.confirmModal.confirmText)/2
	m.screen.DrawString(confirmX, buttonY, m.confirmModal.confirmText, confirmStyle)
}

func (m *model) drawFormModal() {
	if !m.formModal.visible {
		return
	}

	width, height := m.formModal.container.GetSize()
	x := (m.width - width) / 2
	y := (m.height - height) / 2

	// Draw backdrop
	m.drawModalBackdrop(x, y, width, height)

	// Draw container
	m.formModal.container.Draw(m.screen, x, y, &m.theme)

	// Draw form fields
	labelStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)

	fieldY := y + 3
	fieldSpacing := 4

	for i, label := range m.formModal.labels {
		// Draw label
		m.screen.DrawString(x+3, fieldY, label, labelStyle)

		// Update input width
		m.formModal.inputs[i].SetWidth(width - 8)

		// Draw input
		m.formModal.inputs[i].Draw(m.screen, x+3, fieldY+1, &m.theme)

		fieldY += fieldSpacing
	}

	// Draw help text
	helpStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Surface)
	helpText := "Tab: Next field | Enter: Submit | Escape: Cancel"
	helpX := x + (width-len(helpText))/2
	m.screen.DrawString(helpX, y+height-3, helpText, helpStyle)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
