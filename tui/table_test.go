package tui

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable()

	if len(table.columns) != 0 {
		t.Error("New table should have no columns")
	}
	if len(table.rows) != 0 {
		t.Error("New table should have no rows")
	}
	if table.selectedRow != 0 || table.selectedCol != 0 {
		t.Error("New table should have selection at (0,0)")
	}
	if !table.editable {
		t.Error("New table should be editable by default")
	}
	if table.IsFocused() {
		t.Error("New table should not be focused")
	}
}

func TestTableSetData(t *testing.T) {
	table := NewTable()

	// Set columns
	columns := []TableColumn{
		{Title: "Name", Width: 10},
		{Title: "Age", Width: 5},
		{Title: "City", Width: 15},
	}
	table.SetColumns(columns)

	if len(table.columns) != 3 {
		t.Errorf("Expected 3 columns, got %d", len(table.columns))
	}

	// Set rows
	rows := []TableRow{
		{"Alice", "30", "New York"},
		{"Bob", "25", "London"},
		{"Charlie", "35", "Tokyo"},
	}
	table.SetRows(rows)

	if len(table.rows) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(table.rows))
	}

	// Add a row
	table.AddRow(TableRow{"David", "28", "Paris"})
	if len(table.rows) != 4 {
		t.Errorf("Expected 4 rows after adding, got %d", len(table.rows))
	}
}

func TestTableNavigation(t *testing.T) {
	table := setupTestTable()

	// Test arrow key navigation
	table.HandleInput("down")
	if table.selectedRow != 1 {
		t.Errorf("Expected row 1 selected, got %d", table.selectedRow)
	}

	table.HandleInput("right")
	if table.selectedCol != 1 {
		t.Errorf("Expected col 1 selected, got %d", table.selectedCol)
	}

	// Test boundary conditions
	table.selectedRow = 2
	table.HandleInput("down") // Should not go past last row
	if table.selectedRow != 2 {
		t.Error("Should not navigate past last row")
	}

	table.selectedCol = 2
	table.HandleInput("right") // Should not go past last column
	if table.selectedCol != 2 {
		t.Error("Should not navigate past last column")
	}
}

func TestTableEditing(t *testing.T) {
	table := setupTestTable()
	table.Focus()

	// Start editing
	table.HandleInput("enter")
	if !table.editingCell {
		t.Error("Should be editing after Enter")
	}

	// Edit the cell
	table.HandleInput("backspace") // Remove last char
	table.HandleInput("backspace")
	table.HandleInput("backspace")
	table.HandleInput("backspace")
	table.HandleInput("backspace") // Clear "Alice"
	table.HandleInput("E")
	table.HandleInput("v")
	table.HandleInput("e")

	// Confirm edit
	table.HandleInput("enter")
	if table.editingCell {
		t.Error("Should not be editing after confirming")
	}

	if table.rows[0][0] != "Eve" {
		t.Errorf("Expected cell to be 'Eve', got %q", table.rows[0][0])
	}

	// Test cancel edit
	table.HandleInput("enter") // Start editing again
	table.HandleInput("X")
	table.HandleInput("escape") // Cancel

	if table.rows[0][0] != "Eve" {
		t.Error("Cell should not change when edit is cancelled")
	}
}

func TestTableDraw(t *testing.T) {
	screen := NewScreenSimulation(50, 15)
	theme := NewTestTheme()
	table := setupTestTable()

	table.Draw(screen.Screen, 0, 0, 50, 15, theme)

	// Check headers are drawn
	AssertTextExists(t, screen, "Name")
	AssertTextExists(t, screen, "Age")
	AssertTextExists(t, screen, "City")

	// Check data is drawn
	AssertTextExists(t, screen, "Alice")
	AssertTextExists(t, screen, "Bob")
	AssertTextExists(t, screen, "New York")
}

func TestTableScrolling(t *testing.T) {
	table := NewTable()
	table.SetHeight(3) // Small height to force scrolling

	// Add many rows
	columns := []TableColumn{{Title: "Data", Width: 10}}
	table.SetColumns(columns)

	rows := make([]TableRow, 10)
	for i := 0; i < 10; i++ {
		rows[i] = TableRow{string(rune('A' + i))}
	}
	table.SetRows(rows)

	// Navigate down past visible area
	for i := 0; i < 5; i++ {
		table.HandleInput("down")
	}

	// Should have scrolled
	if table.scrollOffset == 0 {
		t.Error("Expected table to scroll")
	}
}

func TestTableFocus(t *testing.T) {
	table := NewTable()

	// Test Focus
	table.Focus()
	if !table.IsFocused() {
		t.Error("Table should be focused after Focus()")
	}

	// Test Blur
	table.Blur()
	if table.IsFocused() {
		t.Error("Table should not be focused after Blur()")
	}
}

func TestTableEditable(t *testing.T) {
	table := setupTestTable()
	table.Focus()

	// Test with editable = true (default)
	table.HandleInput("enter")
	if !table.editingCell {
		t.Error("Should be able to edit when editable=true")
	}
	table.HandleInput("escape") // Cancel edit

	// Test with editable = false
	table.editable = false
	table.HandleInput("enter")
	if table.editingCell {
		t.Error("Should not be able to edit when editable=false")
	}
}

func TestTableUnicodeContent(t *testing.T) {
	table := NewTable()
	columns := []TableColumn{
		{Title: "名前", Width: 10},
		{Title: "年齢", Width: 5},
	}
	table.SetColumns(columns)

	rows := []TableRow{
		{"田中", "30"},
		{"佐藤", "25"},
	}
	table.SetRows(rows)

	screen := NewScreenSimulation(30, 10)
	theme := NewTestTheme()
	table.Draw(screen.Screen, 0, 0, 30, 10, theme)

	// Check unicode content is drawn
	AssertTextExists(t, screen, "名前")
	AssertTextExists(t, screen, "田中")
}

func TestTableSetValue(t *testing.T) {
	table := setupTestTable()

	// Update a specific cell
	table.SetValue(1, 2, "Berlin")

	if table.rows[1][2] != "Berlin" {
		t.Errorf("Expected cell to be 'Berlin', got %q", table.rows[1][2])
	}

	// Test out of bounds update (should not crash)
	table.SetValue(10, 0, "Invalid") // Row out of bounds
	table.SetValue(0, 10, "Invalid") // Col out of bounds
}

func TestTableGetValue(t *testing.T) {
	table := setupTestTable()

	// Get value at specific position
	value := table.GetValue(0, 0)
	if value != "Alice" {
		t.Errorf("Expected value 'Alice', got %q", value)
	}

	// Get another value
	value = table.GetValue(1, 1)
	if value != "25" {
		t.Errorf("Expected value '25', got %q", value)
	}

	// Test out of bounds
	value = table.GetValue(10, 0)
	if value != "" {
		t.Error("Out of bounds should return empty string")
	}
}

// Helper function to create a test table
func setupTestTable() *Table {
	table := NewTable()

	columns := []TableColumn{
		{Title: "Name", Width: 10},
		{Title: "Age", Width: 5},
		{Title: "City", Width: 15},
	}
	table.SetColumns(columns)

	rows := []TableRow{
		{"Alice", "30", "New York"},
		{"Bob", "25", "London"},
		{"Charlie", "35", "Tokyo"},
	}
	table.SetRows(rows)

	return table
}
