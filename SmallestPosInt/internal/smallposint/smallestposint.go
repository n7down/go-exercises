package smallestposint

import (
	"sort"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SmallestPosInt(a []int) int {
	sort.Ints(a)
	values := removeDuplicateValues(a)
	var smallestPos int = values[0]

	for contains(values, smallestPos) || smallestPos < 1 {
		smallestPos = smallestPos + 1
	}
	return smallestPos
}
