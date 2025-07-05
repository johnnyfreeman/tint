# Modal Examples

This example demonstrates various modal patterns in the Tint TUI library.

## Modal Types

### 1. Simple Modal
- Single container with content
- Used for displaying information, help text, or simple messages
- Press `1` to show

### 2. Fuzzy Finder Modal
- Three-container layout:
  - Left column: Search and Results (stacked with no gap)
  - Right column: Preview (full height)
- Modal provides backdrop/shadow only
- Containers are drawn directly on the modal background
- Press `2` to show
- Controls:
  - Type to filter items
  - Up/Down or j/k to navigate results
  - Enter to select
  - Escape to close

### 3. Confirm Dialog
- Action confirmation with Yes/No buttons
- Button navigation with Tab or arrow keys
- Press `3` to show

### 4. Form Modal
- Multiple input fields
- Tab navigation between fields
- Press `4` to show

## Key Concepts

### Modal as Backdrop
The modal backdrop should be minimal:
- Shadow with 1-pixel offset (bottom-right) for depth
- No background fill - let containers handle their own backgrounds
- This avoids color inconsistencies between modal and container backgrounds

### Container Layout
Containers are drawn directly on the modal backdrop:
- Each container has its own borders and padding
- No extra wrapping container around the content
- Proper spacing between containers

### Shadow Rendering
```go
// Draw shadow with 1 pixel offset
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
```

## Running the Example

```bash
cd examples/modals
go run .
```

## Controls
- `1-4`: Show different modal types
- `Escape`: Close any modal
- `q`: Quit the application