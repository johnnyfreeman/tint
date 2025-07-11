# Tint Examples

This directory contains examples demonstrating various features and patterns in the Tint TUI library.

## Modal Examples

### [simple-modal](./simple-modal/)
A minimal example showing the basic modal/container pattern. Press space to show/hide a modal with a container inside.

### [modal-test](./modal-test/)
Comprehensive modal testing suite with 6 different modal types:
- Simple modal with text
- Large modal with scrollable content
- Small modal
- Modal with formatted text
- Empty modal (just the surface, no content)
- Modal with multiple containers

### [modal-container-pattern](./modal-container-pattern/)
Demonstrates the proper architectural pattern for modals:
- Information modal
- Confirmation dialog
- Form modal with input fields

### [modals](./modals/)
Advanced modal patterns including:
- Fuzzy finder with three-pane layout
- Confirmation dialogs
- Multi-field forms

## Component Examples

### [container-demo](./container-demo/)
Showcases container border elements:
- Simple titles
- Embedded tabs
- Status indicators
- Icons and badges
- Multi-element borders

### [text-editor](./text-editor/)
A full-featured text editor demo with:
- File explorer sidebar
- Tabbed editing
- Fuzzy file finder (modal)
- Settings dialog (modal)
- Help viewer (modal)
- Status bar

### [screen](./screen/)
Basic screen rendering with:
- Container drawing
- Modal with container overlay
- Theme-aware rendering

### [demo](./demo/)
General showcase including:
- Sidebar navigation
- Tabbed content
- Modal dialogs
- Notifications
- Theme picker
- Status bar

### [api-client](./api-client/)
REST API client interface demonstrating:
- Split views
- Request/response containers
- Header tables
- JSON syntax highlighting

## Unicode Examples

### [unicode-test](./unicode-test/)
Tests Unicode rendering including:
- Wide characters (CJK)
- Emojis
- RTL text
- Combining characters

### [unicode-theme-test](./unicode-theme-test/)
Tests Unicode with different themes to ensure proper rendering across color schemes.

## Running Examples

Each example can be run from its directory:

```bash
cd examples/simple-modal
go run .
```

Or built and run:

```bash
go build -o simple-modal
./simple-modal
```

## Key Patterns

### Modal/Container Pattern
Modals provide elevated surfaces with shadows, while containers provide structure (borders, titles, padding). Always use a container inside a modal:

```go
// Create modal (elevated surface)
modal := tui.NewModal()
modal.SetSize(40, 10)
modal.SetCentered(true)

// Draw modal
modal.Draw(screen, 0, 0, screenWidth, screenHeight, theme)

// Create container (structure)
container := tui.NewContainer()
container.SetTitle("My Dialog")
container.SetSize(40, 10)  // Fill the modal
container.SetContent(content)

// Draw container at modal position
container.Draw(screen, modalX, modalY, 40, 10, theme)
```

### Theme Support
All examples use the theme system for consistent styling:

```go
theme := tui.GetTheme("tokyonight")
screen := tui.NewScreen(width, height, theme)
```

Available themes include:
- tokyonight
- solarized
- dracula
- catppuccin
- rosepine
- tokyonight-light
- solarized-light