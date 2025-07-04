package tui

// ContainerBuilder provides a fluent interface for building containers
type ContainerBuilder struct {
	container *Container
}

// NewContainerBuilder creates a new container builder
func NewContainerBuilder() *ContainerBuilder {
	return &ContainerBuilder{
		container: NewContainer(),
	}
}

// WithContent sets the container content
func (b *ContainerBuilder) WithContent(content Component) *ContainerBuilder {
	b.container.SetContent(content)
	return b
}

// WithTitle adds a title to the top-left border
func (b *ContainerBuilder) WithTitle(title string) *ContainerBuilder {
	b.container.SetTitle(title)
	return b
}

// WithTabs adds tabs to the top border
func (b *ContainerBuilder) WithTabs(tabs []string, activeTab int) *ContainerBuilder {
	tabsElement := NewTabsElement(tabs)
	tabsElement.SetActiveTab(activeTab)
	b.container.AddBorderElement(tabsElement, BorderTop, BorderAlignLeft)
	return b
}

// WithStatus adds a status indicator to the specified position
func (b *ContainerBuilder) WithStatus(status string, position BorderPosition, alignment BorderAlignment) *ContainerBuilder {
	statusElement := NewStatusElement(status)
	b.container.AddBorderElement(statusElement, position, alignment)
	return b
}

// WithBadge adds a badge to the specified position
func (b *ContainerBuilder) WithBadge(text string, position BorderPosition, alignment BorderAlignment) *ContainerBuilder {
	badge := NewBadgeElement(text)
	b.container.AddBorderElement(badge, position, alignment)
	return b
}

// WithIcon adds an icon to the border
func (b *ContainerBuilder) WithIcon(icon rune, position BorderPosition, alignment BorderAlignment) *ContainerBuilder {
	iconElement := NewIconElement(icon)
	b.container.AddBorderElement(iconElement, position, alignment)
	return b
}

// WithPadding sets the container padding
func (b *ContainerBuilder) WithPadding(padding Margin) *ContainerBuilder {
	b.container.SetPadding(padding)
	return b
}

// WithBorderStyle sets the border style
func (b *ContainerBuilder) WithBorderStyle(style string) *ContainerBuilder {
	b.container.SetBorderStyle(style)
	return b
}

// NoBorder removes the border
func (b *ContainerBuilder) NoBorder() *ContainerBuilder {
	b.container.SetShowBorder(false)
	return b
}

// Build returns the configured container
func (b *ContainerBuilder) Build() *Container {
	return b.container
}

// TitledPanel creates a container with a title
func TitledPanel(title string, content Component) *Container {
	return NewContainerBuilder().
		WithTitle(title).
		WithContent(content).
		Build()
}

// TabbedPanel creates a container with tabs
func TabbedPanel(tabs []string, activeTab int, content Component) *Container {
	return NewContainerBuilder().
		WithTabs(tabs, activeTab).
		WithContent(content).
		Build()
}

// StatusPanel creates a container with a status indicator
func StatusPanel(title string, status string, content Component) *Container {
	return NewContainerBuilder().
		WithTitle(title).
		WithStatus(status, BorderTop, BorderAlignRight).
		WithContent(content).
		Build()
}

// IconPanel creates a container with an icon and title
func IconPanel(icon rune, title string, content Component) *Container {
	container := NewContainer()
	container.AddBorderElement(NewIconElement(icon), BorderTop, BorderAlignLeft)
	container.AddBorderElementWithOffset(NewTextElement(title), BorderTop, BorderAlignLeft, 3)
	container.SetContent(content)
	return container
}

// MultiElementPanel demonstrates complex border composition
func MultiElementPanel(title string, tabs []string, activeTab int, status string, content Component) *Container {
	container := NewContainer()
	
	// Title on the left
	container.SetTitle(title)
	
	// Tabs in the center
	tabsElement := NewTabsElement(tabs)
	tabsElement.SetActiveTab(activeTab)
	container.AddBorderElement(tabsElement, BorderTop, BorderAlignCenter)
	
	// Status on the right
	container.AddBorderElement(NewStatusElement(status), BorderTop, BorderAlignRight)
	
	// Add a badge to bottom-right
	container.AddBorderElement(NewBadgeElement("v1.0"), BorderBottom, BorderAlignRight)
	
	container.SetContent(content)
	return container
}