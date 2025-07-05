# Tint

A refreshingly simple terminal UI (TUI) component library for Go, built on top of [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss). Theme your TUIs with ease!

## Why Tint?

1. **Cell-based rendering** - Prevents ANSI escape sequence corruption that plagues other TUI libraries
2. **Theme-first design** - Consistent, beautiful UIs out of the box with multiple built-in themes
3. **Terminal-native layouts** - Designed specifically for terminal constraints, not web concepts
4. **Production-ready** - Used in real applications with complex requirements
5. **Easy to learn** - Consistent patterns across all components

## Features

- **Cell-based rendering** - Prevents ANSI escape sequence corruption
- **Comprehensive theming system** - Multiple built-in themes with state-based styling
- **Rich component library** - Containers, inputs, tables, text areas, viewers, modals, tabs, and notifications
- **Constraint-based layout system** - Terminal-native layouts with flexible sizing
- **Full Unicode support** - Proper handling of emojis, CJK characters, and combining marks
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

## Layout System

Tint provides a powerful constraint-based layout system designed specifically for terminal UIs:

### Linear Layout (HBox/VBox)

Arrange components horizontally or vertically with flexible sizing:

```go
// Horizontal layout
header := tui.HBox().
    AddFixed(logo, 20).           // Fixed 20 columns
    AddFlex(title, 1).            // Flexible space
    AddFixed(clock, 10)           // Fixed 10 columns

// Vertical layout  
sidebar := tui.VBox().
    AddFixed(header, 3).          // Fixed 3 rows
    AddFlex(content, 1).          // Takes remaining space
    AddFixed(statusBar, 1)        // Fixed 1 row
```

### Split Layout

Create resizable split panes:

```go
// Vertical split (side by side)
split := tui.NewVSplit().
    SetFixed(200).                // Fixed left panel
    SetFirst(sidebar).
    SetSecond(mainContent)

// Horizontal split (top/bottom)
split := tui.NewHSplit().
    SetPercentage(0.7).           // Top takes 70%
    SetFirst(editor).
    SetSecond(terminal)
```

### Stack Layout

Layer components on top of each other:

```go
// Create layered UI with modals
stack := tui.NewStack().
    AddLayer(mainView).           // Base layer
    AddLayer(modal, tui.StackItem{
        X: tui.Percentage(0.25),  // Center at 25%
        Y: tui.Percentage(0.25),
        Width: tui.Percentage(0.5),
        Height: tui.Percentage(0.5),
    })
```

### Conditional Layout

Responsive layouts based on terminal size:

```go
layout := tui.NewConditional().
    AddCondition(func(w, h int) bool {
        return w < 80  // Mobile view
    }, mobileLayout).
    AddCondition(func(w, h int) bool {
        return w < 120 // Tablet view  
    }, tabletLayout).
    SetFallback(desktopLayout)
```

## Core Components

### Container

The fundamental building block for creating bordered regions:

```go
// Create a container with padding and title
container := tui.NewContainer().
    SetTitle("Settings").
    SetBorderStyle(tui.BorderStyleRounded).
    SetPadding(tui.NewPadding(1, 2, 1, 2)). // top, right, bottom, left
    SetContent(myComponent)

// Or use builder methods
container := tui.BoxWithTitle("User Profile", userForm).
    WithPadding(2).
    WithStyle(tui.BorderStyleDouble)
```

### Screen

The foundation of the rendering system with theme support:

```go
// Create themed screen
theme := tui.GetTheme("tokyonight")
screen := tui.NewScreen(width, height, theme)

// Draw with proper Unicode support
screen.DrawString(x, y, "Hello 世界! 🎉", style)
screen.Render() // Returns the final output
```

### Input

Single-line text input with full Unicode support:

```go
input := tui.NewInput()
input.SetPlaceholder("Enter name...")
input.SetWidth(40)
input.Focus()

// Wrap in a container for better presentation
container := tui.BoxWithTitle("Username", input)
container.Draw(screen, x, y, theme)
```

### TextArea

Multi-line text editor with Unicode-aware cursor navigation:

```go
textarea := tui.NewTextArea()
textarea.SetSize(80, 20)
textarea.SetValue("Initial content")
textarea.Focus()

// Draw with container
container := tui.BoxWithTitle("Editor", textarea).
    WithPadding(1)
container.Draw(screen, x, y, theme)
```

### Table

Editable table with cell navigation and scrolling:

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

