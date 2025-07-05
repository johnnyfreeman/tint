package tui

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// TableColumn defines a column in the table
type TableColumn struct {
	Title string
	Width int
}

// TableRow represents a row of data
type TableRow []string

// Table represents a simple table component
type Table struct {
	columns      []TableColumn
	rows         []TableRow
	selectedRow  int
	selectedCol  int
	focused      bool
	editable     bool
	editingCell  bool
	editValue    string
	editCursor   int
	height       int // Maximum visible rows
	scrollOffset int
}

// NewTable creates a new table
func NewTable() *Table {
	return &Table{
		columns:     []TableColumn{},
		rows:        []TableRow{},
		selectedRow: 0,
		selectedCol: 0,
		focused:     false,
		editable:    true,
		height:      5,
	}
}

// SetColumns defines the table columns
func (t *Table) SetColumns(columns []TableColumn) {
	t.columns = columns
}

// SetRows sets the table data
func (t *Table) SetRows(rows []TableRow) {
	t.rows = rows
	t.adjustSelection()
}

// AddRow adds a new row to the table
func (t *Table) AddRow(row TableRow) {
	t.rows = append(t.rows, row)
}

// RemoveRow removes a row at the given index
func (t *Table) RemoveRow(index int) {
	if index >= 0 && index < len(t.rows) {
		t.rows = append(t.rows[:index], t.rows[index+1:]...)
		t.adjustSelection()
	}
}

// GetValue returns the value at the specified row and column
func (t *Table) GetValue(row, col int) string {
	if row >= 0 && row < len(t.rows) && col >= 0 && col < len(t.rows[row]) {
		return t.rows[row][col]
	}
	return ""
}

// SetValue sets the value at the specified row and column
func (t *Table) SetValue(row, col int, value string) {
	if row >= 0 && row < len(t.rows) && col >= 0 && col < len(t.rows[row]) {
		t.rows[row][col] = value
	}
}

// Focus sets the focus state
func (t *Table) Focus() {
	t.focused = true
}

// Blur removes focus
func (t *Table) Blur() {
	t.focused = false
	t.editingCell = false
}

// IsFocused returns whether the table is focused
func (t *Table) IsFocused() bool {
	return t.focused
}

// SetHeight sets the maximum visible rows
func (t *Table) SetHeight(height int) {
	t.height = height
}

// HandleInput processes keyboard input
func (t *Table) HandleInput(key string) {
	if t.editingCell {
		t.handleEditKey(key)
		return
	}

	switch key {
	case "up", "k":
		if t.selectedRow > 0 {
			t.selectedRow--
			t.adjustScroll()
		}
	case "down", "j":
		if t.selectedRow < len(t.rows)-1 {
			t.selectedRow++
			t.adjustScroll()
		}
	case "left", "h":
		if t.selectedCol > 0 {
			t.selectedCol--
		}
	case "right", "l":
		if t.selectedCol < len(t.columns)-1 {
			t.selectedCol++
		}
	case "enter":
		if t.editable && t.selectedRow < len(t.rows) {
			t.editingCell = true
			t.editValue = t.GetValue(t.selectedRow, t.selectedCol)
			t.editCursor = len(t.editValue)
		}
	case "n":
		// Add new row
		if t.editable {
			newRow := make(TableRow, len(t.columns))
			t.AddRow(newRow)
			t.selectedRow = len(t.rows) - 1
			t.adjustScroll()
		}
	case "d":
		// Delete current row
		if t.editable && len(t.rows) > 0 {
			t.RemoveRow(t.selectedRow)
		}
	}
}

func (t *Table) handleEditKey(key string) {
	switch key {
	case "enter":
		// Save the edit
		t.SetValue(t.selectedRow, t.selectedCol, t.editValue)
		t.editingCell = false
	case "esc":
		// Cancel the edit
		t.editingCell = false
	case "left", "ctrl+b":
		t.moveCursorLeftTable()
	case "right", "ctrl+f":
		t.moveCursorRightTable()
	case "home", "ctrl+a":
		t.editCursor = 0
	case "end", "ctrl+e":
		t.editCursor = StringWidth(t.editValue)
	case "backspace", "ctrl+h":
		t.deleteBeforeCursorTable()
	case "delete", "ctrl+d":
		t.deleteAtCursorTable()
	default:
		// Handle regular character input
		if len(key) == 1 && key[0] >= 32 && key[0] < 127 {
			t.insertAtCursorTable(key)
		}
	}
}

