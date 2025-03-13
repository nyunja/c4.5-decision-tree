package internal

/*
* StringArrayToOrdinal - converts an array of strings to an array of integers
* the integers will be assigned as first come basis. All strings that are equal will have the same integer value
* input:
*  an array of strings
*
* output:
*  an array of integers with substituted integer values
 */
func StringArrayToOrdinal(arr []string) []int {
	index := 1
	m := make(map[string]int)
	norminal := []int{}

	for _, item := range arr {
		if v, ok := m[item]; ok {
			norminal = append(norminal, v)
		} else {
			m[item] = index
			index++
		}
	}
	return norminal
}
