package tui

// TextArea unicode helper methods

// moveCursorLeft moves the cursor one visual position to the left
func (t *TextArea) moveCursorLeft() {
	if t.cursorCol > 0 {
		// Find the byte position of current visual column
		line := t.lines[t.cursorRow]
		currentBytePos := GetByteOffset(line, t.cursorCol)
		// Move to previous character
		prevBytePos := GetPrevCharBoundary(line, currentBytePos)
		// Convert back to visual column
		t.cursorCol = GetVisualColumn(line, prevBytePos)
	} else if t.cursorRow > 0 {
		// Move to end of previous line
		t.cursorRow--
		t.cursorCol = StringWidth(t.lines[t.cursorRow])
	}
}

// moveCursorRight moves the cursor one visual position to the right
func (t *TextArea) moveCursorRight() {
	line := t.lines[t.cursorRow]
	lineWidth := StringWidth(line)

	if t.cursorCol < lineWidth {
		// Find the byte position of current visual column
		currentBytePos := GetByteOffset(line, t.cursorCol)
		// Move to next character
		nextBytePos := GetNextCharBoundary(line, currentBytePos)
		// Convert back to visual column
		t.cursorCol = GetVisualColumn(line, nextBytePos)
	} else if t.cursorRow < len(t.lines)-1 {
		// Move to beginning of next line
		t.cursorRow++
		t.cursorCol = 0
	}
}

// insertAtCursor inserts text at the current cursor position
func (t *TextArea) insertAtCursor(text string) {
	line := t.lines[t.cursorRow]
	bytePos := GetByteOffset(line, t.cursorCol)

	// Insert text at byte position
	newLine := line[:bytePos] + text + line[bytePos:]
	t.lines[t.cursorRow] = newLine

	// Move cursor by the visual width of inserted text
	t.cursorCol += StringWidth(text)
}

// deleteBeforeCursor deletes one character before the cursor
func (t *TextArea) deleteBeforeCursor() {
	if t.cursorCol > 0 {
		line := t.lines[t.cursorRow]
		currentBytePos := GetByteOffset(line, t.cursorCol)
		prevBytePos := GetPrevCharBoundary(line, currentBytePos)

		// Delete the character
		t.lines[t.cursorRow] = line[:prevBytePos] + line[currentBytePos:]

		// Update cursor position
		t.cursorCol = GetVisualColumn(line, prevBytePos)
	} else if t.cursorRow > 0 {
		// Join with previous line
		prevLine := t.lines[t.cursorRow-1]
		currentLine := t.lines[t.cursorRow]
		t.cursorCol = StringWidth(prevLine)
		t.lines[t.cursorRow-1] = prevLine + currentLine
		// Remove current line
		t.lines = append(t.lines[:t.cursorRow], t.lines[t.cursorRow+1:]...)
		t.cursorRow--
	}
}

// deleteAtCursor deletes one character at the cursor position
func (t *TextArea) deleteAtCursor() {
	line := t.lines[t.cursorRow]
	lineWidth := StringWidth(line)

	if t.cursorCol < lineWidth {
		currentBytePos := GetByteOffset(line, t.cursorCol)
		nextBytePos := GetNextCharBoundary(line, currentBytePos)

		// Delete the character
		t.lines[t.cursorRow] = line[:currentBytePos] + line[nextBytePos:]
	} else if t.cursorRow < len(t.lines)-1 {
		// Join with next line
		t.lines[t.cursorRow] = line + t.lines[t.cursorRow+1]
		// Remove next line
		t.lines = append(t.lines[:t.cursorRow+1], t.lines[t.cursorRow+2:]...)
	}
}

// splitLineAtCursor splits the current line at cursor position
func (t *TextArea) splitLineAtCursor() {
	line := t.lines[t.cursorRow]
	bytePos := GetByteOffset(line, t.cursorCol)

	before := line[:bytePos]
	after := line[bytePos:]

	// Update current line and insert new line
	t.lines[t.cursorRow] = before
	newLines := append(t.lines[:t.cursorRow+1], append([]string{after}, t.lines[t.cursorRow+1:]...)...)
	t.lines = newLines

	// Move cursor to beginning of new line
	t.cursorRow++
	t.cursorCol = 0
}

// getVisibleLine returns the visible portion of a line based on scroll offset
func (t *TextArea) getVisibleLine(line string) string {
	if t.offsetCol >= StringWidth(line) {
		return ""
	}

	startCol := t.offsetCol
	endCol := t.offsetCol + t.width

	return SafeSliceByVisual(line, startCol, endCol)
}

// getCursorScreenCol returns the screen column position of the cursor
func (t *TextArea) getCursorScreenCol() int {
	return t.cursorCol - t.offsetCol
}
