package tui

import (
	"strings"
	"testing"
	"github.com/charmbracelet/lipgloss"
)

func BenchmarkScreenRender(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	style := lipgloss.NewStyle()
	
	// Fill screen with content
	for y := 0; y < 25; y++ {
		screen.DrawString(0, y, "This is a line of text that fills the screen width", style)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = screen.Render()
	}
}

func BenchmarkScreenDrawString(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	style := lipgloss.NewStyle()
	text := "Hello, World!"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		screen.DrawString(10, 10, text, style)
	}
}

func BenchmarkScreenDrawUnicode(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	style := lipgloss.NewStyle()
	text := "こんにちは世界🌍"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		screen.DrawString(10, 10, text, style)
	}
}

func BenchmarkCellRender(b *testing.B) {
	cell := NewCell('A').WithStyle(
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#0000FF")).
			Bold(true),
	)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cell.Render()
	}
}

func BenchmarkStringWidth(b *testing.B) {
	testStrings := []string{
		"Hello, World!",
		"你好世界",
		"🚀🌟😀 Emoji test",
		"Mixed 文字 test",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			_ = StringWidth(s)
		}
	}
}

func BenchmarkTruncate(b *testing.B) {
	longString := strings.Repeat("This is a very long string ", 10)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Truncate(longString, 50)
	}
}

func BenchmarkInputHandleKey(b *testing.B) {
	input := NewInput()
	input.SetValue("Initial text")
	
	keys := []string{"a", "left", "right", "backspace", "end", "home"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := keys[i%len(keys)]
		input.HandleInput(key)
	}
}

func BenchmarkTextAreaInsert(b *testing.B) {
	ta := NewTextArea()
	ta.SetValue(strings.Repeat("Line of text\n", 100))
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ta.HandleInput("a")
		if i%100 == 0 {
			ta.HandleInput("enter")
		}
	}
}

func BenchmarkTableDraw(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	theme := NewTestTheme()
	
	table := NewTable()
	table.SetColumns([]TableColumn{
		{Title: "ID", Width: 10},
		{Title: "Name", Width: 20},
		{Title: "Description", Width: 30},
	})
	
	// Add 20 rows
	rows := make([]TableRow, 20)
	for i := 0; i < 20; i++ {
		rows[i] = TableRow{
			string(rune('0' + i)),
			"Name " + string(rune('A' + i)),
			"Description for item " + string(rune('A' + i)),
		}
	}
	table.SetRows(rows)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table.Draw(screen, 0, 0, theme)
		screen.Clear()
	}
}

func BenchmarkContainerDraw(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	theme := NewTestTheme()
	
	container := NewContainer()
	container.SetSize(60, 20)
	container.SetTitle("Benchmark Container")
	
	input := NewInput()
	input.SetValue("Benchmark content")
	container.SetContent(input)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container.Draw(screen, 10, 2, theme)
		screen.Clear()
	}
}

func BenchmarkScreenClearArea(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	style := lipgloss.NewStyle().Background(lipgloss.Color("#FF0000"))
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ClearArea(screen, 10, 5, 40, 15, style)
	}
}

func BenchmarkWideCharacterHandling(b *testing.B) {
	screen := NewDefaultScreen(80, 25)
	style := lipgloss.NewStyle()
	
	// Mix of regular and wide characters
	texts := []string{
		"Regular text",
		"Wide 文字 text",
		"Emoji 🎉 test",
		"Mixed 你好 world 🌍",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		y := i % 25
		text := texts[i%len(texts)]
		screen.DrawString(0, y, text, style)
	}
}