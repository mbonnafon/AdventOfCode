package helpers

import (
	"strconv"
)

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

// Rotate will rotate a n*n matrix of integers clockwise
// implemented https://stackoverflow.com/a/35438327 as it's O(1)
func Rotate(m [][]int) {
	size := len(m)
	for x := 0; x <= size/2; x++ {
		last := size - x - 1
		for y := x; y < last; y++ {
			offset := y - x
			top := m[x][y]

			m[x][y] = m[last-offset][x]
			m[last-offset][x] = m[last][last-offset]
			m[last][last-offset] = m[y][last]
			m[y][last] = top
		}
	}
}

// RotateString will rotate a n*n matrix of strings clockwise
func RotateString(m [][]string) {
	size := len(m)
	for x := 0; x <= size/2; x++ {
		last := size - x - 1
		for y := x; y < last; y++ {
			offset := y - x
			top := m[x][y]

			m[x][y] = m[last-offset][x]
			m[last-offset][x] = m[last][last-offset]
			m[last][last-offset] = m[y][last]
			m[y][last] = top
		}
	}
}

// FlipHorizontally will flip horizontally a n*n matrix of integers
func FlipHorizontally(s [][]int) {
	var buff []int
	for i := 0; i < len(s)/2; i++ {
		buff = s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = buff
	}
}

// FlipStringHorizontally will flip horizontally a n*n matrix of strings
func FlipStringHorizontally(s [][]string) {
	var buff []string
	for i := 0; i < len(s)/2; i++ {
		buff = s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = buff
	}
}

// FlipVertically will flip vertically a n*n matrix of integers
func FlipVertically(s [][]int) {
	var buff int
	for x := 0; x < len(s); x++ {
		for y := 0; y < len(s)/2; y++ {
			buff = s[x][y]
			s[x][y] = s[x][len(s)-1-y]
			s[x][len(s)-1-y] = buff
		}
	}
}

// FlipStringVertically will flip vertically a n*n matrix of strings
func FlipStringVertically(s [][]string) {
	var buff string
	for x := 0; x < len(s); x++ {
		for y := 0; y < len(s)/2; y++ {
			buff = s[x][y]
			s[x][y] = s[x][len(s)-1-y]
			s[x][len(s)-1-y] = buff
		}
	}
}
