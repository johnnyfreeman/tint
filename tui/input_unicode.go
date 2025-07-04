package tui

// Input unicode helper methods

// moveCursorLeftInput moves the cursor one visual position to the left
func (i *Input) moveCursorLeft() {
	if i.cursor > 0 {
		// Find the byte position of current visual column
		currentBytePos := GetByteOffset(i.value, i.cursor)
		// Move to previous character
		prevBytePos := GetPrevCharBoundary(i.value, currentBytePos)
		// Convert back to visual column
		i.cursor = GetVisualColumn(i.value, prevBytePos)
	}
}

// moveCursorRightInput moves the cursor one visual position to the right
func (i *Input) moveCursorRight() {
	valueWidth := StringWidth(i.value)
	if i.cursor < valueWidth {
		// Find the byte position of current visual column
		currentBytePos := GetByteOffset(i.value, i.cursor)
		// Move to next character
		nextBytePos := GetNextCharBoundary(i.value, currentBytePos)
		// Convert back to visual column
		i.cursor = GetVisualColumn(i.value, nextBytePos)
	}
}

// insertAtCursorInput inserts text at the current cursor position
func (i *Input) insertAtCursorInput(text string) {
	bytePos := GetByteOffset(i.value, i.cursor)
	
	// Insert text at byte position
	i.value = i.value[:bytePos] + text + i.value[bytePos:]
	
	// Move cursor by the visual width of inserted text
	i.cursor += StringWidth(text)
}

// deleteBeforeCursorInput deletes one character before the cursor
func (i *Input) deleteBeforeCursorInput() {
	if i.cursor > 0 {
		currentBytePos := GetByteOffset(i.value, i.cursor)
		prevBytePos := GetPrevCharBoundary(i.value, currentBytePos)
		
		// Delete the character
		i.value = i.value[:prevBytePos] + i.value[currentBytePos:]
		
		// Update cursor position
		i.cursor = GetVisualColumn(i.value, prevBytePos)
	}
}

// deleteAtCursorInput deletes one character at the cursor position
func (i *Input) deleteAtCursorInput() {
	valueWidth := StringWidth(i.value)
	
	if i.cursor < valueWidth {
		currentBytePos := GetByteOffset(i.value, i.cursor)
		nextBytePos := GetNextCharBoundary(i.value, currentBytePos)
		
		// Delete the character
		i.value = i.value[:currentBytePos] + i.value[nextBytePos:]
	}
}

// killToEndOfLine deletes from cursor to end of line
func (i *Input) killToEndOfLine() {
	bytePos := GetByteOffset(i.value, i.cursor)
	i.value = i.value[:bytePos]
}

// killToBeginningOfLine deletes from beginning to cursor
func (i *Input) killToBeginningOfLine() {
	bytePos := GetByteOffset(i.value, i.cursor)
	i.value = i.value[bytePos:]
	i.cursor = 0
}

// getVisibleValuePortion returns the visible portion of the value based on offset
func (i *Input) getVisibleValuePortion() string {
	valueWidth := StringWidth(i.value)
	if i.offset >= valueWidth {
		return ""
	}
	
	startCol := i.offset
	endCol := i.offset + i.width
	
	return SafeSliceByVisual(i.value, startCol, endCol)
}