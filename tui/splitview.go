package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// SplitView represents a split pane container with two panels
type SplitView struct {
	left      Component
	right     Component
	vertical  bool    // true for left|right, false for top/bottom
	split     float64 // 0.0-1.0 for percentage, >1 for fixed pixels
	showBorder bool
	focused    bool
	focusSide  string // "left" or "right"
	minSize    int    // minimum size for each pane
}

// NewSplitView creates a new split view
func NewSplitView(vertical bool) *SplitView {
	return &SplitView{
		vertical:   vertical,
		split:      0.5, // 50% by default
		showBorder: true,
		focusSide:  "left",
		minSize:    5,
	}
}

// SetLeft sets the left (or top) component
func (sv *SplitView) SetLeft(component Component) {
	sv.left = component
}

// SetRight sets the right (or bottom) component
func (sv *SplitView) SetRight(component Component) {
	sv.right = component
}

// SetSplit sets the split position
// Use 0.0-1.0 for percentage, or >1 for fixed pixel size
func (sv *SplitView) SetSplit(split float64) {
	sv.split = split
}

// SetShowBorder sets whether to show a border between panes
func (sv *SplitView) SetShowBorder(show bool) {
	sv.showBorder = show
}

// SetMinSize sets the minimum size for each pane
func (sv *SplitView) SetMinSize(size int) {
	sv.minSize = size
}

// Draw renders the split view to the screen
func (sv *SplitView) Draw(screen *Screen, x, y int, theme *Theme) {
	sv.draw(screen, x, y, screen.Width()-x, screen.Height()-y, theme)
}

// DrawWithSize draws the split view with specific dimensions
func (sv *SplitView) draw(screen *Screen, x, y, width, height int, theme *Theme) {
	if width < sv.minSize*2 || height < sv.minSize*2 {
		// Too small to split
		return
	}

	// Calculate split position
	var leftWidth, leftHeight, rightX, rightY, rightWidth, rightHeight int
	borderSize := 0
	if sv.showBorder {
		borderSize = 1
	}

	if sv.vertical {
		// Vertical split (left|right)
		splitX := int(sv.split)
		if sv.split <= 1.0 {
			// Percentage
			splitX = int(float64(width) * sv.split)
		}
		
		// Ensure minimum sizes
		if splitX < sv.minSize {
			splitX = sv.minSize
		} else if splitX > width-sv.minSize-borderSize {
			splitX = width - sv.minSize - borderSize
		}

		leftWidth = splitX
		leftHeight = height
		rightX = x + splitX + borderSize
		rightY = y
		rightWidth = width - splitX - borderSize
		rightHeight = height

		// Draw border if enabled
		if sv.showBorder {
			borderStyle := lipgloss.NewStyle().
				Foreground(theme.Palette.Border).
				Background(theme.Palette.Background)
			
			for i := 0; i < height; i++ {
				screen.DrawRune(x+splitX, y+i, '│', borderStyle)
			}
		}
	} else {
		// Horizontal split (top/bottom)
		splitY := int(sv.split)
		if sv.split <= 1.0 {
			// Percentage
			splitY = int(float64(height) * sv.split)
		}
		
		// Ensure minimum sizes
		if splitY < sv.minSize {
			splitY = sv.minSize
		} else if splitY > height-sv.minSize-borderSize {
			splitY = height - sv.minSize - borderSize
		}

		leftWidth = width
		leftHeight = splitY
		rightX = x
		rightY = y + splitY + borderSize
		rightWidth = width
		rightHeight = height - splitY - borderSize

		// Draw border if enabled
		if sv.showBorder {
			borderStyle := lipgloss.NewStyle().
				Foreground(theme.Palette.Border).
				Background(theme.Palette.Background)
			
			for i := 0; i < width; i++ {
				screen.DrawRune(x+i, y+splitY, '─', borderStyle)
			}
		}
	}

	// Draw left/top component
	if sv.left != nil {
		// Create a sub-screen for clipping
		leftScreen := &Screen{
			cells:  screen.cells,
			width:  screen.width,
			height: screen.height,
		}
		
		// Draw with clipping bounds
		if drawer, ok := sv.left.(interface {
			DrawWithBounds(*Screen, int, int, int, int, *Theme)
		}); ok {
			drawer.DrawWithBounds(leftScreen, x, y, leftWidth, leftHeight, theme)
		} else {
			sv.left.Draw(leftScreen, x, y, theme)
		}
	}

	// Draw right/bottom component
	if sv.right != nil {
		// Create a sub-screen for clipping
		rightScreen := &Screen{
			cells:  screen.cells,
			width:  screen.width,
			height: screen.height,
		}
		
		// Draw with clipping bounds
		if drawer, ok := sv.right.(interface {
			DrawWithBounds(*Screen, int, int, int, int, *Theme)
		}); ok {
			drawer.DrawWithBounds(rightScreen, rightX, rightY, rightWidth, rightHeight, theme)
		} else {
			sv.right.Draw(rightScreen, rightX, rightY, theme)
		}
	}
}

