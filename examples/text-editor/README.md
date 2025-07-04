# Tint Text Editor UI Demo

A UI demonstration showcasing Tint's components in a text editor interface. This is a mock UI that demonstrates component composition and interactions without implementing actual file operations.

## Features

### UI Components Demonstrated

1. **Fuzzy File Finder** (`p`)
   - Modal overlay with search input
   - Live filtering as you type
   - Up/Down arrow navigation through results
   - Escape to close

2. **File Explorer** (`e`)
   - Toggle sidebar with file tree
   - Table component showing file structure
   - Visual file/folder icons

3. **Settings Screen** (`s`)
   - Modal with toggle options
   - Different setting types (toggle, select, input)
   - Visual feedback for on/off states
   - Keyboard navigation

4. **Multi-tab Interface**
   - TabsComponent for open files
   - Tab key to cycle through files
   - w to close tabs

5. **Status Bar**
   - Current mode, file info, cursor position
   - Right-aligned battery/system info

## Running the Demo

```bash
cd examples/text-editor
go run .
```

## Keyboard Shortcuts

### Global
- `q` - Quit the application
- `p` - Open fuzzy file finder
- `e` - Toggle file explorer
- `s` - Open settings
- `Tab` - Cycle through open tabs
- `w` - Close current tab
- `Escape` - Close any modal

### Fuzzy Finder
- Type to filter files
- `↑/↓` - Navigate results
- `Enter` - Open selected file
- `Escape` - Cancel

### Settings
- `↑/↓` or `j/k` - Navigate options
- `Space/Enter` - Toggle option
- `Escape` - Close settings

## Implementation Details

This demo showcases how to:

1. **Compose Multiple Components** - Combines Table, TextArea, Tabs, Modal, and Input components

2. **Manage Focus States** - Properly routes keyboard input to the active component

3. **Create Custom Components** - Extends base components with additional behavior:
   ```go
   type fuzzyFinderComponent struct {
       *tui.Modal
       input    *tui.Input
       results  *tui.Table
       allFiles []string
       filtered []string
   }
   ```

4. **Handle Modal Overlays** - Settings and fuzzy finder draw over the main UI

5. **Theme Integration** - Uses Tint's theme system throughout:
   ```go
   style := lipgloss.NewStyle().
       Foreground(m.theme.Palette.Text).
       Background(m.theme.Palette.Surface)
   ```

## Components Used

- **Screen** - Main rendering surface
- **TextArea** - Main editor view
- **Table** - File explorer and search results
- **TabsComponent** - File tabs
- **Modal** - Settings and fuzzy finder containers
- **Input** - Search input field
- **Cell-based rendering** - Custom status bar

## Extending the Demo

This demo provides a foundation for building more complex UIs:

- Add more modal dialogs (find/replace, goto line)
- Implement split panes
- Add context menus
- Create custom color themes
- Add more interactive settings
- Implement command palette with actions

## Notes

- This is a UI demo only - no actual files are read or written
- All file content is hardcoded for demonstration
- The fuzzy finder uses a predefined list of files
- Settings changes are visual only and don't persist