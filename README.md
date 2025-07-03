# Tint

A refreshingly simple terminal UI (TUI) component library for Go, built on top of [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss). Theme your TUIs with ease!

## Features

- **Cell-based rendering** - Prevents ANSI escape sequence corruption
- **Comprehensive theming system** - Multiple built-in themes with state-based styling
- **Rich component library** - Input fields, tables, text areas, viewers, modals, tabs, and notifications
- **Consistent component interface** - All components implement a standard interface
- **Focus management** - Proper keyboard navigation between components
- **No rendering artifacts** - Clean, flicker-free updates

## Installation

```bash
go get github.com/johnnyfreeman/tint
```

## Component Interface

All components implement a consistent interface for easy integration:

```go
type Component interface {
    // Draw renders the component to the screen at the specified position
    Draw(screen *Screen, x, y int, theme *Theme)
    
    // Focus management
    Focus()
    Blur()
    IsFocused() bool
    
    // Keyboard input handling
    HandleKey(key string) bool
}
```

## Core Components

### Screen
The foundation of the rendering system. Manages a 2D grid of cells.

```go
screen := tui.NewScreen(width, height)
screen.DrawString(x, y, "Hello", style)
screen.Render() // Returns the final output
```

### Input
Single-line text input with cursor and scrolling support.

```go
input := tui.NewInput()
input.SetPlaceholder("Enter text...")
input.SetWidth(40)
input.Focus()

// Draw it
input.Draw(screen, x, y, theme)
```

### TextArea
Multi-line text editor with cursor navigation.

```go
textarea := tui.NewTextArea()
textarea.SetSize(80, 20)
textarea.SetValue("Initial content")
textarea.Focus()

// Draw with border
textarea.DrawWithBorder(screen, x, y, theme, "Editor")
```

### Table
Editable table with cell navigation and scrolling.

```go
table := tui.NewTable()
table.SetColumns([]tui.TableColumn{
    {Title: "Name", Width: 20},
    {Title: "Value", Width: 40},
})
table.SetRows([]tui.TableRow{
    {"Key1", "Value1"},
    {"Key2", "Value2"},
})

// Draw in a box
table.DrawInBox(screen, x, y, width, height, "Data", theme)
```

### Viewer
Read-only scrollable text viewer with scroll indicators.

```go
viewer := tui.NewViewer()
viewer.SetContent(longText)
viewer.SetWrapText(true)
viewer.SetSize(80, 20)

// Draw it
viewer.Draw(screen, x, y, theme)
```

### Modal
Modal dialog component with customizable content.

```go
modal := tui.NewModal()
modal.SetTitle("Confirm")
modal.SetContent("Are you sure you want to proceed?")
modal.SetSize(40, 10)
modal.Show()

// Draw centered by default
modal.Draw(screen, 0, 0, theme)
```

### Tabs
Tabbed container for organizing content.

```go
tabs := tui.NewTabs()
tabs.AddTab("General", "General settings content")
tabs.AddTab("Advanced", advancedComponent)
tabs.SetSize(60, 20)

// Draw it
tabs.Draw(screen, x, y, theme)
```

### Notification
Toast-style notifications with different types.

```go
notif := tui.NewNotification()
notif.SetPosition(tui.NotificationBottomRight)
notif.SetDuration(3 * time.Second)

// Show different types
notif.ShowSuccess("Operation completed!")
notif.ShowError("Something went wrong")
notif.ShowWarning("Please check your input")
notif.ShowInfo("New update available")

// Draw it (auto-positions based on settings)
notif.Draw(screen, 0, 0, theme)
```

## Theming

The library includes a comprehensive theming system with multiple built-in themes:

- **Tokyo Night** - A clean, dark theme with vibrant colors
- **Rosé Pine** - Soho vibes with muted colors
- **Catppuccin** - Soothing pastel theme
- **Monochrome** - High contrast black and white

```go
theme := tui.GetTheme("tokyonight")

// Themes provide semantic colors
theme.Palette.Primary    // Main brand color
theme.Palette.Text       // Primary text
theme.Palette.Background // Background color

// And component-specific styles
theme.Components.Interactive.Selected // Selected item style
theme.Components.Container.Border.Focused // Focused border
theme.Components.Tab.Active.Focused // Active tab style
```

## Examples

### Basic Example

```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/johnnyfreeman/tint/tui"
)

type model struct {
    screen *tui.Screen
    input  *tui.Input
    width  int
    height int
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
        // Pass key to focused component
        m.input.HandleKey(msg.String())
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.screen = tui.NewScreen(m.width, m.height)
    }
    return m, nil
}

func (m model) View() string {
    theme := tui.GetTheme("tokyonight")
    m.screen.Clear()
    
    // Draw input field
    m.input.Draw(m.screen, 10, 5, theme)
    
    return m.screen.Render()
}
```

### Component Integration

All components follow the same pattern:

```go
// Create and configure
component := tui.NewTextArea()
component.SetSize(60, 10)
component.Focus()

// Handle input when focused
if component.IsFocused() {
    handled := component.HandleKey(key)
}

// Draw at position
component.Draw(screen, x, y, theme)

// Or draw with border
component.DrawWithBorder(screen, x, y, theme, "Title")
```

### Complete Examples

Check out the `examples/` directory for complete applications:

- **demo** - Comprehensive demo showcasing all components
- **api-client** - A REST API client UI demonstrating real-world usage

## Why Cell-Based Rendering?

Traditional TUI libraries often suffer from rendering artifacts when components overlap or when ANSI escape sequences get corrupted. This library solves these issues by:

1. **Separating content from presentation** - Each cell stores its character and style separately
2. **Atomic updates** - The entire screen is rendered in one pass
3. **Proper layering** - Components can safely overlap without corruption

## Unicode Support

Note: The library currently assumes one character per cell. For proper Unicode support (including wide characters and combining characters), consider using ASCII characters for UI elements like borders and indicators.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT