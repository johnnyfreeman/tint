package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type tickMsg time.Time

type model struct {
	screen          *tui.Screen
	sidebar         *Sidebar
	tabs            *tui.TabsComponent
	modal           *tui.Modal
	modalContainer  *tui.Container
	modalContent    *tui.TextArea
	notification    *tui.Notification
	themePicker     *ThemePicker
	statusBar       *tui.StatusBar
	width           int
	height          int
	focus           string // "sidebar", "tabs", "modal", "themepicker"
}

func initialModel() model {
	// Set up status bar
	statusBar := tui.NewStatusBar()
	statusBar.AddSegment("Tab: focus", "left")
	statusBar.AddSegment("Ctrl+B: sidebar | 1-3: tabs | m: modal | n/s/w/e: notif | t: theme | q: quit", "right")

	// Create modal with container and content (Modal → Container → Content pattern)
	modal, modalContainer, modalContent := createDemoModal()

	return model{
		screen:         tui.NewDefaultScreen(80, 24),
		sidebar:        NewSidebar(),
		tabs:           createDemoTabs(),
		modal:          modal,
		modalContainer: modalContainer,
		modalContent:   modalContent,
		notification:   tui.NewNotification(),
		themePicker:    NewThemePicker(),
		statusBar:      statusBar,
		width:          80,
		height:         24,
		focus:          "tabs",
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tickCmd(),
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle theme picker controls when focused
		if m.themePicker.IsVisible() && m.focus == "themepicker" {
			switch msg.String() {
			case "t", "esc":
				m.themePicker.Toggle()
				// Return focus to previous component
				if m.sidebar.IsVisible() {
					m.focus = "sidebar"
				} else {
					m.focus = "tabs"
				}
			case "up":
				m.themePicker.MoveUp()
			case "down":
				m.themePicker.MoveDown()
			case "enter":
				oldTheme := m.themePicker.GetSelectedTheme()
				m.themePicker.Select()
				newTheme := m.themePicker.GetSelectedTheme()
				if oldTheme != newTheme {
					themeName := tui.GetTheme(newTheme).Name
					m.notification.ShowSuccess("Theme changed to " + themeName)
				}
				// Return focus after selecting
				if m.sidebar.IsVisible() {
					m.focus = "sidebar"
				} else {
					m.focus = "tabs"
				}
			}
			return m, nil
		}

		// Handle modal controls when focused
		if m.modal.IsVisible() && m.focus == "modal" {
			switch msg.String() {
			case "m", "esc":
				m.modal.Toggle()
				m.modalContainer.Blur()
				// Return focus to previous component
				if m.sidebar.IsVisible() {
					m.focus = "sidebar"
				} else {
					m.focus = "tabs"
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			// Toggle focus between sidebar and tabs (only if no modals are open)
			if !m.modal.IsVisible() && !m.themePicker.IsVisible() && m.sidebar.IsVisible() {
				if m.focus == "sidebar" {
					m.focus = "tabs"
				} else {
					m.focus = "sidebar"
				}
			}
		case "ctrl+b":
			m.sidebar.Toggle()
			if m.sidebar.IsVisible() {
				m.focus = "sidebar"
			} else {
				m.focus = "tabs"
			}
		case "m":
			m.modal.Toggle()
			if m.modal.IsVisible() {
				m.focus = "modal"
				m.modal.Focus()
				m.modalContainer.Focus()
			} else {
				m.modalContainer.Blur()
			}
		case "n":
			m.notification.ShowInfo("This is a sample notification!")
		case "s":
			m.notification.ShowSuccess("Operation completed successfully!")
		case "w":
			m.notification.ShowWarning("Warning: Check your settings")
		case "e":
			m.notification.ShowError("Error: Something went wrong")
		case "t":
			m.themePicker.Toggle()
			if m.themePicker.IsVisible() {
				m.focus = "themepicker"
			}
		case "1":
			if m.focus == "tabs" {
				m.tabs.SetActive(0)
			}
		case "2":
			if m.focus == "tabs" {
				m.tabs.SetActive(1)
			}
		case "3":
			if m.focus == "tabs" {
				m.tabs.SetActive(2)
			}
		case "up":
			if m.focus == "sidebar" {
				m.sidebar.MoveUp()
			}
		case "down":
			if m.focus == "sidebar" {
				m.sidebar.MoveDown()
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		theme := tui.GetTheme(m.themePicker.GetPreviewTheme())
		m.screen = tui.NewScreen(m.width, m.height, theme)

	case tickMsg:
		m.notification.Update()
		return m, tickCmd()
	}

	return m, nil
}

func (m model) View() string {
	// Get current theme (preview theme if hovering)
	theme := tui.GetTheme(m.themePicker.GetPreviewTheme())

	// Recreate screen if theme changed
	if m.screen.Theme().Name != theme.Name {
		m.screen = tui.NewScreen(m.width, m.height, theme)
	}

	// Clear screen (now uses theme background automatically)
	m.screen.Clear()

	// Calculate layout
	sidebarWidth := m.sidebar.Width()
	tabsX := sidebarWidth
	if sidebarWidth > 0 {
		tabsX++ // Add spacing
	}
	tabsWidth := m.width - tabsX

	// Calculate content height (minus status bar)
	contentHeight := m.height - 1

	// Draw sidebar
	m.sidebar.DrawWithTheme(m.screen, 0, 0, contentHeight, theme, m.focus == "sidebar")

	// Fill gap between sidebar and tabs if sidebar is visible
	if sidebarWidth > 0 && tabsX > sidebarWidth {
		gapStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
		for y := 0; y < contentHeight; y++ {
			m.screen.DrawRune(sidebarWidth, y, ' ', gapStyle)
		}
	}

	// Draw tabs
	if m.focus == "tabs" {
		m.tabs.Focus()
	} else {
		m.tabs.Blur()
	}
	m.tabs.SetSize(tabsWidth, contentHeight)
	m.tabs.Draw(m.screen, tabsX, 0, &theme)

	// Draw modal
	if m.focus == "modal" {
		m.modal.Focus()
	} else {
		m.modal.Blur()
	}

	if m.modal.IsVisible() {
		// Draw modal surface (provides backdrop and elevation)
		m.modal.Draw(m.screen, 0, 0, &theme)

		// Get modal position for container placement
		modalWidth, modalHeight := m.modal.GetSize()
		modalX := (m.width - modalWidth) / 2
		modalY := (m.height - modalHeight) / 2

		// Focus the container when modal is focused
		if m.focus == "modal" {
			m.modalContainer.Focus()
		} else {
			m.modalContainer.Blur()
		}

		// Draw container filling the entire modal surface
		m.modalContainer.Draw(m.screen, modalX, modalY, &theme)
	}

	// Draw notification (always on top)
	m.notification.Draw(m.screen, 0, 0, &theme)

	// Draw theme picker (always on top)
	m.themePicker.DrawWithTheme(m.screen, &theme, m.focus == "themepicker")

	// Draw status bar at the bottom
	m.statusBar.Draw(m.screen, 0, m.height-1, &theme)

	// Render screen to string
	return m.screen.Render()
}

func createDemoTabs() *tui.TabsComponent {
	themeCount := len(tui.GetAvailableThemes())
	tabs := tui.NewTabs()
	tabs.AddTab("Overview", fmt.Sprintf("Welcome to the theme demo!\n\nNotice how each UI element uses\nstate-based colors from the theme.\n\nFocused elements are highlighted.\nHover states provide feedback.\n\nPress 't' to browse %d themes!", themeCount))
	tabs.AddTab("Themes", "Available Themes:\n\n• Tokyo Night (from glamour)\n• Rosé Pine (3 variants)\n• Catppuccin (4 variants)\n• Monochrome (custom)\n\nThemes are imported from\nofficial packages for\nconsistency and accuracy!")
	tabs.AddTab("Logs", "System Activity:\n\n[10:23:45] App started...\n[10:23:46] Theme loaded\n[10:23:47] Colors applied\n[10:23:48] Ready!\n\nEach theme brings its own\nunique personality and\ncolor combinations.")
	return tabs
}

func createDemoModal() (*tui.Modal, *tui.Container, *tui.TextArea) {
	// Create modal
	modal := tui.NewModal()
	modal.SetSize(40, 12)
	modal.SetCentered(true)

	// Create container that fills the modal
	container := tui.NewContainer()
	container.SetTitle("Sample Modal")
	container.SetSize(40, 12) // Fill the entire modal surface
	container.SetPadding(tui.NewMargin(1))
	container.SetUseSurface(true) // Use surface color for modal

	// Create content for the container
	textarea := tui.NewTextArea()
	textarea.SetValue("This is a modal dialog.\n\nPress 'm' again to close it.\n\nNotice the shadow effect\nbehind the modal.")
	textarea.SetSize(36, 8) // Account for border and padding
	container.SetContent(textarea)

	return modal, container, textarea
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
