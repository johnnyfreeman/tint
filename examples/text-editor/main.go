package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/johnnyfreeman/tint/tui"
)

// UI Layout constants
const (
	defaultWidth  = 80
	defaultHeight = 24
	explorerWidth = 28
	
	// Modal dimensions
	fuzzyFinderWidth  = 60
	fuzzyFinderHeight = 20
	settingsWidth     = 50
	settingsHeight    = 13
	
	// UI constraints
	maxFuzzyResults = 14
	modalPadding    = 2
)

// Demo data
var (
	demoFiles = []string{
		"main.go", "config.go", "utils.go", "server.go", "client.go",
		"auth.go", "database.go", "middleware.go", "routes.go", "models.go",
		"main_test.go", "config_test.go", "utils_test.go",
		"README.md", "CONTRIBUTING.md", "LICENSE",
	}
	
	explorerFiles = []string{
		"src/",
		"  main.go",
		"  config.go", 
		"  utils.go",
		"tests/",
		"  main_test.go",
		"docs/",
		"  README.md",
		"go.mod",
		"go.sum",
	}
	
	demoContent = map[string]string{
		"main.go": `package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("Tint UI Demo")
    
    // This is just a demo
    app := NewApplication()
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}`,
		"config.go": `package main

type Config struct {
    Theme      string
    FontSize   int
    TabSize    int
    WordWrap   bool
    AutoSave   bool
}

func DefaultConfig() *Config {
    return &Config{
        Theme:    "tokyonight",
        FontSize: 14,
        TabSize:  4,
        WordWrap: false,
        AutoSave: true,
    }
}`,
		"README.md": `# Tint Text Editor Demo

This is a demonstration of Tint's UI components.

## Features

- File explorer (Ctrl+E)
- Fuzzy file finder (Ctrl+P)
- Settings screen (Ctrl+,)
- Multiple tabs
- Syntax highlighting (simulated)

## Keyboard Shortcuts

- **p** - Open fuzzy finder
- **e** - Toggle file explorer
- **s** - Open settings
- **Tab** - Switch between files
- **w** - Close current tab
- **Escape** - Close modals
- **q** - Quit`,
	}
	
	helpContent = `Tint Text Editor - Keyboard Shortcuts

GLOBAL SHORTCUTS
================
  q         Quit application
  ?         Show/hide this help
  Tab       Switch between open files
  w         Close current tab

FILE OPERATIONS
===============
  p         Open fuzzy file finder
  e         Toggle file explorer
  s         Open settings

NAVIGATION
==========
  ↑/↓       Move cursor up/down
  ←/→       Move cursor left/right
  Home      Go to beginning of line
  End       Go to end of line
  PgUp      Scroll up one page
  PgDn      Scroll down one page

EDITING
=======
  Type      Insert text
  Backspace Delete character before cursor
  Delete    Delete character at cursor
  Enter     Insert new line

MODAL CONTROLS
==============
  Escape    Close any modal/dialog
  Enter     Confirm selection
  ↑/↓       Navigate options

FUZZY FINDER
============
  Type      Filter files
  ↑/↓       Select file
  Enter     Open selected file
  Escape    Cancel

SETTINGS
========
  ↑/↓       Navigate settings
  Space     Toggle setting
  Enter     Toggle setting
  Escape    Close settings

Tips:
- Auto-save triggers 5 seconds after you stop typing
- The status bar shows cursor position and file info
- Tabs show unsaved changes with an asterisk (*)
`
	
	fileTypeMap = map[string]string{
		".go":   "Go",
		".md":   "Markdown",
		".js":   "JavaScript",
		".ts":   "TypeScript",
		".py":   "Python",
		".rs":   "Rust",
		".java": "Java",
		".cpp":  "C++",
		".c":    "C",
		".h":    "C/C++ Header",
	}
)

