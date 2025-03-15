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
