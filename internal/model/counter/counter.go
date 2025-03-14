package counter

import "math"

// ClassCounter is a helper struct to efficiently count class occurrences
type ClassCounter struct {
	Counts map[string]int
	Total  int
}

// NewClassCounter creates a new ClassCounter
func NewClassCounter() *ClassCounter {
	return &ClassCounter{
		Counts: make(map[string]int),
	}
}

// Add adds a class to the counter
func (c *ClassCounter) Add(class string) {
	c.Counts[class]++
	c.Total++
}

// GetEntropy calculates the entropy of the class distribution
func (c *ClassCounter) GetEntropy() float64 {
	if c.Total == 0 {
		return 0
	}

	entropy := 0.0
	totalFloat := float64(c.Total)

	for _, count := range c.Counts {
		probability := float64(count) / totalFloat
		entropy -= probability * math.Log2(probability)
	}

	return entropy
}

// GetMajorityClass returns the most frequent class
func (c *ClassCounter) GetMajorityClass() string {
	majorityClass := ""
	maxCount := 0

	for class, count := range c.Counts {
		if count > maxCount {
			maxCount = count
			majorityClass = class
		}
	}

	return majorityClass
}