type model struct {
	screen *tui.Screen
	width  int
	height int
	theme  tui.Theme
	
	// UI state
	activeView string // "editor", "explorer", "fuzzy", "settings", "help"
	
	// Components
	editor         *tui.TextArea
	fileExplorer   *tui.Table
	fuzzyFinder    *fuzzyFinderComponent
	settings       *settingsComponent
	helpViewer     *helpViewerComponent
	tabs           *tui.TabsComponent
	notification   *tui.Notification
	statusBar      string
	
	// Demo data
	openFiles []string
	activeTab int
	unsavedFiles map[string]bool
	
	// Auto-save timer
	lastTypedTimer int
}

type fuzzyFinderComponent struct {
	*tui.Modal
	input    *tui.Input
	results  *tui.Table
	allFiles []string
	filtered []string
	selectedIdx int
}

type settingsComponent struct {
	*tui.Modal
	options      []settingOption
	selectedIdx  int
}

type settingOption struct {
	name  string
	value string
	kind  string // "toggle", "select", "input"
}

type helpViewerComponent struct {
	*tui.Modal
	viewer *tui.Viewer
}

func initialModel() *model {
	theme := tui.GetTheme("tokyonight")
	
	// Create main editor
	editor := tui.NewTextArea()
	editor.SetValue(`// Welcome to Tint Text Editor Demo
// This is a UI demonstration of Tint's components

func main() {
    fmt.Println("Hello, Tint!")
    
    // Try these shortcuts:
    // p - Fuzzy file finder
    // e - Toggle file explorer  
    // s - Open settings
    // Tab - Switch between open files
    // q - Quit
}`)
	editor.Focus() // Make sure editor is focused initially
	
	// Create file explorer
	fileExplorer := tui.NewTable()
	fileExplorer.SetColumns([]tui.TableColumn{
		{Title: "📁 Files", Width: 25},
	})
	fileExplorer.SetRows([]tui.TableRow{
		{"📁 src/"},
		{"  📄 main.go"},
		{"  📄 config.go"},
		{"  📄 utils.go"},
		{"📁 tests/"},
		{"  📄 main_test.go"},
		{"📁 docs/"},
		{"  📄 README.md"},
		{"📄 go.mod"},
		{"📄 go.sum"},
	})
	
	// Create fuzzy finder
	fuzzyInput := tui.NewInput()
	fuzzyInput.SetPlaceholder("Type to search files...")
	fuzzyInput.SetWidth(fuzzyFinderWidth - 2*modalPadding)
	
	fuzzyModal := tui.NewModal()
	fuzzyModal.SetTitle("🔍 Find File")
	fuzzyModal.SetSize(fuzzyFinderWidth, fuzzyFinderHeight)
	fuzzyModal.SetCentered(true)
	
	fuzzyResults := tui.NewTable()
	fuzzyResults.SetColumns([]tui.TableColumn{
		{Title: "File", Width: 40},
		{Title: "Path", Width: 18},
	})
	
	fuzzyFinder := &fuzzyFinderComponent{
		Modal:    fuzzyModal,
		input:    fuzzyInput,
		results:  fuzzyResults,
		allFiles: demoFiles,
		filtered: []string{},
	}
	
	// Create settings
	settingsModal := tui.NewModal()
	settingsModal.SetTitle("Settings")
	settingsModal.SetSize(settingsWidth, settingsHeight)
	settingsModal.SetCentered(true)
	
	settings := &settingsComponent{
		Modal: settingsModal,
		options: []settingOption{
			{name: "Theme", value: "Tokyo Night", kind: "select"},
			{name: "Font Size", value: "14", kind: "input"},
			{name: "Line Numbers", value: "On", kind: "toggle"},
			{name: "Word Wrap", value: "Off", kind: "toggle"},
			{name: "Auto Save", value: "On", kind: "toggle"},
			{name: "Tab Size", value: "4", kind: "input"},
			{name: "Minimap", value: "On", kind: "toggle"},
			{name: "Bracket Matching", value: "On", kind: "toggle"},
		},
	}
	
	// Create tabs
	tabs := tui.NewTabs()
	tabs.AddTab("main.go", "")
	tabs.AddTab("config.go", "")
	tabs.AddTab("README.md", "")
	
	// Create notification
	notification := tui.NewNotification()
	notification.SetPosition(tui.NotificationBottomRight)
	notification.SetDuration(3 * time.Second)
	notification.SetSize(20, 3)
	
	// Create help viewer
	helpModal := tui.NewModal()
	helpModal.SetTitle("Help - Keyboard Shortcuts")
	helpModal.SetSize(70, 30)
	helpModal.SetCentered(true)
	
	viewer := tui.NewViewer()
	viewer.SetContent(helpContent)
	viewer.SetWrapText(true)
	viewer.SetSize(66, 26) // Account for modal borders
	
	helpViewer := &helpViewerComponent{
		Modal:  helpModal,
		viewer: viewer,
	}
	
	return &model{
		screen:       tui.NewScreen(defaultWidth, defaultHeight),
		width:        defaultWidth,
		height:       defaultHeight,
		theme:        theme,
		activeView:   "editor",
		editor:       editor,
		fileExplorer: fileExplorer,
		fuzzyFinder:  fuzzyFinder,
		settings:     settings,
		helpViewer:   helpViewer,
		tabs:         tabs,
		notification: notification,
		statusBar:    " NORMAL | main.go | Go | Ln 5, Col 12 | UTF-8 | Spaces: 4",
		openFiles:    []string{"main.go", "config.go", "README.md"},
		activeTab:    0,
		unsavedFiles: make(map[string]bool),
	}
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("Tint Text Editor Demo"),
		tea.WindowSize(), // Request initial window size
		tickCmd(),        // Start the tick timer
	)
}

