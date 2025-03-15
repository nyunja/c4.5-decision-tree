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

// Should verify that multiple calls to NewClassCounter create distinct instances
func TestNewClassCounter_DistinctInstances(t *testing.T) {
	counter1 := NewClassCounter()
	counter2 := NewClassCounter()

	if counter1 == counter2 {
		t.Error("NewClassCounter should create distinct instances, but they are the same")
	}

	counter1.Add("class1")
	if counter2.Total != 0 {
		t.Errorf("Changes to counter1 should not affect counter2. Expected counter2.Total to be 0, but got %d", counter2.Total)
	}

	if len(counter2.Counts) != 0 {
		t.Errorf("Counter2 should have empty Counts, but has %d items", len(counter2.Counts))
	}
}

// Should confirm that the returned ClassCounter has no pre-existing entries in Counts
func TestNewClassCounter_NoPreExistingEntries(t *testing.T) {
	counter := NewClassCounter()
	if counter == nil {
		t.Fatal("NewClassCounter returned nil")
	}
	if len(counter.Counts) != 0 {
		t.Errorf("Expected empty Counts map, got %d entries", len(counter.Counts))
	}
	if counter.Total != 0 {
		t.Errorf("Expected Total to be 0, got %d", counter.Total)
	}
}
