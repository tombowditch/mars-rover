package rover

// util function used in rover Move to make sure rover cannot go off the edge of the planet
func max(i int, max int) int {
	if i > max {
		return max
	} else {
		return i
	}
}

// util function used in rover Move to make sure rover cannot go off the edge of the planet
func min(i int, min int) int {
	if i < min {
		return min
	} else {
		return i
	}
}
