package tui

// ConstraintType defines the type of constraint
type ConstraintType int

const (
	// Length is a fixed size in cells
	Length ConstraintType = iota
	// Percentage is a percentage of the parent's size
	Percentage
	// Ratio is a ratio relative to other components
	Ratio
	// Min is a minimum size constraint
	Min
	// Max is a maximum size constraint
	Max
)

// Constraint represents a size constraint for layouts
type Constraint struct {
	Type  ConstraintType
	Value float64
}

// NewLength creates a fixed length constraint
func NewLength(cells int) Constraint {
	return Constraint{Type: Length, Value: float64(cells)}
}

// NewPercentage creates a percentage constraint (0.0 to 1.0)
func NewPercentage(pct float64) Constraint {
	if pct < 0 {
		pct = 0
	} else if pct > 1 {
		pct = 1
	}
	return Constraint{Type: Percentage, Value: pct}
}

// NewRatio creates a ratio constraint
func NewRatio(ratio float64) Constraint {
	if ratio < 0 {
		ratio = 0
	}
	return Constraint{Type: Ratio, Value: ratio}
}

// NewMin creates a minimum size constraint
func NewMin(min int) Constraint {
	if min < 0 {
		min = 0
	}
	return Constraint{Type: Min, Value: float64(min)}
}

// NewMax creates a maximum size constraint
func NewMax(max int) Constraint {
	if max < 0 {
		max = 0
	}
	return Constraint{Type: Max, Value: float64(max)}
}

// ConstraintSet combines multiple constraints for a single dimension
type ConstraintSet struct {
	Base Constraint   // Primary constraint (Length, Percentage, or Ratio)
	Min  *Constraint  // Optional minimum
	Max  *Constraint  // Optional maximum
}

// NewConstraintSet creates a constraint set with just a base constraint
func NewConstraintSet(base Constraint) ConstraintSet {
	return ConstraintSet{Base: base}
}

// WithMin adds a minimum constraint
func (cs ConstraintSet) WithMin(min int) ConstraintSet {
	minConstraint := NewMin(min)
	cs.Min = &minConstraint
	return cs
}

// WithMax adds a maximum constraint
func (cs ConstraintSet) WithMax(max int) ConstraintSet {
	maxConstraint := NewMax(max)
	cs.Max = &maxConstraint
	return cs
}

// Calculate computes the actual size given a parent size
func (cs ConstraintSet) Calculate(parentSize int, ratioTotal float64) int {
	var size int
	
	switch cs.Base.Type {
	case Length:
		size = int(cs.Base.Value)
	case Percentage:
		size = int(float64(parentSize) * cs.Base.Value)
	case Ratio:
		if ratioTotal > 0 {
			size = int(float64(parentSize) * (cs.Base.Value / ratioTotal))
		}
	}
	
	// Apply min/max constraints
	if cs.Min != nil && size < int(cs.Min.Value) {
		size = int(cs.Min.Value)
	}
	if cs.Max != nil && size > int(cs.Max.Value) {
		size = int(cs.Max.Value)
	}
	
	return size
}

// CalculateConstraints calculates sizes for multiple constraint sets
func CalculateConstraints(constraints []ConstraintSet, totalSize int) []int {
	sizes := make([]int, len(constraints))
	
	// First pass: calculate fixed and percentage constraints
	remainingSize := totalSize
	ratioTotal := 0.0
	ratioIndices := []int{}
	
	for i, cs := range constraints {
		switch cs.Base.Type {
		case Length:
			size := cs.Calculate(totalSize, 0)
			sizes[i] = size
			remainingSize -= size
		case Percentage:
			size := cs.Calculate(totalSize, 0)
			sizes[i] = size
			remainingSize -= size
		case Ratio:
			ratioTotal += cs.Base.Value
			ratioIndices = append(ratioIndices, i)
		}
	}
	
	// Second pass: distribute remaining size among ratio constraints
	if len(ratioIndices) > 0 && remainingSize > 0 {
		for _, i := range ratioIndices {
			sizes[i] = constraints[i].Calculate(remainingSize, ratioTotal)
		}
	}
	
	// Third pass: ensure we don't exceed total size
	totalUsed := 0
	for _, size := range sizes {
		totalUsed += size
	}
	
	// If we've exceeded, scale down proportionally
	if totalUsed > totalSize {
		scale := float64(totalSize) / float64(totalUsed)
		for i := range sizes {
			sizes[i] = int(float64(sizes[i]) * scale)
		}
	}
	
	return sizes
}

// Direction represents layout direction
type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

// Alignment represents how items are aligned
type Alignment int

const (
	AlignStart Alignment = iota
	AlignCenter
	AlignEnd
	AlignStretch
)

// LayoutConfig holds common layout configuration
type LayoutConfig struct {
	Direction Direction
	Alignment Alignment
	Spacing   int
	Padding   Margin
}

// NewLayoutConfig creates a new layout configuration
func NewLayoutConfig(direction Direction) LayoutConfig {
	return LayoutConfig{
		Direction: direction,
		Alignment: AlignStart,
		Spacing:   0,
		Padding:   NewMargin(0),
	}
}