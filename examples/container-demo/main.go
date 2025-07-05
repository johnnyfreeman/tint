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

	containers []*tui.Container
	activeTab  int
}

func initialModel() *model {
	theme := tui.GetTheme("tokyonight")

	// Create various container examples
	containers := []*tui.Container{}

	// Example 1: Simple titled container
	input1 := tui.NewInput()
	input1.SetPlaceholder("Enter text...")
	container1 := tui.TitledPanel("Simple Title", input1)
	containers = append(containers, container1)

	// Example 2: Container with tabs
	viewer := tui.NewViewer()
	viewer.SetContent("This container has tabs embedded in the border.\n\nThe active tab is highlighted.")
	container2 := tui.TabbedPanel([]string{"Code", "Preview", "Console"}, 1, viewer)
	containers = append(containers, container2)

	// Example 3: Container with status
	table := tui.NewTable()
	table.SetColumns([]tui.TableColumn{
		{Title: "Name", Width: 20},
		{Title: "Status", Width: 15},
	})
	table.SetRows([]tui.TableRow{
		{"Service A", "Running"},
		{"Service B", "Stopped"},
	})
	container3 := tui.StatusPanel("Services", "Connected", table)
	containers = append(containers, container3)

	// Example 4: Container with icon
	viewer2 := tui.NewViewer()
	viewer2.SetContent("This container has an icon in the border.\n\nIcons help identify the purpose of a container.")
	container4 := tui.IconPanel('🚀', "Launch Control", viewer2)
	containers = append(containers, container4)

	// Example 5: Complex multi-element container
	textArea := tui.NewTextArea()
	textArea.SetValue("This container demonstrates multiple border elements:\n- Title on the left\n- Tabs in the center\n- Status on the right\n- Version badge on bottom")
	container5 := tui.MultiElementPanel(
		"Editor",
		[]string{"main.go", "test.go", "README.md"},
		0,
		"Ready",
		textArea,
	)
	containers = append(containers, container5)

	// Example 6: Custom container with mixed elements
	container6 := tui.NewContainer()
	container6.SetTitle("Mixed Elements")
	container6.AddBorderElement(tui.NewIconElement('📁'), tui.BorderTop, tui.BorderAlignRight)
	container6.AddBorderElement(tui.NewBadgeElement("New"), tui.BorderTop, tui.BorderAlignRight)
	container6.AddBorderElement(tui.NewStatusElement("OK"), tui.BorderBottom, tui.BorderAlignLeft)
	container6.AddBorderElement(tui.NewTextElement("Footer Text"), tui.BorderBottom, tui.BorderAlignCenter)

	input2 := tui.NewInput()
	input2.SetPlaceholder("Search...")
	container6.SetContent(input2)
	containers = append(containers, container6)

	// Set sizes for all containers
	for _, c := range containers {
		c.SetSize(40, 12)
		c.SetPadding(tui.NewMargin(1))
	}

	return &model{
		screen:     tui.NewScreen(80, 24, theme),
		width:      80,
		height:     24,
		theme:      theme,
		containers: containers,
		activeTab:  0,
	}
}

func (m *model) Init() tea.Cmd {
	return tea.SetWindowTitle("Tint Container Demo")
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
		case "tab":
			m.activeTab = (m.activeTab + 1) % len(m.containers)
			// Update focus
			for i, c := range m.containers {
				if i == m.activeTab {
					c.Focus()
				} else {
					c.Blur()
				}
			}
		case "1", "2", "3", "4", "5", "6":
			idx := int(msg.String()[0] - '1')
			if idx < len(m.containers) {
				m.activeTab = idx
				for i, c := range m.containers {
					if i == m.activeTab {
						c.Focus()
					} else {
						c.Blur()
					}
				}
			}
		}
	}

	return m, nil
}

func (m *model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	m.screen.Clear()

	// Draw title
	titleStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Primary).
		Bold(true)
	title := "Container Border Elements Demo"
	m.screen.DrawString((m.width-len(title))/2, 0, title, titleStyle)

	// Draw subtitle
	subtitleStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted)
	subtitle := "Press Tab or 1-6 to switch containers, q to quit"
	m.screen.DrawString((m.width-len(subtitle))/2, 1, subtitle, subtitleStyle)

	// Calculate grid layout (2x3)
	cols := 2
	containerWidth := 38
	containerHeight := 10
	startY := 3

	// Draw containers in a grid
	for i, container := range m.containers {
		row := i / cols
		col := i % cols

		x := col*(containerWidth+2) + 1
		y := startY + row*(containerHeight+1)

		// Set border style based on focus
		if i == m.activeTab {
			container.SetBorderStyle("heavy")
		} else {
			container.SetBorderStyle("single")
		}

		container.DrawWithBounds(m.screen, x, y, containerWidth, containerHeight, &m.theme)
	}

	// Draw example labels
	labelStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Pine).
		Bold(true)

	labels := []string{
		"1. Simple Title",
		"2. Tabs",
		"3. Status",
		"4. Icon + Title",
		"5. Multi-Element",
		"6. Custom Mix",
	}

	for i, label := range labels {
		row := i / cols
		col := i % cols
		x := col*(containerWidth+2) + 1
		y := startY + row*(containerHeight+1) - 1

		if i == m.activeTab {
			label = "▶ " + label
		} else {
			label = "  " + label
		}
		m.screen.DrawString(x, y, label, labelStyle)
	}

	// Draw help text at bottom
	helpStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted)
	helpText := "Border elements can be titles, tabs, status indicators, icons, badges, and more!"
	m.screen.DrawString((m.width-len(helpText))/2, m.height-2, helpText, helpStyle)

	return m.screen.Render()
}

func main() {
	// Set focus on first container
	m := initialModel()
	if len(m.containers) > 0 {
		m.containers[0].Focus()
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