// Focus gives keyboard focus to this component
func (sv *SplitView) Focus() {
	sv.focused = true
	sv.focusChild()
}

// Blur removes keyboard focus from this component
func (sv *SplitView) Blur() {
	sv.focused = false
	if sv.left != nil {
		if focusable, ok := sv.left.(Focusable); ok {
			focusable.Blur()
		}
	}
	if sv.right != nil {
		if focusable, ok := sv.right.(Focusable); ok {
			focusable.Blur()
		}
	}
}

// IsFocused returns whether this component currently has focus
func (sv *SplitView) IsFocused() bool {
	return sv.focused
}

// HandleKey processes keyboard input when focused
func (sv *SplitView) HandleKey(key string) bool {
	// Handle switching between panes
	switch key {
	case "tab":
		sv.switchFocus()
		return true
	case "left", "h":
		if sv.vertical {
			sv.focusSide = "left"
			sv.focusChild()
			return true
		}
	case "right", "l":
		if sv.vertical {
			sv.focusSide = "right"
			sv.focusChild()
			return true
		}
	case "up", "k":
		if !sv.vertical {
			sv.focusSide = "left" // top
			sv.focusChild()
			return true
		}
	case "down", "j":
		if !sv.vertical {
			sv.focusSide = "right" // bottom
			sv.focusChild()
			return true
		}
	}

	// Pass key to focused child
	if sv.focusSide == "left" && sv.left != nil {
		if handler, ok := sv.left.(interface{ HandleKey(string) bool }); ok {
			return handler.HandleKey(key)
		}
	} else if sv.focusSide == "right" && sv.right != nil {
		if handler, ok := sv.right.(interface{ HandleKey(string) bool }); ok {
			return handler.HandleKey(key)
		}
	}

	return false
}

// SetSize sets the width and height of the component
func (sv *SplitView) SetSize(width, height int) {
	// Split view manages its own sizing based on split ratio
}

// GetSize returns the current width and height
func (sv *SplitView) GetSize() (width, height int) {
	// Return 0,0 as split view doesn't have fixed size
	return 0, 0
}

// DrawWithBorder draws the component with a border and optional title
func (sv *SplitView) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	// Split view manages its own borders
	sv.Draw(screen, x, y, theme)
}

// switchFocus switches focus between panes
func (sv *SplitView) switchFocus() {
	if sv.focusSide == "left" {
		sv.focusSide = "right"
	} else {
		sv.focusSide = "left"
	}
	sv.focusChild()
}

// focusChild focuses the currently selected child
func (sv *SplitView) focusChild() {
	// Blur both first
	if sv.left != nil {
		if focusable, ok := sv.left.(Focusable); ok {
			focusable.Blur()
		}
	}
	if sv.right != nil {
		if focusable, ok := sv.right.(Focusable); ok {
			focusable.Blur()
		}
	}

	// Focus the selected one
	if sv.focusSide == "left" && sv.left != nil {
		if focusable, ok := sv.left.(Focusable); ok {
			focusable.Focus()
		}
	} else if sv.focusSide == "right" && sv.right != nil {
		if focusable, ok := sv.right.(Focusable); ok {
			focusable.Focus()
		}
	}
}

// GetFocusedSide returns which side currently has focus
func (sv *SplitView) GetFocusedSide() string {
	return sv.focusSide
}

// SetFocusedSide sets which side should have focus
func (sv *SplitView) SetFocusedSide(side string) {
	if side == "left" || side == "right" {
		sv.focusSide = side
		if sv.focused {
			sv.focusChild()
		}
	}
}