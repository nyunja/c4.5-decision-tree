package c45

func Train(data [][]string, labelIndex int) *Node {
	return buildTree(data, labelIndex)
}

// Build decision tree recursively
func buildTree(data [][]string, labelIndex int) *Node {
	if allSameClass(data, labelIndex) {
		return &Node{IsLeaf: true, Class: data[0][labelIndex]}
	}
}

func allSameClass(data [][]string, labelIndex int) bool {
	if len(data) == 0 {
		return false
	}
	firstClass := data[0][labelIndex]
	for _, row := range data {
        if row[labelIndex] != firstClass {
            return false
        }
    }
	return true
}
