package helpers

import "strconv"

func ToIntSlice(s []string) []int {
	var intSlice []int
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func Reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
