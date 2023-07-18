package day1

func BinarySearch(haystack []int, needle int) bool {
	lo, hi := 0, len(haystack)

	for lo < hi {
		m := lo + (hi-lo)/2
		v := haystack[m]

		if v == needle {
			return true
		} else if v > needle {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
