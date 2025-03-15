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

// Should check that the returned ClassCounter's Counts map has an initial length of zero
func TestNewClassCounter_InitLenZero(t *testing.T) {
	counter := NewClassCounter()
	if len(counter.Counts) != 0 {
		t.Errorf("Expected initial Counts map length to be 0, but got %d", len(counter.Counts))
	}
}

// Increment the count for a new cless
func TestClassCounter_Add(t *testing.T) {
	counter := NewClassCounter()

	counter.Add("class1")

	if counter.Counts["class1"] != 1 {
		t.Errorf("Expected count for class1 to be 1, but got %d", counter.Counts["class1"])
	}

	if counter.Total != 1 {
		t.Errorf("Expected total count to be 1, but got %d", counter.Total)
	}
}

// Should increment the count for an existing class
func TestClassCounter_AddExistingClass(t *testing.T) {
	counter := NewClassCounter()
	counter.Add("A")
	counter.Add("A")

	if count, exists := counter.Counts["A"]; !exists || count != 2 {
		t.Errorf("Expected count for class 'A' to be 2, got %d", count)
	}

	if counter.Total != 2 {
		t.Errorf("Expected total count to be 2, got %d", counter.Total)
	}
}

// Should return 0 when Total is 0
func TestGetEntropy_ZeroReturn(t *testing.T) {
	c := NewClassCounter()

	entropy := c.GetEntropy()

	if entropy != 0 {
		t.Errorf("Expected entropy to be 0 when Total is 0, but got %f", entropy)
	}
}