func (t *Table) adjustSelection() {
	if t.selectedRow >= len(t.rows) {
		t.selectedRow = len(t.rows) - 1
	}
	if t.selectedRow < 0 {
		t.selectedRow = 0
	}
	if t.selectedCol >= len(t.columns) {
		t.selectedCol = len(t.columns) - 1
	}
	if t.selectedCol < 0 {
		t.selectedCol = 0
	}
	t.adjustScroll()
}

func (t *Table) adjustScroll() {
	// Scroll up if selected row is above visible area
	if t.selectedRow < t.scrollOffset {
		t.scrollOffset = t.selectedRow
	}
	// Scroll down if selected row is below visible area
	if t.selectedRow >= t.scrollOffset+t.height {
		t.scrollOffset = t.selectedRow - t.height + 1
	}
	// Don't scroll past the beginning
	if t.scrollOffset < 0 {
		t.scrollOffset = 0
	}
}

// Draw renders the table to the screen
func (t *Table) Draw(screen *Screen, x, y int, theme *Theme) {
	// Clear the entire table area with theme background
	tableWidth := t.getTableWidth()
	tableHeight := t.height + 2 // +2 for header and separator
	ClearComponentArea(screen, x, y, tableWidth, tableHeight, theme)

	// Header style
	headerStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Primary).
		Background(theme.Palette.Background).
		Bold(true)

	// Draw column headers
	currentX := x
	for colIdx, col := range t.columns {
		header := col.Title
		if StringWidth(header) > col.Width {
			header = TruncateWithEllipsis(header, col.Width)
		}
		// Pad to column width
		headerWidth := StringWidth(header)
		if headerWidth < col.Width {
			header = header + strings.Repeat(" ", col.Width-headerWidth)
		}
		screen.DrawString(currentX, y, header, headerStyle)

		// Draw bold column separator (except after last column)
		if colIdx < len(t.columns)-1 {
			separatorStyle := lipgloss.NewStyle().
				Foreground(theme.Palette.Border).
				Background(theme.Palette.Background)
			screen.DrawRune(currentX+col.Width, y, '┃', separatorStyle)
		}

		currentX += col.Width + 1 // +1 for separator
	}

	// Draw bold horizontal line under headers
	lineStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Border).
		Background(theme.Palette.Background)
	for i := 0; i < t.getTableWidth(); i++ {
		screen.DrawRune(x+i, y+1, '━', lineStyle)
	}

	// Draw rows
	visibleRows := t.height
	if len(t.rows)-t.scrollOffset < visibleRows {
		visibleRows = len(t.rows) - t.scrollOffset
	}

	for i := 0; i < visibleRows; i++ {
		rowIndex := t.scrollOffset + i
		if rowIndex >= len(t.rows) {
			break
		}

		row := t.rows[rowIndex]
		rowY := y + i + 2 // +2 for header and separator line

		currentX = x
		for colIndex, col := range t.columns {
			var cellValue string
			if colIndex < len(row) {
				cellValue = row[colIndex]
			}

			// Determine cell style
			var cellStyle lipgloss.Style
			isSelected := t.focused && rowIndex == t.selectedRow && colIndex == t.selectedCol

			if t.editingCell && isSelected {
				// Editing mode - show edit value
				cellValue = t.editValue
				cellStyle = lipgloss.NewStyle().
					Foreground(theme.Palette.Text).
					Background(theme.Palette.Surface).
					Underline(true)
			} else if isSelected {
				// Selected cell
				cellStyle = lipgloss.NewStyle().
					Foreground(theme.Components.Interactive.Selected.Text).
					Background(theme.Palette.Background).
					Bold(true)
			} else {
				// Normal cell
				cellStyle = lipgloss.NewStyle().
					Foreground(theme.Palette.Text).
					Background(theme.Palette.Background)
			}

			// Truncate if too long
			displayValue := cellValue
			if StringWidth(displayValue) > col.Width {
				displayValue = TruncateWithEllipsis(displayValue, col.Width)
			}
			// Pad to column width
			displayWidth := StringWidth(displayValue)
			if displayWidth < col.Width {
				displayValue = displayValue + strings.Repeat(" ", col.Width-displayWidth)
			}

			screen.DrawString(currentX, rowY, displayValue, cellStyle)

			// Draw cursor if editing this cell
			if t.editingCell && isSelected {
				cursorX := currentX + t.editCursor
				if cursorX < currentX+col.Width {
					cursorStyle := lipgloss.NewStyle().
						Foreground(theme.Palette.Surface).
						Background(theme.Palette.Text)

					var cursorChar rune = ' '
					if t.editCursor < len(t.editValue) {
						cursorChar = rune(t.editValue[t.editCursor])
					}
					screen.DrawRune(cursorX, rowY, cursorChar, cursorStyle)
				}
			}

			// Draw column separator (except after last column)
			if colIndex < len(t.columns)-1 {
				separatorStyle := lipgloss.NewStyle().
					Foreground(theme.Palette.Border).
					Background(theme.Palette.Background)
				screen.DrawRune(currentX+col.Width, rowY, '┃', separatorStyle)
			}

			currentX += col.Width + 1
		}
	}

	// Draw empty row indicator if no rows
	if len(t.rows) == 0 {
		emptyStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background).
			Italic(true)
		screen.DrawString(x, y+2, "No data. Press 'n' to add a row.", emptyStyle)
	}

	// Draw scroll indicator if needed
	if len(t.rows) > t.height {
		scrollStyle := lipgloss.NewStyle().
			Foreground(theme.Palette.TextMuted).
			Background(theme.Palette.Background)

		scrollText := ""
		if t.scrollOffset > 0 {
			scrollText += "↑ "
		}
		scrollText += "[" + string(rune(t.scrollOffset+1+'0')) + "/" + string(rune(len(t.rows)+'0')) + "]"
		if t.scrollOffset+t.height < len(t.rows) {
			scrollText += " ↓"
		}

		scrollX := x + t.getTableWidth() - StringWidth(scrollText)
		screen.DrawString(scrollX, y, scrollText, scrollStyle)
	}
}

