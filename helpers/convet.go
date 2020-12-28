package helpers

import "strconv"

func StringSliceToIntSlice(s []string) []int {
	var intSlice []int
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		intSlice = append(intSlice, i)
	}
	return intSlice
}
