package model

import "testing"

func TestStringArrayToOrdinal(t *testing.T) {
	table := []struct {
		Input    []string
		Expected []int
		Error    string
	}{
		{
			Input:    []string{"red", "blue", "red", "blue", "red", "blue", "green", "green", "green"},
			Expected: []int{1, 2, 1, 2, 1, 2, 3, 3, 3},
			Error:    "test 1: color classes test failure",
		},
		{
			Input:    []string{"red", "blue", "green"},
			Expected: []int{1, 2, 3},
			Error:    "test 1: color classes test failure",
		},
		{
			Input:    []string{"red", "blue", "blue"},
			Expected: []int{1, 2, 2},
			Error:    "test 1: color classes test failure",
		},
	}

	for _, test := range table {
		_, classes := StringArrayToOrdinal(test.Input)
		equal := compareSlices(classes, test.Expected)
		if !equal {
			t.Errorf(test.Error)
		}
	}
}

func compareSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
