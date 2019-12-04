package main

import (
	"testing"
)

func TestPass(t *testing.T) {
	got := Part2(111111, 111111)
	if got != 1 {
		t.Errorf("got = %d; want 1", got)
	}
	got1 := Part2(122223, 122223)
	if got1 != 1 {
		t.Errorf("got = %d; want 1", got1)
	}
	got2 := Part2(112233, 112233)
	if got2 != 1 {
		t.Errorf("got2 = %d; want 1", got2)
	}
	got3 := Part2(123455, 123455)
	if got3 != 1 {
		t.Errorf("got2 = %d; want 1", got3)
	}
	got4 := Part2(111122, 111122)
	if got4 != 1 {
		t.Errorf("got4 = %d; want 1", got4)
	}
	got5 := Part2(112345, 112345)
	if got5 != 1 {
		t.Errorf("got4 = %d; want 1", got5)
	}
}

func TestFail(t *testing.T) {
	got0 := Part2(123456, 123456)
	if got0 != 1 {
		t.Errorf("got1 = %d; want 1", got0)
	}
	got := Part2(654321, 654321)
	if got != 0 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
	got2 := Part2(111110, 111110)
	if got2 != 0 {
		t.Errorf("Abs(-1) = %d; want 1", got2)
	}
	got3 := Part2(011111, 011111)
	if got3 != 0 {
		t.Errorf("Abs(-1) = %d; want 1", got3)
	}
	got4 := Part2(123444, 123444)
	if got4 != 0 {
		t.Errorf("Abs(-1) = %d; want 1", got4)
	}
	got5 := Part2(223450, 223450)
	if got5 != 0 {
		t.Errorf("Abs(-1) = %d; want 1", got5)
	}
	got6 := Part2(123789, 123789)
	if got6 != 0 {
		t.Errorf("Abs(-1) = %d; want 1", got6)
	}
}
