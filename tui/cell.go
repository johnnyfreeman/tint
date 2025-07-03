package tui

import (
	"sync"
	
	"github.com/charmbracelet/lipgloss"
)

// Style cache to reduce allocations
var (
	styleCache = make(map[uint64]lipgloss.Style)
	styleMutex sync.RWMutex
)

type Cell struct {
	Rune       rune
	Foreground lipgloss.TerminalColor
	Background lipgloss.TerminalColor
	Bold       bool
	Italic     bool
	Underline  bool
	Dim        bool
}

func NewCell(r rune) Cell {
	return Cell{
		Rune:       r,
		Foreground: lipgloss.NoColor{},
		Background: lipgloss.NoColor{},
	}
}

func (c Cell) WithStyle(style lipgloss.Style) Cell {
	// Extract style attributes
	c.Foreground = style.GetForeground()
	c.Background = style.GetBackground()
	c.Bold = style.GetBold()
	c.Italic = style.GetItalic()
	c.Underline = style.GetUnderline()
	// Note: lipgloss doesn't have a direct Dim attribute, we'll handle it separately
	return c
}

// Generate a cache key for this cell's style attributes
func (c Cell) cacheKey() uint64 {
	var key uint64
	
	// Simple hash combining style attributes
	if c.Bold {
		key |= 1 << 0
	}
	if c.Italic {
		key |= 1 << 1
	}
	if c.Underline {
		key |= 1 << 2
	}
	if c.Dim {
		key |= 1 << 3
	}
	
	// For common cases (no color), use special values
	_, fgIsNoColor := c.Foreground.(lipgloss.NoColor)
	_, bgIsNoColor := c.Background.(lipgloss.NoColor)
	
	if fgIsNoColor && bgIsNoColor {
		return key // Just the attribute flags
	}
	
	// For now, don't cache colored styles (they're less common)
	return 0
}

func (c Cell) Render() string {
	if c.Rune == 0 {
		return " "
	}

	// Try to use cached style for common cases
	cacheKey := c.cacheKey()
	if cacheKey != 0 {
		styleMutex.RLock()
		if cachedStyle, ok := styleCache[cacheKey]; ok {
			styleMutex.RUnlock()
			return cachedStyle.Render(string(c.Rune))
		}
		styleMutex.RUnlock()
	}

	// Build style
	style := lipgloss.NewStyle()
	
	if _, isNoColor := c.Foreground.(lipgloss.NoColor); !isNoColor {
		style = style.Foreground(c.Foreground)
	}
	if _, isNoColor := c.Background.(lipgloss.NoColor); !isNoColor {
		style = style.Background(c.Background)
	}
	if c.Bold {
		style = style.Bold(true)
	}
	if c.Italic {
		style = style.Italic(true)
	}
	if c.Underline {
		style = style.Underline(true)
	}
	if c.Dim {
		// Apply dimming using the Faint style attribute
		style = style.Faint(true)
	}

	// Cache common styles
	if cacheKey != 0 && len(styleCache) < 100 { // Limit cache size
		styleMutex.Lock()
		styleCache[cacheKey] = style
		styleMutex.Unlock()
	}

	return style.Render(string(c.Rune))
}

// Merge overlays one cell on top of another
func (c Cell) Merge(overlay Cell) Cell {
	// If overlay has no content, return original
	if overlay.Rune == 0 || overlay.Rune == ' ' {
		// But if overlay has background color, apply it
		if _, isNoColor := overlay.Background.(lipgloss.NoColor); !isNoColor {
			c.Background = overlay.Background
		}
		// If overlay is dimmed, dim the original
		if overlay.Dim {
			c.Dim = true
		}
		return c
	}
	// Otherwise return the overlay
	return overlay
}

// IsDefault checks if this is an empty/default cell
func (c Cell) IsDefault() bool {
	_, fgIsNoColor := c.Foreground.(lipgloss.NoColor)
	_, bgIsNoColor := c.Background.(lipgloss.NoColor)
	
	return c.Rune == ' ' && 
		fgIsNoColor && 
		bgIsNoColor && 
		!c.Bold && 
		!c.Italic && 
		!c.Underline &&
		!c.Dim
}