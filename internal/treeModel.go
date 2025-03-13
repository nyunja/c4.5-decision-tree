package internal

type DecisionTree struct {
	Feature    string                `json:"feature,omitempty"`    // The feature to split on
	Threshold  float64               `json:"threshold,omitempty"`  // The threshold for numeric splits (only for continuous features)
	SplitType  string                `json:"split_type,omitempty"` // "categorical" or "continuous"
	Categories map[any]*categorylabel `json:"categories,omitempty"`
	Left       *DecisionTree         `json:"left,omitempty"`     // Left child (if applicable)
	Right      *DecisionTree         `json:"right,omitempty"`    // Right child (if applicable)
	Label      string                `json:"label,omitempty"`   
	Metadata   *Metadata             `json:"metadata"` 
}

type categorylabel struct {
	Label string `json:"label"`
}

type Metadata struct {
	Features []string          `json:"features"` // List of feature names
	Types    map[string]string `json:"types"`    // Feature data types ("categorical" or "continuous")
}
