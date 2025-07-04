package tui

import (
	"testing"
	
	"github.com/charmbracelet/lipgloss"
)

func TestStringWidth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Empty string", "", 0},
		{"ASCII", "Hello", 5},
		{"ASCII with spaces", "Hello World", 11},
		{"Single emoji", "😀", 2},
		{"Multiple emojis", "🚀🌟😀", 6},
		{"Chinese", "你好", 4},
		{"Japanese", "こんにちは", 10},
		{"Korean", "안녕하세요", 10},
		{"Mixed ASCII and CJK", "Hello你好", 9},
		{"Mixed with emoji", "Test🎉Party", 11},
		{"Zero-width joiner", "👨‍👩‍👧‍👦", 2}, // Family emoji
		{"Combining marks", "café", 4},
		{"Tab character", "a\tb", 2}, // runewidth counts tab as 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringWidth(tt.input)
			if result != tt.expected {
				t.Errorf("StringWidth(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		width    int
		expected string
	}{
		{"No truncation needed", "Hello", 10, "Hello"},
		{"Truncate ASCII", "Hello World", 5, "Hello"},
		{"Truncate at emoji boundary", "Hi😀Test", 4, "Hi😀"},
		{"Truncate before emoji", "Hi😀Test", 3, "Hi"},
		{"Truncate CJK", "你好世界", 6, "你好世"},
		{"Truncate mixed", "Test测试", 6, "Test测"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.input, tt.width)
			if result != tt.expected {
				t.Errorf("Truncate(%q, %d) = %q, want %q", tt.input, tt.width, result, tt.expected)
			}
		})
	}
}

func TestGetVisualColumn(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		byteOffset int
		expected   int
	}{
		{"Start of string", "Hello", 0, 0},
		{"Middle of ASCII", "Hello", 3, 3},
		{"After emoji", "😀Hi", 4, 2}, // Emoji is 4 bytes
		{"After CJK", "你好", 3, 2},    // First Chinese char is 3 bytes
		{"End of mixed", "Hi你", 5, 4}, // "Hi" (2) + "你" (2 width, 3 bytes)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetVisualColumn(tt.input, tt.byteOffset)
			if result != tt.expected {
				t.Errorf("GetVisualColumn(%q, %d) = %d, want %d", tt.input, tt.byteOffset, result, tt.expected)
			}
		})
	}
}

func TestGetByteOffset(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		visualCol int
		expected  int
	}{
		{"Start of string", "Hello", 0, 0},
		{"Middle of ASCII", "Hello", 3, 3},
		{"After emoji width", "😀Hi", 2, 4}, // Skip 4-byte emoji
		{"After CJK width", "你好", 2, 3},    // Skip 3-byte character
		{"Middle of wide char", "你好", 1, 0}, // Can't position in middle
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetByteOffset(tt.input, tt.visualCol)
			if result != tt.expected {
				t.Errorf("GetByteOffset(%q, %d) = %d, want %d", tt.input, tt.visualCol, result, tt.expected)
			}
		})
	}
}

func TestCellContinuation(t *testing.T) {
	// Test that wide characters create continuation cells
	screen := NewDefaultScreen(10, 1)
	
	// Draw a wide character
	style := lipgloss.NewStyle()
	screen.DrawString(0, 0, "你", style)
	
	// Check that first cell has the character
	if screen.cells[0][0].Rune != '你' {
		t.Errorf("Expected '你' at position 0, got %c", screen.cells[0][0].Rune)
	}
	if screen.cells[0][0].Width != 2 {
		t.Errorf("Expected width 2 at position 0, got %d", screen.cells[0][0].Width)
	}
	
	// Check that second cell is a continuation
	if !screen.cells[0][1].IsContinuation() {
		t.Error("Expected continuation cell at position 1")
	}
	
	// Test overwriting part of a wide character
	screen.DrawString(1, 0, "X", style)
	
	// Original wide character should be cleared
	if screen.cells[0][0].Rune != ' ' {
		t.Errorf("Expected space at position 0 after overwrite, got %c", screen.cells[0][0].Rune)
	}
	// New character should be at position 1
	if screen.cells[0][1].Rune != 'X' {
		t.Errorf("Expected 'X' at position 1, got %c", screen.cells[0][1].Rune)
	}
}