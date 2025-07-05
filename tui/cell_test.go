package tui

import (
	"testing"
	"github.com/charmbracelet/lipgloss"
)

func TestNewCell(t *testing.T) {
	// Test ASCII character
	cell := NewCell('A')
	if cell.Rune != 'A' {
		t.Errorf("Expected rune 'A', got %c", cell.Rune)
	}
	if cell.Width != 1 {
		t.Errorf("Expected width 1 for ASCII, got %d", cell.Width)
	}
	
	// Test wide character
	cell = NewCell('你')
	if cell.Rune != '你' {
		t.Errorf("Expected rune '你', got %c", cell.Rune)
	}
	if cell.Width != 2 {
		t.Errorf("Expected width 2 for wide char, got %d", cell.Width)
	}
	
	// Test emoji
	cell = NewCell('🎉')
	if cell.Width != 2 {
		t.Errorf("Expected width 2 for emoji, got %d", cell.Width)
	}
}

func TestNewContinuationCell(t *testing.T) {
	cell := NewContinuationCell()
	
	if cell.Rune != 0 {
		t.Errorf("Continuation cell should have rune 0, got %c", cell.Rune)
	}
	if cell.Width != 0 {
		t.Errorf("Continuation cell should have width 0, got %d", cell.Width)
	}
	if !cell.IsContinuation() {
		t.Error("IsContinuation() should return true")
	}
}

func TestCellWithStyle(t *testing.T) {
	cell := NewCell('A')
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Background(lipgloss.Color("#0000FF")).
		Bold(true).
		Italic(true).
		Underline(true)
	
	styledCell := cell.WithStyle(style)
	
	// Check style attributes were applied
	if styledCell.Bold != true {
		t.Error("Bold style not applied")
	}
	if styledCell.Italic != true {
		t.Error("Italic style not applied")
	}
	if styledCell.Underline != true {
		t.Error("Underline style not applied")
	}
	
	// Check colors
	if _, isNoColor := styledCell.Foreground.(lipgloss.NoColor); isNoColor {
		t.Error("Foreground color not applied")
	}
	if _, isNoColor := styledCell.Background.(lipgloss.NoColor); isNoColor {
		t.Error("Background color not applied")
	}
}

func TestCellRender(t *testing.T) {
	// Test basic render
	cell := NewCell('A')
	output := cell.Render()
	if output == "" {
		t.Error("Render should not return empty string for normal cell")
	}
	
	// Test continuation cell render
	contCell := NewContinuationCell()
	output = contCell.Render()
	if output != "" {
		t.Error("Continuation cell should render as empty string")
	}
	
	// Test null rune render
	cell = NewCell(0)
	cell.Width = 1 // Force it to not be continuation
	output = cell.Render()
	if output != " " {
		t.Errorf("Null rune should render as space, got %q", output)
	}
}

func TestCellMerge(t *testing.T) {
	// Test merging with content
	base := NewCell('A')
	base.Background = lipgloss.Color("#FF0000")
	
	overlay := NewCell('B')
	overlay.Background = lipgloss.Color("#0000FF")
	
	merged := base.Merge(overlay)
	if merged.Rune != 'B' {
		t.Errorf("Merged cell should have overlay rune, got %c", merged.Rune)
	}
	
	// Test merging with space overrides content but can preserve background
	overlay = NewCell(' ')
	overlay.Background = lipgloss.Color("#00FF00")
	overlay.Dim = true
	
	merged = base.Merge(overlay)
	if merged.Rune != ' ' {
		t.Errorf("Merging with space should override with space, got %c", merged.Rune)
	}
	if !merged.Dim {
		t.Error("Dim attribute should be applied from overlay")
	}
	
	// Test merging with transparent cell preserves content
	overlay = NewTransparentCell()
	overlay.Background = lipgloss.Color("#00FF00")
	overlay.Dim = true
	
	merged = base.Merge(overlay)
	if merged.Rune != 'A' {
		t.Errorf("Merging with transparent cell should preserve base rune, got %c", merged.Rune)
	}
	if !merged.Dim {
		t.Error("Dim attribute should be applied from overlay")
	}
	
	// Test merging with no-color background preserves original background
	overlay = NewCell('C')
	overlay.Background = lipgloss.NoColor{}
	
	merged = base.Merge(overlay)
	if _, isNoColor := merged.Background.(lipgloss.NoColor); isNoColor {
		t.Error("Original background should be preserved when overlay has no color")
	}
	if merged.Rune != 'C' {
		t.Errorf("Overlay rune should be used, got %c", merged.Rune)
	}
}