// tickMsg is sent every second
type tickMsg struct{}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.screen = tui.NewScreen(m.width, m.height)
		
	case tickMsg:
		// Update notification (handles auto-hide)
		m.notification.Update()
		
		// Increment typing timer
		if m.lastTypedTimer > 0 {
			m.lastTypedTimer++
		}
		
		// Trigger auto-save after 5 seconds of no typing
		if m.lastTypedTimer == 5 {
			m.notification.ShowSuccess("Auto-saved")
			m.lastTypedTimer = 0
		}
		return m, tickCmd()
		
	case tea.KeyMsg:
		// Handle escape key first, before routing to components
		if msg.String() == "escape" || msg.String() == "esc" || msg.Type == tea.KeyEsc {
			return m.handleEscape(), nil
		}
		
		// Handle view-specific input
		switch m.activeView {
		case "fuzzy":
			m.handleFuzzyInput(msg)
			return m, nil
			
		case "settings":
			m.handleSettingsInput(msg)
			return m, nil
			
		case "help":
			// Pass input to viewer for scrolling
			m.helpViewer.viewer.HandleKey(msg.String())
			return m, nil
		}
		
		// Global shortcuts
		switch msg.String() {
		case "q":
			return m, tea.Quit
			
		case "?":
			if m.activeView == "help" {
				m.helpViewer.Hide()
				m.activeView = "editor"
			} else {
				m.activeView = "help"
				m.helpViewer.Show()
				m.helpViewer.viewer.Focus()
			}
			return m, nil
			
		case "p":
			m.activeView = "fuzzy"
			m.fuzzyFinder.Show()
			m.fuzzyFinder.input.Focus()
			m.fuzzyFinder.input.SetValue("")
			m.updateFuzzyResults("")
			return m, nil
			
		case "e":
			if m.activeView == "explorer" {
				m.activeView = "editor"
			} else {
				m.activeView = "explorer"
			}
			return m, nil
			
		case "s":
			m.activeView = "settings"
			m.settings.Show()
			return m, nil
			
		case "tab":
			m.activeTab = (m.activeTab + 1) % len(m.openFiles)
			m.tabs.SetActive(m.activeTab)
			m.updateEditorContent()
			return m, nil
			
		case "w":
			// Simulate closing a tab
			if len(m.openFiles) > 1 {
				m.openFiles = append(m.openFiles[:m.activeTab], m.openFiles[m.activeTab+1:]...)
				if m.activeTab >= len(m.openFiles) {
					m.activeTab = len(m.openFiles) - 1
				}
				m.updateTabs()
			}
			return m, nil
		}
		
		// Route input to active view
		switch m.activeView {
		case "explorer":
			// Explorer navigation would be handled here in a real app
		case "editor":
			m.editor.HandleInput(msg.String())
			// Start typing timer when user types in editor
			if len(msg.String()) == 1 || msg.Type == tea.KeyBackspace || msg.Type == tea.KeyDelete {
				m.lastTypedTimer = 1
			}
		}
	}
	
	return m, nil
}

