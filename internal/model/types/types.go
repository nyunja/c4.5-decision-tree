package model

// Node represents a node in the decision tree
type Node struct {
	Feature    string      `json:"feature,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	IsLeaf     bool        `json:"is_leaf"`
	Class      string      `json:"class,omitempty"`
	Children   []*Node     `json:"children,omitempty"`
	Continuous bool        `json:"continuous,omitempty"`
	Threshold  float64     `json:"threshold,omitempty"`
}

// Model represents the trained decision tree model
type Model struct {
	Root         *Node             `json:"root"`
	FeatureTypes map[string]string `json:"feature_types"` // categorical, numerical, date, timestamp
	FeatureNames []string          `json:"feature_names"`
	TargetName   string            `json:"target_name"`
}

type Instance map[string]interface{}

// ColumnStats stores statistics about a column
type ColumnStats struct {
	Min           float64
	Max           float64
	Sum           float64
	Count         int
	UniqueValues  map[string]int
	IsNumeric     bool
	IsDate        bool
	IsTimestamp   bool
	IsCategorical bool
}

// DatasetStats stores statistics about the entire dataset
type DatasetStats struct {
	RowCount    int
	ColumnStats map[string]*ColumnStats
}
