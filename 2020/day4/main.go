package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
}

func main() {
	s, _ := helpers.StringGroupOfLines("./input.txt")
	fmt.Printf("Part1. %d passports are valid\n", pt1(s))
	fmt.Printf("Part2. %d passports are valid\n", pt2(s))
}

func pt1(s []string) int {
	var count int
	for _, l := range s {
		if genPassport(l).validate() {
			count++
		}
	}
	return count
}

func pt2(s []string) int {
	var count int
	for _, l := range s {
		if genPassport(l).validateFields() {
			count++
		}
	}
	return count
}

func genPassport(s string) Passport {
	field := make(map[string]string, len(s))
	for _, p := range strings.Split(s, " ") {
		k, v := helpers.SplitKeyValue(p, ":")
		field[k] = v
	}
	return Passport{
		BirthYear:      field["byr"],
		IssueYear:      field["iyr"],
		ExpirationYear: field["eyr"],
		Height:         field["hgt"],
		HairColor:      field["hcl"],
		EyeColor:       field["ecl"],
		PassportID:     field["pid"],
	}
}

func (p Passport) validate() bool {
	if p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != "" {
		return true
	}
	return false
}

func (p Passport) validateFields() bool {
	if !validateBirthYear(p.BirthYear) ||
		!validateIssueYear(p.IssueYear) ||
		!validateExpirationYear(p.ExpirationYear) ||
		!validateHeight(p.Height) ||
		!validateHairColor(p.HairColor) ||
		!validateEyeColor(p.EyeColor) ||
		!validatePassportID(p.PassportID) {
		return false
	}
	return true
}

func validateBirthYear(s string) bool {
	return helpers.StringInInterval(1920, 2002, s)
}

func validateIssueYear(s string) bool {
	return helpers.StringInInterval(2010, 2020, s)
}

func validateExpirationYear(s string) bool {
	return helpers.StringInInterval(2020, 2030, s)
}

func validateHeight(s string) bool {
	var size int
	var unit string
	fmt.Sscanf(s, "%d%s", &size, &unit)
	if unit == "cm" {
		return helpers.IntInInterval(150, 193, size)
	} else if unit == "in" {
		return helpers.IntInInterval(59, 76, size)
	}
	return false
}

func validateHairColor(s string) bool {
	return regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(s)
}

func validateEyeColor(s string) bool {
	validHairColor := [...]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range validHairColor {
		if s == v {
			return true
		}
	}
	return false
}

func validatePassportID(s string) bool {
	return regexp.MustCompile("^[0-9]{9}$").MatchString(s)
}
