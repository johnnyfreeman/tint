# Tint

A refreshingly simple terminal UI (TUI) component library for Go, built on top of [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss). Build beautiful, themeable terminal applications without fighting ANSI escape sequences.

## Why Tint?

Traditional TUI libraries often suffer from rendering artifacts, ANSI escape sequence corruption, and inconsistent theming. Tint solves these problems with a unique cell-based rendering approach that ensures your UI always looks perfect.

**Key Benefits:**
- **No more corrupted output** - Cell-based rendering prevents ANSI escape sequence conflicts
- **Beautiful by default** - Ships with carefully crafted themes (Tokyo Night, Catppuccin, Rosé Pine, and more)
- **Terminal-native layouts** - Constraint-based system designed for terminal dimensions, not web concepts
- **Full Unicode support** - Handles emojis, CJK characters, and complex text properly
- **Consistent patterns** - Every component follows the same interface

## Installation

```bash
go get github.com/johnnyfreeman/tint
```

## Quick Start

```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/johnnyfreeman/tint/tui"
)

type model struct {
    screen *tui.Screen
    input  *tui.Input
    theme  *tui.Theme
}

func initialModel() model {
    input := tui.NewInput()
    input.SetPlaceholder("Type something...")
    input.Focus()
    
    return model{
        input: input,
        theme: tui.GetTheme("tokyonight"),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "esc":
            return m, tea.Quit
        default:
            m.input.HandleKey(msg.String())
        }
    case tea.WindowSizeMsg:
        m.screen = tui.NewScreen(msg.Width, msg.Height, *m.theme)
    }
    return m, nil
}

func (m model) View() string {
    if m.screen == nil {
        return ""
    }
    
    m.screen.Clear()
    
    // Draw input in a nice container
    container := tui.BoxWithTitle("Input Example", m.input)
    container.Draw(m.screen, 2, 2, m.theme)
    
    return m.screen.Render()
}

func main() {
    if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
        panic(err)
    }
}
```

## Layout System

Tint provides a powerful constraint-based layout system that makes it easy to build responsive terminal UIs:

### Linear Layouts (HBox/VBox)

```go
// Create a typical application layout
appLayout := tui.VBox().
    AddFixed(header, 3).          // Fixed height header
    AddFlex(tui.HBox().          // Main content area
        AddFixed(sidebar, 30).    // Fixed width sidebar
        AddFlex(content, 1),      // Content takes remaining space
    1).
    AddFixed(statusBar, 1)        // Fixed height status bar
```

### Split Panes

```go
// Create resizable split views
split := tui.NewVSplit().
    SetRatio(0.3).               // 30/70 split
    SetFirst(fileTree).
    SetSecond(editor)
```

### Stacked Layers

```go
// Layer components for modals and overlays
stack := tui.NewStack().
    AddLayer(mainView).
    AddLayer(modal, tui.StackItem{
        X: tui.Percentage(0.25),
        Y: tui.Percentage(0.25),
        Width: tui.Percentage(0.5),
        Height: tui.Percentage(0.5),
    })
```

### Responsive Layouts

```go
// Different layouts for different terminal sizes
layout := tui.NewConditional().
    When(func(w, h int) bool { return w < 80 }, mobileLayout).
    When(func(w, h int) bool { return w < 120 }, tabletLayout).
    Otherwise(desktopLayout)
```

## Components

### Container

The fundamental building block for creating structured UIs:

```go
// Simple box with title
box := tui.BoxWithTitle("Settings", settingsForm)

// Customized container
container := tui.NewContainer().
    SetTitle("User Profile").
    SetBorderStyle(tui.BorderStyleDouble).
    SetPadding(tui.NewPadding(1, 2, 1, 2)).
    SetContent(profileView)
```

### Input

Unicode-aware text input with placeholder support:

```go
input := tui.NewInput()
input.SetPlaceholder("Search...")
input.SetWidth(40)
input.OnChange(func(value string) {
    // Handle input changes
})
```

### TextArea

Multi-line text editor with full Unicode support:

```go
editor := tui.NewTextArea()
editor.SetSize(80, 24)
editor.SetSyntaxHighlighting("go") // Coming soon
editor.SetLineNumbers(true)         // Coming soon
```

### Table

Data tables with navigation and editing:

