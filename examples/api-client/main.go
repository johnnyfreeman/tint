package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

type model struct {
	width  int
	height int
	screen *tui.Screen

	// UI state
	focus string // "url", "headers", "body", "history", "response"

	// API Client state
	method       string
	urlInput     *tui.Input
	headersTable *tui.Table
	bodyTextArea *tui.TextArea
	
	// Response state
	responseStatus  string
	responseHeaders []Header
	responseBody    string
	responseTab     int // 0: Body, 1: Headers, 2: Status
	responseViewer  *tui.Viewer

	// History
	history        []HistoryItem
	historyIndex   int
	historyVisible bool

	// Current theme
	currentTheme string
	
	// New components
	statusBar *tui.StatusBar
	
	// Containers for different sections
	historyContainer  *tui.Container
	urlContainer      *tui.Container
	headersContainer  *tui.Container
	bodyContainer     *tui.Container
	responseContainer *tui.Container
	
	// Tab elements for dynamic updates
	responseTabs      *tui.TabsElement
	headerCountBadge  *tui.BadgeElement
}

type Header struct {
	Key   string
	Value string
}

type HistoryItem struct {
	Method string
	URL    string
	Time   string
}

func initialModel() model {
	urlInput := tui.NewInput()
	urlInput.SetValueWithCursorAtStart("https://api.github.com/users/github")
	urlInput.SetPlaceholder("Enter URL...")
	urlInput.SetWidth(50)
	urlInput.Focus()
	
	headersTable := tui.NewTable()
	headersTable.SetColumns([]tui.TableColumn{
		{Title: "Header", Width: 20},
		{Title: "Value", Width: 40},
	})
	headersTable.SetRows([]tui.TableRow{
		{"Accept", "application/json"},
		{"User-Agent", "API-Client/1.0"},
	})
	
	bodyTextArea := tui.NewTextArea()
	bodyTextArea.SetPlaceholder("Request body (JSON, XML, etc.)")
	bodyTextArea.SetSize(60, 4)
	
	responseViewer := tui.NewViewer()
	responseViewer.SetWrapText(true)
	
	// Create status bar
	statusBar := tui.NewStatusBar()
	statusBar.AddSegment("Tab: focus", "left")
	statusBar.AddSegment("H: history | R: send | M: method | 1-3: tabs | n/d: table | arrows: nav | q: quit", "right")
	
	// Layout is handled manually in View() for more control
	
	// Create containers with enhanced border elements
	historyContainer := tui.NewContainer()
	historyContainer.SetTitle("History")
	historyContainer.SetPadding(tui.NewMargin(1))
	// Add a count badge for history items
	historyContainer.AddBorderElement(tui.NewBadgeElement("3"), tui.BorderTop, tui.BorderAlignRight)
	
	urlContainer := tui.NewContainer()
	urlContainer.SetTitle("Request URL")
	urlContainer.SetPadding(tui.NewMargin(1))
	urlContainer.SetContent(urlInput)
	
	headersContainer := tui.NewContainer()
	headersContainer.SetTitle("Headers")
	headersContainer.SetPadding(tui.NewMargin(1))
	headersContainer.SetContent(headersTable)
	// Add a count badge showing number of headers
	headerCountBadge := tui.NewBadgeElement("2")
	headersContainer.AddBorderElement(headerCountBadge, tui.BorderTop, tui.BorderAlignRight)
	
	bodyContainer := tui.NewContainer()
	bodyContainer.SetTitle("Body")
	bodyContainer.SetPadding(tui.NewMargin(1))
	bodyContainer.SetContent(bodyTextArea)
	// Add format indicator
	bodyContainer.AddBorderElement(tui.NewTextElement("JSON"), tui.BorderTop, tui.BorderAlignRight)
	
	responseContainer := tui.NewContainer()
	responseContainer.SetTitle("Response")
	responseContainer.SetPadding(tui.NewMargin(1))
	responseContainer.SetContent(responseViewer)
	// Add tabs for different response views
	responseTabs := tui.NewTabsElement([]string{"Body", "Headers", "Status"})
	responseTabs.SetActiveTab(0)
	responseContainer.AddBorderElement(responseTabs, tui.BorderTop, tui.BorderAlignCenter)
	// Add status indicator
	responseContainer.AddBorderElement(tui.NewStatusElement("Ready"), tui.BorderTop, tui.BorderAlignRight)
	
	return model{
		width:         80,
		height:        24,
		screen:        tui.NewScreen(80, 24, tui.GetTheme("tokyonight")),
		focus:         "url",
		method:        "GET",
		urlInput:      urlInput,
		headersTable:   headersTable,
		bodyTextArea:   bodyTextArea,
		responseTab:    0,
		responseTabs:   responseTabs,
		responseViewer: responseViewer,
		history: []HistoryItem{
			{Method: "GET", URL: "https://api.github.com/users/github", Time: "10:23:45"},
			{Method: "POST", URL: "https://api.example.com/users", Time: "10:15:32"},
			{Method: "GET", URL: "https://jsonplaceholder.typicode.com/posts/1", Time: "09:45:12"},
		},
		historyIndex:   0,
		historyVisible: true,
		currentTheme:      "tokyonight",
		statusBar:         statusBar,
		historyContainer:  historyContainer,
		urlContainer:      urlContainer,
		headersContainer:  headersContainer,
		bodyContainer:     bodyContainer,
		responseContainer: responseContainer,
		headerCountBadge:  headerCountBadge,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// If URL input is focused, let it handle most keys
		if m.focus == "url" && m.urlInput.IsFocused() {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "tab", "enter":
				// Move to next field
				m.urlInput.Blur()
				m.focus = "headers"
			default:
				// Let the input handle the key
				m.urlInput.HandleKey(msg.String())
			}
			return m, nil
		}
		
		// If headers table is focused, let it handle navigation keys
		if m.focus == "headers" && m.headersTable.IsFocused() {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "tab":
				// Continue to next section
				m.headersTable.Blur()
				if m.method != "GET" && m.method != "DELETE" {
					m.focus = "body"
					m.bodyTextArea.Focus()
				} else {
					m.focus = "response"
					m.responseViewer.Focus()
				}
			default:
				// Let the table handle the key
				m.headersTable.HandleKey(msg.String())
			}
			return m, nil
		}
		
		// If body text area is focused, let it handle most keys
		if m.focus == "body" && m.bodyTextArea.IsFocused() {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "tab":
				// Move to next field
				m.bodyTextArea.Blur()
				m.focus = "response"
			default:
				// Let the text area handle the key
				m.bodyTextArea.HandleKey(msg.String())
			}
			return m, nil
		}
		
		// If response viewer is focused, let it handle navigation keys
		if m.focus == "response" && m.responseTab == 0 && m.responseViewer.IsFocused() {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "tab":
				// Move to next section
				m.responseViewer.Blur()
				if m.historyVisible {
					m.focus = "history"
				} else {
					m.focus = "url"
					m.urlInput.Focus()
				}
			case "1", "2", "3":
				// Switch tabs
				switch msg.String() {
				case "1":
					m.responseTab = 0
				case "2":
					m.responseTab = 1
				case "3":
					m.responseTab = 2
				}
			default:
				// Let the viewer handle navigation
				m.responseViewer.HandleKey(msg.String())
			}
			return m, nil
		}
		
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// Update focus states when changing focus
			if m.focus == "url" {
				m.urlInput.Blur()
			} else if m.focus == "headers" {
				m.headersTable.Blur()
			} else if m.focus == "body" {
				m.bodyTextArea.Blur()
			}
			
			// Cycle through focus areas
			switch m.focus {
			case "url":
				m.focus = "headers"
				m.headersTable.Focus()
			case "headers":
				if m.method != "GET" && m.method != "DELETE" {
					m.focus = "body"
					m.bodyTextArea.Focus()
				} else {
					m.focus = "response"
					m.responseViewer.Focus()
				}
			case "body":
				m.focus = "response"
				m.responseViewer.Focus()
			case "response":
				m.responseViewer.Blur()
				if m.historyVisible {
					m.focus = "history"
				} else {
					m.focus = "url"
					m.urlInput.Focus()
				}
			case "history":
				m.focus = "url"
				m.urlInput.Focus()
			}
		case "ctrl+h", "H":
			// Toggle history sidebar
			m.historyVisible = !m.historyVisible
			if !m.historyVisible && m.focus == "history" {
				m.focus = "url"
			}
		case "ctrl+r", "R":
			// Send request (mock for now)
			m.responseStatus = "200 OK"
			m.responseBody = `{
  "login": "github",
  "id": 9919,
  "avatar_url": "https://avatars.githubusercontent.com/u/9919?v=4",
  "type": "Organization",
  "name": "GitHub",
  "blog": "https://github.com/about",
  "location": "San Francisco, CA",
  "email": null,
  "twitter_username": "github",
  "is_verified": true,
  "has_organization_projects": true,
  "has_repository_projects": true,
  "public_repos": 2000,
  "public_gists": 1248,
  "followers": 0,
  "following": 0,
  "html_url": "https://github.com/github",
  "created_at": "2008-01-22T23:15:41Z",
  "updated_at": "2024-01-01T00:00:00Z"
}`
			m.responseHeaders = []Header{
				{Key: "Content-Type", Value: "application/json; charset=utf-8"},
				{Key: "X-RateLimit-Limit", Value: "60"},
				{Key: "X-RateLimit-Remaining", Value: "59"},
			}
			// Update the viewer with the response
			m.responseViewer.SetContent(m.responseBody)
		case "up":
			if m.focus == "history" && m.historyIndex > 0 {
				m.historyIndex--
			}
		case "down":
			if m.focus == "history" && m.historyIndex < len(m.history)-1 {
				m.historyIndex++
			}
		case "enter":
			if m.focus == "history" {
				// Load selected history item
				item := m.history[m.historyIndex]
				m.method = item.Method
				m.urlInput.SetValueWithCursorAtStart(item.URL)
				m.focus = "url"
				m.urlInput.Focus()
			}
		case "ctrl+m", "M":
			// Cycle through HTTP methods
			methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
			currentIndex := 0
			for i, method := range methods {
				if m.method == method {
					currentIndex = i
					break
				}
			}
			m.method = methods[(currentIndex+1)%len(methods)]
			
			// Show/hide body section based on method
			if m.method == "GET" || m.method == "DELETE" {
				// If we're focused on body, move focus to response
				if m.focus == "body" {
					m.bodyTextArea.Blur()
					m.focus = "response"
					m.responseViewer.Focus()
				}
			}
		case "1", "2", "3":
			// Switch response tabs
			if m.focus == "response" {
				switch msg.String() {
				case "1":
					m.responseTab = 0
				case "2":
					m.responseTab = 1
				case "3":
					m.responseTab = 2
				}
				// Update the tabs element to reflect the change
				if m.responseTabs != nil {
					m.responseTabs.SetActiveTab(m.responseTab)
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.screen = tui.NewScreen(m.width, m.height, tui.GetTheme(m.currentTheme))
	}

	return m, nil
}