func (m *model) handleEscape() *model {
	switch m.activeView {
	case "fuzzy":
		m.fuzzyFinder.Hide()
	case "settings":
		m.settings.Hide()
	case "help":
		m.helpViewer.Hide()
	case "explorer":
		// Explorer doesn't have a hide method, just switch view
	}
	
	// Always return to editor and focus it
	if m.activeView != "editor" {
		m.activeView = "editor"
		if editor := m.currentEditor(); editor != nil {
			editor.Focus()
		}
	}
	
	return m
}

func (m *model) View() string {
	// Don't render if we don't have a proper size yet
	if m.width == 0 || m.height == 0 {
		return ""
	}
	
	m.screen.Clear()
	
	// Calculate layout
	explorerW := 0
	if m.activeView == "explorer" {
		explorerW = explorerWidth
	}
	
	editorX := explorerW
	editorWidth := m.width - explorerW
	contentHeight := m.height - 2 // Leave room for tabs and status bar
	
	// Ensure minimum sizes
	if contentHeight < 1 {
		contentHeight = 1
	}
	if editorWidth < 1 {
		editorWidth = 1
	}
	
	// Draw file explorer if visible
	if m.activeView == "explorer" {
		// Draw file explorer header
		headerStyle := lipgloss.NewStyle().
			Background(m.theme.Palette.Surface).
			Foreground(m.theme.Palette.Primary).
			Bold(true)
		// Draw header with consistent spacing
		header := "Files"
		// Clear the header line first
		for i := 0; i < explorerW-1; i++ {
			m.screen.SetCell(i, 0, tui.Cell{
				Rune:       ' ',
				Background: m.theme.Palette.Surface,
			})
		}
		m.screen.DrawString(1, 0, header, headerStyle)
		
		// Draw file list manually (simpler than using Table)
		fileStyle := lipgloss.NewStyle().
			Foreground(m.theme.Palette.Text).
			Background(m.theme.Palette.Background)
		selectedStyle := lipgloss.NewStyle().
			Foreground(m.theme.Palette.Background).
			Background(m.theme.Palette.Primary)
		
		files := explorerFiles
		
		for i, file := range files {
			y := i + 1 // Start after header
			if y >= contentHeight {
				break
			}
			
			// Pad file name to fill width
			paddedFile := file
			if len(paddedFile) < explorerW-1 {
				paddedFile += strings.Repeat(" ", explorerW-1-len(paddedFile))
			}
			
			style := fileStyle
			// Simple selection highlight (would track in real app)
			if i == 1 && m.activeView == "explorer" { // Highlight second item as example
				style = selectedStyle
			}
			
			m.screen.DrawString(0, y, paddedFile, style)
		}
		
		// Draw separator
		for y := 0; y < contentHeight; y++ {
			m.screen.SetCell(explorerW-1, y, tui.Cell{
				Rune:       '│',
				Foreground: m.theme.Palette.Border,
			})
		}
	}
	
	// Draw custom tab bar
	m.drawTabBar(editorX, 0, editorWidth)
	
	// Draw editor (start at line 1 after tabs)
	m.editor.SetSize(editorWidth, contentHeight-1)
	m.editor.Draw(m.screen, editorX, 1, &m.theme)
	
	// Draw status bar
	m.drawStatusBar()
	
	// Draw overlays
	if m.activeView == "fuzzy" {
		m.drawFuzzyFinder()
	} else if m.activeView == "settings" {
		m.drawSettings()
	} else if m.activeView == "help" {
		m.drawHelp()
	}
	
	// Draw notification
	m.notification.Draw(m.screen, 0, 0, &m.theme)
	
	return m.screen.Render()
}

