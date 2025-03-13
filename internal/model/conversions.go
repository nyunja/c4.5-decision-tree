package model

/*
* StringArrayToOrdinal - converts an array of strings to an array of integers
* the integers will be assigned as first come basis. All strings that are equal will have the same integer value
* input:
*  an array of strings
*
* output:
*  an array of integers with substituted integer values
 */
func StringArrayToOrdinal(arr []string) (map[string]int, []int) {
	index := 1
	stringToIntegerMapping := make(map[string]int)
	norminal := []int{}

	for _, item := range arr {
		if v, ok := stringToIntegerMapping[item]; ok {
			norminal = append(norminal, v)
		} else {
			stringToIntegerMapping[item] = index
			norminal = append(norminal, index)
			index++
		}
	}
	return stringToIntegerMapping, norminal
}