func (m model) View() string {
	theme := tui.GetTheme(m.currentTheme)
	
	// Recreate screen if theme changed
	if m.screen.Theme().Name != theme.Name {
		m.screen = tui.NewScreen(m.width, m.height, theme)
	}
	
	// Clear screen (now uses theme background automatically)
	m.screen.Clear()

	// Calculate layout
	historyWidth := 0
	if m.historyVisible {
		historyWidth = 30
	}
	
	mainX := historyWidth
	if historyWidth > 0 {
		mainX++ // Add spacing
	}
	mainWidth := m.width - mainX

	// Draw components
	if m.historyVisible {
		m.drawHistory(theme, historyWidth)
	}

	// Fill gap between history and main area
	if historyWidth > 0 && mainX > historyWidth {
		gapStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
		for y := 0; y < m.height; y++ {
			m.screen.DrawRune(historyWidth, y, ' ', gapStyle)
		}
	}

	// Draw main area (accounting for status bar)
	m.drawMainArea(theme, mainX, mainWidth)

	// Draw status bar at the bottom
	m.statusBar.Draw(m.screen, 0, m.height-1, &theme)
	
	// Render screen to string
	return m.screen.Render()
}

func (m *model) drawHistory(theme tui.Theme, width int) {
	// Update container focus state
	if m.focus == "history" {
		m.historyContainer.Focus()
	} else {
		m.historyContainer.Blur()
	}
	
	// Set container size
	m.historyContainer.SetSize(width, m.height-1) // -1 for status bar
	
	// Draw container background and border
	m.historyContainer.DrawWithBounds(m.screen, 0, 0, width, m.height-1, &theme)

	// Draw history items
	itemY := 2
	for i, item := range m.history {
		if itemY >= m.height-1 {
			break
		}

		var itemStyle lipgloss.Style
		if i == m.historyIndex && m.focus == "history" {
			itemStyle = lipgloss.NewStyle().
				Foreground(theme.Components.Interactive.Selected.Text).
				Background(theme.Palette.Background).
				Bold(true)
		} else {
			itemStyle = lipgloss.NewStyle().
				Foreground(theme.Palette.Text).
				Background(theme.Palette.Background)
		}

		// Method
		methodStyle := lipgloss.NewStyle().
			Foreground(m.getMethodColor(item.Method, theme)).
			Background(theme.Palette.Background).
			Bold(true)
		m.screen.DrawString(2, itemY, item.Method, methodStyle)
		
		// URL (truncated)
		urlStart := 2 + len(item.Method) + 1
		maxURLLen := width - urlStart - 3
		url := item.URL
		if len(url) > maxURLLen {
			url = url[:maxURLLen-3] + "..."
		}
		m.screen.DrawString(urlStart, itemY, url, itemStyle)
		
		itemY++
		
		// Time
		timeStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background)
		m.screen.DrawString(2, itemY, "  "+item.Time, timeStyle)
		
		itemY += 2
	}
}

