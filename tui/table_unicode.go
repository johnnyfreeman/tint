package tui

// Table unicode helper methods

// moveCursorLeftTable moves the edit cursor one visual position to the left
func (t *Table) moveCursorLeftTable() {
	if t.editCursor > 0 {
		// Find the byte position of current visual column
		currentBytePos := GetByteOffset(t.editValue, t.editCursor)
		// Move to previous character
		prevBytePos := GetPrevCharBoundary(t.editValue, currentBytePos)
		// Convert back to visual column
		t.editCursor = GetVisualColumn(t.editValue, prevBytePos)
	}
}

// moveCursorRightTable moves the edit cursor one visual position to the right
func (t *Table) moveCursorRightTable() {
	valueWidth := StringWidth(t.editValue)
	if t.editCursor < valueWidth {
		// Find the byte position of current visual column
		currentBytePos := GetByteOffset(t.editValue, t.editCursor)
		// Move to next character
		nextBytePos := GetNextCharBoundary(t.editValue, currentBytePos)
		// Convert back to visual column
		t.editCursor = GetVisualColumn(t.editValue, nextBytePos)
	}
}

// insertAtCursorTable inserts text at the current edit cursor position
func (t *Table) insertAtCursorTable(text string) {
	bytePos := GetByteOffset(t.editValue, t.editCursor)
	
	// Insert text at byte position
	t.editValue = t.editValue[:bytePos] + text + t.editValue[bytePos:]
	
	// Move cursor by the visual width of inserted text
	t.editCursor += StringWidth(text)
}

// deleteBeforeCursorTable deletes one character before the edit cursor
func (t *Table) deleteBeforeCursorTable() {
	if t.editCursor > 0 {
		currentBytePos := GetByteOffset(t.editValue, t.editCursor)
		prevBytePos := GetPrevCharBoundary(t.editValue, currentBytePos)
		
		// Delete the character
		t.editValue = t.editValue[:prevBytePos] + t.editValue[currentBytePos:]
		
		// Update cursor position
		t.editCursor = GetVisualColumn(t.editValue, prevBytePos)
	}
}

// deleteAtCursorTable deletes one character at the edit cursor position
func (t *Table) deleteAtCursorTable() {
	valueWidth := StringWidth(t.editValue)
	
	if t.editCursor < valueWidth {
		currentBytePos := GetByteOffset(t.editValue, t.editCursor)
		nextBytePos := GetNextCharBoundary(t.editValue, currentBytePos)
		
		// Delete the character
		t.editValue = t.editValue[:currentBytePos] + t.editValue[nextBytePos:]
	}
}