```go
table := tui.NewTable()
table.SetColumns([]tui.TableColumn{
    {Title: "Name", Width: 30},
    {Title: "Status", Width: 15},
    {Title: "Last Updated", Width: 20},
})
table.SetData(myData)
table.SetEditable(true)
```

### Modal

Elevated surfaces for dialogs and overlays:

```go
// Create modal content
confirmDialog := tui.VBox().
    AddFixed(tui.NewViewer().SetContent("Delete this file?"), 3).
    AddFixed(buttonRow, 3)

// Wrap in modal for elevation
modal := tui.NewModal().
    SetSize(40, 10).
    SetContent(tui.BoxWithTitle("Confirm", confirmDialog))
```

### Additional Components

- **Viewer** - Scrollable read-only text display
- **Tabs** - Tabbed container for organizing content
- **StatusBar** - Information display bar
- **Notification** - Toast-style notifications
- **FocusChain** - Keyboard navigation between components

## Theming

Tint ships with beautiful themes and a powerful theming system:

```go
// Built-in themes
theme := tui.GetTheme("tokyonight")    // Default
theme := tui.GetTheme("catppuccin")    // Soothing pastels
theme := tui.GetTheme("rose-pine")     // Soho vibes
theme := tui.GetTheme("brutalist")     // Bold minimalism
theme := tui.GetTheme("monochrome")    // High contrast

// Access theme colors
theme.Palette.Primary
theme.Palette.Background
theme.Palette.Text

// Component-specific styles
theme.Components.Container.Border.Focused
theme.Components.Interactive.Selected
```

## Component Interface

All components implement a consistent interface:

```go
type Component interface {
    Draw(screen *Screen, x, y int, theme *Theme)
    Focus()
    Blur()
    IsFocused() bool
    HandleKey(key string) bool
}
```

## Full Example

Here's a more complete example showing layouts, multiple components, and focus management:

```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/johnnyfreeman/tint/tui"
)

type app struct {
    screen     *tui.Screen
    layout     *tui.LinearLayout
    searchBox  *tui.Input
    resultList *tui.Table
    details    *tui.Viewer
    focusChain *tui.FocusChain
    theme      *tui.Theme
}

func newApp() app {
    // Create components
    search := tui.NewInput()
    search.SetPlaceholder("Search users...")
    
    results := tui.NewTable()
    results.SetColumns([]tui.TableColumn{
        {Title: "Username", Width: 20},
        {Title: "Email", Width: 30},
    })
    
    details := tui.NewViewer()
    details.SetContent("Select a user to view details")
    
    // Build layout
    layout := tui.VBox().
        AddFixed(tui.BoxWithTitle("Search", search), 5).
        AddFlex(tui.HBox().
            AddFlex(tui.BoxWithTitle("Results", results), 1).
            AddFixed(tui.BoxWithTitle("Details", details), 40),
        1).
        SetPadding(tui.NewPadding(1))
    
    // Setup focus navigation
    focus := tui.NewFocusChain()
    focus.Add(search, results)
    focus.Next() // Focus search by default
    
    return app{
        layout:     layout,
        searchBox:  search,
        resultList: results,
        details:    details,
        focusChain: focus,
        theme:      tui.GetTheme("tokyonight"),
    }
}

func (a app) Init() tea.Cmd {
    return nil
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return a, tea.Quit
        case "tab":
            a.focusChain.Next()
        case "shift+tab":
            a.focusChain.Previous()
        default:
            // Send to focused component
            focused := a.focusChain.Current()
            if focused != nil {
                focused.HandleKey(msg.String())
            }
        }
    case tea.WindowSizeMsg:
        a.screen = tui.NewScreen(msg.Width, msg.Height, *a.theme)
        a.layout.SetSize(msg.Width, msg.Height)
    }
    return a, nil
}

func (a app) View() string {
    if a.screen == nil {
        return ""
    }
    
    a.screen.Clear()
    a.layout.Draw(a.screen, 0, 0, a.theme)
    return a.screen.Render()
}

func main() {
    if _, err := tea.NewProgram(newApp()).Run(); err != nil {
        panic(err)
    }
}
```

## Examples

Explore the `examples/` directory for more complete applications:

- **demo** - Comprehensive component showcase
- **api-client** - REST API client with request/response viewer
- **text-editor** - Text editor with fuzzy file finder
- **layouts** - Interactive layout demonstrations
- **container-demo** - Container styling examples
- **modals** - Modal dialog patterns

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT