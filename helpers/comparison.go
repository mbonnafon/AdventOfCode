package helpers

func Min(s []int) int {
	var min int
	for i, v := range s {
		if i == 0 {
			min = v
			continue
		}
		if v < min {
			min = v
		}
	}
	return min
}

func Max(s []int) int {
	var max int
	for i, v := range s {
		if i == 0 {
			max = v
			continue
		}
		if v > max {
			max = v
		}
	}
	return max
}
