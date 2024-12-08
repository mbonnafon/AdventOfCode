package helpers

import (
	"math"
	"sort"
	"strconv"
)

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

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

func Median(n []int) int {
	sort.Ints(n)
	i := len(n) / 2
	if IsOdd(n[i]) {
		return n[i]
	}
	return (n[i-1] + n[i]) / 2
}

func Mean(n []int) int {
	var total float64
	for _, v := range n {
		total += float64(v)
	}
	return int(math.Round((total/float64(len(n)) - 0.5)))
}

func IsOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

// Transpose slice From (n x m) To (m x n)
func Transpose(s [][]string) [][]string {
	result := make([][]string, len(s[0]))
	for i := 0; i < len(s[0]); i++ {
		result[i] = make([]string, len(s))
		for j := 0; j < len(s); j++ {
			result[i][j] = s[j][i]
		}
	}
	return result
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