func (m *model) handleFuzzyInput(msg tea.KeyMsg) {
	switch msg.String() {
	case "up":
		// Navigate results up
		if m.fuzzyFinder.selectedIdx > 0 {
			m.fuzzyFinder.selectedIdx--
		}
	case "down":
		// Navigate results down
		if m.fuzzyFinder.selectedIdx < len(m.fuzzyFinder.filtered)-1 {
			m.fuzzyFinder.selectedIdx++
		}
	case "enter":
		// Open selected file
		if m.fuzzyFinder.selectedIdx < len(m.fuzzyFinder.filtered) {
			filename := m.fuzzyFinder.filtered[m.fuzzyFinder.selectedIdx]
			m.openFile(filename)
			m.fuzzyFinder.Hide()
			m.activeView = "editor"
		}
	case "escape", "esc":
		// Don't pass escape to input
		return
	default:
		m.fuzzyFinder.input.HandleInput(msg.String())
		m.updateFuzzyResults(m.fuzzyFinder.input.Value())
	}
}

func (m *model) updateFuzzyResults(query string) {
	m.fuzzyFinder.filtered = []string{}
	
	if query == "" {
		m.fuzzyFinder.filtered = m.fuzzyFinder.allFiles[:8] // Show first 8 files
	} else {
		query = strings.ToLower(query)
		for _, file := range m.fuzzyFinder.allFiles {
			if strings.Contains(strings.ToLower(file), query) {
				m.fuzzyFinder.filtered = append(m.fuzzyFinder.filtered, file)
				if len(m.fuzzyFinder.filtered) >= 8 {
					break
				}
			}
		}
	}
	
	// Update results table
	var rows []tui.TableRow
	for _, file := range m.fuzzyFinder.filtered {
		path := "src/"
		if strings.HasSuffix(file, "_test.go") {
			path = "tests/"
		} else if strings.HasSuffix(file, ".md") {
			path = "docs/"
		}
		rows = append(rows, tui.TableRow{file, path})
	}
	m.fuzzyFinder.results.SetRows(rows)
	// Reset selection to first row
	m.fuzzyFinder.selectedIdx = 0
}

func (m *model) handleSettingsInput(msg tea.KeyMsg) {
	switch msg.String() {
	case "up", "k":
		if m.settings.selectedIdx > 0 {
			m.settings.selectedIdx--
		}
	case "down", "j":
		if m.settings.selectedIdx < len(m.settings.options)-1 {
			m.settings.selectedIdx++
		}
	case "enter", "space":
		opt := &m.settings.options[m.settings.selectedIdx]
		if opt.kind == "toggle" {
			if opt.value == "On" {
				opt.value = "Off"
			} else {
				opt.value = "On"
			}
		}
	}
}

func (m *model) openFile(filename string) {
	// Check if already open
	for i, f := range m.openFiles {
		if f == filename {
			m.activeTab = i
			m.tabs.SetActive(i)
			return
		}
	}
	
	// Add to open files
	m.openFiles = append(m.openFiles, filename)
	m.activeTab = len(m.openFiles) - 1
	m.updateTabs()
	m.updateEditorContent()
}

