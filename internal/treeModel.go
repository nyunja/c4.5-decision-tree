package internal

type DecisionTree struct {
	Feature    string                     `json:"feature,omitempty"`    // The feature to split on
	Threshold  float64                    `json:"threshold,omitempty"`  // Threshold for numeric splits (only for continuous features)
	SplitType  string                      `json:"split_type,omitempty"` // "categorical" or "continuous"
	Categories map[string]*DecisionTree    `json:"categories,omitempty"` // Categorical splits (maps category values to child nodes)
	Left       *DecisionTree               `json:"left,omitempty"`       // Left child (for continuous splits)
	Right      *DecisionTree               `json:"right,omitempty"`      // Right child (for continuous splits)
	Label      string                      `json:"label,omitempty"`      // If a node has a Label, it's a leaf node
	Metadata   *Metadata                   `json:"metadata,omitempty"`   // Metadata, should be set only in root
}

type Metadata struct {
	Features []string          `json:"features"` // List of feature names
	Types    map[string]string `json:"types"`    // Feature data types ("categorical" or "continuous")
}

