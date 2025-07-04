package tui

import (
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type Screen struct {
	width  int
	height int
	cells  [][]Cell
	theme  Theme
}

func NewScreen(width, height int, theme Theme) *Screen {
	s := &Screen{
		width:  width,
		height: height,
		theme:  theme,
	}
	s.cells = make([][]Cell, height)
	for i := range s.cells {
		s.cells[i] = make([]Cell, width)
	}
	s.Clear() // Clear with theme background
	return s
}

// NewDefaultScreen creates a new screen with the default theme
func NewDefaultScreen(width, height int) *Screen {
	return NewScreen(width, height, DefaultTheme)
}

func (s *Screen) SetCell(x, y int, cell Cell) {
	if x >= 0 && x < s.width && y >= 0 && y < s.height {
		// Check if we're overwriting a continuation cell
		if x > 0 && s.cells[y][x].IsContinuation() {
			// Find the start of the wide character and clear it
			for i := x - 1; i >= 0 && i > x - 3; i-- {
				if !s.cells[y][i].IsContinuation() {
					// Found the wide character, clear it
					s.cells[y][i] = NewCell(' ')
					break
				}
			}
		}
		
		// Check if we're overwriting the start of a wide character
		if x+1 < s.width && s.cells[y][x+1].IsContinuation() {
			// Clear continuation cells
			for i := x + 1; i < s.width && s.cells[y][i].IsContinuation(); i++ {
				s.cells[y][i] = NewCell(' ')
			}
		}
		
		s.cells[y][x] = cell
		
		// For wide characters, set continuation cells
		if cell.Width > 1 {
			continuationCell := NewContinuationCell()
			continuationCell.Foreground = cell.Foreground
			continuationCell.Background = cell.Background
			continuationCell.Bold = cell.Bold
			continuationCell.Italic = cell.Italic
			continuationCell.Underline = cell.Underline
			continuationCell.Dim = cell.Dim
			
			for i := 1; i < cell.Width && x+i < s.width; i++ {
				s.cells[y][x+i] = continuationCell
			}
		}
	}
}

func (s *Screen) DrawRune(x, y int, r rune, style lipgloss.Style) {
	if x >= 0 && x < s.width && y >= 0 && y < s.height {
		cell := NewCell(r).WithStyle(style)
		s.SetCell(x, y, cell)
	}
}

func (s *Screen) DrawString(x, y int, str string, style lipgloss.Style) {
	xOffset := 0
	for _, r := range str {
		s.DrawRune(x+xOffset, y, r, style)
		xOffset += RuneWidth(r)
	}
}

func (s *Screen) Clear() {
	style := lipgloss.NewStyle().
		Background(s.theme.Palette.Background).
		Foreground(s.theme.Palette.Text)
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

func (s *Screen) Theme() Theme {
	return s.theme
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
	titleLen := StringWidth(titleWithSpaces)
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

// DrawBrutalistBox draws a box with heavy borders
func (s *Screen) DrawBrutalistBox(x, y, width, height int, style lipgloss.Style) {
	// Top border (heavy)
	s.DrawRune(x, y, '┏', style)
	for i := 1; i < width-1; i++ {
		s.DrawRune(x+i, y, '━', style)
	}
	s.DrawRune(x+width-1, y, '┓', style)

	// Sides (heavy)
	for i := 1; i < height-1; i++ {
		s.DrawRune(x, y+i, '┃', style)
		s.DrawRune(x+width-1, y+i, '┃', style)
	}

	// Bottom border (heavy)
	s.DrawRune(x, y+height-1, '┗', style)
	for i := 1; i < width-1; i++ {
		s.DrawRune(x+i, y+height-1, '━', style)
	}
	s.DrawRune(x+width-1, y+height-1, '┛', style)
}

// DrawBrutalistBoxWithTitle draws a brutalist box with a title
func (s *Screen) DrawBrutalistBoxWithTitle(x, y, width, height int, title string, borderStyle, titleStyle lipgloss.Style) {
	// Top left corner
	s.DrawRune(x, y, '┏', borderStyle)
	
	// Calculate title position (centered)
	titleWithSpaces := " " + title + " "
	titleLen := StringWidth(titleWithSpaces)
	titleStart := x + (width-titleLen)/2
	
	// Draw left border segment
	for i := x + 1; i < titleStart; i++ {
		s.DrawRune(i, y, '━', borderStyle)
	}
	
	// Draw title
	s.DrawString(titleStart, y, titleWithSpaces, titleStyle)
	
	// Draw right border segment
	for i := titleStart + titleLen; i < x+width-1; i++ {
		s.DrawRune(i, y, '━', borderStyle)
	}
	
	// Top right corner
	s.DrawRune(x+width-1, y, '┓', borderStyle)

	// Sides (heavy)
	for i := 1; i < height-1; i++ {
		s.DrawRune(x, y+i, '┃', borderStyle)
		s.DrawRune(x+width-1, y+i, '┃', borderStyle)
	}

	// Bottom border (heavy)
	s.DrawRune(x, y+height-1, '┗', borderStyle)
	for i := 1; i < width-1; i++ {
		s.DrawRune(x+i, y+height-1, '━', borderStyle)
	}
	s.DrawRune(x+width-1, y+height-1, '┛', borderStyle)
}

// DrawBlockShadow draws a solid block shadow for a given area
func (s *Screen) DrawBlockShadow(x, y, width, height int, shadowStyle lipgloss.Style, offsetX, offsetY int) {
	// Draw shadow using solid blocks
	
	// Draw bottom shadow (full width, offset down)
	for i := 0; i < width; i++ {
		for j := 0; j < offsetY; j++ {
			s.DrawRune(x+offsetX+i, y+height+j, '█', shadowStyle)
		}
	}
	
	// Draw right shadow (full height, offset right)
	for i := 0; i < height; i++ {
		for j := 0; j < offsetX; j++ {
			s.DrawRune(x+width+j, y+offsetY+i, '█', shadowStyle)
		}
	}
	
	// Draw corner shadow (fills the gap)
	for i := 0; i < offsetX; i++ {
		for j := 0; j < offsetY; j++ {
			s.DrawRune(x+width+i, y+height+j, '█', shadowStyle)
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