package set

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Fatal("New() returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("New set should be empty, got length %d", s.Len())
	}
}

func TestOf(t *testing.T) {
	s := Of(1, 2, 3, 2, 1)
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Error("Set should contain 1, 2, and 3")
	}
}

func TestAdd(t *testing.T) {
	s := New[string]()
	s.Add("hello")
	if !s.Contains("hello") {
		t.Error("Set should contain 'hello'")
	}
	if s.Len() != 1 {
		t.Errorf("Expected length 1, got %d", s.Len())
	}

	// Adding duplicate
	s.Add("hello")
	if s.Len() != 1 {
		t.Errorf("Duplicate add should not increase length, got %d", s.Len())
	}
}

func TestRemove(t *testing.T) {
	s := Of(1, 2, 3)
	s.Remove(2)
	if s.Contains(2) {
		t.Error("Set should not contain 2 after removal")
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	// Removing non-existent element
	s.Remove(99)
	if s.Len() != 2 {
		t.Error("Removing non-existent element should not change length")
	}
}

func TestContains(t *testing.T) {
	s := Of("a", "b", "c")
	if !s.Contains("a") {
		t.Error("Set should contain 'a'")
	}
	if s.Contains("z") {
		t.Error("Set should not contain 'z'")
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}
}

func TestClear(t *testing.T) {
	s := Of(1, 2, 3, 4, 5)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Set should be empty after Clear()")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", s.Len())
	}
}

func TestToSlice(t *testing.T) {
	s := Of(1, 2, 3)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	// Check all elements are present (order not guaranteed)
	found := make(map[int]bool)
	for _, v := range slice {
		found[v] = true
	}
	if !found[1] || !found[2] || !found[3] {
		t.Error("Slice should contain all set elements")
	}
}

func TestClone(t *testing.T) {
	s1 := Of(1, 2, 3)
	s2 := s1.Clone()

	if !s1.Equal(s2) {
		t.Error("Clone should be equal to original")
	}

	// Modify clone
	s2.Add(4)
	if s1.Contains(4) {
		t.Error("Modifying clone should not affect original")
	}
}

func TestUnion(t *testing.T) {
	s1 := Of(1, 2, 3)
	s2 := Of(3, 4, 5)
	result := s1.Union(s2)

	if result.Len() != 5 {
		t.Errorf("Expected union length 5, got %d", result.Len())
	}
	for i := 1; i <= 5; i++ {
		if !result.Contains(i) {
			t.Errorf("Union should contain %d", i)
		}
	}
}

func TestIntersection(t *testing.T) {
	s1 := Of(1, 2, 3, 4)
	s2 := Of(3, 4, 5, 6)
	result := s1.Intersection(s2)

	if result.Len() != 2 {
		t.Errorf("Expected intersection length 2, got %d", result.Len())
	}
	if !result.Contains(3) || !result.Contains(4) {
		t.Error("Intersection should contain 3 and 4")
	}
	if result.Contains(1) || result.Contains(5) {
		t.Error("Intersection should not contain 1 or 5")
	}
}

func TestDifference(t *testing.T) {
	s1 := Of(1, 2, 3, 4)
	s2 := Of(3, 4, 5, 6)
	result := s1.Difference(s2)

	if result.Len() != 2 {
		t.Errorf("Expected difference length 2, got %d", result.Len())
	}
	if !result.Contains(1) || !result.Contains(2) {
		t.Error("Difference should contain 1 and 2")
	}
	if result.Contains(3) || result.Contains(4) {
		t.Error("Difference should not contain 3 or 4")
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := Of(1, 2, 3)
	s2 := Of(3, 4, 5)
	result := s1.SymmetricDifference(s2)

	if result.Len() != 4 {
		t.Errorf("Expected symmetric difference length 4, got %d", result.Len())
	}
	if !result.Contains(1) || !result.Contains(2) || !result.Contains(4) || !result.Contains(5) {
		t.Error("Symmetric difference should contain 1, 2, 4, and 5")
	}
	if result.Contains(3) {
		t.Error("Symmetric difference should not contain 3")
	}
}

func TestIsSubset(t *testing.T) {
	s1 := Of(1, 2)
	s2 := Of(1, 2, 3, 4)

	if !s1.IsSubset(s2) {
		t.Error("s1 should be a subset of s2")
	}
	if s2.IsSubset(s1) {
		t.Error("s2 should not be a subset of s1")
	}

	s3 := Of(1, 2)
	if !s1.IsSubset(s3) {
		t.Error("Equal sets should be subsets of each other")
	}
}

func TestIsSuperset(t *testing.T) {
	s1 := Of(1, 2, 3, 4)
	s2 := Of(1, 2)

	if !s1.IsSuperset(s2) {
		t.Error("s1 should be a superset of s2")
	}
	if s2.IsSuperset(s1) {
		t.Error("s2 should not be a superset of s1")
	}
}

func TestEqual(t *testing.T) {
	s1 := Of(1, 2, 3)
	s2 := Of(3, 2, 1)
	s3 := Of(1, 2, 3, 4)

	if !s1.Equal(s2) {
		t.Error("s1 should equal s2")
	}
	if s1.Equal(s3) {
		t.Error("s1 should not equal s3")
	}
}

func TestForEach(t *testing.T) {
	s := Of(1, 2, 3)
	sum := 0
	s.ForEach(func(n int) {
		sum += n
	})
	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestFilter(t *testing.T) {
	s := Of(1, 2, 3, 4, 5, 6)
	evens := s.Filter(func(n int) bool {
		return n%2 == 0
	})

	if evens.Len() != 3 {
		t.Errorf("Expected 3 even numbers, got %d", evens.Len())
	}
	if !evens.Contains(2) || !evens.Contains(4) || !evens.Contains(6) {
		t.Error("Filter should contain 2, 4, and 6")
	}
}

func TestStringSet(t *testing.T) {
	s := Of("apple", "banana", "cherry")
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}
	if !s.Contains("banana") {
		t.Error("Set should contain 'banana'")
	}
}

func BenchmarkAdd(b *testing.B) {
	s := New[int]()
	b.ResetTimer()
	for i := range b.N {
		s.Add(i)
	}
}

func BenchmarkContains(b *testing.B) {
	s := New[int]()
	for i := range 1000 {
		s.Add(i)
	}
	b.ResetTimer()
	for i := range b.N {
		s.Contains(i % 1000)
	}
}

func BenchmarkUnion(b *testing.B) {
	s1 := New[int]()
	s2 := New[int]()
	for i := range b.N {
		s1.Add(i)
		s2.Add(i + 250)
	}
	b.ResetTimer()
	for b.Loop() {
		s1.Union(s2)
	}
}