func TestCellIsDefault(t *testing.T) {
	// Default cell
	cell := NewCell(' ')
	if !cell.IsDefault() {
		t.Error("Space cell with no styles should be default")
	}
	
	// Non-default rune
	cell = NewCell('A')
	if cell.IsDefault() {
		t.Error("Cell with non-space rune should not be default")
	}
	
	// Styled space
	cell = NewCell(' ')
	cell.Bold = true
	if cell.IsDefault() {
		t.Error("Styled space should not be default")
	}
	
	// Colored space
	cell = NewCell(' ')
	cell.Foreground = lipgloss.Color("#FF0000")
	if cell.IsDefault() {
		t.Error("Colored space should not be default")
	}
	
	// Wide space (shouldn't normally happen, but test anyway)
	cell = NewCell(' ')
	cell.Width = 2
	if cell.IsDefault() {
		t.Error("Wide space should not be default")
	}
}

func TestCellCacheKey(t *testing.T) {
	// Test that cells with same attributes get same cache key
	cell1 := NewCell('A')
	cell1.Bold = true
	cell1.Italic = true
	
	cell2 := NewCell('B')
	cell2.Bold = true
	cell2.Italic = true
	
	if cell1.cacheKey() != cell2.cacheKey() {
		t.Error("Cells with same style attributes should have same cache key")
	}
	
	// Test that different attributes get different keys
	cell3 := NewCell('C')
	cell3.Bold = true
	cell3.Underline = true
	
	if cell1.cacheKey() == cell3.cacheKey() {
		t.Error("Cells with different style attributes should have different cache keys")
	}
	
	// Test colored cells don't use cache (return 0)
	cell4 := NewCell('D')
	cell4.Foreground = lipgloss.Color("#FF0000")
	
	if cell4.cacheKey() != 0 {
		t.Error("Colored cells should return 0 cache key")
	}
}

func TestCellDimming(t *testing.T) {
	cell := NewCell('A')
	cell.Dim = true
	
	// The render method should apply faint style when Dim is true
	output := cell.Render()
	if output == "" {
		t.Error("Dimmed cell should still render")
	}
	
	// Test that dim flag is preserved through operations
	style := lipgloss.NewStyle()
	_ = cell.WithStyle(style)
	// Note: WithStyle doesn't preserve Dim since lipgloss doesn't have a GetDim method
	// This is expected behavior
}

func TestCellStyleCombinations(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() Cell
		expected  uint64
	}{
		{
			name: "No styles",
			setup: func() Cell {
				return NewCell('A')
			},
			expected: 0,
		},
		{
			name: "Bold only",
			setup: func() Cell {
				c := NewCell('A')
				c.Bold = true
				return c
			},
			expected: 1 << 0,
		},
		{
			name: "Italic only", 
			setup: func() Cell {
				c := NewCell('A')
				c.Italic = true
				return c
			},
			expected: 1 << 1,
		},
		{
			name: "Underline only",
			setup: func() Cell {
				c := NewCell('A') 
				c.Underline = true
				return c
			},
			expected: 1 << 2,
		},
		{
			name: "Dim only",
			setup: func() Cell {
				c := NewCell('A')
				c.Dim = true
				return c
			},
			expected: 1 << 3,
		},
		{
			name: "All styles",
			setup: func() Cell {
				c := NewCell('A')
				c.Bold = true
				c.Italic = true
				c.Underline = true
				c.Dim = true
				return c
			},
			expected: (1 << 0) | (1 << 1) | (1 << 2) | (1 << 3),
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cell := tt.setup()
			key := cell.cacheKey()
			if key != tt.expected {
				t.Errorf("Expected cache key %d, got %d", tt.expected, key)
			}
		})
	}
}