// Wrap in container
container := tui.BoxWithTitle("Data", table)
container.Draw(screen, x, y, theme)
```

### Modal

Modal dialogs that work as elevated surfaces:

```go
// Create modal content
content := tui.VBox().
    AddFixed(tui.NewViewer().SetContent("Are you sure?"), 3).
    AddFixed(buttons, 3)

// Create modal (just provides elevation)
modal := tui.NewModal().SetSize(40, 10)

// Wrap content in container for structure
container := tui.BoxWithTitle("Confirm", content).
    WithStyle(tui.BorderStyleDouble)

// Set container as modal content
modal.SetContent(container)
modal.Show()

// Draw modal (centers automatically)
modal.Draw(screen, 0, 0, theme)
```

## Theming

The library includes a comprehensive theming system with multiple built-in themes:

- **Tokyo Night** - A clean, dark theme with vibrant colors
- **Rosé Pine** - Soho vibes with muted colors
- **Catppuccin** - Soothing pastel theme
- **Monochrome** - High contrast black and white
- **Brutalist** - Bold, stark design aesthetic

```go
// Get a theme
theme := tui.GetTheme("tokyonight")

// Themes provide semantic colors
theme.Palette.Primary    // Main brand color
theme.Palette.Text       // Primary text
theme.Palette.Background // Background color

// Component-specific styles
theme.Components.Interactive.Selected // Selected item style
theme.Components.Container.Border.Focused // Focused border
theme.Components.Tab.Active.Focused // Active tab style

// Inline elements for special UI elements
theme.InlineElements.CodeBlock    // For code snippets
theme.InlineElements.Strong       // Bold/important text
theme.InlineElements.Link         // Hyperlinks
```

## Unicode Support

Tint provides comprehensive Unicode support:

- **Wide characters** - Proper handling of CJK characters and emojis
- **Combining marks** - Correct rendering of diacritics
- **Zero-width joiners** - Emoji sequences render correctly
- **Grapheme clusters** - Multi-codepoint characters work seamlessly

```go
// All components handle Unicode properly
input.SetValue("Hello 世界! 👨‍👩‍👧‍👦")
textarea.SetValue("Café résumé naïve")
table.SetCell(0, 0, "🇺🇸 🇯🇵 🇰🇷")
```

## Complete Example

```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/johnnyfreeman/tint/tui"
)

type model struct {
    screen     *tui.Screen
    layout     *tui.LinearLayout
    input      *tui.Input
    textarea   *tui.TextArea
    focusChain *tui.FocusChain
    theme      *tui.Theme
    width      int
    height     int
}

func initialModel() model {
    // Create components
    input := tui.NewInput()
    input.SetPlaceholder("Enter command...")
    
    textarea := tui.NewTextArea()
    textarea.SetSize(60, 10)
    
    // Create layout
    layout := tui.VBox().
        AddFixed(tui.BoxWithTitle("Input", input), 5).
        AddFlex(tui.BoxWithTitle("Output", textarea), 1).
        SetPadding(tui.NewPadding(1))
    
    // Setup focus chain
    focusChain := tui.NewFocusChain()
    focusChain.Add(input)
    focusChain.Add(textarea)
    focusChain.Next() // Focus first component
    
    return model{
        layout:     layout,
        input:      input,
        textarea:   textarea,
        focusChain: focusChain,
        theme:      tui.GetTheme("tokyonight"),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "tab":
            m.focusChain.Next()
        case "shift+tab":
            m.focusChain.Previous()
        default:
            // Let focused component handle the key
            if m.input.IsFocused() {
                m.input.HandleKey(msg.String())
            } else if m.textarea.IsFocused() {
                m.textarea.HandleKey(msg.String())
            }
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.screen = tui.NewScreen(m.width, m.height, *m.theme)
        m.layout.SetSize(m.width, m.height)
    }
    return m, nil
}

func (m model) View() string {
    if m.screen == nil {
        return ""
    }
    
    m.screen.Clear()
    m.layout.Draw(m.screen, 0, 0, m.theme)
    return m.screen.Render()
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        panic(err)
    }
}
```

## Examples

Check out the `examples/` directory for complete applications:

- **demo** - Comprehensive demo showcasing all components
- **api-client** - REST API client demonstrating real-world usage
- **text-editor** - Full-featured text editor with fuzzy finder
- **layouts** - Interactive demonstration of all layout types
- **container-demo** - Container styling and theming examples
- **modals** - Various modal dialog patterns

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT