# Unicode UI Elements Reference

Now that the TUI library has full unicode support, here are some nice characters that can be used for UI elements:

## Selection Indicators
- `◉` / `○` - Radio buttons (filled/empty)
- `▶` / `▷` - Arrows (filled/empty)
- `■` / `□` - Checkboxes (checked/unchecked)
- `✓` / `✗` - Checkmarks and crosses
- `●` / `○` - Bullets (filled/empty)
- `▸` / `▹` - Small arrows
- `►` / `▻` - Medium arrows

## Status Indicators
- `✓` - Success/Done
- `✗` - Error/Failed  
- `⚠` - Warning
- `ℹ` - Info
- `⏳` - Loading/In Progress
- `🔄` - Refresh/Sync
- `⚡` - Fast/Quick
- `🔒` / `🔓` - Locked/Unlocked

## Decorative Elements
- `─` `━` `│` `┃` - Box drawing (thin/thick)
- `┌` `┐` `└` `┘` - Corners (thin)
- `┏` `┓` `┗` `┛` - Corners (thick)
- `├` `┤` `┬` `┴` - Connectors
- `═` `║` `╔` `╗` `╚` `╝` - Double lines

## Progress Bars
- `▏` `▎` `▍` `▌` `▋` `▊` `▉` `█` - Block elements
- `⣀` `⣄` `⣤` `⣦` `⣶` `⣷` `⣿` - Braille patterns
- `◐` `◓` `◑` `◒` - Spinning indicators

## Arrows
- `←` `→` `↑` `↓` - Basic arrows
- `⇐` `⇒` `⇑` `⇓` - Double arrows
- `↖` `↗` `↘` `↙` - Diagonal arrows
- `⟵` `⟶` `⟷` - Long arrows

## Examples in Use

### Theme Picker (Updated)
```
┌────────── 🎨 Choose Theme ───────────┐
│ ○ Tokyo Night                        │
│ ◉ Rosé Pine         (selected)       │
│ ○ Catppuccin                         │
└──────────────────────────────────────┘
```

### Sidebar (Updated)
```
┌─── Navigation ───┐
│ ▶ Dashboard      │
│   Settings       │
│   Profile        │
└──────────────────┘
```

### Status Messages
```
✓ Operation completed successfully
⚠ Warning: Low disk space
✗ Error: Connection failed
ℹ Info: Update available
```

All these characters now render properly with correct width calculations!