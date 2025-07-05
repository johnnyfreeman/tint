package tui

import (
	"testing"
)

func TestConstraints(t *testing.T) {
	t.Run("Length constraint", func(t *testing.T) {
		c := NewLength(20)
		if c.Type != Length {
			t.Errorf("Expected Length type, got %v", c.Type)
		}
		if c.Value != 20 {
			t.Errorf("Expected value 20, got %v", c.Value)
		}
	})
	
	t.Run("Percentage constraint", func(t *testing.T) {
		c := NewPercentage(0.5)
		if c.Type != Percentage {
			t.Errorf("Expected Percentage type, got %v", c.Type)
		}
		if c.Value != 0.5 {
			t.Errorf("Expected value 0.5, got %v", c.Value)
		}
		
		// Test clamping
		c = NewPercentage(1.5)
		if c.Value != 1.0 {
			t.Errorf("Expected value clamped to 1.0, got %v", c.Value)
		}
		
		c = NewPercentage(-0.5)
		if c.Value != 0.0 {
			t.Errorf("Expected value clamped to 0.0, got %v", c.Value)
		}
	})
	
	t.Run("Ratio constraint", func(t *testing.T) {
		c := NewRatio(2.5)
		if c.Type != Ratio {
			t.Errorf("Expected Ratio type, got %v", c.Type)
		}
		if c.Value != 2.5 {
			t.Errorf("Expected value 2.5, got %v", c.Value)
		}
		
		// Test negative clamping
		c = NewRatio(-1)
		if c.Value != 0 {
			t.Errorf("Expected value clamped to 0, got %v", c.Value)
		}
	})
}

func TestConstraintSet(t *testing.T) {
	t.Run("Calculate Length", func(t *testing.T) {
		cs := NewConstraintSet(NewLength(30))
		size := cs.Calculate(100, 0)
		if size != 30 {
			t.Errorf("Expected size 30, got %d", size)
		}
	})
	
	t.Run("Calculate Percentage", func(t *testing.T) {
		cs := NewConstraintSet(NewPercentage(0.25))
		size := cs.Calculate(100, 0)
		if size != 25 {
			t.Errorf("Expected size 25, got %d", size)
		}
	})
	
	t.Run("Calculate Ratio", func(t *testing.T) {
		cs := NewConstraintSet(NewRatio(2))
		size := cs.Calculate(100, 5) // 2/5 of 100
		if size != 40 {
			t.Errorf("Expected size 40, got %d", size)
		}
	})
	
	t.Run("With Min constraint", func(t *testing.T) {
		cs := NewConstraintSet(NewPercentage(0.1)).WithMin(20)
		size := cs.Calculate(100, 0) // 10% of 100 = 10, but min is 20
		if size != 20 {
			t.Errorf("Expected size 20 (min), got %d", size)
		}
	})
	
	t.Run("With Max constraint", func(t *testing.T) {
		cs := NewConstraintSet(NewPercentage(0.5)).WithMax(30)
		size := cs.Calculate(100, 0) // 50% of 100 = 50, but max is 30
		if size != 30 {
			t.Errorf("Expected size 30 (max), got %d", size)
		}
	})
	
	t.Run("With Min and Max", func(t *testing.T) {
		cs := NewConstraintSet(NewLength(5)).WithMin(10).WithMax(20)
		size := cs.Calculate(100, 0) // 5 is below min
		if size != 10 {
			t.Errorf("Expected size 10 (min), got %d", size)
		}
		
		cs = NewConstraintSet(NewLength(25)).WithMin(10).WithMax(20)
		size = cs.Calculate(100, 0) // 25 is above max
		if size != 20 {
			t.Errorf("Expected size 20 (max), got %d", size)
		}
	})
}

func TestCalculateConstraints(t *testing.T) {
	t.Run("Mixed constraints", func(t *testing.T) {
		constraints := []ConstraintSet{
			NewConstraintSet(NewLength(20)),      // Fixed 20
			NewConstraintSet(NewPercentage(0.3)), // 30% of 100 = 30
			NewConstraintSet(NewRatio(1)),        // Remaining 50
		}
		
		sizes := CalculateConstraints(constraints, 100)
		
		if len(sizes) != 3 {
			t.Fatalf("Expected 3 sizes, got %d", len(sizes))
		}
		
		if sizes[0] != 20 {
			t.Errorf("Expected first size 20, got %d", sizes[0])
		}
		if sizes[1] != 30 {
			t.Errorf("Expected second size 30, got %d", sizes[1])
		}
		if sizes[2] != 50 {
			t.Errorf("Expected third size 50, got %d", sizes[2])
		}
	})
	
	t.Run("Multiple ratios", func(t *testing.T) {
		constraints := []ConstraintSet{
			NewConstraintSet(NewRatio(1)), // 1/3
			NewConstraintSet(NewRatio(2)), // 2/3
		}
		
		sizes := CalculateConstraints(constraints, 90)
		
		if sizes[0] != 30 {
			t.Errorf("Expected first size 30, got %d", sizes[0])
		}
		if sizes[1] != 60 {
			t.Errorf("Expected second size 60, got %d", sizes[1])
		}
	})
	
	t.Run("Overflow handling", func(t *testing.T) {
		constraints := []ConstraintSet{
			NewConstraintSet(NewLength(60)),
			NewConstraintSet(NewLength(60)),
		}
		
		sizes := CalculateConstraints(constraints, 100)
		
		// Should scale down proportionally
		total := sizes[0] + sizes[1]
		if total > 100 {
			t.Errorf("Total size %d exceeds available 100", total)
		}
	})
	
	t.Run("All fixed sizes", func(t *testing.T) {
		constraints := []ConstraintSet{
			NewConstraintSet(NewLength(20)),
			NewConstraintSet(NewLength(30)),
			NewConstraintSet(NewLength(40)),
		}
		
		sizes := CalculateConstraints(constraints, 100)
		
		if sizes[0] != 20 {
			t.Errorf("Expected first size 20, got %d", sizes[0])
		}
		if sizes[1] != 30 {
			t.Errorf("Expected second size 30, got %d", sizes[1])
		}
		if sizes[2] != 40 {
			t.Errorf("Expected third size 40, got %d", sizes[2])
		}
	})
}