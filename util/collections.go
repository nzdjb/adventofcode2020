package util

// Index Find index of first match in slice.
func Index(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

// Include Determine whether needle exists in haystack.
func Include(haystack []int, needle int) bool {
	return Index(haystack, needle) >= 0
}
