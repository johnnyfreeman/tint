package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/example/tint/tui"
)

func main() {
	// Initialize theme
	theme := tui.GetTheme("monochrome")

	// Create screen
	screen := tui.NewScreen(80, 30)

	// Test strings with various unicode characters
	testStrings := []struct {
		name string
		text string
	}{
		{"ASCII", "Hello World!"},
		{"Emoji", "🚀 Rocket 🌟 Star 😀 Face"},
		{"CJK", "你好世界 こんにちは 안녕하세요"},
		{"Mixed", "Test 测试 🎉 Party!"},
		{"Flags", "🇺🇸 🇬🇧 🇯🇵 🇰🇷 🇨🇳"},
		{"Complex", "👨‍👩‍👧‍👦 Family 👩‍💻 Coder"},
		{"Math", "∑ ∏ ∫ ∂ ∇ √ ∞"},
		{"Diacritics", "café naïve résumé"},
	}

	// Clear screen
	fmt.Print("\033[2J\033[H")

	// Draw title
	titleStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Primary).
		Bold(true)
	screen.DrawString(25, 1, "Unicode Width Test", titleStyle)

	// Create text style once
	textStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Text)

	// Draw test strings
	y := 3
	for _, test := range testStrings {
		// Draw label
		labelStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted)
		screen.DrawString(2, y, fmt.Sprintf("%-12s:", test.name), labelStyle)

		// Draw test string
		screen.DrawString(16, y, test.text, textStyle)

		// Draw width info
		width := tui.StringWidth(test.text)
		widthStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.Gold)
		screen.DrawString(60, y, fmt.Sprintf("Width: %d", width), widthStyle)

		y++
	}

	// Draw a box with unicode content
	y += 2
	boxStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Primary)
	titleStyle = lipgloss.NewStyle().
		Foreground(theme.Palette.Text).
		Bold(true)
	screen.DrawBoxWithTitle(2, y, 40, 8, "Unicode Box 📦", boxStyle, titleStyle)

	// Test overwriting wide characters
	y += 2
	screen.DrawString(4, y, "Overwrite test:", textStyle)
	y++
	screen.DrawString(4, y, "Original: 你好世界", textStyle)
	y++
	// Overwrite middle of wide character
	screen.DrawString(4, y, "你好世界", textStyle)
	screen.DrawString(6, y, "XX", lipgloss.NewStyle().Foreground(theme.Palette.Love))

	// Create interactive components
	y = 16
	input := tui.NewInput()
	input.SetWidth(30)
	input.SetPlaceholder("Type emojis here... 🌟")
	input.Focus()
	input.DrawInBox(screen, 45, y, "Unicode Input", &theme)

	// Render initial screen
	fmt.Print(screen.Render())

	// Display for 2 seconds then exit
	time.Sleep(2 * time.Second)
	
	// Clear screen and exit
	fmt.Print("\033[2J\033[H")
	fmt.Println("Unicode test completed successfully!")
}