package tui

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// Unicode text measurement and manipulation functions

// StringWidth returns the display width of a string, properly handling
// unicode characters, emojis, and CJK characters.
func StringWidth(s string) int {
	return runewidth.StringWidth(s)
}

// RuneWidth returns the display width of a single rune.
func RuneWidth(r rune) int {
	return runewidth.RuneWidth(r)
}

// Truncate truncates a string to fit within the given width, properly
// handling unicode characters. Returns the truncated string.
func Truncate(s string, width int) string {
	return runewidth.Truncate(s, width, "")
}

// TruncateWithEllipsis truncates a string to fit within the given width,
// adding an ellipsis if truncation occurs.
func TruncateWithEllipsis(s string, width int) string {
	if StringWidth(s) <= width {
		return s
	}
	if width <= 3 {
		return Truncate(s, width)
	}
	return runewidth.Truncate(s, width-3, "") + "..."
}

// FillRight pads a string with spaces on the right to reach the given width.
func FillRight(s string, width int) string {
	return runewidth.FillRight(s, width)
}

// FillLeft pads a string with spaces on the left to reach the given width.
func FillLeft(s string, width int) string {
	return runewidth.FillLeft(s, width)
}

// Wrap wraps a string to fit within the given width, returning a slice of lines.
func Wrap(s string, width int) []string {
	if width <= 0 || s == "" {
		return []string{s}
	}

	// Use runewidth's wrap function which returns a string with newlines
	wrapped := runewidth.Wrap(s, width)
	if wrapped == "" {
		return []string{}
	}

	// Split by newlines to get individual lines
	return strings.Split(wrapped, "\n")
}

// Position conversion functions

// GetVisualColumn returns the visual column position for a given byte offset in a string
func GetVisualColumn(s string, byteOffset int) int {
	if byteOffset <= 0 {
		return 0
	}
	if byteOffset >= len(s) {
		return StringWidth(s)
	}

	visualCol := 0
	currentByte := 0

	for _, r := range s {
		if currentByte >= byteOffset {
			break
		}
		visualCol += RuneWidth(r)
		currentByte += len(string(r))
	}

	return visualCol
}

// GetByteOffset returns the byte offset for a given visual column position
func GetByteOffset(s string, visualCol int) int {
	if visualCol <= 0 {
		return 0
	}

	currentCol := 0
	byteOffset := 0

	for _, r := range s {
		if currentCol >= visualCol {
			break
		}
		width := RuneWidth(r)
		if currentCol+width > visualCol {
			// We're in the middle of a wide character, return current position
			break
		}
		currentCol += width
		byteOffset += len(string(r))
	}

	return byteOffset
}

// Character boundary functions

// GetPrevCharBoundary returns the byte offset of the previous character boundary
func GetPrevCharBoundary(s string, byteOffset int) int {
	if byteOffset <= 0 {
		return 0
	}
	if byteOffset > len(s) {
		byteOffset = len(s)
	}

	// Find the start of the current rune
	for byteOffset > 0 && !isRuneStart(s[byteOffset-1]) {
		byteOffset--
	}

	// Move to previous rune
	if byteOffset > 0 {
		byteOffset--
		for byteOffset > 0 && !isRuneStart(s[byteOffset-1]) {
			byteOffset--
		}
	}

	return byteOffset
}

// GetNextCharBoundary returns the byte offset of the next character boundary
func GetNextCharBoundary(s string, byteOffset int) int {
	if byteOffset < 0 {
		return 0
	}
	if byteOffset >= len(s) {
		return len(s)
	}

	// Skip current rune
	for byteOffset < len(s) && !isRuneStart(s[byteOffset]) {
		byteOffset++
	}
	if byteOffset < len(s) {
		byteOffset++
		for byteOffset < len(s) && !isRuneStart(s[byteOffset]) {
			byteOffset++
		}
	}

	return byteOffset
}

// isRuneStart checks if a byte is the start of a UTF-8 rune
func isRuneStart(b byte) bool {
	return (b & 0xC0) != 0x80
}

// Visual string operations

// SafeSliceByVisual slices a string based on visual column positions
func SafeSliceByVisual(s string, startCol, endCol int) string {
	if startCol < 0 {
		startCol = 0
	}

	startByte := GetByteOffset(s, startCol)

	if endCol < 0 {
		return s[startByte:]
	}

	endByte := GetByteOffset(s, endCol)
	if endByte > len(s) {
		endByte = len(s)
	}

	return s[startByte:endByte]
}

// GetCharAtVisualCol returns the rune at the given visual column position
func GetCharAtVisualCol(s string, visualCol int) (rune, bool) {
	currentCol := 0

	for _, r := range s {
		width := RuneWidth(r)
		if currentCol <= visualCol && visualCol < currentCol+width {
			return r, true
		}
		currentCol += width
	}

	return ' ', false
}