func (t *Table) getTableWidth() int {
	width := 0
	for _, col := range t.columns {
		width += col.Width + 1
	}
	return width - 1 // Remove last separator
}

// DrawInBox renders the table inside a box with a title
func (t *Table) DrawInBox(screen *Screen, x, y, width, height int, title string, theme *Theme) {
	var borderColors, titleColors StateColors
	if t.focused {
		borderColors = theme.Components.Container.Border.Focused
		titleColors = theme.Components.Container.Title.Focused
	} else {
		borderColors = theme.Components.Container.Border.Unfocused
		titleColors = theme.Components.Container.Title.Unfocused
	}

	borderStyle := lipgloss.NewStyle().
		Foreground(borderColors.Border).
		Background(theme.Palette.Background)
	titleStyle := lipgloss.NewStyle().
		Foreground(titleColors.Text).
		Background(theme.Palette.Background)

	// Fill background
	bgStyle := lipgloss.NewStyle().Background(theme.Palette.Background)
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			screen.DrawRune(x+dx, y+dy, ' ', bgStyle)
		}
	}

	// Draw box with title - use heavy borders when focused
	if t.focused {
		screen.DrawBrutalistBoxWithTitle(x, y, width, height, title, borderStyle, titleStyle)
	} else {
		screen.DrawBoxWithTitle(x, y, width, height, title, borderStyle, titleStyle)
	}

	// Set table height based on box height (minus borders and header)
	t.SetHeight(height - 3) // -2 for borders, -1 for header

	// Draw the table inside the box
	t.Draw(screen, x+2, y+1, theme)
}

// HandleKey processes keyboard input when focused
func (t *Table) HandleKey(key string) bool {
	if !t.focused {
		return false
	}
	t.HandleInput(key)
	return true
}

// GetSize returns the current width and height
func (t *Table) GetSize() (width, height int) {
	totalWidth := 0
	for _, col := range t.columns {
		totalWidth += col.Width + 3 // +3 for borders and padding
	}
	return totalWidth - 1, t.height
}

// SetSize sets the width and height of the component
func (t *Table) SetSize(width, height int) {
	// Table uses column widths, so we just set height
	t.height = height
}

// DrawWithBorder draws the component with a border and optional title
func (t *Table) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	width, _ := t.GetSize()
	t.DrawInBox(screen, x, y, width+4, t.height+2, title, theme)
}
