package utils

// this should be the fastest (optimized by complier to have no if branches)
func Bool2int(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