func (m *model) updateTabs() {
	// Clear and rebuild tabs
	newTabs := tui.NewTabs()
	for _, file := range m.openFiles {
		newTabs.AddTab(file, "")
	}
	newTabs.SetActive(m.activeTab)
	m.tabs = newTabs
}

func (m *model) currentEditor() *tui.TextArea {
	return m.editor
}

func (m *model) updateEditorContent() {
	// Update editor with demo content based on active file
	filename := m.openFiles[m.activeTab]
	
	if content, ok := demoContent[filename]; ok {
		m.editor.SetValue(content)
	} else {
		m.editor.SetValue(fmt.Sprintf("// %s\n// File content would appear here", filename))
	}
	
	// Update status bar
	m.statusBar = fmt.Sprintf(" NORMAL | %s | %s | Ln 1, Col 1 | UTF-8 | Spaces: 4", 
		filename, getFileType(filename))
}

func (m *model) drawTabBar(x, y, width int) {
	// Tab bar background
	bgStyle := lipgloss.NewStyle().
		Background(m.theme.Palette.Surface).
		Foreground(m.theme.Palette.Text)
	
	// Clear the tab bar line
	for i := x; i < x+width; i++ {
		m.screen.SetCell(i, y, tui.Cell{
			Rune:       ' ',
			Background: m.theme.Palette.Surface,
		})
	}
	
	// Draw tabs
	currentX := x
	for i, file := range m.openFiles {
		tabText := " " + filepath.Base(file) + " "
		if m.unsavedFiles[file] {
			tabText = " " + filepath.Base(file) + " * "
		}
		
		var style lipgloss.Style
		if i == m.activeTab {
			style = lipgloss.NewStyle().
				Background(m.theme.Palette.Primary).
				Foreground(m.theme.Palette.Background).
				Bold(true)
		} else {
			style = bgStyle
		}
		
		if currentX + len(tabText) < x + width {
			m.screen.DrawString(currentX, y, tabText, style)
			currentX += len(tabText)
			
			// Tab separator
			if i < len(m.openFiles)-1 && currentX < x+width {
				m.screen.DrawString(currentX, y, "│", bgStyle)
				currentX++
			}
		}
	}
}

func (m *model) drawStatusBar() {
	y := m.height - 1
	style := lipgloss.NewStyle().
		Background(m.theme.Palette.Surface).
		Foreground(m.theme.Palette.TextMuted)
	
	// Clear the line
	for x := 0; x < m.width; x++ {
		m.screen.SetCell(x, y, tui.Cell{
			Rune:       ' ',
			Background: m.theme.Palette.Surface,
		})
	}
	
	// Draw status text
	m.screen.DrawString(0, y, m.statusBar, style)
	
	// Draw right-aligned info with shortcuts hint
	rightInfo := "?:help p:find e:explore s:settings q:quit"
	m.screen.DrawString(m.width-len(rightInfo)-1, y, rightInfo, style)
}