func (m *model) drawMainArea(theme tui.Theme, x, width int) {
	// Calculate heights
	urlHeight := 3
	headersHeight := 8
	bodyHeight := 0
	if m.method != "GET" && m.method != "DELETE" {
		bodyHeight = 6
	}
	responseY := urlHeight + headersHeight + bodyHeight
	responseHeight := m.height - responseY - 1 // -1 for status bar

	// Draw URL input area
	m.drawURLInput(theme, x, 0, width, urlHeight)
	
	// Draw headers area
	m.drawHeaders(theme, x, urlHeight, width, headersHeight)
	
	// Draw request body (if not GET/DELETE)
	if m.method != "GET" && m.method != "DELETE" {
		m.drawRequestBody(theme, x, urlHeight+headersHeight, width, bodyHeight)
	}
	
	// Draw response area
	m.drawResponse(theme, x, responseY, width, responseHeight)
}

func (m *model) drawURLInput(theme tui.Theme, x, y, width, height int) {
	// Update container focus state
	if m.focus == "url" {
		m.urlContainer.Focus()
	} else {
		m.urlContainer.Blur()
	}
	
	// Set container size and draw
	m.urlContainer.SetSize(width, height)
	m.urlContainer.DrawWithBounds(m.screen, x, y, width, height, &theme)
	
	// Draw method
	methodStyle := lipgloss.NewStyle().
		Foreground(m.getMethodColor(m.method, theme)).
		Background(theme.Palette.Background).
		Bold(true)
	m.screen.DrawString(x+2, y+1, m.method, methodStyle)
	
	// Draw the URL input field
	inputX := x + 2 + len(m.method) + 1
	availableWidth := width - 4 - len(m.method) - 1
	m.urlInput.SetWidth(availableWidth)
	m.urlInput.Draw(m.screen, inputX, y+1, &theme)
}

func (m *model) drawHeaders(theme tui.Theme, x, y, width, height int) {
	// Update container focus state
	if m.focus == "headers" {
		m.headersContainer.Focus()
		m.headersTable.Focus()
	} else {
		m.headersContainer.Blur()
		m.headersTable.Blur()
	}
	
	// Set container size and draw
	m.headersContainer.SetSize(width, height)
	m.headersContainer.DrawWithBounds(m.screen, x, y, width, height, &theme)
}

func (m *model) drawRequestBody(theme tui.Theme, x, y, width, height int) {
	// Update container focus state
	if m.focus == "body" {
		m.bodyContainer.Focus()
		m.bodyTextArea.Focus()
	} else {
		m.bodyContainer.Blur()
		m.bodyTextArea.Blur()
	}
	
	// Set container size and draw
	m.bodyContainer.SetSize(width, height)
	m.bodyContainer.DrawWithBounds(m.screen, x, y, width, height, &theme)
}

func (m *model) drawResponse(theme tui.Theme, x, y, width, height int) {
	var borderColors tui.StateColors
	if m.focus == "response" {
		borderColors = theme.Components.Container.Border.Focused
	} else {
		borderColors = theme.Components.Container.Border.Unfocused
	}
	
	borderStyle := lipgloss.NewStyle().
		Foreground(borderColors.Border).
		Background(theme.Palette.Background)

	// Fill background
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			m.screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}
	
	// Draw top border with tabs
	m.screen.DrawRune(x, y, '┌', borderStyle)
	
	// Draw tabs
	tabs := []string{"Body", "Headers", "Status"}
	currentX := x + 1
	for i, tab := range tabs {
		if i > 0 {
			m.screen.DrawRune(currentX, y, '─', borderStyle)
			currentX++
		}
		
		title := " " + tab + " "
		var tabStyle lipgloss.Style
		
		if i == m.responseTab {
			if m.focus == "response" {
				tabStyle = lipgloss.NewStyle().
					Foreground(theme.Components.Tab.Active.Focused.Text).
					Background(theme.Palette.Background).
					Bold(true)
			} else {
				tabStyle = lipgloss.NewStyle().
					Foreground(theme.Components.Tab.Active.Unfocused.Text).
					Background(theme.Palette.Background)
			}
		} else {
			tabStyle = lipgloss.NewStyle().
				Foreground(theme.Components.Tab.Inactive.Text).
				Background(theme.Palette.Background)
		}
		
		m.screen.DrawString(currentX, y, title, tabStyle)
		currentX += len(title)
	}
	
	// Fill rest of top border
	for currentX < x+width-1 {
		m.screen.DrawRune(currentX, y, '─', borderStyle)
		currentX++
	}
	m.screen.DrawRune(x+width-1, y, '┐', borderStyle)
	
	// Draw sides
	for i := 1; i < height-1; i++ {
		m.screen.DrawRune(x, y+i, '│', borderStyle)
		m.screen.DrawRune(x+width-1, y+i, '│', borderStyle)
	}
	
	// Draw bottom border
	m.screen.DrawRune(x, y+height-1, '└', borderStyle)
	for i := 1; i < width-1; i++ {
		m.screen.DrawRune(x+i, y+height-1, '─', borderStyle)
	}
	m.screen.DrawRune(x+width-1, y+height-1, '┘', borderStyle)
	
	// Draw content based on active tab
	contentStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Background(theme.Palette.Background)
	
	// Set viewer focus based on current state
	if m.focus == "response" && m.responseTab == 0 {
		m.responseViewer.Focus()
	} else {
		m.responseViewer.Blur()
	}
	
	switch m.responseTab {
	case 0: // Body
		if m.responseBody != "" {
			// Use the viewer to display the response body
			// The viewer needs the inner area (excluding the border we already drew)
			m.responseViewer.SetSize(width-4, height-2)
			m.responseViewer.Draw(m.screen, x+2, y+1, &theme)
		} else {
			noDataStyle := lipgloss.NewStyle().
				Foreground(theme.Palette.TextMuted).
				Background(theme.Palette.Background).
				Italic(true)
			m.screen.DrawString(x+2, y+2, "No response yet. Press Ctrl+R to send request.", noDataStyle)
		}
		
	case 1: // Headers
		if len(m.responseHeaders) > 0 {
			headerY := y + 1
			for _, header := range m.responseHeaders {
				if headerY >= y+height-1 {
					break
				}
				
				keyStyle := lipgloss.NewStyle().
					Foreground(theme.Palette.Primary).
					Background(theme.Palette.Background).
					Bold(true)
				valueStyle := contentStyle
				
				m.screen.DrawString(x+2, headerY, header.Key+":", keyStyle)
				m.screen.DrawString(x+2+len(header.Key)+2, headerY, header.Value, valueStyle)
				headerY++
			}
		}
		
	case 2: // Status
		if m.responseStatus != "" {
			statusStyle := lipgloss.NewStyle().
				Foreground(theme.Palette.Pine). // Green for 200
				Background(theme.Palette.Background).
				Bold(true)
			m.screen.DrawString(x+2, y+2, "Status: "+m.responseStatus, statusStyle)
		}
	}
}

func (m *model) getMethodColor(method string, theme tui.Theme) lipgloss.TerminalColor {
	switch method {
	case "GET":
		return theme.Palette.Primary // Blue
	case "POST":
		return theme.Palette.Pine // Green
	case "PUT":
		return theme.Palette.Gold // Yellow
	case "DELETE":
		return theme.Palette.Love // Red
	default:
		return theme.Palette.Text
	}
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}