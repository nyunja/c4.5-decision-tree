package counter

import "testing"

// Should return a new ClassCounter instance with an empty Counts map
func TestNewClassCounter_ReturnNewCounterClass(t *testing.T) {
	counter := NewClassCounter()
	if counter == nil {
		t.Fatal("NewClassCounter returned nil")
	}
	if counter.Counts == nil {
		t.Error("Counts field is nil")
	}
}

// Should ensure the Counts map is initialized and ready for use
func TestNewClassCounter_InitCountMap(t *testing.T) {
	counter := NewClassCounter()

	if counter == nil {
		t.Fatal("Expected NewClassCounter to return a non-nil ClassCounter")
	}

	if counter.Counts == nil {
		t.Error("Expected Counts map to be initialized")
	}

	if len(counter.Counts) != 0 {
		t.Errorf("Expected Counts map to be empty, but got %d elements", len(counter.Counts))
	}

	if counter.Total != 0 {
		t.Errorf("Expected Total to be 0, but got %d", counter.Total)
	}
}
