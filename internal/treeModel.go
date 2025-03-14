package internal

/*
* TreeNode - represents a node in the decision tree
* Attribute: Index of the attribute used for splitting
* Threshold: Threshold for continuous attributes
* Children: Child nodes (key: attribute value or threshold comparison result)
* Class: Class label for leaf nodes
 */
type TreeNode struct {
	Attribute int                       // Index of the attribute used for splitting
	Threshold float64                   // Threshold for continuous attributes
	Children  map[interface{}]*TreeNode // Child nodes (key: attribute value or threshold comparison result)
	Class     int                       // Class label for leaf nodes
}

/*
* JSONTreeNode - represents a node in the decision tree for JSON output
* SplitFeature: Name of the feature used for splitting
* SplitType: Type of split (continuous or categorical)
* Children: Child nodes (key: attribute value or threshold comparison result)
* ClassDistribution: Class distribution for leaf nodes
 */
type JSONTreeNode struct {
	SplitFeature      string                   `json:"split_feature"`
	SplitType         string                   `json:"split_type"`
	Children          map[string]*JSONTreeNode `json:"children"`
	ClassDistribution map[string]int           `json:"class_distribution"`
}

/*
* Subset - represents a subset of data and labels
* Dataset: Subset of the original dataset
* Labels: Corresponding class labels
 */
type Subset struct {
	Dataset [][]float64
	Labels  []int
}
