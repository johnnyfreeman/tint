# TUI Component Migration Guide

## SplitView -> Split Migration

The old `SplitView` component has been replaced with the new constraint-based `Split` component. Here's how to migrate:

### Old SplitView API:
```go
splitView := tui.NewSplitView(true) // vertical split
splitView.SetSplit(30) // 30 pixels for left pane
splitView.SetLeft(leftComponent)
splitView.SetRight(rightComponent)
```

### New Split API:
```go
split := tui.NewVSplit() // or tui.NewSplit(true)
split.SetFixed(30) // 30 pixels for first pane
split.SetFirst(leftComponent)
split.SetSecond(rightComponent)
```

### Key Differences:

1. **Naming**: `SetLeft/SetRight` → `SetFirst/SetSecond`
2. **Constraints**: New Split supports constraints with min/max bounds
3. **API**: More flexible constraint-based sizing

### Advanced Usage with Constraints:
```go
// Percentage with min/max
split.SetConstraint(
    tui.NewConstraintSet(tui.NewPercentage(0.3)).
        WithMin(100).
        WithMax(300),
)

// Ratio-based splitting
split.SetRatio(1.0) // First pane gets 1:remaining ratio
```

## FlexLayout -> LinearLayout Migration

### Old FlexLayout API:
```go
flex := tui.NewFlexLayout(tui.LayoutHorizontal)
flex.Spacing = 5
flex.Padding = 10
```

### New LinearLayout API:
```go
layout := tui.HBox() // or tui.NewLinearLayout(tui.Horizontal)
layout.SetSpacing(5)
layout.SetPadding(tui.NewMargin(10))

// Add components with constraints
layout.AddFixed(component1, 100)    // Fixed 100 pixels
layout.AddFlex(component2, 1)       // Flex ratio 1
layout.AddPercentage(component3, 0.25) // 25% of available space
```

## Benefits of New Layout System

1. **Constraint-based**: More flexible sizing with min/max bounds
2. **Component interface**: All layouts implement the standard Component interface
3. **Composable**: Layouts can be nested easily
4. **Responsive**: Conditional layouts support responsive design
5. **Consistent API**: All layouts follow the same patterns