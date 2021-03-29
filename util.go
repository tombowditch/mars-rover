package rover

// util function to check int cannot go out of bounds of min,max
func boundsCheck(i int, min int, max int) int {
	// if i is bigger than max, return max
	if i > max {
		return max
	}

	// if i is smaller than min, return min
	if i < min {
		return min
	}

	return i
}