func (m *model) drawFuzzyFinder() {
	modalX := (m.width - fuzzyFinderWidth) / 2
	modalY := (m.height - fuzzyFinderHeight) / 2
	
	// Draw modal
	m.fuzzyFinder.Modal.SetSize(fuzzyFinderWidth, fuzzyFinderHeight)
	m.fuzzyFinder.Modal.Draw(m.screen, modalX, modalY, &m.theme)
	
	// Draw input
	m.fuzzyFinder.input.SetWidth(fuzzyFinderWidth - 2*modalPadding)
	m.fuzzyFinder.input.Draw(m.screen, modalX+modalPadding, modalY+modalPadding, &m.theme)
	
	// Draw results manually with selection
	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)
	selectedStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Background).
		Background(m.theme.Palette.Primary)
	
	startY := modalY + 4
	
	for i, file := range m.fuzzyFinder.filtered {
		if i >= maxFuzzyResults {
			break
		}
		
		// Create full-width selection bar
		if i == m.fuzzyFinder.selectedIdx {
			// Fill entire row with selection background
			for x := modalX + 1; x < modalX + fuzzyFinderWidth - 1; x++ {
				m.screen.SetCell(x, startY+i, tui.Cell{
					Rune:       ' ',
					Background: m.theme.Palette.Primary,
				})
			}
		}
		
		style := textStyle
		if i == m.fuzzyFinder.selectedIdx {
			style = selectedStyle
		}
		
		// Draw file name and path
		path := "src/"
		if strings.HasSuffix(file, "_test.go") {
			path = "tests/"
		} else if strings.HasSuffix(file, ".md") {
			path = "docs/"
		}
		
		// Draw file name on the left
		m.screen.DrawString(modalX+modalPadding, startY+i, file, style)
		
		// Draw path on the right
		pathX := modalX + fuzzyFinderWidth - modalPadding - len(path) - modalPadding
		m.screen.DrawString(pathX, startY+i, path, style)
	}
}

func (m *model) drawSettings() {
	modalX := (m.width - settingsWidth) / 2
	modalY := (m.height - settingsHeight) / 2
	
	// Draw modal
	m.settings.Modal.SetSize(settingsWidth, settingsHeight)
	m.settings.Modal.Draw(m.screen, modalX, modalY, &m.theme)
	
	// Draw settings options
	textStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Text).
		Background(m.theme.Palette.Surface)
	selectedStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.Background).
		Background(m.theme.Palette.Primary)
	
	for i, opt := range m.settings.options {
		y := modalY + 2 + i  // Start right after title bar
		
		// Create full-width selection bar
		style := textStyle
		if i == m.settings.selectedIdx {
			// Fill entire row with selection background
			for x := modalX + 1; x < modalX + settingsWidth - 1; x++ {
				m.screen.SetCell(x, y, tui.Cell{
					Rune:       ' ',
					Background: m.theme.Palette.Primary,
				})
			}
			style = selectedStyle
		}
		
		// Draw option name
		label := opt.name
		if len(label) > 25 {
			label = label[:25]
		}
		m.screen.DrawString(modalX+modalPadding, y, label, style)
		
		// Draw value aligned to the right
		valueStyle := style
		if opt.kind == "toggle" {
			if opt.value == "On" {
				valueStyle = lipgloss.NewStyle().
					Foreground(m.theme.Palette.Pine).
					Background(style.GetBackground())
			} else {
				valueStyle = lipgloss.NewStyle().
					Foreground(m.theme.Palette.TextMuted).
					Background(style.GetBackground())
			}
		}
		
		value := opt.value
		if len(value) > 12 {
			value = value[:12]
		}
		// Right-align the value
		valueX := modalX + settingsWidth - modalPadding - len(value)
		m.screen.DrawString(valueX, y, value, valueStyle)
	}
	
	// Draw help text at bottom of modal
	helpStyle := lipgloss.NewStyle().
		Foreground(m.theme.Palette.TextMuted).
		Background(m.theme.Palette.Surface)
	helpText := "↑↓ Navigate  Space/Enter Toggle  Esc Close"
	// Center the help text
	helpX := modalX + (settingsWidth - len(helpText)) / 2
	m.screen.DrawString(helpX, modalY + settingsHeight - 2, helpText, helpStyle)
}

func (m *model) drawHelp() {
	modalX := (m.width - 70) / 2
	modalY := (m.height - 30) / 2
	
	// Draw modal
	m.helpViewer.Modal.Draw(m.screen, modalX, modalY, &m.theme)
	
	// Draw viewer inside modal
	m.helpViewer.viewer.Draw(m.screen, modalX+2, modalY+2, &m.theme)
}

func getFileType(filename string) string {
	ext := filepath.Ext(filename)
	if fileType, ok := fileTypeMap[ext]; ok {
		return fileType
	}
	return "Plain Text"
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}