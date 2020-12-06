package helpers

import "strconv"

func IntInInterval(min, max, target int) bool {
	if (target >= min) && (target <= max) {
		return true
	}
	return false
}

func StringInInterval(min, max int, target string) bool {
	t, _ := strconv.Atoi(target)
	if (t >= min) && (t <= max) {
		return true
	}
	return false
}
