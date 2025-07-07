package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// NotificationType represents the type of notification
type NotificationType int

const (
	NotificationSuccess NotificationType = iota
	NotificationWarning
	NotificationError
	NotificationInfo
)

// NotificationPosition represents where to show the notification
type NotificationPosition int

const (
	NotificationBottomRight NotificationPosition = iota
	NotificationBottomLeft
	NotificationTopRight
	NotificationTopLeft
)

// Notification represents a toast-style notification component
type Notification struct {
	visible   bool
	message   string
	notifType NotificationType
	timestamp time.Time
	duration  time.Duration
	width     int
	height    int
	position  NotificationPosition
	focused   bool
}

// NewNotification creates a new notification
func NewNotification() *Notification {
	return &Notification{
		visible:   false,
		notifType: NotificationInfo,
		duration:  3 * time.Second,
		width:     30,
		height:    4,
		position:  NotificationBottomRight,
		focused:   false,
	}
}

// SetPosition sets where the notification appears
func (n *Notification) SetPosition(pos NotificationPosition) {
	n.position = pos
}

// SetDuration sets how long the notification stays visible
func (n *Notification) SetDuration(duration time.Duration) {
	n.duration = duration
}

// Show displays a notification with the given message and type
func (n *Notification) Show(message string, notifType NotificationType) {
	n.message = message
	n.notifType = notifType
	n.timestamp = time.Now()
	n.visible = true
}

// ShowSuccess displays a success notification
func (n *Notification) ShowSuccess(message string) {
	n.Show(message, NotificationSuccess)
}

// ShowWarning displays a warning notification
func (n *Notification) ShowWarning(message string) {
	n.Show(message, NotificationWarning)
}

// ShowError displays an error notification
func (n *Notification) ShowError(message string) {
	n.Show(message, NotificationError)
}

// ShowInfo displays an info notification
func (n *Notification) ShowInfo(message string) {
	n.Show(message, NotificationInfo)
}

// Hide immediately hides the notification
func (n *Notification) Hide() {
	n.visible = false
}

// Update checks if the notification should auto-hide
func (n *Notification) Update() {
	if n.visible && n.duration > 0 && time.Since(n.timestamp) > n.duration {
		n.visible = false
	}
}

// IsVisible returns whether the notification is visible
func (n *Notification) IsVisible() bool {
	return n.visible
}

// Draw renders the notification to the screen
func (n *Notification) Draw(screen *Screen, x, y int, theme *Theme) {
	if !n.visible {
		return
	}

	// Calculate position based on position setting
	actualX, actualY := n.calculatePosition(screen)

	// Get notification style based on type
	var notifStyle NotificationStyle
	switch n.notifType {
	case NotificationSuccess:
		notifStyle = NotificationStyles.Success
	case NotificationWarning:
		notifStyle = NotificationStyles.Warning
	case NotificationError:
		notifStyle = NotificationStyles.Error
	default:
		notifStyle = NotificationStyles.Info
	}

	// Draw block shadow with neo-brutalist style (offset by 1 cell)
	shadowStyle := lipgloss.NewStyle().
		Foreground(theme.Palette.Shadow).
		Background(theme.Palette.Background)
	shadowOffsetX := 1
	shadowOffsetY := 1

	// Use the new DrawBlockShadow method
	screen.DrawBlockShadow(actualX, actualY, n.width, n.height, shadowStyle, shadowOffsetX, shadowOffsetY)

	// Clear the entire notification area with surface color
	surfaceStyle := lipgloss.NewStyle().Background(theme.Palette.Surface)
	ClearArea(screen, actualX, actualY, n.width, n.height, surfaceStyle)

	// Create title with icon
	title := notifStyle.Icon + " Notification"

	// Create a temporary container for the notification
	container := NewContainer()
	container.SetTitle(title)
	container.SetSize(n.width, n.height)
	container.SetPadding(NewMargin(1))
	
	// Create a viewer for the message content
	viewer := NewViewer()
	viewer.SetContent(n.message)
	viewer.SetWrapText(true)
	container.SetContent(viewer)
	
	// Set focus state
	if n.focused {
		container.Focus()
	} else {
		container.Blur()
	}
	
	// Draw the container
	container.Draw(screen, actualX, actualY, theme)
}

// calculatePosition determines where to draw the notification
func (n *Notification) calculatePosition(screen *Screen) (x, y int) {
	margin := 2

	switch n.position {
	case NotificationBottomRight:
		x = screen.Width() - n.width - margin
		y = screen.Height() - n.height - 1
	case NotificationBottomLeft:
		x = margin
		y = screen.Height() - n.height - 1
	case NotificationTopRight:
		x = screen.Width() - n.width - margin
		y = margin
	case NotificationTopLeft:
		x = margin
		y = margin
	}

	return x, y
}

// wrapText wraps text to fit within the given width
func (n *Notification) wrapText(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 > width {
			if currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = word
			} else {
				// Word is longer than width, truncate it
				if len(word) > width {
					lines = append(lines, word[:width])
					currentLine = ""
				} else {
					lines = append(lines, word)
					currentLine = ""
				}
			}
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

// Focus gives keyboard focus to this component
func (n *Notification) Focus() {
	n.focused = true
}

// Blur removes keyboard focus from this component
func (n *Notification) Blur() {
	n.focused = false
}

// IsFocused returns whether this component currently has focus
func (n *Notification) IsFocused() bool {
	return n.focused
}

// HandleKey processes keyboard input when focused
func (n *Notification) HandleKey(key string) bool {
	switch key {
	case "esc", "enter":
		n.Hide()
		return true
	}
	return false
}

// SetSize sets the width and height of the component
func (n *Notification) SetSize(width, height int) {
	n.width = width
	n.height = height
}

// GetSize returns the current width and height
func (n *Notification) GetSize() (width, height int) {
	return n.width, n.height
}

// DrawWithBorder draws the component with a border and optional title
func (n *Notification) DrawWithBorder(screen *Screen, x, y int, theme *Theme, title string) {
	// Notifications always draw themselves with borders, so just call Draw
	n.Draw(screen, x, y, theme)
}
