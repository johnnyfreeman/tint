# Modal/Container Pattern

This document explains the proper pattern for using modals and containers in Tint.

## Architecture

The modal and container components have distinct responsibilities:

### Modal Component
The `Modal` component provides:
- **Elevated surface** - Different background color (Surface color from theme)
- **Drop shadow** - Neo-brutalist style shadow offset by 1 cell
- **Positioning logic** - Can be centered or positioned manually
- **Visibility management** - Show/Hide methods

The Modal does NOT provide:
- Borders
- Title bars  
- Padding
- Content management

### Container Component
The `Container` component provides:
- **Structure** - Border, title, padding
- **Content management** - Holds child components
- **Focus handling** - Manages keyboard focus
- **Border elements** - Tabs, status indicators, badges, etc.

## Usage Pattern

Always use a Container inside a Modal to provide structure:

```go
// 1. Create the modal (elevated surface)
modal := tui.NewModal()
modal.SetSize(50, 20)
modal.SetCentered(true)
modal.Show()

// 2. Draw the modal
modal.Draw(screen, 0, 0, theme)

// 3. Get modal position for container placement
modalWidth, modalHeight := modal.GetSize()
modalX := (screenWidth - modalWidth) / 2
modalY := (screenHeight - modalHeight) / 2

// 4. Create container that fills the modal
container := tui.NewContainer()
container.SetTitle("My Dialog")
container.SetSize(modalWidth, modalHeight)  // Same size as modal
container.SetPadding(tui.NewMargin(1))

// 5. Add content to container
content := tui.NewTextArea()
content.SetValue("Dialog content goes here")
content.SetSize(modalWidth-4, modalHeight-4)  // Account for border + padding
container.SetContent(content)

// 6. Draw container at modal position
container.Draw(screen, modalX, modalY, theme)
```

## Examples

### Simple Information Modal

```go
func drawInfoModal(screen *tui.Screen, theme *tui.Theme) {
    // Create and show modal
    modal := tui.NewModal()
    modal.SetSize(40, 10)
    modal.SetCentered(true)
    modal.Show()
    modal.Draw(screen, 0, 0, theme)
    
    // Calculate position
    modalX := (screen.Width() - 40) / 2
    modalY := (screen.Height() - 10) / 2
    
    // Create container with content
    container := tui.NewContainer()
    container.SetTitle("Information")
    container.SetSize(40, 10)
    container.SetPadding(tui.NewMargin(1))
    
    viewer := tui.NewViewer()
    viewer.SetContent("This is an information dialog.\n\nPress ESC to close.")
    viewer.SetSize(36, 6)
    container.SetContent(viewer)
    
    container.Draw(screen, modalX, modalY, theme)
}
```

### Confirmation Dialog

```go
func drawConfirmModal(screen *tui.Screen, theme *tui.Theme) {
    modal := tui.NewModal()
    modal.SetSize(40, 8)
    modal.SetCentered(true)
    modal.Show()
    modal.Draw(screen, 0, 0, theme)
    
    modalX := (screen.Width() - 40) / 2
    modalY := (screen.Height() - 8) / 2
    
    container := tui.NewContainer()
    container.SetTitle("Confirm")
    container.SetSize(40, 8)
    container.Draw(screen, modalX, modalY, theme)
    
    // Draw custom content for yes/no buttons
    question := "Are you sure?"
    screen.DrawString(modalX + 20 - len(question)/2, modalY + 3, question, textStyle)
    
    // Draw buttons
    screen.DrawString(modalX + 10, modalY + 5, "[Y] Yes", yesStyle)
    screen.DrawString(modalX + 25, modalY + 5, "[N] No", noStyle)
}
```

### Form Modal

```go
func drawFormModal(screen *tui.Screen, theme *tui.Theme) {
    modal := tui.NewModal()
    modal.SetSize(60, 20)
    modal.SetCentered(true)
    modal.Show()
    modal.Draw(screen, 0, 0, theme)
    
    modalX := (screen.Width() - 60) / 2
    modalY := (screen.Height() - 20) / 2
    
    container := tui.NewContainer()
    container.SetTitle("User Registration")
    container.SetSize(60, 20)
    container.SetPadding(tui.NewMargin(2))
    container.Draw(screen, modalX, modalY, theme)
    
    // Create form fields...
}
```

## Best Practices

1. **Always use Container inside Modal** - The modal is just the elevated surface; the container provides the structure.

2. **Size the Container to fill the Modal** - Set the container size equal to the modal size for a cohesive appearance.

3. **Account for padding in content size** - When sizing content inside a container, subtract border (2) and padding from the container size.

4. **Use theme colors consistently** - The modal uses Surface color, containers use theme-appropriate border colors.

5. **Handle focus properly** - Containers manage focus for their content; modals just control visibility.

## Common Mistakes

❌ **Don't add padding between modal and container**
```go
// Wrong - adds unnecessary spacing
container.Draw(screen, modalX + 2, modalY + 2, theme)
```

✅ **Do fill the modal completely**
```go
// Correct - container fills modal
container.Draw(screen, modalX, modalY, theme)
```

❌ **Don't try to add borders to modals**
```go
// Wrong - modals don't have borders
modal.SetBorderStyle("rounded")  // This method doesn't exist
```

✅ **Do use containers for borders**
```go
// Correct - containers provide borders
container.SetBorderStyle("rounded")
```

## See Also

- [Simple Modal Example](../examples/simple-modal/main.go)
- [Modal/Container Pattern Demo](../examples/modal-container-pattern/main.go)
- [Modal Test Suite](../examples/modal-test/main.go)