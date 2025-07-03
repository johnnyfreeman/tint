package tui

import (
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type Screen struct {
	width  int
	height int
	cells  [][]Cell
}

func NewScreen(width, height int) *Screen {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
		for j := range cells[i] {
			cells[i][j] = NewCell(' ')
		}
	}
	return &Screen{
		width:  width,
		height: height,
		cells:  cells,
	}
}

func (s *Screen) SetCell(x, y int, cell Cell) {
	if x >= 0 && x < s.width && y >= 0 && y < s.height {
		s.cells[y][x] = cell
	}
}

func (s *Screen) DrawRune(x, y int, r rune, style lipgloss.Style) {
	if x >= 0 && x < s.width && y >= 0 && y < s.height {
		s.cells[y][x] = NewCell(r).WithStyle(style)
	}
}

func (s *Screen) DrawString(x, y int, str string, style lipgloss.Style) {
	for i, r := range str {
		s.DrawRune(x+i, y, r, style)
	}
}

func (s *Screen) Clear() {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			s.cells[y][x] = NewCell(' ')
		}
	}
}

func (s *Screen) ClearWithStyle(style lipgloss.Style) {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			s.cells[y][x] = NewCell(' ').WithStyle(style)
		}
	}
}

func (s *Screen) Width() int {
	return s.width
}

func (s *Screen) Height() int {
	return s.height
}

func (s *Screen) DrawBox(x, y, width, height int, style lipgloss.Style) {
	// Top border
	s.DrawRune(x, y, '┌', style)
	for i := 1; i < width-1; i++ {
		s.DrawRune(x+i, y, '─', style)
	}
	s.DrawRune(x+width-1, y, '┐', style)

	// Sides
	for i := 1; i < height-1; i++ {
		s.DrawRune(x, y+i, '│', style)
		s.DrawRune(x+width-1, y+i, '│', style)
	}

	// Bottom border
	s.DrawRune(x, y+height-1, '└', style)
	for i := 1; i < width-1; i++ {
		s.DrawRune(x+i, y+height-1, '─', style)
	}
	s.DrawRune(x+width-1, y+height-1, '┘', style)
}

func (s *Screen) DrawBoxWithTitle(x, y, width, height int, title string, borderStyle, titleStyle lipgloss.Style) {
	// Top left corner
	s.DrawRune(x, y, '┌', borderStyle)
	
	// Calculate title position (centered)
	titleWithSpaces := " " + title + " "
	titleLen := len(titleWithSpaces)
	titleStart := x + (width-titleLen)/2
	
	// Draw left border segment
	for i := x + 1; i < titleStart; i++ {
		s.DrawRune(i, y, '─', borderStyle)
	}
	
	// Draw title
	s.DrawString(titleStart, y, titleWithSpaces, titleStyle)
	
	// Draw right border segment
	for i := titleStart + titleLen; i < x+width-1; i++ {
		s.DrawRune(i, y, '─', borderStyle)
	}
	
	// Top right corner
	s.DrawRune(x+width-1, y, '┐', borderStyle)

	// Sides
	for i := 1; i < height-1; i++ {
		s.DrawRune(x, y+i, '│', borderStyle)
		s.DrawRune(x+width-1, y+i, '│', borderStyle)
	}

	// Bottom border
	s.DrawRune(x, y+height-1, '└', borderStyle)
	for i := 1; i < width-1; i++ {
		s.DrawRune(x+i, y+height-1, '─', borderStyle)
	}
	s.DrawRune(x+width-1, y+height-1, '┘', borderStyle)
}

func (s *Screen) DimArea(x, y, width, height int) {
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			cellX := x + dx
			cellY := y + dy
			if cellX >= 0 && cellX < s.width && cellY >= 0 && cellY < s.height {
				// Apply dimming by setting the Dim flag
				s.cells[cellY][cellX].Dim = true
			}
		}
	}
}

func (s *Screen) Render() string {
	var builder strings.Builder
	builder.Grow(s.width * s.height * 2) // Pre-allocate space
	
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			builder.WriteString(s.cells[y][x].Render())
		}
		if y < s.height-1 {
			builder.WriteByte('\n')
		}
	}
	return builder.String()
}

// DrawRegion draws a sub-region of one screen onto another
func (s *Screen) DrawRegion(x, y int, src *Screen, srcX, srcY, width, height int) {
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			if srcY+dy < src.height && srcX+dx < src.width {
				cell := src.cells[srcY+dy][srcX+dx]
				s.SetCell(x+dx, y+dy, cell)
			}
		}
	}
}