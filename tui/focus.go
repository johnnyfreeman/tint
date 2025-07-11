package tui


// FocusManager manages focus state between multiple components
type FocusManager struct {
	components []FocusableComponent
	current    int
	wrapAround bool
}

// FocusableComponent pairs a component with its identifier
type FocusableComponent struct {
	ID        string
	Component Focusable
}

// NewFocusManager creates a new focus manager
func NewFocusManager() *FocusManager {
	return &FocusManager{
		components: []FocusableComponent{},
		current:    -1,
		wrapAround: true,
	}
}

// Add registers a component with the focus manager
func (fm *FocusManager) Add(id string, component Focusable) {
	fm.components = append(fm.components, FocusableComponent{
		ID:        id,
		Component: component,
	})

	// If this is the first component, focus it
	if len(fm.components) == 1 {
		fm.current = 0
		component.Focus()
	}
}

// Remove unregisters a component from the focus manager
func (fm *FocusManager) Remove(id string) {
	for i, fc := range fm.components {
		if fc.ID == id {
			// Blur the component if it's focused
			if i == fm.current {
				fc.Component.Blur()
			}

			// Remove from slice
			fm.components = append(fm.components[:i], fm.components[i+1:]...)

			// Adjust current index
			if fm.current >= len(fm.components) && len(fm.components) > 0 {
				fm.current = len(fm.components) - 1
			} else if len(fm.components) == 0 {
				fm.current = -1
			}

			// Focus the new current component
			if fm.current >= 0 && fm.current < len(fm.components) {
				fm.components[fm.current].Component.Focus()
			}

			break
		}
	}
}

// Focus sets focus to a specific component by ID
func (fm *FocusManager) Focus(id string) {
	for i, fc := range fm.components {
		if fc.ID == id {
			fm.setFocus(i)
			break
		}
	}
}

// FocusNext moves focus to the next component
func (fm *FocusManager) FocusNext() {
	if len(fm.components) == 0 {
		return
	}

	newIndex := fm.current + 1
	if newIndex >= len(fm.components) {
		if fm.wrapAround {
			newIndex = 0
		} else {
			return
		}
	}

	fm.setFocus(newIndex)
}

// FocusPrevious moves focus to the previous component
func (fm *FocusManager) FocusPrevious() {
	if len(fm.components) == 0 {
		return
	}

	newIndex := fm.current - 1
	if newIndex < 0 {
		if fm.wrapAround {
			newIndex = len(fm.components) - 1
		} else {
			return
		}
	}

	fm.setFocus(newIndex)
}

// GetFocused returns the currently focused component and its ID
func (fm *FocusManager) GetFocused() (id string, component Focusable) {
	if fm.current >= 0 && fm.current < len(fm.components) {
		fc := fm.components[fm.current]
		return fc.ID, fc.Component
	}
	return "", nil
}

// GetFocusedID returns just the ID of the currently focused component
func (fm *FocusManager) GetFocusedID() string {
	if fm.current >= 0 && fm.current < len(fm.components) {
		return fm.components[fm.current].ID
	}
	return ""
}

// SetWrapAround sets whether focus should wrap around at the ends
func (fm *FocusManager) SetWrapAround(wrap bool) {
	fm.wrapAround = wrap
}

// HandleKey processes common focus navigation keys
// Returns true if the key was handled
func (fm *FocusManager) HandleKey(key string) bool {
	switch key {
	case "tab":
		fm.FocusNext()
		return true
	case "shift+tab":
		fm.FocusPrevious()
		return true
	}
	return false
}

// Clear removes all components and resets focus
func (fm *FocusManager) Clear() {
	// Blur current component
	if fm.current >= 0 && fm.current < len(fm.components) {
		fm.components[fm.current].Component.Blur()
	}

	fm.components = []FocusableComponent{}
	fm.current = -1
}

// setFocus is an internal method to change focus
func (fm *FocusManager) setFocus(index int) {
	if index < 0 || index >= len(fm.components) {
		return
	}

	// Blur current component
	if fm.current >= 0 && fm.current < len(fm.components) {
		fm.components[fm.current].Component.Blur()
	}

	// Focus new component
	fm.current = index
	fm.components[fm.current].Component.Focus()
}

// FocusGroup manages multiple focus managers for complex layouts
type FocusGroup struct {
	managers   map[string]*FocusManager
	activeID   string
	groupOrder []string
}

// NewFocusGroup creates a new focus group
func NewFocusGroup() *FocusGroup {
	return &FocusGroup{
		managers:   make(map[string]*FocusManager),
		groupOrder: []string{},
	}
}

// AddManager adds a focus manager to the group
func (fg *FocusGroup) AddManager(id string, manager *FocusManager) {
	fg.managers[id] = manager
	fg.groupOrder = append(fg.groupOrder, id)

	// If this is the first manager, make it active
	if len(fg.managers) == 1 {
		fg.activeID = id
	}
}

// SwitchTo switches to a different focus manager
func (fg *FocusGroup) SwitchTo(id string) {
	if _, exists := fg.managers[id]; exists {
		fg.activeID = id
	}
}

// HandleKey routes key events to the active focus manager
func (fg *FocusGroup) HandleKey(key string) bool {
	if manager, exists := fg.managers[fg.activeID]; exists {
		return manager.HandleKey(key)
	}
	return false
}
