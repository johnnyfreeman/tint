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
	Width      int  // 0 for continuation cell, 1-2 for actual character
	Foreground lipgloss.TerminalColor
	Background lipgloss.TerminalColor
	Bold       bool
	Italic     bool
	Underline  bool
	Dim        bool
	HasContent bool // true if this cell has content to draw
}

func NewCell(r rune) Cell {
	return Cell{
		Rune:       r,
		Width:      RuneWidth(r),
		Foreground: lipgloss.NoColor{},
		Background: lipgloss.NoColor{},
		HasContent: true,
	}
}

// NewContinuationCell creates a continuation cell for wide characters
func NewContinuationCell() Cell {
	return Cell{
		Rune:       0,
		Width:      0, // Continuation cell
		Foreground: lipgloss.NoColor{},
		Background: lipgloss.NoColor{},
		HasContent: true, // Continuation cells have content
	}
}

// NewTransparentCell creates a cell that preserves what's underneath
func NewTransparentCell() Cell {
	return Cell{
		Rune:       0,
		Width:      1,
		Foreground: lipgloss.NoColor{},
		Background: lipgloss.NoColor{},
		HasContent: false, // No content - preserves what's underneath
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
	// Continuation cells should not render anything
	if c.Width == 0 {
		return ""
	}
	
	// Cells without content shouldn't happen in render (they should be merged away)
	// But if they do, render as space
	if !c.HasContent || c.Rune == 0 {
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
	// If overlay has no content, preserve the original cell
	if !overlay.HasContent {
		// But still apply style attributes if they're set
		result := c
		
		// Apply background if it's set
		if _, isNoColor := overlay.Background.(lipgloss.NoColor); !isNoColor {
			result.Background = overlay.Background
		}
		
		// Apply foreground if it's set
		if _, isNoColor := overlay.Foreground.(lipgloss.NoColor); !isNoColor {
			result.Foreground = overlay.Foreground
		}
		
		// Apply style attributes
		if overlay.Bold {
			result.Bold = true
		}
		if overlay.Italic {
			result.Italic = true
		}
		if overlay.Underline {
			result.Underline = true
		}
		if overlay.Dim {
			result.Dim = true
		}
		
		return result
	}
	
	// Otherwise, overlay has content - use it but preserve background if not set
	result := overlay
	
	// If overlay has no background color, preserve the original background
	if _, isNoColor := overlay.Background.(lipgloss.NoColor); isNoColor {
		result.Background = c.Background
	}
	
	return result
}

// IsDefault checks if this is an empty/default cell
func (c Cell) IsDefault() bool {
	_, fgIsNoColor := c.Foreground.(lipgloss.NoColor)
	_, bgIsNoColor := c.Background.(lipgloss.NoColor)
	
	return c.Rune == ' ' && 
		c.Width == 1 &&
		c.HasContent &&
		fgIsNoColor && 
		bgIsNoColor && 
		!c.Bold && 
		!c.Italic && 
		!c.Underline &&
		!c.Dim
}

// IsContinuation checks if this is a continuation cell
func (c Cell) IsContinuation() bool {
	return c.Width == 0
}